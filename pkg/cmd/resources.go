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

//go:generate go run github.com/sacloud/usacloud/tools/gen-commands/
package cmd

import (
	"github.com/sacloud/usacloud/pkg/cmd/commands/archive"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/cmd/root"
	_ "github.com/sacloud/usacloud/pkg/cmd/version"
)

var Resources = core.Resources{
	// libsacloud services
	archive.Resource,
}

func initCommands() {
	for _, r := range Resources {
		cmd := r.CLICommand()
		if len(cmd.Commands()) > 0 {
			root.Command.AddCommand(cmd)
		}
	}
	core.BuildRootCommandsUsage(root.Command, Resources.CategorizedResources())
}
