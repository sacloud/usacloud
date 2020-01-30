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
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/sacloud/libsacloud/utils/server"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/helper/printer"
)

var serverSSHMutex = sync.Mutex{}

func ServerSshExec(ctx command.Context, params *params.SshExecServerParam) error {

	// run as serialized
	serverSSHMutex.Lock()
	defer serverSSHMutex.Unlock()

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerSshExec is failed: %s", e)
	}

	// has NIC?
	if len(p.Interfaces) == 0 {
		return fmt.Errorf("ServerSshExec is failed: server has no network interfaces")
	}

	// file exists?
	keyPath := params.Key
	if keyPath == "" {
		p, err := getSSHPrivateKeyStorePath(p.ID)
		if err != nil {
			return fmt.Errorf("ServerSshExec is failed: getting HomeDir is failed: %s", e)
		}
		keyPath = p
	}

	// collect server IPAddress
	ip := p.Interfaces[0].IPAddress
	if ip == "" {
		ip = p.Interfaces[0].UserIPAddress
	}
	if ip == "" {
		return fmt.Errorf("ServerSshExec is failed: collecting IPAddress from server is failed: %#v", p)
	}

	// collect username
	user := params.User
	if user == "" {
		if user == "" {
			sshUser, err := getSSHDefaultUserName(client, p.ID)
			if err != nil {
				return fmt.Errorf("ServerSshExec is failed: get default ssh username is failed: %s", err)
			}
			if sshUser == "" {
				sshUser = "root"
			}
			user = sshUser
		}
	}

	displayName := fmt.Sprintf("%s(%d)", p.Name, p.ID)
	sshParam := &server.SSHClientParams{
		DisplayName:    displayName,
		UserName:       user,
		Password:       params.Password,
		Host:           ip,
		Port:           params.Port,
		PrivateKeyPath: keyPath,
		Out:            command.GlobalOption.Progress,
		Quiet:          params.Quiet,
	}
	conn, err := server.CreateSSHClient(sshParam)
	if err != nil {
		return fmt.Errorf("ServerSshExec is failed: creating ssh-client is failed: %s", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return fmt.Errorf("ServerSshExec is failed: opening session is failed: %s", err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	var args []string
	if ctx.NArgs() > 0 {
		args = append(args, ctx.Args()[1:]...)
	}

	if !params.Quiet {
		printer.Fprintf(command.GlobalOption.Progress, color.New(color.FgHiGreen), "=== start | %s ===\n", displayName)
	}
	err = session.Run(strings.Join(args, " "))
	if !params.Quiet {
		printer.Fprintf(command.GlobalOption.Progress, color.New(color.FgHiGreen), "=== end | %s ===\n", displayName)
	}

	if err != nil {
		return fmt.Errorf("ServerSshExec is failed: %s", err)
		//if ee, ok := err.(*ssh.ExitError); ok {
		//	return ee.ExitStatus()
		//}
	}

	return nil
}
