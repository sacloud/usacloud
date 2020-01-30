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
	"os/exec"
	"strings"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerSsh(ctx command.Context, params *params.SshServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerSsh is failed: %s", e)
	}

	// has NIC?
	if len(p.Interfaces) == 0 {
		return fmt.Errorf("ServerSsh is failed: server has no network interfaces")
	}

	// file exists?
	keyPath := params.Key
	if keyPath == "" {
		p, err := getSSHPrivateKeyStorePath(p.ID)
		if err != nil {
			return fmt.Errorf("ServerSsh is failed: getting HomeDir is failed: %s", e)
		}
		keyPath = p
	}
	_, err := os.Stat(keyPath)
	fileExists := err == nil

	// collect server IPAddress
	ip := p.Interfaces[0].IPAddress
	if ip == "" {
		ip = p.Interfaces[0].UserIPAddress
	}
	if ip == "" {
		return fmt.Errorf("ServerSsh is failed: collecting IPAddress from server is failed: %#v", p)
	}

	// collect username
	user := params.User
	if user == "" {
		if user == "" {
			sshUser, _ := getSSHDefaultUserName(client, p.ID)
			//if err != nil {
			//	return fmt.Errorf("ServerSsh is failed: get default ssh username is failed: %s", err)
			//}
			if sshUser == "" {
				sshUser = "root"
			}
			user = sshUser
		}
	}

	// exec/spawn a ssh session
	args := []string{fmt.Sprintf("%s@%s", user, ip)}
	if fileExists {
		args = append(args, "-i", keyPath)
	}
	if params.Port != 22 {
		args = append(args, "-p", fmt.Sprintf("%d", params.Port))
	}
	if ctx.NArgs() > 0 {
		args = append(args, ctx.Args()[1:]...)
	}

	// output connect info
	if !params.Quiet {
		fmt.Fprintf(command.GlobalOption.Progress, "connecting server...\n\tcommand: ssh %s \n", strings.Join(args, " "))
	}

	cmd := exec.Command("ssh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()

	return err
}
