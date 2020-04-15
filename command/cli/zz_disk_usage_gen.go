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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-usage'; DO NOT EDIT

package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func diskCommandOrder(cmd *cobra.Command) []*commandSet {
	var commands []*commandSet
	{
		set := &commandSet{
			title: "Basics",
		}
		set.commands = append(set.commands, lookupCmd(cmd, "list"))
		set.commands = append(set.commands, lookupCmd(cmd, "create"))
		set.commands = append(set.commands, lookupCmd(cmd, "read"))
		set.commands = append(set.commands, lookupCmd(cmd, "update"))
		set.commands = append(set.commands, lookupCmd(cmd, "delete"))
		commands = append(commands, set)
	}
	{
		set := &commandSet{
			title: "Disk Edit",
		}
		set.commands = append(set.commands, lookupCmd(cmd, "edit"))
		set.commands = append(set.commands, lookupCmd(cmd, "resize-partition"))
		commands = append(commands, set)
	}
	{
		set := &commandSet{
			title: "Re-Install",
		}
		set.commands = append(set.commands, lookupCmd(cmd, "reinstall-from-archive"))
		set.commands = append(set.commands, lookupCmd(cmd, "reinstall-from-disk"))
		set.commands = append(set.commands, lookupCmd(cmd, "reinstall-to-blank"))
		commands = append(commands, set)
	}
	{
		set := &commandSet{
			title: "Server Connection Management",
		}
		set.commands = append(set.commands, lookupCmd(cmd, "server-connect"))
		set.commands = append(set.commands, lookupCmd(cmd, "server-disconnect"))
		commands = append(commands, set)
	}
	{
		set := &commandSet{
			title: "Monitoring",
		}
		set.commands = append(set.commands, lookupCmd(cmd, "monitor"))
		commands = append(commands, set)
	}
	{
		set := &commandSet{
			title: "Other",
		}
		set.commands = append(set.commands, lookupCmd(cmd, "wait-for-copy"))
		commands = append(commands, set)
	}

	return commands
}

func diskListFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("scope"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("source-archive-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("source-disk-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("storage"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("limit-offset", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("from"))
		fs.AddFlag(cmd.LocalFlags().Lookup("max"))
		sets = append(sets, &flagSet{
			title: "Limit/Offset options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("sort", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("sort"))
		sets = append(sets, &flagSet{
			title: "Sort options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("column"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		sets = append(sets, &flagSet{
			title: "Output options",
			flags: fs,
		})
	}

	return sets
}

func diskCreateFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("disk", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("plan"))
		fs.AddFlag(cmd.LocalFlags().Lookup("connection"))
		fs.AddFlag(cmd.LocalFlags().Lookup("source-archive-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("source-disk-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("size"))
		fs.AddFlag(cmd.LocalFlags().Lookup("distant-from"))
		sets = append(sets, &flagSet{
			title: "Disk options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("common", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("icon-id"))
		sets = append(sets, &flagSet{
			title: "Common options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("column"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		sets = append(sets, &flagSet{
			title: "Output options",
			flags: fs,
		})
	}

	return sets
}

func diskReadFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("column"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		sets = append(sets, &flagSet{
			title: "Output options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func diskUpdateFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("disk", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("connection"))
		sets = append(sets, &flagSet{
			title: "Disk options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("common", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("icon-id"))
		sets = append(sets, &flagSet{
			title: "Common options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("column"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		sets = append(sets, &flagSet{
			title: "Output options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func diskDeleteFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("column"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		sets = append(sets, &flagSet{
			title: "Output options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func diskEditFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("edit", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("hostname"))
		fs.AddFlag(cmd.LocalFlags().Lookup("password"))
		fs.AddFlag(cmd.LocalFlags().Lookup("ssh-key-ids"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disable-password-auth"))
		fs.AddFlag(cmd.LocalFlags().Lookup("ipaddress"))
		fs.AddFlag(cmd.LocalFlags().Lookup("default-route"))
		fs.AddFlag(cmd.LocalFlags().Lookup("nw-masklen"))
		fs.AddFlag(cmd.LocalFlags().Lookup("startup-script-ids"))
		sets = append(sets, &flagSet{
			title: "Edit options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("column"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		sets = append(sets, &flagSet{
			title: "Output options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func diskResizePartitionFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("column"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		sets = append(sets, &flagSet{
			title: "Output options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func diskReinstallFromArchiveFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("install", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("source-archive-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("distant-from"))
		sets = append(sets, &flagSet{
			title: "Install options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func diskReinstallFromDiskFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("install", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("source-disk-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("distant-from"))
		sets = append(sets, &flagSet{
			title: "Install options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func diskReinstallToBlankFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("install", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("distant-from"))
		sets = append(sets, &flagSet{
			title: "Install options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func diskServerConnectFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("connect", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("server-id"))
		sets = append(sets, &flagSet{
			title: "Connect options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func diskServerDisconnectFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func diskMonitorFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("column"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		sets = append(sets, &flagSet{
			title: "Output options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("end"))
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("key-format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("start"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}

func diskWaitForCopyFlagOrder(cmd *cobra.Command) []*flagSet {
	var sets []*flagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("filter", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("selector"))
		sets = append(sets, &flagSet{
			title: "Filter options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("param-template-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameter-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		sets = append(sets, &flagSet{
			title: "Input options",
			flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("default", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("id"))
		sets = append(sets, &flagSet{
			title: "Other options",
			flags: fs,
		})
	}

	return sets
}