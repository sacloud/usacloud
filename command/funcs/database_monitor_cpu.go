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

func DatabaseMonitorCpu(ctx command.Context, params *params.MonitorCpuDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()
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

	res, err := api.MonitorCPU(params.Id, req)
	if err != nil {
		return fmt.Errorf("DatabaseMonitorCpu is failed: %s", err)
	}

	// collect values
	cpuValues, err := res.FlattenCPUTimeValue()
	if err != nil {
		return fmt.Errorf("DatabaseMonitorCpu is failed: %s", err)
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
		return fmt.Errorf("DatabaseMonitorCpu is failed: %s", err)
	}
	key = buf.String()

	// build sortable struct
	list := MonitorValues{}
	for i := range cpuValues {
		list = append(list, MonitorValue{
			"Key":       key,
			"TimeStamp": fmt.Sprintf("%s", cpuValues[i].Time),
			"UnixTime":  fmt.Sprintf("%d", cpuValues[i].Time.Unix()),
			"CPUTime":   fmt.Sprintf("%f", cpuValues[i].Value),
		})
	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.GetOutput().Print(values...)
}
