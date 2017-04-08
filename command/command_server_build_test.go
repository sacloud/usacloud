package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/builder"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServerBuild_CreateBuilder_FromDisk(t *testing.T) {
	param := &BuildServerParam{
		DiskMode:     "create",
		Name:         "withDisk",
		SourceDiskId: 999999999999,
	}

	sb := createServerBuilder(dummyContext, param)
	assert.NotNil(t, sb)

	// builder type should be *builder.CommonServerBuilder
	expectedBuilder := sb.(*builder.CommonServerBuilder)
	assert.NotNil(t, expectedBuilder)

	actualBuilder := builder.ServerFromDisk(dummyContext.GetAPIClient(), param.Name, param.SourceDiskId)
	assert.EqualValues(t, expectedBuilder, actualBuilder)
}

func TestServerBuild_CreateBuilder_FromArchive(t *testing.T) {
	param := &BuildServerParam{
		DiskMode:        "create",
		Name:            "withDisk",
		SourceArchiveId: 999999999999,
	}

	sb := createServerBuilder(dummyContext, param)
	assert.NotNil(t, sb)

	// builder type should be *builder.CommonServerBuilder
	expectedBuilder := sb.(*builder.CommonServerBuilder)
	assert.NotNil(t, expectedBuilder)

	actualBuilder := builder.ServerFromArchive(dummyContext.GetAPIClient(), param.Name, param.SourceArchiveId)
	assert.EqualValues(t, expectedBuilder, actualBuilder)
}

func TestServerBuild_CreateBuilder_FromUnix(t *testing.T) {
	param := &BuildServerParam{
		DiskMode: "create",
		Name:     "withDisk",
		OsType:   "centos",
		Password: "dummy_password",
	}

	sb := createServerBuilder(dummyContext, param)
	assert.NotNil(t, sb)

	// builder type should be *builder.CommonServerBuilder
	expectedBuilder := sb.(*builder.PublicArchiveUnixServerBuilder)
	assert.NotNil(t, expectedBuilder)

	actualBuilder := builder.ServerPublicArchiveUnix(dummyContext.GetAPIClient(), strToOSType(param.OsType), param.Name, param.Password)
	assert.EqualValues(t, expectedBuilder, actualBuilder)
}

func TestServerBuild_CreateBuilder_FromWindows(t *testing.T) {
	param := &BuildServerParam{
		DiskMode: "create",
		Name:     "withDisk",
		OsType:   "windows2016",
		Password: "dummy_password",
	}
	sb := createServerBuilder(dummyContext, param)
	assert.NotNil(t, sb)

	// builder type should be *builder.CommonServerBuilder
	expectedBuilder := sb.(*builder.PublicArchiveWindowsServerBuilder)
	assert.NotNil(t, expectedBuilder)

	actualBuilder := builder.ServerPublicArchiveWindows(dummyContext.GetAPIClient(), strToOSType(param.OsType), param.Name)
	assert.EqualValues(t, expectedBuilder, actualBuilder)
}

func TestServerBuild_HandleParams_FromUnix(t *testing.T) {
	param := &BuildServerParam{
		Core:                    2,
		Memory:                  4,
		DiskMode:                "create",
		OsType:                  "centos",
		DiskPlan:                "hdd",
		DiskConnection:          "virtio",
		DiskSize:                40,
		DistantFrom:             []int64{999999999999},
		IsoImageId:              999999999999,
		NetworkMode:             "switch",
		UseNicVirtio:            true,
		PacketFilterId:          999999999999,
		SwitchId:                999999999999,
		Hostname:                "dummy_hostname",
		Password:                "dummy_password",
		DisablePasswordAuth:     true,
		Ipaddress:               "192.168.2.11",
		NwMasklen:               24,
		DefaultRoute:            "192.168.2.1",
		StartupScriptIds:        []int64{999999999999},
		StartupScriptsEphemeral: true,
		SshKeyMode:              "generate",
		SshKeyName:              "dummy_keyname",
		SshKeyPassPhrase:        "dummy_passphrase",
		SshKeyDescription:       "dummy_description",
		SshKeyEphemeral:         false,
		Name:                    "dummy_name",
		Description:             "dummy_description",
		Tags:                    []string{"dummy1", "dummy2"},
		IconId:                  999999999999,
		DisableBootAfterCreate:  true,
	}

	sb := createServerBuilder(dummyContext, param)
	assert.NotNil(t, sb)
	// handle build processes
	var handlers = []func(interface{}, Context, *BuildServerParam) error{
		handleNetworkParams,
		handleDiskEditParams,
		handleDiskParams,
		handleServerCommonParams,
	}
	for _, handler := range handlers {
		err := handler(sb, dummyContext, param)
		if !assert.NoError(t, err) {
			return
		}
	}

	// builder type should be *builder.CommonServerBuilder
	expectedBuilder := sb.(*builder.PublicArchiveUnixServerBuilder)
	assert.NotNil(t, expectedBuilder)

	actualBuilder := builder.ServerPublicArchiveUnix(dummyContext.GetAPIClient(), strToOSType(param.OsType), param.Name, param.Password)
	var i interface{} = actualBuilder
	{
		b := i.(serverBuilder)
		b.SetCore(param.Core)
		b.SetMemory(param.Memory)
		b.SetServerName(param.Name)
		b.SetDescription(param.Description)
		b.SetTags(param.Tags)
		b.SetIconID(param.IconId)
		b.SetBootAfterCreate(false)
		b.SetISOImageID(param.IsoImageId)
	}
	{
		b := i.(serverDiskParams)
		b.SetDiskPlan(param.DiskPlan)
		b.SetDiskConnection(sacloud.DiskConnectionVirtio)
		b.SetDiskSize(param.DiskSize)
		b.SetDistantFrom(param.DistantFrom)
	}
	{
		b := i.(serverNetworkParams)
		b.SetUseVirtIONetPCI(param.UseNicVirtio)
		b.SetPacketFilterIDs([]int64{param.PacketFilterId})
	}
	{
		b := i.(serverConnectSwitchParamWithEditableDisk)
		b.AddExistsSwitchConnectedNIC(fmt.Sprintf("%d", param.SwitchId), param.Ipaddress, param.NwMasklen, param.DefaultRoute)
	}
	{
		b := i.(serverEditDiskParam)
		b.SetHostName(param.Hostname)
		b.SetPassword(param.Password)
		b.SetDisablePWAuth(param.DisablePasswordAuth)

		for _, v := range param.StartupScriptIds {
			b.AddNoteID(v)
		}
		b.SetNotesEphemeral(param.StartupScriptsEphemeral)

		b.SetSSHKeysEphemeral(param.SshKeyEphemeral)
		b.SetGenerateSSHKeyName(param.SshKeyName)
		b.SetGenerateSSHKeyPassPhrase(param.SshKeyPassPhrase)
		b.SetGenerateSSHKeyDescription(param.SshKeyDescription)
	}

	assert.EqualValues(t, expectedBuilder, actualBuilder)

}
