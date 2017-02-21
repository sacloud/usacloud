package command

import (
	"fmt"
	"io"
	"os"
)

func SSHKeyGenerate(ctx Context, params *GenerateSSHKeyParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSSHKeyAPI()

	// call manipurate functions
	key, err := api.Generate(params.Name, params.PassPhrase, params.Description)
	if err != nil {
		return fmt.Errorf("SSHKeyGenerate is failed: %s", err)
	}

	var w io.Writer
	if params.PrivateKeyOutput != "" {
		// file
		f, err := os.OpenFile(params.PrivateKeyOutput, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return fmt.Errorf("SSHKeyGenerate is failed: %s", err)
		}
		w = f
		defer f.Close()

		_, err = w.Write([]byte(key.PrivateKey))
		if err != nil {
			return fmt.Errorf("SSHKeyGenerate is failed: %s", err)
		}
	}

	err = ctx.GetOutput().Print(key)
	if err != nil {
		return fmt.Errorf("SSHKeyGenerate is failed: %s", err)
	}

	if params.PrivateKeyOutput == "" {
		// output privatekey to os.StdOut
		fmt.Fprintf(GlobalOption.Out, sshPrivateKeyStdOutFormat, key.PrivateKey)
	}

	return err
}

var sshPrivateKeyStdOutFormat = `
=======================================
Please save the following private-key.
=======================================

%s
`
