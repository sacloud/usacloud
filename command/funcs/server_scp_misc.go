// +build !darwin,!linux,!windows

package funcs

import (
	"errors"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerScp(ctx command.Context, params *params.ScpServerParam) error {
	return errors.New("this platform does not support scp")
}
