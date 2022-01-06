// Copyright 2017-2022 The Usacloud Authors
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

package examples

import (
	"strings"

	"github.com/sacloud/usacloud/pkg/vdef"

	"github.com/sacloud/libsacloud/v2/sacloud/pointer"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
)

var (
	ID   = types.ID(123456789012)
	Name = cflag.NameParameter{
		Name: "example",
	}
	Description = cflag.DescParameter{
		Description: "example",
	}
	Tags = cflag.TagsParameter{
		Tags: types.Tags{"tag1=example1", "tag2=example2"},
	}
	IconID = cflag.IconIDParameter{
		IconID: ID,
	}
	NameUpdate = cflag.NameUpdateParameter{
		Name: pointer.NewString("example"),
	}
	DescriptionUpdate = cflag.DescUpdateParameter{
		Description: pointer.NewString("example"),
	}
	TagsUpdate = cflag.TagsUpdateParameter{
		Tags: &Tags.Tags,
	}
	IconIDUpdate = cflag.IconIDUpdateParameter{
		IconID: &ID,
	}
	IPAddress        = "192.0.2.11"
	IPAddresses      = []string{"192.0.2.21", "192.0.2.22"}
	VirtualIPAddress = "192.0.2.101"
	NetworkMaskLen   = 24
	DefaultRoute     = "192.0.2.1"

	SlackNotifyWebhooksURL = "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX"

	ScriptContent = `#!/bin/bash

...`
)

func Zones(zones []string) cflag.ZoneParameter {
	return cflag.ZoneParameter{
		Zone: ZonesString(zones),
	}
}

func ZonesString(zones []string) string {
	var values []string
	for _, z := range zones {
		if z != "all" {
			values = append(values, z)
		}
	}
	return strings.Join(values, " | ")
}

func OptionsString(vdefKey string) string {
	keys, ok := vdef.Keys(vdefKey)
	if !ok {
		return ""
	}
	return strings.Join(keys, " | ")
}
