// +build windows

package vnc

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"os"
)

func OpenVNCClient(vncProxyInfo *sacloud.VNCProxyResponse) error {

	uri := ""

	for uri == "" {
		// create .vnc tmp-file
		f, err := ioutil.TempFile("", "usacloud_open_vnc")
		if err != nil {
			return err
		}
		defer f.Close()
		uri = fmt.Sprintf("%s.vnc", f.Name())
		if _, err := os.Stat(uri); err == nil {
			uri = ""
		}
	}
	ioutil.WriteFile(uri, []byte(vncProxyInfo.VNCFile), 0755)
	return open.Start(uri)
}
