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

package funcs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/sacloud/libsacloud/builder"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/libsacloud/sacloud/ostype"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/internal"
	"github.com/sacloud/usacloud/pkg/params"
)

func ServerBuild(ctx cli.Context, params *params.BuildServerParam) error {

	// validate --- for disk mode params
	errs := validateServerDiskModeParams(ctx, params)
	if len(errs) > 0 {
		return fmt.Errorf("%s", cli.FlattenErrors(errs))
	}

	// select builder
	sb := createServerBuilder(ctx, params)

	// handle build processes
	for _, handler := range serverBuildHandlers {
		err := handler(sb, ctx, params)
		if err != nil {
			return err
		}
	}

	// call Create(id)
	var b = sb.(serverBuilder)
	res, err := b.Build()
	if err != nil {
		return fmt.Errorf("ServerCreate is failed: %s", err)
	}

	// store private key if ssh-key was generated
	if len(res.Disks) > 0 && res.Disks[0].GeneratedSSHKey != nil {
		path := params.SSHKeyPrivateKeyOutput
		if path == "" {
			p, err := getSSHPrivateKeyStorePath(res.Server.ID)
			if err != nil {
				return fmt.Errorf("ServerCreate is failed: getting HomeDir is failed:%s", err)
			}
			path = p
		}
		pKey := res.Disks[0].GeneratedSSHKey.PrivateKey
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0600); err != nil {
			return fmt.Errorf("ServerCreate is failed: creating directory(%s) is failed:%s", dir, err)
		}

		err = ioutil.WriteFile(path, []byte(pKey), os.FileMode(0600))
		if err != nil {
			return fmt.Errorf("ServerCreate is failed: Writing private key to %s is failed:%s", params.SSHKeyPrivateKeyOutput, err)
		}
	}

	return ctx.GetOutput().Print(res.Server)
}

func createServerBuilder(ctx cli.Context, params *params.BuildServerParam) serverBuilder {
	client := builder.NewAPIClient(ctx.GetAPIClient())
	var sb serverBuilder

	switch params.DiskMode {
	case "create":
		if params.SourceDiskId > 0 {
			sb = builder.ServerFromDisk(client, params.Name, params.SourceDiskId)
		} else if params.SourceArchiveId > 0 {
			sb = builder.ServerFromArchive(client, params.Name, params.SourceArchiveId)
		} else {

			if params.OsType == "" {
				sb = builder.ServerBlankDisk(client, params.Name)
			} else {
				// Windows?
				if isWindows(params.OsType) {
					sb = builder.ServerPublicArchiveWindows(client, strToOSType(params.OsType), params.Name)
					// [HACK] ディスク修正可否から判定するのが望ましいが、Sophos/Netwiser/OPNsense以外に対象がないため現状は決め打ち判定
				} else if params.OsType == "sophos-utm" || params.OsType == "netwiser" || params.OsType == "opnsense" {
					sb = builder.ServerPublicArchiveFixedUnix(client, strToOSType(params.OsType), params.Name)
				} else {
					sb = builder.ServerPublicArchiveUnix(client, strToOSType(params.OsType), params.Name, params.Password)
				}
			}
		}
	case "connect":
		sb = builder.ServerFromExistsDisk(client, params.Name, params.DiskId)
	case "diskless":
		sb = builder.ServerDiskless(client, params.Name)
	}
	return sb
}

var serverBuildHandlers = []func(serverBuilder, cli.Context, *params.BuildServerParam) error{
	handleNetworkParams,
	handleDiskEditParams,
	handleDiskParams,
	handleServerCommonParams,
	handleDiskEvents,
	handleServerEvents,
}

func handleNetworkParams(sb serverBuilder, ctx cli.Context, params *params.BuildServerParam) error {
	// validate --- for network params
	errs := validateServerNetworkParams(sb, ctx, params)
	if len(errs) > 0 {
		return fmt.Errorf("%s", cli.FlattenErrors(errs))
	}

	// set network params
	//if sb, ok := sb.(builder.NetworkInterfaceProperty); ok {
	if sb.HasNetworkInterfaceProperty() {
		sb := sb.(builder.NetworkInterfaceProperty)
		switch params.NetworkMode {
		case "shared":
			sb.AddPublicNWConnectedNIC()
		case "switch":
			sb.AddExistsSwitchConnectedNIC(fmt.Sprintf("%d", params.SwitchId))
		case "disconnect":
			sb.AddDisconnectedNIC()
		case "none":
		// noop
		default:
			panic(fmt.Errorf("Unknown NetworkMode : %s", params.NetworkMode))
		}

		if params.PacketFilterId != sacloud.EmptyID {
			sb.SetPacketFilterIDs([]sacloud.ID{params.PacketFilterId})
		}
	}

	return nil
}

func handleDiskEditParams(sb serverBuilder, ctx cli.Context, params *params.BuildServerParam) error {
	// validate --- for disk params
	errs := validateServerDiskEditParams(sb, ctx, params)
	if len(errs) > 0 {
		return fmt.Errorf("%s", cli.FlattenErrors(errs))
	}

	// set disk edit params
	if sb.HasDiskEditProperty() {
		sb := sb.(builder.DiskEditProperty)
		sb.SetHostName(params.Hostname)
		sb.SetPassword(params.Password)
		sb.SetDisablePWAuth(params.DisablePasswordAuth)

		if params.Ipaddress != "" {
			sb.SetIPAddress(params.Ipaddress)
		}
		if params.DefaultRoute != "" {
			sb.SetDefaultRoute(params.DefaultRoute)
		}
		if params.NwMasklen > 0 {
			sb.SetNetworkMaskLen(params.NwMasklen)
		}

		for _, v := range params.StartupScriptIds {
			sb.AddNoteID(v)
		}
		for _, v := range params.StartupScripts {
			sb.AddNote(v)
		}
		sb.SetNotesEphemeral(params.StartupScriptsEphemeral)

		// SSH Key generate params
		switch params.SSHKeyMode {
		case "id":
			for _, v := range params.SSHKeyIds {
				sb.AddSSHKeyID(v)
			}
		case "generate":
			keyName := params.SSHKeyName
			if keyName == "" {
				keyName = fmt.Sprintf("generated-%d", time.Now().UnixNano())
			}
			sb.SetGenerateSSHKeyName(keyName)
			sb.SetGenerateSSHKeyPassPhrase(params.SSHKeyPassPhrase)
			sb.SetGenerateSSHKeyDescription(params.SSHKeyDescription)
			sb.SetSSHKeysEphemeral(params.SSHKeyEphemeral)
		case "upload":
			// pubkey(text)
			for _, v := range params.SSHKeyPublicKeys {
				sb.AddSSHKey(v)
			}
			// pubkey(from file)
			for _, v := range params.SSHKeyPublicKeyFiles {
				b, err := ioutil.ReadFile(v)
				if err != nil {
					return fmt.Errorf("ServerCreate is failed: %s", err)
				}
				sb.AddSSHKey(string(b))

			}
			sb.SetSSHKeysEphemeral(params.SSHKeyEphemeral)
		}

	}
	return nil
}

func handleDiskParams(sb serverBuilder, ctx cli.Context, params *params.BuildServerParam) error {
	// set disk params
	if sb.HasDiskProperty() {
		sb := sb.(builder.DiskProperty)
		sb.SetDiskPlan(params.DiskPlan)
		sb.SetDiskConnection(sacloud.EDiskConnection(params.DiskConnection))
		sb.SetDiskSize(params.DiskSize)
		sb.SetDistantFrom(params.DistantFrom)
	}

	return nil
}

func handleServerCommonParams(sb serverBuilder, ctx cli.Context, params *params.BuildServerParam) error {
	// set common params
	tags := params.GetTags()

	sb.SetCore(params.GetCore())
	sb.SetMemory(params.GetMemory())
	sb.SetCommitment(sacloud.ECommitment(params.Commitment))
	sb.SetPrivateHostID(params.PrivateHostId)
	sb.SetServerName(params.GetName())
	sb.SetDescription(params.GetDescription())
	if params.UsKeyboard {
		tags = append(tags, sacloud.TagKeyboardUS)
	}
	sb.SetTags(tags)
	sb.SetIconID(params.IconId)
	sb.SetBootAfterCreate(!params.DisableBootAfterCreate)
	sb.SetISOImageID(params.GetISOImageId())
	sb.SetInterfaceDriver(sacloud.EInterfaceDriver(params.InterfaceDriver))
	return nil
}

func handleDiskEvents(sb serverBuilder, ctx cli.Context, params *params.BuildServerParam) error {
	// set events
	if sb.HasDiskEventProperty() {
		sb := sb.(builder.DiskEventProperty)
		// create disk
		progCreate := internal.NewProgress(
			"Still creating disk...",
			"Create disk",
			ctx.IO().Progress(),
			ctx.Option().NoColor)
		sb.SetDiskEventHandler(builder.DiskBuildOnCreateDiskBefore, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progCreate.Start()
		})
		sb.SetDiskEventHandler(builder.DiskBuildOnCreateDiskAfter, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progCreate.Stop()
		})

		// cleanup startup script
		progCleanupNotes := internal.NewProgress(
			"Still cleaning StartupScript...",
			"Cleanup StartupScript",
			ctx.IO().Progress(),
			ctx.Option().NoColor)
		sb.SetDiskEventHandler(builder.DiskBuildOnCleanupNoteBefore, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progCleanupNotes.Start()
		})
		sb.SetDiskEventHandler(builder.DiskBuildOnCleanupNoteAfter, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progCleanupNotes.Stop()
		})

		// cleanup ssh key script
		progCleanupSSHKey := internal.NewProgress(
			"Still cleaning SSHKey...",
			"Cleanup SSHKey",
			ctx.IO().Progress(),
			ctx.Option().NoColor)
		sb.SetDiskEventHandler(builder.DiskBuildOnCleanupSSHKeyBefore, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progCleanupSSHKey.Start()
		})
		sb.SetDiskEventHandler(builder.DiskBuildOnCleanupSSHKeyAfter, func(value *builder.DiskBuildValue, result *builder.DiskBuildResult) {
			progCleanupSSHKey.Stop()
		})
	}

	return nil
}

func handleServerEvents(sb serverBuilder, ctx cli.Context, params *params.BuildServerParam) error {
	if sb.HasServerEventProperty() {
		sb := sb.(builder.ServerEventProperty)
		progCreate := internal.NewProgress(
			"Still creating server...",
			"Create server",
			ctx.IO().Progress(),
			ctx.Option().NoColor)

		sb.SetEventHandler(builder.ServerBuildOnCreateServerBefore, func(value *builder.ServerBuildValue, result *builder.ServerBuildResult) {
			progCreate.Start()
		})
		sb.SetEventHandler(builder.ServerBuildOnCreateServerAfter, func(value *builder.ServerBuildValue, result *builder.ServerBuildResult) {
			progCreate.Stop()
		})

		progBoot := internal.NewProgress(
			"Still booting server...",
			"Boot server",
			ctx.IO().Progress(),
			ctx.Option().NoColor)

		sb.SetEventHandler(builder.ServerBuildOnBootBefore, func(value *builder.ServerBuildValue, result *builder.ServerBuildResult) {
			progBoot.Start()
		})
		sb.SetEventHandler(builder.ServerBuildOnBootAfter, func(value *builder.ServerBuildValue, result *builder.ServerBuildResult) {
			progBoot.Stop()
		})

	}
	return nil
}

func validateServerDiskModeParams(ctx cli.Context, params *params.BuildServerParam) []error {

	var errs []error
	var appendErrors = func(e []error) {
		errs = append(errs, e...)
	}
	var validateIfCtxIsSet = func(baseParamName string, baseParamValue interface{}, targetParamName string, targetValue interface{}) {
		if ctx.IsSet(targetParamName) {
			appendErrors(validateConflictValues(baseParamName, baseParamValue, map[string]interface{}{
				targetParamName: targetValue,
			}))
		}
	}

	switch params.DiskMode {
	case "create":
		// check required values
		appendErrors(validateRequired("disk-plan", params.DiskPlan))
		appendErrors(validateRequired("disk-connection", params.DiskConnection))
		appendErrors(validateRequired("disk-size", params.DiskSize))

		if params.SourceDiskId == 0 && params.SourceArchiveId == 0 {
			// Windows?
			if params.OsType != "" && !isWindows(params.OsType) &&
				params.OsType != "sophos-utm" && params.OsType != "netwiser" && params.OsType != "opnsense" {
				appendErrors(validateRequired("password", params.Password))
			}

		} else {
			validateIfCtxIsSet("source-archive-id", params.SourceArchiveId, "os-type", params.OsType)
			validateIfCtxIsSet("source-disk-id", params.SourceArchiveId, "os-type", params.OsType)
		}

		validateIfCtxIsSet("disk-mode", params.DiskMode, "disk-id", params.DiskId)

	case "connect":
		appendErrors(validateRequired("disk-id", params.DiskId))
		validateIfCtxIsSet("disk-mode", params.DiskMode, "disk-plan", params.DiskPlan)
		validateIfCtxIsSet("disk-mode", params.DiskMode, "disk-connection", params.DiskConnection)
		validateIfCtxIsSet("disk-size", params.DiskMode, "disk-size", params.DiskSize)
		validateIfCtxIsSet("disk-size", params.DiskMode, "os-type", params.OsType)

	case "diskless":
		validateIfCtxIsSet("disk-mode", params.DiskMode, "disk-id", params.DiskId)
		validateIfCtxIsSet("disk-mode", params.DiskMode, "disk-plan", params.DiskPlan)
		validateIfCtxIsSet("disk-mode", params.DiskMode, "disk-connection", params.DiskConnection)
		validateIfCtxIsSet("disk-size", params.DiskMode, "disk-size", params.DiskSize)
		validateIfCtxIsSet("disk-size", params.DiskMode, "os-type", params.OsType)
	}

	return errs
}

func validateServerNetworkParams(sb serverBuilder, ctx cli.Context, params *params.BuildServerParam) []error {
	var errs []error
	var appendErrors = func(e []error) {
		errs = append(errs, e...)
	}
	var validateIfCtxIsSet = func(baseParamName string, baseParamValue interface{}, targetParamName string, targetValue interface{}) {
		if ctx.IsSet(targetParamName) {
			appendErrors(validateConflictValues(baseParamName, baseParamValue, map[string]interface{}{
				targetParamName: targetValue,
			}))
		}
	}
	var validateProhibitedIfCtxIsSet = func(paramName string, paramValue interface{}) {
		if ctx.IsSet(paramName) {
			appendErrors(validateSetProhibited(paramName, paramValue))
		}
	}

	if sb.HasNetworkInterfaceProperty() {
		switch params.NetworkMode {
		case "shared", "disconnect", "none":
			validateIfCtxIsSet("network-mode", params.NetworkMode, "switch-id", params.SwitchId)
			validateIfCtxIsSet("network-mode", params.NetworkMode, "ipaddress", params.Ipaddress)
			validateIfCtxIsSet("network-mode", params.NetworkMode, "nw-masklen", params.NwMasklen)
			validateIfCtxIsSet("network-mode", params.NetworkMode, "default-route", params.DefaultRoute)

			if params.NetworkMode == "none" {
				validateIfCtxIsSet("network-mode", params.NetworkMode, "packet-filter-id", params.PacketFilterId)
			}

		case "switch":
			appendErrors(validateRequired("switch-id", params.SwitchId))
		}
	} else {
		validateProhibitedIfCtxIsSet("network-mode", params.NetworkMode)
		validateProhibitedIfCtxIsSet("switch-id", params.SwitchId)
		validateProhibitedIfCtxIsSet("ipaddress", params.Ipaddress)
		validateProhibitedIfCtxIsSet("nw-masklen", params.NwMasklen)
		validateProhibitedIfCtxIsSet("default-route", params.DefaultRoute)
		validateProhibitedIfCtxIsSet("packet-filter-id", params.PacketFilterId)
	}

	return errs
}

func validateServerDiskEditParams(sb serverBuilder, ctx cli.Context, params *params.BuildServerParam) []error {
	var errs []error
	var appendErrors = func(e []error) {
		errs = append(errs, e...)
	}
	var validateIfCtxIsSet = func(baseParamName string, baseParamValue interface{}, targetParamName string, targetValue interface{}) {
		if ctx.IsSet(targetParamName) {
			appendErrors(validateConflictValues(baseParamName, baseParamValue, map[string]interface{}{
				targetParamName: targetValue,
			}))
		}
	}
	var validateProhibitedIfCtxIsSet = func(paramName string, paramValue interface{}) {
		if ctx.IsSet(paramName) {
			appendErrors(validateSetProhibited(paramName, paramValue))
		}
	}

	if sb.HasDiskEditProperty() {
		// SSH Key generate params
		switch params.SSHKeyMode {
		case "id":
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-name", params.SSHKeyName)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-pass-phrase", params.SSHKeyPassPhrase)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-description", params.SSHKeyDescription)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-private-key-output", params.SSHKeyPrivateKeyOutput)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-public-keys", params.SSHKeyPublicKeys)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-public-key-files", params.SSHKeyPublicKeyFiles)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-ephemeral", params.SSHKeyEphemeral)
		case "generate":
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-ids", params.SSHKeyIds)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-private-key-output", params.SSHKeyPrivateKeyOutput)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-public-keys", params.SSHKeyPublicKeys)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-public-key-files", params.SSHKeyPublicKeyFiles)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-ephemeral", params.SSHKeyEphemeral)
		case "upload":

			if len(params.SSHKeyPublicKeys) == 0 && len(params.SSHKeyPublicKeyFiles) == 0 {
				errs = append(errs,
					fmt.Errorf("%q or %q is required when %q is %q",
						"ssh-key-public-keys",
						"ssh-key-public-key-files",
						"ssh-key-mode",
						"upload",
					))
			}
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-ids", params.SSHKeyIds)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-name", params.SSHKeyName)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-pass-phrase", params.SSHKeyPassPhrase)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-description", params.SSHKeyDescription)
			validateIfCtxIsSet("ssh-key-mode", params.SSHKeyMode, "ssh-key-private-key-output", params.SSHKeyPrivateKeyOutput)
		case "none":
			validateProhibitedIfCtxIsSet("ssh-key-mode", params.SSHKeyMode)
			validateProhibitedIfCtxIsSet("ssh-key-ids", params.SSHKeyIds)
			validateProhibitedIfCtxIsSet("ssh-key-name", params.SSHKeyName)
			validateProhibitedIfCtxIsSet("ssh-key-pass-phrase", params.SSHKeyPassPhrase)
			validateProhibitedIfCtxIsSet("ssh-key-description", params.SSHKeyDescription)
			validateProhibitedIfCtxIsSet("ssh-key-private-key-output", params.SSHKeyPrivateKeyOutput)
			validateProhibitedIfCtxIsSet("ssh-key-public-keys", params.SSHKeyPublicKeys)
			validateProhibitedIfCtxIsSet("ssh-key-public-key-files", params.SSHKeyPublicKeyFiles)
			validateProhibitedIfCtxIsSet("ssh-key-ephemeral", params.SSHKeyEphemeral)
		}

	} else {
		validateProhibitedIfCtxIsSet("hostname", params.Hostname)
		validateProhibitedIfCtxIsSet("password", params.Password)
		validateProhibitedIfCtxIsSet("disable-password-auth", params.DisablePasswordAuth)
		validateProhibitedIfCtxIsSet("startup-script-ids", params.StartupScriptIds)
		validateProhibitedIfCtxIsSet("startup-scripts", params.StartupScripts)
		validateProhibitedIfCtxIsSet("startup-scripts-ephemeral", params.StartupScriptsEphemeral)
		validateProhibitedIfCtxIsSet("ssh-key-mode", params.SSHKeyMode)
		validateProhibitedIfCtxIsSet("ssh-key-ids", params.SSHKeyIds)
		validateProhibitedIfCtxIsSet("ssh-key-name", params.SSHKeyName)
		validateProhibitedIfCtxIsSet("ssh-key-pass-phrase", params.SSHKeyPassPhrase)
		validateProhibitedIfCtxIsSet("ssh-key-description", params.SSHKeyDescription)
		validateProhibitedIfCtxIsSet("ssh-key-private-key-output", params.SSHKeyPrivateKeyOutput)
		validateProhibitedIfCtxIsSet("ssh-key-public-keys", params.SSHKeyPublicKeys)
		validateProhibitedIfCtxIsSet("ssh-key-public-key-files", params.SSHKeyPublicKeyFiles)
		validateProhibitedIfCtxIsSet("ssh-key-ephemeral", params.SSHKeyEphemeral)
	}

	return errs
}

func isWindows(osType string) bool {
	return strToOSType(osType).IsWindows()
}

func strToOSType(strOSType string) ostype.ArchiveOSTypes {
	return ostype.StrToOSType(strOSType)
}

type serverBuilder interface {
	builder.Builder
	builder.CommonProperty
}
