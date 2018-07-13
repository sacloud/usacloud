package funcs

import (
	"bytes"
	"fmt"
	"sort"
	"text/template"
	"time"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SIMMonitor(ctx command.Context, params *params.MonitorSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()

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
		return fmt.Errorf("SimMonitor is failed: %s", err)
	}

	// collect values
	uplinkValues, err := res.FlattenUplinkBPSValue()
	if err != nil {
		return fmt.Errorf("SimMonitor is failed: %s", err)
	}
	downlinkValues, err := res.FlattenDownlinkBPSValue()
	if err != nil {
		return fmt.Errorf("SimMonitor is failed: %s", err)
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
		return fmt.Errorf("SimMonitor is failed: %s", err)
	}
	key = buf.String()

	// build sortable struct
	list := MonitorValues{}
	for i := range uplinkValues {
		list = append(list, MonitorValue{
			"Key":         key,
			"TimeStamp":   fmt.Sprintf("%s", uplinkValues[i].Time),
			"UnixTime":    fmt.Sprintf("%d", uplinkValues[i].Time.Unix()),
			"UplinkBPS":   fmt.Sprintf("%f", uplinkValues[i].Value),
			"DownlinkBPS": fmt.Sprintf("%f", downlinkValues[i].Value),
		})
	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.GetOutput().Print(values...)
}
