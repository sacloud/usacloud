package command

import (
	"fmt"
)

func AuthStatusShow(ctx Context, params *ShowAuthStatusParam) error {
	client := ctx.GetAPIClient()

	// auth-status
	auth, err := client.GetAuthStatusAPI().Read()
	if err != nil {
		return fmt.Errorf("AuthStatus is failed: %s", err)
	}

	return ctx.GetOutput().Print(auth)
}
