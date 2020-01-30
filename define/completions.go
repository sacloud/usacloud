// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package define

import (
	"fmt"

	"github.com/sacloud/usacloud/schema"
)

func completeInStrValues(values ...string) schema.CompletionFunc {
	return schema.CompleteInStrValues(values...)
}

func completeInIntValues(values ...int) schema.CompletionFunc {
	return schema.CompleteInIntValues(values...)
}

func completeIconID() schema.CompletionFunc {
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

func completeServerID() schema.CompletionFunc {
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

func completeBridgeID() schema.CompletionFunc {
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

func completeDiskID() schema.CompletionFunc {
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

func completeDatabaseID() schema.CompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetDatabaseAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.Databases {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completeArchiveID() schema.CompletionFunc {
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

func completeISOImageID() schema.CompletionFunc {
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

func completeSwitchID() schema.CompletionFunc {
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

func completePacketFilterID() schema.CompletionFunc {
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

func completeSSHKeyID() schema.CompletionFunc {
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

func completeNoteID() schema.CompletionFunc {
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

func completeLicenseInfoID() schema.CompletionFunc {
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

func completeInterfaceID() schema.CompletionFunc {
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

func completeSubnetID() schema.CompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetSubnetAPI()
		res := []string{}

		result, err := api.Reset().Find()
		if err != nil {
			return res
		}

		for _, v := range result.Subnets {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completeSIMID() schema.CompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetSIMAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.CommonServiceSIMItems {
			res = append(res, fmt.Sprintf("%d", v.ID))
		}
		return res
	}
}

func completeServerCore() schema.CompletionFunc {
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

func completeServerMemory() schema.CompletionFunc {
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

func completePrivateHostID() schema.CompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetPrivateHostAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.PrivateHosts {
			res = append(res, v.GetStrID())
		}
		return res
	}
}

func completeBackupTime() schema.CompletionFunc {
	timeStrings := []string{}

	minutes := []int{0, 15, 30, 45}

	// create list [00:00 ,,, 23:45]
	for hour := 0; hour <= 23; hour++ {
		for _, minute := range minutes {
			timeStrings = append(timeStrings, fmt.Sprintf("%02d:%02d", hour, minute))
		}
	}

	return schema.CompleteInStrValues(timeStrings...)
}

func completeStorageName() schema.CompletionFunc {
	return func(ctx schema.CompletionContext, currentValue string) []string {

		api := ctx.GetAPIClient().GetDiskAPI()
		res := []string{}

		result, err := api.Find()
		if err != nil {
			return res
		}

		for _, v := range result.Disks {
			res = append(res, v.Storage.Name)
		}
		return res
	}
}
