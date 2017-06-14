package funcs

import (
	"bytes"
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"sort"
	"text/template"
	"time"
)

func LoadBalancerMonitor(ctx command.Context, params *params.MonitorLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()

	end := parseDateTimeString(params.End)
	start := end.Add(-1 * time.Hour)
	if params.Start != "" {
		start = parseDateTimeString(params.Start)
	}

	// validate start <= end
	if !(start.Unix() <= end.Unix()) {
		return fmt.Errorf("Invalid Parameter : start(%s) or end(%s) is invalid", start, end)
	}

	req := sacloud.NewResourceMonitorRequest(&start, &end)

	res, err := api.Monitor(params.Id, req)
	if err != nil {
		return fmt.Errorf("LoadBalancerMonitor is failed: %s", err)
	}

	// collect values
	receiveValues, err := res.FlattenPacketReceiveValue()
	if err != nil {
		return fmt.Errorf("LoadBalancerMonitor is failed: %s", err)
	}
	sendValues, err := res.FlattenPacketSendValue()
	if err != nil {
		return fmt.Errorf("LoadBalancerMonitor is failed: %s", err)
	}

	// build key string
	key := ""
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(params.KeyFormat))
	err = t.Execute(buf, map[string]interface{}{
		"ID": params.Id,
	})
	if err != nil {
		return fmt.Errorf("LoadBalancerMonitor is failed: %s", err)
	}
	key = buf.String()

	// build sortable struct
	list := MonitorValues{}
	for i := range receiveValues {
		list = append(list, MonitorValue{
			"Key":       key,
			"TimeStamp": fmt.Sprintf("%s", receiveValues[i].Time),
			"UnixTime":  fmt.Sprintf("%d", receiveValues[i].Time.Unix()),
			"Receive":   fmt.Sprintf("%f", receiveValues[i].Value),
			"Send":      fmt.Sprintf("%f", sendValues[i].Value),
		})
	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.GetOutput().Print(values...)

}
