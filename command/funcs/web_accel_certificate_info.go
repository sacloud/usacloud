package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func WebAccelCertificateInfo(ctx command.Context, params *params.CertificateInfoWebAccelParam) error {

	client := ctx.GetAPIClient()
	api := client.GetWebAccelAPI()
	p, e := api.ReadCertificate(fmt.Sprintf("%d", params.Id))
	if e != nil {
		return fmt.Errorf("WebAccelCertificateInfo is failed: %s", e)
	}

	if p == nil {
		fmt.Fprintf(command.GlobalOption.Err, "Result is empty\n")
		return nil
	}

	return ctx.GetOutput().Print(p)
}
