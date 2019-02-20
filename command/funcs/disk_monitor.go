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

func DiskMonitor(ctx command.Context, params *params.MonitorDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()

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
		return fmt.Errorf("DiskMonitor is failed: %s", err)
	}

	// collect values
	reads, err := res.FlattenDiskReadValue()
	if err != nil {
		return fmt.Errorf("DiskMonitor is failed: %s", err)
	}
	writes, err := res.FlattenDiskWriteValue()
	if err != nil {
		return fmt.Errorf("DiskMonitor is failed: %s", err)
	}

	// sort
	sort.Slice(reads, func(i, j int) bool { return reads[i].Time.Before(reads[j].Time) })
	sort.Slice(writes, func(i, j int) bool { return writes[i].Time.Before(writes[j].Time) })

	// build key string
	key := ""
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(params.KeyFormat))
	err = t.Execute(buf, map[string]interface{}{
		"ID": params.Id,
	})
	if err != nil {
		return fmt.Errorf("DiskMonitor is failed: %s", err)
	}
	key = buf.String()

	// build sortable struct
	list := MonitorValues{}
	for i := range reads {
		list = append(list, MonitorValue{
			"Key":       key,
			"TimeStamp": fmt.Sprintf("%s", reads[i].Time),
			"UnixTime":  fmt.Sprintf("%d", reads[i].Time.Unix()),
			"Read":      fmt.Sprintf("%f", reads[i].Value),
			"Write":     fmt.Sprintf("%f", writes[i].Value),
		})
	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.GetOutput().Print(values...)
}
