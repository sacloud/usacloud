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

func InternetMonitor(ctx command.Context, params *params.MonitorInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()

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
		return fmt.Errorf("InternetMonitor is failed: %s", err)
	}

	// collect values
	inValues, err := res.FlattenInternetInValue()
	if err != nil {
		return fmt.Errorf("InternetMonitor is failed: %s", err)
	}
	outValues, err := res.FlattenInternetOutValue()
	if err != nil {
		return fmt.Errorf("InternetMonitor is failed: %s", err)
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
		return fmt.Errorf("DiskMonitor is failed: %s", err)
	}
	key = buf.String()

	// build sortable struct
	list := MonitorValues{}
	for i := range inValues {
		list = append(list, MonitorValue{
			"Key":       key,
			"TimeStamp": fmt.Sprintf("%s", inValues[i].Time),
			"UnixTime":  fmt.Sprintf("%d", inValues[i].Time.Unix()),
			"In":        fmt.Sprintf("%f", inValues[i].Value),
			"Out":       fmt.Sprintf("%f", outValues[i].Value),
		})
	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.GetOutput().Print(values...)
}
