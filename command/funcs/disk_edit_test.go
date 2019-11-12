// Copyright 2017-2019 The Usacloud Authors
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
	"testing"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command/params"
	"github.com/stretchr/testify/assert"
)

func TestDiskEdit(t *testing.T) {

	params := &params.EditDiskParam{
		Hostname:            "dummy",
		Password:            "dummy",
		SshKeyIds:           []int64{111111111111, 222222222222},
		DisablePasswordAuth: true,
		Ipaddress:           "192.2.0.10",
		DefaultRoute:        "192.2.0.1",
		NwMasklen:           24,
		StartupScriptIds:    []int64{333333333333, 444444444444},
	}

	ctx := &dummyCommandContext{
		outputDest: ioutil.Discard,
		flags: map[string]interface{}{
			"hostname":              params.Hostname,
			"password":              params.Password,
			"ssh-key-ids":           params.SshKeyIds,
			"disable-password-auth": params.DisablePasswordAuth,
			"ipaddress":             params.Ipaddress,
			"default-route":         params.DefaultRoute,
			"nw-masklen":            params.NwMasklen,
			"startup-script-ids":    params.StartupScriptIds,
		},
	}

	p := buildDiskEditValue(ctx, params)

	expects := []struct {
		source interface{}
		dest   interface{}
	}{
		{source: ctx.flags["hostname"], dest: *p.HostName},
		{source: ctx.flags["password"], dest: *p.Password},
		{source: ctx.flags["ssh-key-ids"], dest: extractSSHKeyIDs(p)},
		{source: ctx.flags["disable-password-auth"], dest: *p.DisablePWAuth},
		{source: ctx.flags["ipaddress"], dest: *p.UserIPAddress},
		{source: ctx.flags["default-route"], dest: p.UserSubnet.DefaultRoute},
		{source: fmt.Sprintf("%d", ctx.flags["nw-masklen"]), dest: p.UserSubnet.NetworkMaskLen},
		{source: ctx.flags["startup-script-ids"], dest: extractStartupScriptIDs(p)},
	}

	for _, expect := range expects {
		assert.EqualValues(t, expect.source, expect.dest)
	}
}

func extractSSHKeyIDs(p *sacloud.DiskEditValue) []int64 {
	var ids []int64
	for _, k := range p.SSHKeys {
		ids = append(ids, k.ID)
	}
	return ids
}

func extractStartupScriptIDs(p *sacloud.DiskEditValue) []int64 {
	var ids []int64
	for _, k := range p.Notes {
		ids = append(ids, k.ID)
	}
	return ids
}
