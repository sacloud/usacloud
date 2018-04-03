package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/helper/self"
)

func SelfInfo(ctx command.Context, params *params.InfoSelfParam) error {
	id, err := self.ID()
	if err != nil {
		return fmt.Errorf("SelfInfo is failed: %s", err)
	}
	fmt.Fprint(command.GlobalOption.Out, id)
	return nil
}
