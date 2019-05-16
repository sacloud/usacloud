package funcs

import (
	"fmt"
	"testing"

	"github.com/sacloud/libsacloud/builder"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/stretchr/testify/assert"
)

func TestServerBuild_CreateBuilder_FromDisk(t *testing.T) {
	param := &params.BuildServerParam{
		DiskMode:     "create",
		Name:         "withDisk",
		SourceDiskId: 999999999999,
	}

	sb := createServerBuilder(dummyContext, param)
	assert.NotNil(t, sb)

	expectedBuilder := builder.ServerFromDisk(builder.NewAPIClient(dummyContext.GetAPIClient()), param.Name, param.SourceDiskId)
	actualBuilder := sb.(builder.CommonServerBuilder)
	assert.NotNil(t, actualBuilder)

	assert.EqualValues(t, expectedBuilder, actualBuilder)
}

func TestServerBuild_CreateBuilder_FromArchive(t *testing.T) {
	param := &params.BuildServerParam{
		DiskMode:        "create",
		Name:            "withDisk",
		SourceArchiveId: 999999999999,
	}

	sb := createServerBuilder(dummyContext, param)
	assert.NotNil(t, sb)

	actualBuilder := sb.(builder.CommonServerBuilder)
	assert.NotNil(t, actualBuilder)

	expectedBuilder := builder.ServerFromArchive(builder.NewAPIClient(dummyContext.GetAPIClient()), param.Name, param.SourceArchiveId)
	assert.EqualValues(t, expectedBuilder, actualBuilder)
}

func TestServerBuild_CreateBuilder_FromBlank(t *testing.T) {
	param := &params.BuildServerParam{
		DiskMode: "create",
		Name:     "withDisk",
		// without os-type , source-disk-id , source-archive-id
	}

	sb := createServerBuilder(dummyContext, param)
	assert.NotNil(t, sb)

	actualBuilder := sb.(builder.BlankDiskServerBuilder)
	assert.NotNil(t, actualBuilder)

	expectedBuilder := builder.ServerBlankDisk(builder.NewAPIClient(dummyContext.GetAPIClient()), param.Name)
	assert.EqualValues(t, expectedBuilder, actualBuilder)
}

func TestServerBuild_CreateBuilder_FromUnix(t *testing.T) {
	param := &params.BuildServerParam{
		DiskMode: "create",
		Name:     "withDisk",
		OsType:   "centos",
		Password: "dummy_password",
	}

	sb := createServerBuilder(dummyContext, param)
	assert.NotNil(t, sb)

	actualBuilder := sb.(builder.PublicArchiveUnixServerBuilder)
	assert.NotNil(t, actualBuilder)

	expectedBuilder := builder.ServerPublicArchiveUnix(builder.NewAPIClient(dummyContext.GetAPIClient()), strToOSType(param.OsType), param.Name, param.Password)
	assert.EqualValues(t, expectedBuilder, actualBuilder)
}

func TestServerBuild_CreateBuilder_FromWindows(t *testing.T) {
	param := &params.BuildServerParam{
		DiskMode: "create",
		Name:     "withDisk",
		OsType:   "windows2016",
		Password: "dummy_password",
	}
	sb := createServerBuilder(dummyContext, param)
	assert.NotNil(t, sb)

	actualBuilder := sb.(builder.PublicArchiveWindowsServerBuilder)
	assert.NotNil(t, actualBuilder)

	expectedBuilder := builder.ServerPublicArchiveWindows(builder.NewAPIClient(dummyContext.GetAPIClient()), strToOSType(param.OsType), param.Name)
	assert.EqualValues(t, expectedBuilder, actualBuilder)
}

func TestServerBuild_HandleParams_FromUnix(t *testing.T) {
	param := &params.BuildServerParam{
		Core:                    2,
		Memory:                  4,
		Commitment:              "standard",
		DiskMode:                "create",
		OsType:                  "centos",
		DiskPlan:                "hdd",
		DiskConnection:          "virtio",
		DiskSize:                40,
		DistantFrom:             []int64{999999999999},
		IsoImageId:              999999999999,
		NetworkMode:             "switch",
		InterfaceDriver:         "virtio",
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
	var handlers = []func(serverBuilder, command.Context, *params.BuildServerParam) error{
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

	actualBuilder := sb.(builder.PublicArchiveUnixServerBuilder)
	assert.NotNil(t, actualBuilder)

	expectedBuilder := builder.ServerPublicArchiveUnix(builder.NewAPIClient(dummyContext.GetAPIClient()), strToOSType(param.OsType), param.Name, param.Password)
	var i interface{} = expectedBuilder
	{
		b := i.(serverBuilder)
		b.SetCore(param.Core)
		b.SetMemory(param.Memory)
		b.SetCommitment(sacloud.ECommitment(param.Commitment))
		b.SetServerName(param.Name)
		b.SetDescription(param.Description)
		b.SetTags(param.Tags)
		b.SetIconID(param.IconId)
		b.SetBootAfterCreate(false)
		b.SetISOImageID(param.IsoImageId)
		b.SetInterfaceDriver(sacloud.EInterfaceDriver(param.InterfaceDriver))
	}
	{
		b := i.(builder.DiskProperty)
		b.SetDiskPlan(param.DiskPlan)
		b.SetDiskConnection(sacloud.DiskConnectionVirtio)
		b.SetDiskSize(param.DiskSize)
		b.SetDistantFrom(param.DistantFrom)
	}
	{
		b := i.(builder.NetworkInterfaceProperty)
		b.AddExistsSwitchConnectedNIC(fmt.Sprintf("%d", param.SwitchId))
		b.SetPacketFilterIDs([]int64{param.PacketFilterId})
	}
	{
		b := i.(builder.DiskEditProperty)
		b.SetHostName(param.Hostname)
		b.SetPassword(param.Password)
		b.SetDisablePWAuth(param.DisablePasswordAuth)

		b.SetIPAddress(param.Ipaddress)
		b.SetDefaultRoute(param.DefaultRoute)
		b.SetNetworkMaskLen(param.NwMasklen)

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

func TestServerBuild_CreateBuilder_WithConnect(t *testing.T) {
	param := &params.BuildServerParam{
		DiskMode: "connect",
		DiskId:   999999999999,
		Name:     "connectDisk",
	}

	sb := createServerBuilder(dummyContext, param)
	assert.NotNil(t, sb)

	actualBuilder := sb.(builder.ConnectDiskServerBuilder)
	assert.NotNil(t, actualBuilder)

	expectedBuilder := builder.ServerFromExistsDisk(builder.NewAPIClient(dummyContext.GetAPIClient()), param.Name, 999999999999)
	assert.EqualValues(t, expectedBuilder, actualBuilder)
}

func TestServerBuild_CreateBuilder_FixedUnix(t *testing.T) {
	osTypes := []string{"sophos-utm", "netwiser", "opnsense"}
	for _, ostype := range osTypes {

		t.Run(ostype, func(t *testing.T) {
			param := &params.BuildServerParam{
				DiskMode: "create",
				Name:     "fixedUnix",
				OsType:   ostype,
			}

			sb := createServerBuilder(dummyContext, param)
			assert.NotNil(t, sb)

			// builder type should be builder.CommonServerBuilder
			actualBuilder := sb.(builder.FixedUnixArchiveServerBuilder)
			assert.NotNil(t, actualBuilder)

			expectedBuilder := builder.ServerPublicArchiveFixedUnix(builder.NewAPIClient(dummyContext.GetAPIClient()), strToOSType(param.OsType), param.Name)
			assert.EqualValues(t, expectedBuilder, actualBuilder)
		})
	}
}

func TestServerBuild_CreateBuilder_Diskless(t *testing.T) {
	param := &params.BuildServerParam{
		DiskMode: "diskless",
		Name:     "diskless",
	}

	sb := createServerBuilder(dummyContext, param)
	assert.NotNil(t, sb)

	// builder type should be builder.CommonServerBuilder
	actualBuilder := sb.(builder.DisklessServerBuilder)
	assert.NotNil(t, actualBuilder)

	expectedBuilder := builder.ServerDiskless(builder.NewAPIClient(dummyContext.GetAPIClient()), param.Name)
	assert.EqualValues(t, expectedBuilder, actualBuilder)
}
