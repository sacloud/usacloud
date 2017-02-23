package command

import (
	"fmt"
)

func ServerInterfaceInfo(ctx Context, params *InterfaceInfoServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerInterfaceInfo is failed: %s", e)
	}

	interfaces := p.GetInterfaces()
	if len(interfaces) == 0 {
		fmt.Fprintf(GlobalOption.Err, "Server don't have any interfaces\n")
		return nil
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range interfaces {
		list = append(list, &interfaces[i])
	}

	return ctx.GetOutput().Print(list...)

}
