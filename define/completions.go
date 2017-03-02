package define

import (
	"fmt"
	"github.com/sacloud/usacloud/schema"
)

func completeInStrValues(values ...string) schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {
		return values
	}
}

func completeInIntValues(values ...int) schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {
		res := []string{}
		for _, v := range values {
			res = append(res, fmt.Sprintf("%d", v))
		}
		return res
	}
}

func completeIconID() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetIconAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.Icons {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completeServerID() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetServerAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.Servers {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completeBridgeID() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetBridgeAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.Bridges {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completeDiskID() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetDiskAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.Disks {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completeArchiveID() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetArchiveAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.Archives {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completeISOImageID() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetCDROMAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.CDROMs {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completeSwitchID() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetSwitchAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.Switches {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completePacketFilterID() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetPacketFilterAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.PacketFilters {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completeSSHKeyID() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetSSHKeyAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.SSHKeys {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completeNoteID() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetNoteAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.Notes {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completeLicenseInfoID() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetProductLicenseAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.LicenseInfo {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completeInterfaceID() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetInterfaceAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		ignoreTags := []string{"@appliance-database", "@appliance-loadbalancer", "@appliance-vpcrouter"}
	Outer:
		for _, nic := range result.Interfaces {
			// customize: ignore appliance interface
			for _, t := range ignoreTags {
				if nic.Server.HasTag(t) {
					continue Outer
				}
			}

			res = append(res, fmt.Sprintf("%d", nic.ID))
		}
		return res
	}
}

func completeServerCore() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetProductServerAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		var existsMap = map[int]bool{}

		for _, v := range result.ServerPlans {
			if _, ok := existsMap[v.CPU]; !ok {
				res = append(res, fmt.Sprintf("%d", v.CPU))
				existsMap[v.CPU] = true
			}
		}
		return res
	}
}

func completeServerMemory() schema.SchemaCompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetProductServerAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		var existsMap = map[int]bool{}

		for _, v := range result.ServerPlans {
			gb := v.GetMemoryGB()
			if _, ok := existsMap[gb]; !ok {
				res = append(res, fmt.Sprintf("%d", gb))
				existsMap[gb] = true
			}
		}
		return res
	}
}
