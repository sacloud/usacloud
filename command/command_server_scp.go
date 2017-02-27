package command

import (
	"fmt"
	"github.com/hnakamur/go-scp"
	"github.com/sacloud/usacloud/remote"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func ServerScp(ctx Context, params *ScpServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()

	// collect args
	if ctx.NArgs() != 2 {
		return fmt.Errorf("server scp required 2 args")
	}

	src := ctx.Args()[0]
	dest := ctx.Args()[1]

	srcID, srcTokens, e := parseScpArgs(src)
	if e != nil {
		return fmt.Errorf("ServerScp is failed: %s", e)
	}
	destID, destTokens, e := parseScpArgs(dest)
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
		if runtime.GOOS == "windows" {
			user = os.Getenv("USERNAME")
		} else {
			user = os.Getenv("USER")
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
		return fmt.Errorf("ServerScp is failed: creating ssh-client is failed: %s", err)
	}
	defer conn.Close()

	scpClient := scp.NewSCP(conn)
	if localToRemote {

		// is local path is directory?
		localPath := srcTokens[0]
		remotePath := destTokens[1] // 000000000000:/path/to/remote/location

		// check local file stat
		stat, err := os.Stat(localPath)
		if err != nil {
			return fmt.Errorf("ServerScp is failed: %s", e)
		}
		if stat.IsDir() {

			localPath = filepath.Clean(localPath)
			err := scpClient.SendDir(localPath, remotePath, func(parentDir string, info os.FileInfo) (bool, error) {
				if params.Recursive {
					return true, nil
				}

				current := filepath.Join(parentDir, info.Name())
				return localPath == current || (localPath == parentDir && !info.IsDir()), nil

			})
			if err != nil {
				return fmt.Errorf("ServerScp is failed: %s", e)
			}
		} else {
			if strings.HasSuffix(remotePath, "/") {
				remotePath = fmt.Sprintf("%s%s", remotePath, filepath.Base(localPath))
			}
			err := scpClient.SendFile(localPath, remotePath)
			if err != nil {
				return fmt.Errorf("ServerScp is failed: %s", e)
			}
		}

	} else {
		// is local path is directory?
		localPath := destTokens[0]
		remotePath := srcTokens[1] // 000000000000:/path/to/remote/location

		// create dir
		err := os.MkdirAll(filepath.Dir(localPath), 0755)
		if err != nil {
			return fmt.Errorf("ServerScp is failed: %s", e)
		}
		// first, try copy file
		err = scpClient.ReceiveFile(remotePath, localPath)
		if err != nil {
			// next , try copy directory
			err := scpClient.ReceiveDir(remotePath, localPath, func(parentDir string, info os.FileInfo) (bool, error) {

				if params.Recursive {
					return true, nil
				}

				current := filepath.Join(parentDir, info.Name())
				return localPath == current || (localPath == parentDir && !info.IsDir()), nil

			})
			if err != nil {
				return fmt.Errorf("ServerScp is failed: %s", e)
			}
		}
	}

	return nil
}

func parseScpArgs(arg string) (int64, []string, error) {

	tokens := strings.Split(arg, ":")

	if len(tokens) == 2 {

		strID := tokens[0]
		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			return -1, []string{}, fmt.Errorf("ID is invalid: %s", err)
		}
		if len(fmt.Sprintf("%d", id)) != 12 {
			return -1, []string{}, fmt.Errorf("ID is invalid: %s", err)
		}

		return id, tokens, nil
	}

	return -1, tokens, nil

}
