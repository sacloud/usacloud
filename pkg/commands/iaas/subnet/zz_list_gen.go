// Copyright 2017-2023 The sacloud/usacloud Authors
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

package subnet

import (
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *listParameter) CleanupEmptyValue(fs *pflag.FlagSet) {

}

func (p *listParameter) buildFlags(fs *pflag.FlagSet) {

	fs.StringVarP(&p.Zone, "zone", "", p.Zone, "(*required) ")
	fs.StringVarP(&p.Parameters, "parameters", "", p.Parameters, "Input parameters in JSON format")
	fs.BoolVarP(&p.GenerateSkeleton, "generate-skeleton", "", p.GenerateSkeleton, "Output skeleton of parameters with JSON format (aliases: --skeleton)")
	fs.BoolVarP(&p.Example, "example", "", p.Example, "Output example parameters with JSON format")
	fs.IntVarP(&p.Count, "count", "", p.Count, "(aliases: --max, --limit)")
	fs.IntVarP(&p.From, "from", "", p.From, "(aliases: --offset)")
	fs.StringVarP(&p.OutputType, "output-type", "o", p.OutputType, "Output format options: [table/json/yaml] (aliases: --out)")
	fs.BoolVarP(&p.Quiet, "quiet", "q", p.Quiet, "Output IDs only")
	fs.StringVarP(&p.Format, "format", "", p.Format, "Output format in Go templates (aliases: --fmt)")
	fs.StringVarP(&p.Query, "query", "", p.Query, "Query for JSON output")
	fs.StringVarP(&p.QueryDriver, "query-driver", "", p.QueryDriver, "Name of the driver that handles queries to JSON output options: [jmespath/jq]")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *listParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "skeleton":
		name = "generate-skeleton"
	case "max":
		name = "count"
	case "limit":
		name = "count"
	case "offset":
		name = "from"
	case "out":
		name = "output-type"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func (p *listParameter) buildFlagsUsage(cmd *cobra.Command) {
	var sets []*core.FlagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("limit-offset", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("count"))
		fs.AddFlag(cmd.LocalFlags().Lookup("from"))
		sets = append(sets, &core.FlagSet{
			Title: "Limit/Offset options",
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
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		sets = append(sets, &core.FlagSet{
			Title: "Input options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-driver"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		sets = append(sets, &core.FlagSet{
			Title: "Output options",
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

func (p *listParameter) setCompletionFunc(cmd *cobra.Command) {

}

func (p *listParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
