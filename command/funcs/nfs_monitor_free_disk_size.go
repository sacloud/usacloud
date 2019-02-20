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

func NFSMonitorFreeDiskSize(ctx command.Context, params *params.MonitorFreeDiskSizeNFSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetNFSAPI()

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

	res, err := api.MonitorFreeDiskSize(params.Id, req)
	if err != nil {
		return fmt.Errorf("NFSMonitorFreeDiskSize is failed: %s", err)
	}

	// collect values
	freeValues, err := res.FlattenFreeDiskSizeValue()
	if err != nil {
		return fmt.Errorf("NFSMonitorFreeDiskSize is failed: %s", err)
	}

	// sort
	sort.Slice(freeValues, func(i, j int) bool { return freeValues[i].Time.Before(freeValues[j].Time) })

	// build key string
	key := ""
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(params.KeyFormat))
	err = t.Execute(buf, map[string]interface{}{
		"ID": params.Id,
	})
	if err != nil {
		return fmt.Errorf("NFSMonitorFreeDiskSize is failed: %s", err)
	}
	key = buf.String()

	// build sortable struct
	list := MonitorValues{}
	for i := range freeValues {
		list = append(list, MonitorValue{
			"Key":       key,
			"TimeStamp": fmt.Sprintf("%s", freeValues[i].Time),
			"UnixTime":  fmt.Sprintf("%d", freeValues[i].Time.Unix()),
			"Free":      fmt.Sprintf("%f", freeValues[i].Value),
		})
	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.GetOutput().Print(values...)

}
