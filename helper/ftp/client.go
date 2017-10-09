package ftp

import (
	"errors"
	"fmt"
	"github.com/sacloud/ftps"
	"path/filepath"
	"strings"
)

type Client struct {
	UserName string
	Password string
	Host     string
}

func NewClient(user string, pass string, host string) *Client {
	return &Client{
		UserName: user,
		Password: pass,
		Host:     host,
	}
}

func (c *Client) Upload(filePath string) error {
	ftpsClient := &ftps.FTPS{}
	ftpsClient.TLSConfig.InsecureSkipVerify = true

	err := ftpsClient.Connect(c.Host, 21)
	if err != nil {
		return fmt.Errorf("Connect FTP failed: %#v", err)
	}

	err = ftpsClient.Login(c.UserName, c.Password)
	if err != nil {
		return fmt.Errorf("Auth FTP failed: %#v", err)
	}

	err = ftpsClient.StoreFile(filepath.Base(filePath), filePath)
	if err != nil {
		return fmt.Errorf("Storefile FTP failed: %#v", err)
	}

	err = ftpsClient.Quit()
	if err != nil {
		return fmt.Errorf("Quit FTP failed: %#v", err)
	}

	return nil
}

func (c *Client) Download(filePath string) error {
	ftpsClient := &ftps.FTPS{}
	ftpsClient.TLSConfig.InsecureSkipVerify = true

	err := ftpsClient.Connect(c.Host, 21)
	if err != nil {
		return fmt.Errorf("Connect FTP failed: %#v", err)
	}

	err = ftpsClient.Login(c.UserName, c.Password)
	if err != nil {
		return fmt.Errorf("Auth FTP failed: %#v", err)
	}

	entries, err := ftpsClient.List()
	if err != nil {
		return fmt.Errorf("FTP List Entry failed: %#v", err)
	}

	var serverFilePath string
	for _, e := range entries {
		if e.Type == ftps.EntryTypeFile && !strings.HasPrefix(e.Name, ".") {
			serverFilePath = e.Name
			break
		}
	}
	if serverFilePath == "" {
		return errors.New("FTP retrieve filename failed")
	}

	// download
	err = ftpsClient.RetrieveFile(serverFilePath, filePath)
	if err != nil {
		return fmt.Errorf("FTP download file is failed: %#v", err)
	}

	err = ftpsClient.Quit()
	if err != nil {
		return fmt.Errorf("Quit FTP failed: %#v", err)
	}

	return nil
}
