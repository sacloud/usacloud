// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package funcs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"
)

func setSortBy(target sortable, key string) {
	reverse := strings.HasPrefix(key, "-")
	key = strings.Replace(key, "-", "", -1)
	target.SetSortBy(key, reverse)
}

type sortable interface {
	SetSortBy(key string, reverse bool)
}

func getSSHPrivateKeyStorePath(serverID sacloud.ID) (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("getting HomeDir is failed:%s", err)
	}
	return filepath.Join(homeDir, ".ssh", fmt.Sprintf("sacloud_pkey_%d", serverID)), nil
}

func getSSHDefaultUserName(client *api.Client, serverID sacloud.ID) (string, error) {

	// read server
	server, err := client.GetServerAPI().Read(serverID)
	if err != nil {
		return "", err
	}

	if len(server.Disks) == 0 {
		return "", nil
	}

	return getSSHDefaultUserNameDiskRec(client, server.Disks[0].ID)
}

func getSSHDefaultUserNameDiskRec(client *api.Client, diskID sacloud.ID) (string, error) {

	disk, err := client.GetDiskAPI().Read(diskID)
	if err != nil {
		return "", err
	}

	if disk.SourceDisk != nil {
		return getSSHDefaultUserNameDiskRec(client, disk.SourceDisk.ID)
	}

	if disk.SourceArchive != nil {
		return getSSHDefaultUserNameArchiveRec(client, disk.SourceArchive.ID)

	}

	return "", nil
}

func getSSHDefaultUserNameArchiveRec(client *api.Client, archiveID sacloud.ID) (string, error) {
	// read archive
	archive, err := client.GetArchiveAPI().Read(archiveID)
	if err != nil {
		return "", err
	}

	if archive.Scope == string(sacloud.ESCopeShared) {

		// has ubuntu/coreos tag?
		if archive.HasTag("distro-ubuntu") {
			return "ubuntu", nil
		}

		if archive.HasTag("distro-vyos") {
			return "vyos", nil
		}

		if archive.HasTag("distro-coreos") {
			return "core", nil
		}

		if archive.HasTag("distro-rancheros") {
			return "rancher", nil
		}

		if archive.HasTag("distro-k3os") {
			return "rancher", nil
		}
	}
	if archive.SourceDisk != nil {
		return getSSHDefaultUserNameDiskRec(client, archive.SourceDisk.ID)
	}

	if archive.SourceArchive != nil {
		return getSSHDefaultUserNameArchiveRec(client, archive.SourceArchive.ID)
	}
	return "", nil

}

func parseDateTimeString(strDateTime string) time.Time {
	allowDatetimeFormatList := []string{
		time.RFC3339,
	}

	if strDateTime != "" {
		for _, format := range allowDatetimeFormatList {
			d, err := time.Parse(format, strDateTime)
			if err == nil {
				// success
				return d
			}
		}
	}

	return time.Now()
}

func fileOrStdin(path string) (file *os.File, deferFunc func(), err error) {
	if path == "" || path == "-" {
		file = os.Stdin // TODO ビルドを通すための仮実装
		deferFunc = func() {}
	} else {
		file, err = os.Open(path)
		if err != nil {
			return
		}
		deferFunc = func() {
			file.Close()
		}
	}
	return
}
