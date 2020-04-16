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

// +build darwin linux windows

package funcs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sacloud/libsacloud/sacloud"

	"github.com/hnakamur/go-scp"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/utils/server"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerScp(ctx command.Context, params *params.ScpServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()

	// collect args
	if ctx.NArgs() != 2 {
		return fmt.Errorf("server scp required 2 args")
	}

	src := ctx.Args()[0]
	dest := ctx.Args()[1]

	srcID, srcTokens, e := parseScpArgs(api, src)
	if e != nil {
		return fmt.Errorf("ServerScp is failed: %s", e)
	}
	destID, destTokens, e := parseScpArgs(api, dest)
	if e != nil {
		return fmt.Errorf("ServerScp is failed: %s", e)
	}
	if srcID > 0 && destID > 0 {
		return fmt.Errorf("Server ID can be specified as either source or destination")
	} else if srcID < 0 && destID < 0 {
		return fmt.Errorf("Server ID is required for either source or destination")
	}

	id := srcID
	localToRemote := false
	if destID > 0 {
		id = destID
		localToRemote = true
	}

	p, e := api.Read(id)
	if e != nil {
		return fmt.Errorf("ServerScp is failed: %s", e)
	}

	// has NIC?
	if len(p.Interfaces) == 0 {
		return fmt.Errorf("ServerScp is failed: server has no network interfaces")
	}

	// file exists?
	keyPath := params.Key
	if keyPath == "" {
		p, err := getSSHPrivateKeyStorePath(p.ID)
		if err != nil {
			return fmt.Errorf("ServerScp is failed: getting HomeDir is failed: %s", e)
		}
		keyPath = p
	}

	// collect server IPAddress
	ip := p.Interfaces[0].IPAddress
	if ip == "" {
		ip = p.Interfaces[0].UserIPAddress
	}
	if ip == "" {
		return fmt.Errorf("ServerScp is failed: collecting IPAddress from server is failed: %#v", p)
	}

	// collect username
	user := params.User
	if user == "" {
		sshUser, err := getSSHDefaultUserName(client, p.ID)
		if err != nil {
			return fmt.Errorf("ServerScp is failed: get default ssh username is failed: %s", err)
		}
		if sshUser == "" {
			sshUser = "root"
		}
		user = sshUser
	}

	sshParam := &server.SSHClientParams{
		UserName:       user,
		Password:       params.Password,
		Host:           ip,
		Port:           params.Port,
		PrivateKeyPath: keyPath,
		Out:            ctx.IO().Progress(),
		Quiet:          params.Quiet,
	}
	conn, err := server.CreateSSHClient(sshParam)
	if err != nil {
		return fmt.Errorf("ServerScp is failed: creating ssh-client is failed: %s", err)
	}
	defer conn.Close()

	scpClient := scp.NewSCP(conn)
	if localToRemote {

		// is local path is directory?
		localPath := strings.Join(srcTokens[0:], ":")
		remotePath := strings.Join(destTokens[1:], ":") // 000000000000:/path/to/remote/location

		// check local file stat
		stat, err := os.Stat(localPath)
		if err != nil {
			return fmt.Errorf("ServerScp is failed: %s", err)
		}
		if stat.IsDir() {

			if !params.Recursive {
				return fmt.Errorf("%q is directory. Use -r or --recursive flag", localPath)
			}

			localPath = filepath.Clean(localPath)
			err := scpClient.SendDir(localPath, remotePath, func(parentDir string, info os.FileInfo) (bool, error) {
				if !info.IsDir() {
					fmt.Fprintf(ctx.IO().Progress(), "copy: %s\n", filepath.Join(parentDir, info.Name()))
				}
				return true, nil
			})
			if err != nil {
				return fmt.Errorf("ServerScp is failed: %s", err)
			}
		} else {
			if strings.HasSuffix(remotePath, "/") {
				remotePath = fmt.Sprintf("%s%s", remotePath, filepath.Base(localPath))
			}
			fmt.Fprintf(ctx.IO().Progress(), "copy: %s\n", localPath)
			err := scpClient.SendFile(localPath, remotePath)
			if err != nil {
				return fmt.Errorf("ServerScp is failed: %s", err)
			}
		}

	} else {
		// is local path is directory?
		localPath := strings.Join(destTokens[0:], ":")
		remotePath := strings.Join(srcTokens[1:], ":") // 000000000000:/path/to/remote/location

		// create dir
		err := os.MkdirAll(filepath.Dir(localPath), 0755)
		if err != nil {
			return fmt.Errorf("ServerScp is failed: %s", err)
		}
		// first, try copy file
		err = scpClient.ReceiveFile(remotePath, localPath)
		if err == nil {
			fmt.Fprintf(ctx.IO().Progress(), "copy: %s\n", localPath)
		} else {
			if !params.Recursive {
				return fmt.Errorf("%q isn't readable file or is a directory. try use -r or --recursive flag", remotePath)
			}

			// next , try copy directory
			err := scpClient.ReceiveDir(remotePath, localPath, func(parentDir string, info os.FileInfo) (bool, error) {
				if !info.IsDir() {
					fmt.Fprintf(ctx.IO().Progress(), "copy: %s\n", filepath.Join(parentDir, info.Name()))
				}
				return true, nil
			})
			if err != nil {
				return fmt.Errorf("ServerScp is failed: %s", err)
			}
		}
	}

	return nil
}

func parseScpArgs(api *api.ServerAPI, arg string) (sacloud.ID, []string, error) {

	tokens := strings.Split(arg, ":")

	if len(tokens) > 1 {

		strID := tokens[0]
		id := sacloud.StringID(strID)
		if !id.IsEmpty() {
			api.Reset()
			api.SetNameLike(strID)
			res, err := api.Find()
			if err != nil {
				return -1, []string{}, fmt.Errorf("Find server[%s] is failed: %s", strID, err)
			}

			switch len(res.Servers) {
			case 0:
				return -1, tokens, nil // for windows like path (C://...)
			case 1:
				id = res.Servers[0].ID
			default:
				return -1, []string{}, fmt.Errorf("scp command can't specify more than one server")
			}

			if len(res.Servers) == 0 {
			}

		}
		if len(fmt.Sprintf("%d", id)) != 12 {
			return -1, []string{}, fmt.Errorf("ID[%q] is invalid format", strID)
		}

		return id, tokens, nil
	}

	return -1, tokens, nil

}
