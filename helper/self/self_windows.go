// +build windows

package self

import (
	"fmt"
	"os/exec"
	"regexp"
)

func ID() (string, error) {
	o, err := exec.Command("wmic", "path", "win32_computersystemproduct", "get", "IdentifyingNumber", "/VALUE").Output()
	if err != nil {
		return "", err
	}
	r := regexp.MustCompile(`IdentifyingNumber=([0-9]{12})`)
	groups := r.FindStringSubmatch(string(o))
	if len(groups) > 1 {
		return groups[1], nil
	}
	return "", fmt.Errorf("Can't find IdentifyingNumber(WMI)")
}
