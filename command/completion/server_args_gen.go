package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerListCompleteArgs(ctx command.Context, params *params.ListServerParam, cur, prev, commandName string) {

}

func ServerBuildCompleteArgs(ctx command.Context, params *params.BuildServerParam, cur, prev, commandName string) {

}

func ServerReadCompleteArgs(ctx command.Context, params *params.ReadServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerUpdateCompleteArgs(ctx command.Context, params *params.UpdateServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerDeleteCompleteArgs(ctx command.Context, params *params.DeleteServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerPlanChangeCompleteArgs(ctx command.Context, params *params.PlanChangeServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerBootCompleteArgs(ctx command.Context, params *params.BootServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerShutdownCompleteArgs(ctx command.Context, params *params.ShutdownServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerShutdownForceCompleteArgs(ctx command.Context, params *params.ShutdownForceServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerResetCompleteArgs(ctx command.Context, params *params.ResetServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerWaitForBootCompleteArgs(ctx command.Context, params *params.WaitForBootServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerWaitForDownCompleteArgs(ctx command.Context, params *params.WaitForDownServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerSshCompleteArgs(ctx command.Context, params *params.SshServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerSshExecCompleteArgs(ctx command.Context, params *params.SshExecServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerScpCompleteArgs(ctx command.Context, params *params.ScpServerParam, cur, prev, commandName string) {

}

func ServerVncCompleteArgs(ctx command.Context, params *params.VncServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerVncInfoCompleteArgs(ctx command.Context, params *params.VncInfoServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerVncSendCompleteArgs(ctx command.Context, params *params.VncSendServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerDiskInfoCompleteArgs(ctx command.Context, params *params.DiskInfoServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerDiskConnectCompleteArgs(ctx command.Context, params *params.DiskConnectServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerDiskDisconnectCompleteArgs(ctx command.Context, params *params.DiskDisconnectServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerInterfaceInfoCompleteArgs(ctx command.Context, params *params.InterfaceInfoServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerInterfaceAddForInternetCompleteArgs(ctx command.Context, params *params.InterfaceAddForInternetServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerInterfaceAddForRouterCompleteArgs(ctx command.Context, params *params.InterfaceAddForRouterServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerInterfaceAddForSwitchCompleteArgs(ctx command.Context, params *params.InterfaceAddForSwitchServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerInterfaceAddDisconnectedCompleteArgs(ctx command.Context, params *params.InterfaceAddDisconnectedServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerIsoInfoCompleteArgs(ctx command.Context, params *params.IsoInfoServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerIsoInsertCompleteArgs(ctx command.Context, params *params.IsoInsertServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerIsoEjectCompleteArgs(ctx command.Context, params *params.IsoEjectServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerMonitorCpuCompleteArgs(ctx command.Context, params *params.MonitorCpuServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerMonitorNicCompleteArgs(ctx command.Context, params *params.MonitorNicServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerMonitorDiskCompleteArgs(ctx command.Context, params *params.MonitorDiskServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Servers {
		fmt.Println(res.Servers[i].ID)
		var target interface{} = &res.Servers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func ServerMaintenanceInfoCompleteArgs(ctx command.Context, params *params.MaintenanceInfoServerParam, cur, prev, commandName string) {

}
