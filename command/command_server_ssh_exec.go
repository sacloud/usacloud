package command

import (
	"fmt"
	"github.com/sacloud/usacloud/remote"
	"os"
	"strings"
)

func ServerSshExec(ctx Context, params *SshExecServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerSshExec is failed: %s", e)
	}

	// has NIC?
	if len(p.Interfaces) == 0 {
		return fmt.Errorf("ServerSshExec is failed: server has no network interfaces:")
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

	sshParam := &remote.SSHParams{
		UserName:       user,
		Password:       params.Password,
		Host:           ip,
		Port:           params.Port,
		PrivateKeyPath: keyPath,
		Out:            GlobalOption.Out,
		Quiet:          params.Quiet,
	}
	conn, err := remote.CreateSSHClient(sshParam)
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
		if ctx.Args()[0] == fmt.Sprintf("%d", params.Id) {
			args = append(args, ctx.Args()[1:]...)
		} else {
			args = append(args, ctx.Args()...)
		}
	}

	err = session.Run(strings.Join(args, " "))

	if err != nil {
		return fmt.Errorf("ServerSshExec is failed: %s", err)
		//if ee, ok := err.(*ssh.ExitError); ok {
		//	return ee.ExitStatus()
		//}
	}

	return nil
}
