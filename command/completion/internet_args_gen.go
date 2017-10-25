package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func InternetListCompleteArgs(ctx command.Context, params *params.ListInternetParam, cur, prev, commandName string) {

}

func InternetCreateCompleteArgs(ctx command.Context, params *params.CreateInternetParam, cur, prev, commandName string) {

}

func InternetReadCompleteArgs(ctx command.Context, params *params.ReadInternetParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInternetAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Internet {
		fmt.Println(res.Internet[i].ID)
		var target interface{} = &res.Internet[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InternetUpdateCompleteArgs(ctx command.Context, params *params.UpdateInternetParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInternetAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Internet {
		fmt.Println(res.Internet[i].ID)
		var target interface{} = &res.Internet[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InternetDeleteCompleteArgs(ctx command.Context, params *params.DeleteInternetParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInternetAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Internet {
		fmt.Println(res.Internet[i].ID)
		var target interface{} = &res.Internet[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InternetUpdateBandwidthCompleteArgs(ctx command.Context, params *params.UpdateBandwidthInternetParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInternetAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Internet {
		fmt.Println(res.Internet[i].ID)
		var target interface{} = &res.Internet[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InternetSubnetInfoCompleteArgs(ctx command.Context, params *params.SubnetInfoInternetParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInternetAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Internet {
		fmt.Println(res.Internet[i].ID)
		var target interface{} = &res.Internet[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InternetSubnetAddCompleteArgs(ctx command.Context, params *params.SubnetAddInternetParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInternetAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Internet {
		fmt.Println(res.Internet[i].ID)
		var target interface{} = &res.Internet[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InternetSubnetDeleteCompleteArgs(ctx command.Context, params *params.SubnetDeleteInternetParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInternetAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Internet {
		fmt.Println(res.Internet[i].ID)
		var target interface{} = &res.Internet[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InternetSubnetUpdateCompleteArgs(ctx command.Context, params *params.SubnetUpdateInternetParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInternetAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Internet {
		fmt.Println(res.Internet[i].ID)
		var target interface{} = &res.Internet[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InternetIpv6InfoCompleteArgs(ctx command.Context, params *params.Ipv6InfoInternetParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInternetAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Internet {
		fmt.Println(res.Internet[i].ID)
		var target interface{} = &res.Internet[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InternetIpv6EnableCompleteArgs(ctx command.Context, params *params.Ipv6EnableInternetParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInternetAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Internet {
		fmt.Println(res.Internet[i].ID)
		var target interface{} = &res.Internet[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InternetIpv6DisableCompleteArgs(ctx command.Context, params *params.Ipv6DisableInternetParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInternetAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Internet {
		fmt.Println(res.Internet[i].ID)
		var target interface{} = &res.Internet[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func InternetMonitorCompleteArgs(ctx command.Context, params *params.MonitorInternetParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInternetAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Internet {
		fmt.Println(res.Internet[i].ID)
		var target interface{} = &res.Internet[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
