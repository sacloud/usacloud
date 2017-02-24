package command

import (
	"bufio"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/mattn/go-tty"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"
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
	_, err := os.Stat(keyPath)
	fileExists := err == nil

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
		if runtime.GOOS == "windows" {
			user = os.Getenv("USERNAME")
		} else {
			user = os.Getenv("USER")
		}
	}

	// collect option
	strOpt := ""
	if fileExists {
		strOpt = fmt.Sprintf(" -i %s", keyPath)
	}

	// build command string
	targetHost := fmt.Sprintf("%s:%d", ip, params.Port)

	// output connect info
	fmt.Fprintf(GlobalOption.Out, "\nconnecting server...\n\tcommand: ssh %s@%s%s\n\n", user, targetHost, strOpt)

	cnf := &ssh.ClientConfig{
		User:    user,
		Timeout: 10 * time.Second,
	}

	// build auth methods

	var authMethods []ssh.AuthMethod

	// add ssh-agent
	//sshSock := os.ExpandEnv("$SSH_AUTH_SOCK")
	//if sshSock != "" {
	//	addr, _ := net.ResolveUnixAddr("unix", sshSock)
	//	agentConn, _ := net.DialUnix("unix", nil, addr)
	//	ag := agent.NewClient(agentConn)
	//	authMethods = append(authMethods, ssh.PublicKeysCallback(ag.Signers))
	//}

	// private key
	if fileExists {

		signer, err := getSigners(keyPath, params.Password)
		if err != nil {
			return fmt.Errorf("ServerSshExec is failed: parse private-key(%s) is failed: %s", keyPath, err)
		}

		authMethods = append(authMethods, ssh.PublicKeys(signer...))

	}

	// password prompt
	authMethods = append(authMethods, ssh.PasswordCallback(func() (string, error) {
		if params.Password == "" {
			return pprompt("password: ")
		}
		return params.Password, nil
	}))

	cnf.Auth = authMethods
	var conn *ssh.Client

	if params.Proxy != "" {
		proxyUrl, err := url.Parse(params.Proxy)
		if err != nil {
			return fmt.Errorf("ServerSshExec is failed: opening session is failed: %s", err)
		}
		tcp, err := net.Dial("tcp", proxyUrl.Host)
		if err != nil {
			return fmt.Errorf("ServerSshExec is failed: opening session is failed: %s", err)
		}
		connReq := &http.Request{
			Method: "CONNECT",
			URL:    &url.URL{Path: targetHost},
			Host:   targetHost,
			Header: make(http.Header),
		}
		if proxyUrl.User != nil {
			if p, ok := proxyUrl.User.Password(); ok {
				connReq.SetBasicAuth(proxyUrl.User.Username(), p)
			}
		}
		connReq.Write(tcp)
		resp, err := http.ReadResponse(bufio.NewReader(tcp), connReq)
		if err != nil {
			return fmt.Errorf("ServerSshExec is failed: opening session is failed: %s", err)
		}
		defer resp.Body.Close()

		c, chans, reqs, err := ssh.NewClientConn(tcp, targetHost, cnf)
		if err != nil {
			return fmt.Errorf("ServerSshExec is failed: opening session is failed: %s", err)
		}
		conn = ssh.NewClient(c, chans, reqs)
	} else {
		conn, err = ssh.Dial("tcp", targetHost, cnf)
		if err != nil {
			return fmt.Errorf("ServerSshExec is failed: connecting(%s) is failed: %s", targetHost, err)
		}
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
func pprompt(prompt string) (string, error) {
	t, err := tty.Open()
	if err != nil {
		return "", err
	}
	defer t.Close()
	fmt.Print(prompt)
	defer t.Output().WriteString("\r" + strings.Repeat(" ", runewidth.StringWidth(prompt)) + "\r")
	return t.ReadPasswordClear()
}

func getSigners(keyfile string, password string) ([]ssh.Signer, error) {
	buf, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}

	b, _ := pem.Decode(buf)
	if x509.IsEncryptedPEMBlock(b) {
		pass := password
		if pass == "" {
			p, err := pprompt("pass-phrase: ")
			if err != nil {
				return nil, fmt.Errorf("ServerSsh is failed: collecting is failed: %s", err)
			}
			pass = p
		}
		buf, err = x509.DecryptPEMBlock(b, []byte(pass))
		if err != nil {
			return nil, err
		}
		pk, err := x509.ParsePKCS1PrivateKey(buf)
		if err != nil {
			return nil, err
		}
		k, err := ssh.NewSignerFromKey(pk)
		if err != nil {
			return nil, err
		}
		return []ssh.Signer{k}, nil
	}
	k, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		return nil, err
	}
	return []ssh.Signer{k}, nil
}
