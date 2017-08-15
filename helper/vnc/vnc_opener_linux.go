// +build linux

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

	body := fmt.Sprintf(vncFileFormat,
		vncProxyInfo.Host,
		vncProxyInfo.Port,
		vncProxyInfo.Password,
	)

	ioutil.WriteFile(uri, []byte(body), 0755)
	return open.Start(uri)
}

var vncFileFormat = `[Connection]
Host=%s
Port=%s
Password=%s
`
