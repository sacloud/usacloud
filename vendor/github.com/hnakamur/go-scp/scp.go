package scp

import "golang.org/x/crypto/ssh"

// SCP is the type for the SCP client.
type SCP struct {
	client *ssh.Client
}

// NewSCP creates the SCP client.
func NewSCP(client *ssh.Client) *SCP {
	return &SCP{
		client: client,
	}
}
