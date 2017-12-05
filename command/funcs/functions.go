package funcs

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func setSortBy(target sortable, key string) {
	reverse := strings.HasPrefix(key, "-")
	key = strings.Replace(key, "-", "", -1)
	target.SetSortBy(key, reverse)
}

type sortable interface {
	SetSortBy(key string, reverse bool)
}

func getSSHPrivateKeyStorePath(serverID int64) (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("getting HomeDir is failed:%s", err)
	}
	return filepath.Join(homeDir, ".ssh", fmt.Sprintf("sacloud_pkey_%d", serverID)), nil
}

func getSSHDefaultUserName(client *api.Client, serverID int64) (string, error) {

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

func getSSHDefaultUserNameDiskRec(client *api.Client, diskID int64) (string, error) {

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

func getSSHDefaultUserNameArchiveRec(client *api.Client, archiveID int64) (string, error) {
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
		file = command.GlobalOption.In
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
