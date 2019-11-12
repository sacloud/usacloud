// Copyright 2017-2019 The Usacloud Authors
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
	"github.com/sacloud/usacloud/schema"
)

func SelfResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"info": {
			Type:             schema.CommandCustom,
			Params:           selfInfoParam(),
			UseCustomCommand: true,
			NoOutput:         true,
			NeedlessConfirm:  true,
			SkipAuth:         true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		DefaultCommand:   "info",
		ResourceCategory: CategoryOther,
		Usage:            "Show self info",
	}
}

func selfInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}
