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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands'; DO NOT EDIT

package disk

import (
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *connectToServerParameter) CleanupEmptyValue(fs *pflag.FlagSet) {

}

func (p *connectToServerParameter) buildFlags(fs *pflag.FlagSet) {

	fs.StringVarP(&p.Zone, "zone", "", p.Zone, "(*required) ")
	fs.StringVarP(&p.Parameters, "parameters", "", p.Parameters, "Input parameters in JSON format")
	fs.BoolVarP(&p.GenerateSkeleton, "generate-skeleton", "", p.GenerateSkeleton, "Output skeleton of parameters with JSON format (aliases: --skeleton)")
	fs.BoolVarP(&p.Example, "example", "", p.Example, "Output example parameters with JSON format")
	fs.BoolVarP(&p.AssumeYes, "assumeyes", "y", p.AssumeYes, "Assume that the answer to any question which would be asked is yes")
	fs.VarP(core.NewIDFlag(&p.ServerID, &p.ServerID), "server-id", "", "(*required) ")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *connectToServerParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "skeleton":
		name = "generate-skeleton"
	}
	return pflag.NormalizedName(name)
}

func (p *connectToServerParameter) buildFlagsUsage(cmd *cobra.Command) {
	var sets []*core.FlagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("disk", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("server-id"))
		sets = append(sets, &core.FlagSet{
			Title: "Disk-specific options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("zone", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("zone"))
		sets = append(sets, &core.FlagSet{
			Title: "Zone options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("input", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		sets = append(sets, &core.FlagSet{
			Title: "Input options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("example", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("example"))
		sets = append(sets, &core.FlagSet{
			Title: "Parameter example",
			Flags: fs,
		})
	}

	core.BuildFlagsUsage(cmd, sets)
}

func (p *connectToServerParameter) setCompletionFunc(cmd *cobra.Command) {

}

func (p *connectToServerParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
