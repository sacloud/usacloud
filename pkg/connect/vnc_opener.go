// Copyright 2017-2022 The sacloud/usacloud Authors
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

//go:build !windows && !linux
// +build !windows,!linux

package connect

import (
	"fmt"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/skratchdot/open-golang/open"
)

// StartDefaultVNCClient starts OS's default VNC client
func StartDefaultVNCClient(vncProxyInfo *iaas.VNCProxyInfo) error {
	host := vncProxyInfo.Host
	if host == "localhost" {
		host = vncProxyInfo.IOServerHost
	}
	uri := fmt.Sprintf("vnc://:%s@%s:%s",
		vncProxyInfo.Password,
		host,
		vncProxyInfo.Port)

	err := open.Start(uri)
	time.Sleep(time.Second)
	return err
}
