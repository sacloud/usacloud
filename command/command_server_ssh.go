package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ServerSsh(ctx Context, params *SshServerParam) error {

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
		fmt.Fprintf(GlobalOption.Progress, "connecting server...\n\tcommand: ssh %s \n", strings.Join(args, " "))
	}

	cmd := exec.Command("ssh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()

	return err
}
