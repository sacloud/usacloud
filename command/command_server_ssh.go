package command

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ServerSsh(ctx Context, params *SshServerParam) error {

	if runtime.GOOS == "windows" {
		return fmt.Errorf("This command can't be executed on Windows. Please use ssh-exec command.")
	}

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerSsh is failed: %s", e)
	}

	// has NIC?
	if len(p.Interfaces) == 0 {
		return fmt.Errorf("ServerSsh is failed: server has no network interfaces:")
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
		user = os.Getenv("USER")
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
		if ctx.Args()[0] == fmt.Sprintf("%d", params.Id) {
			args = append(args, ctx.Args()[1:]...)
		} else {
			args = append(args, ctx.Args()...)
		}
	}

	// output connect info
	fmt.Fprintf(GlobalOption.Out, "connecting server...\n\tcommand: %sssh \n", strings.Join(args, " "))

	cmd := exec.Command("ssh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()

	return err
}
