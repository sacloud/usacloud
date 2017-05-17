package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"strings"
	"time"
)

func SummaryShow(ctx Context, params *ShowSummaryParam) error {

	client := ctx.GetAPIClient()

	// auth-status
	auth, err := client.GetAuthStatusAPI().Read()
	if err != nil {
		return fmt.Errorf("SummaryShow is failed: %s", err)
	}
	accountID := sacloud.NewResourceByStringID(auth.Account.ID).ID
	if accountID == 0 {
		return fmt.Errorf("SummaryShow is failed: %s", "invalid account id")
	}

	if !strings.Contains(auth.ExternalPermission, "bill") {
		return fmt.Errorf("Don't have permission to view bills")
	}

	// call Find()
	res, err := client.GetBillAPI().ByContract(accountID)
	if err != nil || len(res.Bills) == 0 {
		return fmt.Errorf("SummaryShow is failed: %s", err)
	}

	// use latest
	bill := res.Bills[0]

	// get service classes(public-price API)
	priceRes, err := client.GetPublicPriceAPI().Find()
	if err != nil {
		return fmt.Errorf("SummaryShow is failed: %s", err)
	}
	services := priceRes.ServiceClasses
	funcFindServiceClass := func(id int64) *sacloud.PublicPrice {
		for i := range services {
			if services[i].ServiceClassID == int(id) {
				return &services[i]
			}
		}
		return nil
	}

	// get&build BillDetails
	detailRes, err := client.GetBillAPI().GetDetail(auth.Member.Code, bill.BillID)
	if err != nil {
		return fmt.Errorf("SummaryShow is failed: %s", err)
	}
	var details []*billDetail
	now := time.Now()
	for _, d := range detailRes.BillDetails {
		if !d.IsContractEnded(now) {
			service := funcFindServiceClass(d.ServiceClassID)
			if service == nil {
				continue
			}
			details = append(details, &billDetail{
				BillDetail:       d,
				ServiceClassPath: service.ServiceClassPath,
			})
		}
	}

	// build
	for _, d := range details {
		summaryValues.addBillDetail(d)
	}

	return ctx.GetOutput().Print(summaryValues.buildSummaryOutputs()...)
}

type billDetail struct {
	*sacloud.BillDetail
	ServiceClassPath string
}

type summaryOutput struct {
	ResourceName string
	tk1a         string
	is1a         string
	is1b         string
	Total        string
}

type summaryTarget struct {
	name      string
	classPath string
	isGlobal  bool
	tk1a      int
	is1a      int
	is1b      int
	total     int
}

func (s *summaryTarget) buildSummaryOutput() *summaryOutput {

	tk1a := "-"
	is1a := "-"
	is1b := "-"

	if !s.isGlobal {
		tk1a = fmt.Sprintf("%d", s.tk1a)
		is1a = fmt.Sprintf("%d", s.is1a)
		is1b = fmt.Sprintf("%d", s.is1b)
	}

	return &summaryOutput{
		ResourceName: s.name,
		tk1a:         tk1a,
		is1a:         is1a,
		is1b:         is1b,
		Total:        fmt.Sprintf("%d", s.total),
	}
}

func (s *summaryTarget) addBillDetail(d *billDetail) {

	if s.isGlobal {
		s.total++
		return
	}

	switch d.Zone {
	case "tk1a":
		s.tk1a++
	case "is1a":
		s.is1a++
	case "is1b":
		s.is1b++
	}
}

type summaryTargets []summaryTarget

func (s summaryTargets) addBillDetail(d *billDetail) {

	for _, t := range []summaryTarget(s) {
		if strings.HasSuffix(d.ServiceClassPath, t.classPath) {
			t.addBillDetail(d)
		}
	}

}

func (s summaryTargets) buildSummaryOutputs() []interface{} {
	list := []interface{}{}
	for _, t := range []summaryTarget(s) {
		list = append(list, t.buildSummaryOutput())
	}
	return list
}

var summaryValues = summaryTargets([]summaryTarget{
	{
		name:      "Server",
		classPath: "cloud/plan/",
		isGlobal:  false,
	},
	{
		name:      "Disk",
		classPath: "cloud/disk/",
		isGlobal:  false,
	},
	{
		name:      "Archive",
		classPath: "cloud/archive/",
		isGlobal:  false,
	},
	{
		name:      "ISO-Image",
		classPath: "cloud/iso/",
		isGlobal:  false,
	},
	{
		name:      "Switch",
		classPath: "cloud/switch",
		isGlobal:  false,
	},
	{
		name:      "Internet",
		classPath: "cloud/internet/router/",
		isGlobal:  false,
	},
	{
		name:      "Bridge",
		classPath: "cloud/bridge",
		isGlobal:  true,
	},
	{
		name:      "LoadBalancer",
		classPath: "cloud/appliance/loadbalancer",
		isGlobal:  false,
	},
	{
		name:      "VPCRouter",
		classPath: "cloud/appliance/vpc/",
		isGlobal:  false,
	},
	{
		name:      "Database",
		classPath: "cloud/appliance/database/",
		isGlobal:  false,
	},
	{
		name:      "GSLB",
		classPath: "cloud/gslb",
		isGlobal:  true,
	},
	{
		name:      "DNS",
		classPath: "cloud/dns",
		isGlobal:  true,
	},
	{
		name:      "SimpleMonitor",
		classPath: "cloud/simplemon/",
		isGlobal:  true,
	},
	{
		name:      "License",
		classPath: "cloud/os/",
		isGlobal:  true,
	},
	{
		name:      "ObjectStorage",
		classPath: "cloud/ostorage/bucket",
		isGlobal:  true,
	},
	{
		name:      "WebAccelerator",
		classPath: "cloud/cdn/webaccel",
		isGlobal:  true,
	},
})
