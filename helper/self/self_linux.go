// +build linux

package self

import (
	"io/ioutil"
)

const serialFilePath = "/sys/devices/virtual/dmi/id/product_serial"

func ID() (string, error) {

	data, err := ioutil.ReadFile(serialFilePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
