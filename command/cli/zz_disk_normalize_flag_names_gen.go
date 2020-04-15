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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-normalize-flag-name'; DO NOT EDIT

package cli

import (
	"github.com/spf13/pflag"
)

func diskListNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "selector":
		name = "tags"
	case "offset":
		name = "from"
	case "limit":
		name = "max"
	case "out":
		name = "output-type"
	case "col":
		name = "column"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func diskCreateNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "desc":
		name = "description"
	case "out":
		name = "output-type"
	case "col":
		name = "column"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func diskReadNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "out":
		name = "output-type"
	case "col":
		name = "column"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func diskUpdateNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "desc":
		name = "description"
	case "out":
		name = "output-type"
	case "col":
		name = "column"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func diskDeleteNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "out":
		name = "output-type"
	case "col":
		name = "column"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func diskEditNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "disable-pw-auth":
		name = "disable-password-auth"
	case "ip":
		name = "ipaddress"
	case "gateway":
		name = "default-route"
	case "network-masklen":
		name = "nw-masklen"
	case "note-ids":
		name = "startup-script-ids"
	case "out":
		name = "output-type"
	case "col":
		name = "column"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func diskResizePartitionNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "out":
		name = "output-type"
	case "col":
		name = "column"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func diskReinstallFromArchiveNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	return pflag.NormalizedName(name)
}

func diskReinstallFromDiskNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	return pflag.NormalizedName(name)
}

func diskReinstallToBlankNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	return pflag.NormalizedName(name)
}

func diskServerConnectNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	return pflag.NormalizedName(name)
}

func diskServerDisconnectNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	return pflag.NormalizedName(name)
}

func diskMonitorNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "out":
		name = "output-type"
	case "col":
		name = "column"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func diskWaitForCopyNormalizeFlagNames(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	return pflag.NormalizedName(name)
}