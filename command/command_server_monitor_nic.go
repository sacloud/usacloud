package command

import (
	"bytes"
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"sort"
	"text/template"
	"time"
)

func ServerMonitorNic(ctx Context, params *MonitorNicServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()

	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerMonitorNic is failed: %s", e)
	}

	end := parseDateTimeString(params.End)
	start := end.Add(-1 * time.Hour)
	if params.Start != "" {
		start = parseDateTimeString(params.Start)
	}

	// validate start <= end
	if !(start.Unix() <= end.Unix()) {
		return fmt.Errorf("Invalid Parameter : start(%s) or end(%s) is invalid", start, end)
	}

	// validate nic indexes
	for _, v := range params.Index {
		if v > int64(len(p.Interfaces))-1 {
			return fmt.Errorf("Invalid NIC index: NIC(#%d) is not exists", v)
		}
	}
	if len(params.Index) == 0 {
		for i := range p.Interfaces {
			params.Index = append(params.Index, int64(i))
		}
	}

	req := sacloud.NewResourceMonitorRequest(&start, &end)

	list := MonitorValues{}
	for _, nicIndex := range params.Index {

		res, err := client.Interface.Monitor(p.Interfaces[nicIndex].ID, req)
		if err != nil {
			return fmt.Errorf("ServerMonitorNic is failed: %s", err)
		}

		// collect values
		receiveValues, err := res.FlattenPacketReceiveValue()
		if err != nil {
			return fmt.Errorf("ServerMonitorNic is failed: %s", err)
		}
		sendValues, err := res.FlattenPacketSendValue()
		if err != nil {
			return fmt.Errorf("ServerMonitorNic is failed: %s", err)
		}

		// build key string
		key := ""
		buf := bytes.NewBufferString("")
		t := template.New("t")
		template.Must(t.Parse(params.KeyFormat))
		err = t.Execute(buf, map[string]interface{}{
			"ID":    params.Id,
			"Index": nicIndex,
		})
		if err != nil {
			return fmt.Errorf("ServerMonitorNic is failed: %s", err)
		}
		key = buf.String()

		// build sortable struct
		for i := range receiveValues {
			list = append(list, MonitorValue{
				"Index":     fmt.Sprintf("%d", nicIndex),
				"Key":       key,
				"NicID":     fmt.Sprintf("%d", p.Interfaces[nicIndex].ID),
				"TimeStamp": fmt.Sprintf("%s", receiveValues[i].Time),
				"UnixTime":  fmt.Sprintf("%d", receiveValues[i].Time.Unix()),
				"Send":      fmt.Sprintf("%f", receiveValues[i].Value),
				"Receive":   fmt.Sprintf("%f", sendValues[i].Value),
			})
		}

	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.GetOutput().Print(values...)

}
