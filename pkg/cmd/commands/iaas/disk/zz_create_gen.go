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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands'; DO NOT EDIT

package disk

import (
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *createParameter) CleanupEmptyValue(fs *pflag.FlagSet) {

}

func (p *createParameter) buildFlags(fs *pflag.FlagSet) {

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
	fs.StringVarP(&p.Name, "name", "", p.Name, "(*required) ")
	fs.StringVarP(&p.Description, "description", "", p.Description, "")
	fs.StringSliceVarP(&p.Tags, "tags", "", p.Tags, "")
	fs.VarP(core.NewIDFlag(&p.IconID, &p.IconID), "icon-id", "", "")
	fs.StringVarP(&p.DiskPlan, "disk-plan", "", p.DiskPlan, "(*required) options: [ssd/hdd]")
	fs.IntVarP(&p.SizeGB, "size", "", p.SizeGB, "")
	fs.StringVarP(&p.Connection, "connector", "", p.Connection, "(*required) options: [virtio/ide] (aliases: --connection)")
	fs.StringVarP(&p.OSType, "os-type", "", p.OSType, "options: [almalinux/rockylinux/miraclelinux/centos8stream/ubuntu/debian/rancheros/k3os/...]")
	fs.VarP(core.NewIDFlag(&p.SourceDiskID, &p.SourceDiskID), "source-disk-id", "", "")
	fs.VarP(core.NewIDFlag(&p.SourceArchiveID, &p.SourceArchiveID), "source-archive-id", "", "")
	fs.VarP(core.NewIDFlag(&p.ServerID, &p.ServerID), "server-id", "", "")
	fs.VarP(core.NewIDSliceFlag(&p.DistantFrom, &p.DistantFrom), "distant-from", "", "")
	fs.StringVarP(&p.EditDisk.HostName, "edit-disk-host-name", "", p.EditDisk.HostName, "")
	fs.StringVarP(&p.EditDisk.Password, "edit-disk-password", "", p.EditDisk.Password, "")
	fs.StringVarP(&p.EditDisk.IPAddress, "edit-disk-ip-address", "", p.EditDisk.IPAddress, "")
	fs.IntVarP(&p.EditDisk.NetworkMaskLen, "edit-disk-netmask", "", p.EditDisk.NetworkMaskLen, "(aliases: --network-mask-len)")
	fs.StringVarP(&p.EditDisk.DefaultRoute, "edit-disk-gateway", "", p.EditDisk.DefaultRoute, "(aliases: --default-route)")
	fs.BoolVarP(&p.EditDisk.DisablePWAuth, "edit-disk-disable-pw-auth", "", p.EditDisk.DisablePWAuth, "")
	fs.BoolVarP(&p.EditDisk.EnableDHCP, "edit-disk-enable-dhcp", "", p.EditDisk.EnableDHCP, "")
	fs.BoolVarP(&p.EditDisk.ChangePartitionUUID, "edit-disk-change-partition-uuid", "", p.EditDisk.ChangePartitionUUID, "")
	fs.StringSliceVarP(&p.EditDisk.SSHKeys, "edit-disk-ssh-keys", "", p.EditDisk.SSHKeys, "")
	fs.VarP(core.NewIDSliceFlag(&p.EditDisk.SSHKeyIDs, &p.EditDisk.SSHKeyIDs), "edit-disk-ssh-key-ids", "", "")
	fs.BoolVarP(&p.EditDisk.IsSSHKeysEphemeral, "edit-disk-make-ssh-keys-ephemeral", "", p.EditDisk.IsSSHKeysEphemeral, "")
	fs.VarP(core.NewIDSliceFlag(&p.EditDisk.NoteIDs, &p.EditDisk.NoteIDs), "edit-disk-note-ids", "", "")
	fs.StringVarP(&p.EditDisk.NotesData, "edit-disk-notes", "", p.EditDisk.NotesData, "")
	fs.BoolVarP(&p.EditDisk.IsNotesEphemeral, "edit-disk-make-notes-ephemeral", "", p.EditDisk.IsNotesEphemeral, "")
	fs.BoolVarP(&p.NoWait, "no-wait", "", p.NoWait, "")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *createParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "skeleton":
		name = "generate-skeleton"
	case "out":
		name = "output-type"
	case "fmt":
		name = "format"
	case "connection":
		name = "connector"
	case "network-mask-len":
		name = "edit-disk-netmask"
	case "default-route":
		name = "edit-disk-gateway"
	}
	return pflag.NormalizedName(name)
}

func (p *createParameter) buildFlagsUsage(cmd *cobra.Command) {
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
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-plan"))
		fs.AddFlag(cmd.LocalFlags().Lookup("size"))
		fs.AddFlag(cmd.LocalFlags().Lookup("connector"))
		sets = append(sets, &core.FlagSet{
			Title: "Plan options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("disk", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("distant-from"))
		fs.AddFlag(cmd.LocalFlags().Lookup("server-id"))
		sets = append(sets, &core.FlagSet{
			Title: "Disk-specific options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("diskedit", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-host-name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-password"))
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-ip-address"))
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-netmask"))
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-gateway"))
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-disable-pw-auth"))
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-enable-dhcp"))
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-change-partition-uuid"))
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-ssh-keys"))
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-ssh-key-ids"))
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-make-ssh-keys-ephemeral"))
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-note-ids"))
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-notes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("edit-disk-make-notes-ephemeral"))
		sets = append(sets, &core.FlagSet{
			Title: "Edit disk options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("source", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("os-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("source-disk-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("source-archive-id"))
		sets = append(sets, &core.FlagSet{
			Title: "Source options",
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

func (p *createParameter) setCompletionFunc(cmd *cobra.Command) {
	cmd.RegisterFlagCompletionFunc("disk-plan", util.FlagCompletionFunc("ssd", "hdd"))
	cmd.RegisterFlagCompletionFunc("connector", util.FlagCompletionFunc("virtio", "ide"))
	cmd.RegisterFlagCompletionFunc("os-type", util.FlagCompletionFunc("centos", "centos8stream", "centos7", "almalinux", "rockylinux", "miracle", "miraclelinux", "ubuntu", "ubuntu2004", "ubuntu1804", "debian", "debian10", "debian11", "rancheros", "k3os", "kusanagi", "windows2016", "windows2016-rds", "windows2016-rds-office", "windows2016-sql-web", "windows2016-sql-standard", "windows2016-sql-standard-all", "windows2016-sql2017-standard", "windows2016-sql2017-enterprise", "windows2016-sql2017-standard-all", "windows2019", "windows2019-rds", "windows2019-rds-office2019", "windows2019-sql2017-web", "windows2019-sql2019-web", "windows2019-sql2017-standard", "windows2019-sql2019-standard", "windows2019-sql2017-enterprise", "windows2019-sql2019-enterprise", "windows2019-sql2017-standard-all", "windows2019-sql2019-standard-all"))

}

func (p *createParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}