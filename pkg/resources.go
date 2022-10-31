// Copyright 2017-2022 The sacloud/usacloud Authors
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

//go:generate go run github.com/sacloud/usacloud/tools/gen-commands/
package pkg

import (
	"github.com/sacloud/usacloud/pkg/commands/completion"
	"github.com/sacloud/usacloud/pkg/commands/config"
	"github.com/sacloud/usacloud/pkg/commands/iaas"
	"github.com/sacloud/usacloud/pkg/commands/rest"
	"github.com/sacloud/usacloud/pkg/commands/root"
	updateSelf "github.com/sacloud/usacloud/pkg/commands/update-self"
	"github.com/sacloud/usacloud/pkg/commands/version"
	"github.com/sacloud/usacloud/pkg/commands/webaccel/webaccelerator"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/spf13/cobra"
)

var (
	MiscResources = core.Resources{
		config.Resource,
		rest.Resource,
		webaccelerator.Resource,
	}

	RootCommands = []*cobra.Command{
		completion.Command,
		version.Command,
		updateSelf.Command,
	}
)

func Resources() core.Resources {
	rs := core.Resources{}
	rs = append(rs, iaas.Resources...)
	rs = append(rs, MiscResources...)
	return rs
}

func initCommands() {
	initIaasCommands()
	initMiscCommands()
	initRootCommands()
}

func initIaasCommands() {
	iaas.InitCommand(root.Command)
}

func initMiscCommands() {
	for _, r := range MiscResources {
		cmd := r.CLICommand()
		if len(cmd.Commands()) > 0 {
			root.Command.AddCommand(cmd)
		}
	}
}

func initRootCommands() {
	root.Command.AddCommand(RootCommands...)
}
