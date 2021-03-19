// Copyright 2017-2021 The Libsacloud Authors
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

// +build linux windows

package connect

import (
	"fmt"
	"os"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/skratchdot/open-golang/open"
)

// StartDefaultVNCClient starts OS's default VNC client
func StartDefaultVNCClient(vncProxyInfo *sacloud.VNCProxyInfo) error {
	uri := ""

	for uri == "" {
		// create .vnc tmp-file
		f, err := os.CreateTemp("", "libsacloud_open_vnc")
		if err != nil {
			return err
		}
		defer f.Close()
		defer os.Remove(f.Name())
		uri = fmt.Sprintf("%s.vnc", f.Name())
		if _, err := os.Stat(uri); err == nil {
			uri = ""
		}
	}
	host := vncProxyInfo.Host
	if host == "localhost" {
		host = vncProxyInfo.IOServerHost
	}
	body := fmt.Sprintf(vncFileFormat,
		host,
		vncProxyInfo.Port,
		vncProxyInfo.Password,
	)

	if err := os.WriteFile(uri, []byte(body), 0700); err != nil {
		return err
	}
	defer os.Remove(uri)

	err := open.Start(uri)
	time.Sleep(time.Second)
	return err
}

var vncFileFormat = `[Connection]
Host=%s
Port=%s
Password=%s
`
