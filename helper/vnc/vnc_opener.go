// +build !windows,!linux

package vnc

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/skratchdot/open-golang/open"
)

func OpenVNCClient(vncProxyInfo *sacloud.VNCProxyResponse) error {
	uri := fmt.Sprintf("vnc://:%s@%s:%s",
		vncProxyInfo.Password,
		vncProxyInfo.Host,
		vncProxyInfo.Port)
	return open.Start(uri)
}
