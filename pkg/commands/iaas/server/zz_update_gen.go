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

package server

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
	if !fs.Changed("cpu") {
		p.CPU = nil
	}
	if !fs.Changed("memory") {
		p.Memory = nil
	}
	if !fs.Changed("gpu") {
		p.GPU = nil
	}
	if !fs.Changed("cpu-model") {
		p.CPUModel = nil
	}
	if !fs.Changed("commitment") {
		p.Commitment = nil
	}
	if !fs.Changed("generation") {
		p.Generation = nil
	}
	if !fs.Changed("interface-driver") {
		p.InterfaceDriver = nil
	}
	if !fs.Changed("cdrom-id") {
		p.CDROMID = nil
	}
	if !fs.Changed("private-host-id") {
		p.PrivateHostID = nil
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
	if p.CPU == nil {
		p.CPU = pointer.NewInt(0)
	}
	if p.Memory == nil {
		p.Memory = pointer.NewInt(0)
	}
	if p.GPU == nil {
		p.GPU = pointer.NewInt(0)
	}
	if p.CPUModel == nil {
		p.CPUModel = pointer.NewString("")
	}
	if p.Commitment == nil {
		p.Commitment = pointer.NewString("")
	}
	if p.Generation == nil {
		p.Generation = pointer.NewString("")
	}
	if p.InterfaceDriver == nil {
		p.InterfaceDriver = pointer.NewString("")
	}
	if p.CDROMID == nil {
		v := types.ID(0)
		p.CDROMID = &v
	}
	if p.PrivateHostID == nil {
		v := types.ID(0)
		p.PrivateHostID = &v
	}
	fs.StringVarP(&p.Zone, "zone", "", p.Zone, "(*required) ")
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
	fs.IntVarP(p.CPU, "cpu", "", 0, "(aliases: --core)")
	fs.IntVarP(p.Memory, "memory", "", 0, "")
	fs.IntVarP(p.GPU, "gpu", "", 0, "")
	fs.StringVarP(p.CPUModel, "cpu-model", "", "", "")
	fs.StringVarP(p.Commitment, "commitment", "", "", "options: [standard/dedicatedcpu]")
	fs.StringVarP(p.Generation, "generation", "", "", "options: [default/g100/g200]")
	fs.StringVarP(p.InterfaceDriver, "interface-driver", "", "", "options: [virtio/e1000]")
	fs.VarP(core.NewIDFlag(p.CDROMID, p.CDROMID), "cdrom-id", "", "(aliases: --iso-image-id)")
	fs.VarP(core.NewIDFlag(p.PrivateHostID, p.PrivateHostID), "private-host-id", "", "")
	fs.StringVarP(&p.NetworkInterfaceData, "network-interfaces", "", p.NetworkInterfaceData, "")
	fs.StringVarP(&p.DisksData, "disks", "", p.DisksData, "")
	fs.BoolVarP(&p.NoWait, "no-wait", "", p.NoWait, "")
	fs.BoolVarP(&p.ForceShutdown, "force-shutdown", "", p.ForceShutdown, "")
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
	case "core":
		name = "cpu"
	case "iso-image-id":
		name = "cdrom-id"
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
		fs = pflag.NewFlagSet("plan", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("cpu"))
		fs.AddFlag(cmd.LocalFlags().Lookup("memory"))
		fs.AddFlag(cmd.LocalFlags().Lookup("gpu"))
		fs.AddFlag(cmd.LocalFlags().Lookup("cpu-model"))
		fs.AddFlag(cmd.LocalFlags().Lookup("commitment"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generation"))
		sets = append(sets, &core.FlagSet{
			Title: "Plan options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("server", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("cdrom-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disks"))
		fs.AddFlag(cmd.LocalFlags().Lookup("force-shutdown"))
		fs.AddFlag(cmd.LocalFlags().Lookup("interface-driver"))
		fs.AddFlag(cmd.LocalFlags().Lookup("network-interfaces"))
		fs.AddFlag(cmd.LocalFlags().Lookup("private-host-id"))
		sets = append(sets, &core.FlagSet{
			Title: "Server-specific options",
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
		fs = pflag.NewFlagSet("wait", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("no-wait"))
		sets = append(sets, &core.FlagSet{
			Title: "Wait options",
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
	cmd.RegisterFlagCompletionFunc("commitment", util.FlagCompletionFunc("standard", "dedicatedcpu"))
	cmd.RegisterFlagCompletionFunc("generation", util.FlagCompletionFunc("default", "g100", "g200"))
	cmd.RegisterFlagCompletionFunc("interface-driver", util.FlagCompletionFunc("virtio", "e1000"))

}

func (p *updateParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
