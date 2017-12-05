// +build !windows,!linux

package self

import "fmt"

func ID() (string, error) {
	return "", fmt.Errorf("self is only supported windows/linux")
}
