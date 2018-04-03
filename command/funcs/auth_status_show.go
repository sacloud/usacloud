package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func AuthStatusShow(ctx command.Context, params *params.ShowAuthStatusParam) error {
	client := ctx.GetAPIClient()

	// auth-status
	auth, err := client.GetAuthStatusAPI().Read()
	if err != nil {
		return fmt.Errorf("AuthStatus is failed: %s", err)
	}

	return ctx.GetOutput().Print(auth)
}
