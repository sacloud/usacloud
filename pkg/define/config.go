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
	"github.com/sacloud/usacloud/pkg/schema"
)

func ConfigResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:             schema.CommandCustom,
			Aliases:          []string{"ls"},
			Params:           emptyParam(),
			UseCustomCommand: true,
			NeedlessConfirm:  true,
			NoOutput:         true,
		},
		"edit": {
			Type:             schema.CommandCustom,
			Params:           configEditParam(),
			UseCustomCommand: true,
			NeedlessConfirm:  true,
			NoOutput:         true,
		},
		"use": {
			Type:             schema.CommandCustom,
			Params:           emptyParam(),
			UseCustomCommand: true,
			NeedlessConfirm:  true,
			NoOutput:         true,
		},
		"show": {
			Type:             schema.CommandCustom,
			Params:           emptyParam(),
			UseCustomCommand: true,
			NeedlessConfirm:  true,
			NoOutput:         true,
		},
		"current": {
			Type:             schema.CommandCustom,
			Params:           emptyParam(),
			UseCustomCommand: true,
			NeedlessConfirm:  true,
			NoOutput:         true,
		},
		"delete": {
			Type:             schema.CommandCustom,
			Aliases:          []string{"rm"},
			Params:           emptyParam(),
			UseCustomCommand: true,
			NoOutput:         true,
		},
		"migrate": {
			Type:             schema.CommandCustom,
			Params:           emptyParam(),
			UseCustomCommand: true,
			NeedlessConfirm:  true,
			NoOutput:         true,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		Aliases:             []string{"profile"},
		DefaultCommand:      "edit",
		Usage:               "A manage command of APIKey settings",
		ResourceCategory:    CategoryConfig,
		SkipApplyConfigFile: true,
	}
}

var AllowZones = []string{"is1a", "is1b", "tk1a", "tk1v"}
var AllowOutputTypes = []string{"table", "json", "yaml", "csv", "tsv"}

func configEditParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"token": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "API Token of SakuraCloud",
			Category:    "config",
			Order:       10,
		},
		"secret": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "API Secret of SakuraCloud",
			Category:    "config",
			Order:       20,
		},
		"zone": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "Target zone of SakuraCloud",
			ValidateFunc: validateInStrValues(AllowZones...),
			Category:     "config",
			Order:        30,
		},
		"default-output-type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "Default output format type",
			ValidateFunc: validateInStrValues(AllowOutputTypes...),
			Category:     "config",
			Order:        40,
		},
	}
}
