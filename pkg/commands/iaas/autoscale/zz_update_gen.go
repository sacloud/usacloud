// Copyright 2017-2025 The sacloud/usacloud Authors
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

package autoscale

import (
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/packages-go/pointer"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *updateParameter) CleanupEmptyValue(fs *pflag.FlagSet) {
	if !fs.Changed("name") {
		p.Name = nil
	}
	if !fs.Changed("description") {
		p.Description = nil
	}
	if !fs.Changed("tags") {
		p.Tags = nil
	}
	if !fs.Changed("icon-id") {
		p.IconID = nil
	}
	if !fs.Changed("zones") {
		p.Zones = nil
	}
	if !fs.Changed("config") {
		p.Config = nil
	}
	if !fs.Changed("trigger-type") {
		p.TriggerType = nil
	}
	if !fs.Changed("cpu-threshold-scaling-server-prefix") {
		p.CPUThresholdScaling.ServerPrefix = nil
	}
	if !fs.Changed("cpu-threshold-scaling-up") {
		p.CPUThresholdScaling.Up = nil
	}
	if !fs.Changed("cpu-threshold-scaling-down") {
		p.CPUThresholdScaling.Down = nil
	}
	if !fs.Changed("router-threshold-scaling-router-prefix") {
		p.RouterThresholdScaling.RouterPrefix = nil
	}
	if !fs.Changed("router-threshold-scaling-direction") {
		p.RouterThresholdScaling.Direction = nil
	}
	if !fs.Changed("router-threshold-scaling-mbps") {
		p.RouterThresholdScaling.Mbps = nil
	}
	if !fs.Changed("schedule-scaling") {
		p.ScheduleScalingData = nil
	}
}

func (p *updateParameter) buildFlags(fs *pflag.FlagSet) {
	if p.Name == nil {
		p.Name = pointer.NewString("")
	}
	if p.Description == nil {
		p.Description = pointer.NewString("")
	}
	if p.Tags == nil {
		p.Tags = pointer.NewStringSlice([]string{})
	}
	if p.IconID == nil {
		v := types.ID(0)
		p.IconID = &v
	}
	if p.Zones == nil {
		p.Zones = pointer.NewStringSlice([]string{})
	}
	if p.Config == nil {
		p.Config = pointer.NewString("")
	}
	if p.TriggerType == nil {
		p.TriggerType = pointer.NewString("")
	}
	if p.CPUThresholdScaling.ServerPrefix == nil {
		p.CPUThresholdScaling.ServerPrefix = pointer.NewString("")
	}
	if p.CPUThresholdScaling.Up == nil {
		p.CPUThresholdScaling.Up = pointer.NewInt(0)
	}
	if p.CPUThresholdScaling.Down == nil {
		p.CPUThresholdScaling.Down = pointer.NewInt(0)
	}
	if p.RouterThresholdScaling.RouterPrefix == nil {
		p.RouterThresholdScaling.RouterPrefix = pointer.NewString("")
	}
	if p.RouterThresholdScaling.Direction == nil {
		p.RouterThresholdScaling.Direction = pointer.NewString("")
	}
	if p.RouterThresholdScaling.Mbps == nil {
		p.RouterThresholdScaling.Mbps = pointer.NewInt(0)
	}
	if p.ScheduleScalingData == nil {
		p.ScheduleScalingData = pointer.NewString("")
	}
	fs.StringVarP(&p.Parameters, "parameters", "", p.Parameters, "Input parameters in JSON format")
	fs.BoolVarP(&p.GenerateSkeleton, "generate-skeleton", "", p.GenerateSkeleton, "Output skeleton of parameters with JSON format (aliases: --skeleton)")
	fs.BoolVarP(&p.Example, "example", "", p.Example, "Output example parameters with JSON format")
	fs.BoolVarP(&p.AssumeYes, "assumeyes", "y", p.AssumeYes, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&p.OutputType, "output-type", "o", p.OutputType, "Output format options: [table/json/yaml] (aliases: --out)")
	fs.BoolVarP(&p.Quiet, "quiet", "q", p.Quiet, "Output IDs only")
	fs.StringVarP(&p.Format, "format", "", p.Format, "Output format in Go templates (aliases: --fmt)")
	fs.StringVarP(&p.Query, "query", "", p.Query, "Query for JSON output")
	fs.StringVarP(&p.QueryDriver, "query-driver", "", p.QueryDriver, "Name of the driver that handles queries to JSON output options: [jmespath/jq]")
	fs.StringVarP(p.Name, "name", "", "", "")
	fs.StringVarP(p.Description, "description", "", "", "")
	fs.StringSliceVarP(p.Tags, "tags", "", nil, "")
	fs.VarP(core.NewIDFlag(p.IconID, p.IconID), "icon-id", "", "")
	fs.StringSliceVarP(p.Zones, "zones", "", nil, "(*required) ")
	fs.StringVarP(p.Config, "config", "", "", "(*required) ")
	fs.BoolVarP(&p.Disabled, "disabled", "", p.Disabled, "")
	fs.StringVarP(p.TriggerType, "trigger-type", "", "", "options: [cpu/router/schedule]")
	fs.StringVarP(p.CPUThresholdScaling.ServerPrefix, "cpu-threshold-scaling-server-prefix", "", "", "")
	fs.IntVarP(p.CPUThresholdScaling.Up, "cpu-threshold-scaling-up", "", 0, "")
	fs.IntVarP(p.CPUThresholdScaling.Down, "cpu-threshold-scaling-down", "", 0, "")
	fs.StringVarP(p.RouterThresholdScaling.RouterPrefix, "router-threshold-scaling-router-prefix", "", "", "")
	fs.StringVarP(p.RouterThresholdScaling.Direction, "router-threshold-scaling-direction", "", "", "")
	fs.IntVarP(p.RouterThresholdScaling.Mbps, "router-threshold-scaling-mbps", "", 0, "")
	fs.StringVarP(p.ScheduleScalingData, "schedule-scaling", "", "", "")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *updateParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "skeleton":
		name = "generate-skeleton"
	case "out":
		name = "output-type"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func (p *updateParameter) buildFlagsUsage(cmd *cobra.Command) {
	var sets []*core.FlagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("common", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("icon-id"))
		sets = append(sets, &core.FlagSet{
			Title: "Common options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("auto-scale", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("config"))
		fs.AddFlag(cmd.LocalFlags().Lookup("cpu-threshold-scaling-down"))
		fs.AddFlag(cmd.LocalFlags().Lookup("cpu-threshold-scaling-server-prefix"))
		fs.AddFlag(cmd.LocalFlags().Lookup("cpu-threshold-scaling-up"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("router-threshold-scaling-direction"))
		fs.AddFlag(cmd.LocalFlags().Lookup("router-threshold-scaling-mbps"))
		fs.AddFlag(cmd.LocalFlags().Lookup("router-threshold-scaling-router-prefix"))
		fs.AddFlag(cmd.LocalFlags().Lookup("schedule-scaling"))
		fs.AddFlag(cmd.LocalFlags().Lookup("trigger-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("zones"))
		sets = append(sets, &core.FlagSet{
			Title: "Auto-Scale-specific options",
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

func (p *updateParameter) setCompletionFunc(cmd *cobra.Command) {
	cmd.RegisterFlagCompletionFunc("trigger-type", util.FlagCompletionFunc("cpu", "router", "schedule"))

}

func (p *updateParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
