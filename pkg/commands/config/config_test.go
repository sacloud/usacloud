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

package config

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/saclient-go"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/config"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/validate"
	"github.com/stretchr/testify/require"
)

type testIO struct {
	in       *os.File
	out      bytes.Buffer
	progress bytes.Buffer
	err      bytes.Buffer
}

func newTestIO(t *testing.T) *testIO {
	t.Helper()
	f, err := os.Open("/dev/null")
	require.NoError(t, err)
	t.Cleanup(func() { _ = f.Close() })
	return &testIO{in: f}
}

func (io *testIO) In() *os.File        { return io.in }
func (io *testIO) Out() io.Writer      { return &io.out }
func (io *testIO) Progress() io.Writer { return &io.progress }
func (io *testIO) Err() io.Writer      { return &io.err }

type testContext struct {
	option   *config.Config
	io       cli.IO
	client   saclient.ClientAPI
	args     []string
	resource cli.ResourceContext
}

func newTestContext(t *testing.T, client saclient.ClientAPI, args ...string) *testContext {
	t.Helper()
	return &testContext{
		option: &config.Config{},
		io:     newTestIO(t),
		client: client,
		args:   args,
	}
}

func (c *testContext) Option() *config.Config       { return c.option }
func (c *testContext) Output() output.Output        { return output.NewDiscardOutput() }
func (c *testContext) Client() interface{}          { return nil }
func (c *testContext) IO() cli.IO                   { return c.io }
func (c *testContext) Args() []string               { return c.args }
func (c *testContext) Saclient() saclient.ClientAPI { return c.client }
func (c *testContext) PlatformName() string         { return "test" }
func (c *testContext) ResourceName() string         { return "config" }
func (c *testContext) CommandName() string          { return "test" }
func (c *testContext) WithResource(id, zone string, resource interface{}) cli.Context {
	return c
}
func (c *testContext) ID() string                        { return c.resource.ID }
func (c *testContext) Zone() string                      { return c.resource.Zone }
func (c *testContext) Resource() interface{}             { return c.resource.Resource }
func (c *testContext) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (c *testContext) Done() <-chan struct{}             { return nil }
func (c *testContext) Err() error                        { return nil }
func (c *testContext) Value(key interface{}) interface{} { return nil }

func TestMain(m *testing.M) {
	validate.InitializeValidator(iaas.SakuraCloudZones)
	os.Exit(m.Run())
}

func setupTestClient(t *testing.T) (string, saclient.ClientAPI) {
	t.Helper()
	dir := t.TempDir()
	var client saclient.Client
	require.NoError(t, client.SetEnviron([]string{"SAKURACLOUD_PROFILE_DIR=" + dir}))
	require.NoError(t, client.Populate())
	op, err := client.ProfileOp()
	require.NoError(t, err)
	require.NoError(t, op.Create(&saclient.Profile{Name: "default"}))
	require.NoError(t, op.SetCurrentName("default"))
	return dir, &client
}

func createTestProfile(t *testing.T, op saclient.ProfileAPI, name, token, secret string) {
	t.Helper()
	p := &saclient.Profile{
		Name: name,
		Attributes: map[string]any{
			"AccessToken":       token,
			"AccessTokenSecret": secret,
			"Zone":              "is1b",
		},
	}
	require.NoError(t, op.Create(p))
}

func TestCreateProfile(t *testing.T) {
	t.Run("new profile", func(t *testing.T) {
		_, client := setupTestClient(t)
		op, err := client.ProfileOp()
		require.NoError(t, err)

		ctx := newTestContext(t, client)
		p := &createParameter{
			EditParameter: EditParameter{
				Name:              "foo",
				AccessToken:       "token-foo",
				AccessTokenSecret: "secret-foo",
				Zone:              "is1b",
			},
		}

		require.NoError(t, createCommand.ValidateFunc(ctx, p))
		_, err = createProfile(ctx, p)
		require.NoError(t, err)

		read, err := op.Read("foo")
		require.NoError(t, err)
		require.Equal(t, "token-foo", read.Attributes["AccessToken"])
		require.Equal(t, "is1b", read.Attributes["Zone"])

		// current should remain default
		current, err := op.GetCurrentName()
		require.NoError(t, err)
		require.Equal(t, "default", current)
	})

	t.Run("with --use", func(t *testing.T) {
		_, client := setupTestClient(t)
		op, err := client.ProfileOp()
		require.NoError(t, err)

		ctx := newTestContext(t, client)
		p := &createParameter{
			EditParameter: EditParameter{
				Name:              "bar",
				AccessToken:       "token-bar",
				AccessTokenSecret: "secret-bar",
				Zone:              "tk1a",
				Use:               true,
			},
		}

		_, err = createProfile(ctx, p)
		require.NoError(t, err)

		current, err := op.GetCurrentName()
		require.NoError(t, err)
		require.Equal(t, "bar", current)
	})

	t.Run("already exists", func(t *testing.T) {
		_, client := setupTestClient(t)
		op, err := client.ProfileOp()
		require.NoError(t, err)
		createTestProfile(t, op, "dupe", "token", "secret")

		ctx := newTestContext(t, client)
		p := &createParameter{
			EditParameter: EditParameter{Name: "dupe"},
		}
		err = createCommand.ValidateFunc(ctx, p)
		require.Error(t, err)
		require.Contains(t, err.Error(), "already exists")
	})
}

func TestDeleteProfile(t *testing.T) {
	t.Run("delete non-current", func(t *testing.T) {
		_, client := setupTestClient(t)
		op, err := client.ProfileOp()
		require.NoError(t, err)
		createTestProfile(t, op, "foo", "token", "secret")
		createTestProfile(t, op, "bar", "token", "secret")
		require.NoError(t, op.SetCurrentName("foo"))

		ctx := newTestContext(t, client)
		p := &deleteParameter{
			ProfileParameter: ProfileParameter{Name: "bar"},
			ConfirmParameter: cflag.ConfirmParameter{AssumeYes: true},
		}

		require.NoError(t, deleteCommand.ValidateFunc(ctx, p))
		_, err = deleteFunc(ctx, p)
		require.NoError(t, err)

		_, err = op.Read("bar")
		require.Error(t, err)

		current, err := op.GetCurrentName()
		require.NoError(t, err)
		require.Equal(t, "foo", current)
	})

	t.Run("delete current resets to default", func(t *testing.T) {
		_, client := setupTestClient(t)
		op, err := client.ProfileOp()
		require.NoError(t, err)
		createTestProfile(t, op, "foo", "token", "secret")
		require.NoError(t, op.SetCurrentName("foo"))

		ctx := newTestContext(t, client)
		p := &deleteParameter{
			ProfileParameter: ProfileParameter{Name: "foo"},
			ConfirmParameter: cflag.ConfirmParameter{AssumeYes: true},
		}

		_, err = deleteFunc(ctx, p)
		require.NoError(t, err)

		_, err = op.Read("foo")
		require.Error(t, err)

		current, err := op.GetCurrentName()
		require.NoError(t, err)
		require.Equal(t, "default", current)
	})

	t.Run("delete non-existent", func(t *testing.T) {
		_, client := setupTestClient(t)

		ctx := newTestContext(t, client)
		p := &deleteParameter{
			ProfileParameter: ProfileParameter{Name: "missing"},
			ConfirmParameter: cflag.ConfirmParameter{AssumeYes: true},
		}
		err := deleteCommand.ValidateFunc(ctx, p)
		require.Error(t, err)
		require.Contains(t, err.Error(), "not exists")
	})
}

func TestUseProfile(t *testing.T) {
	_, client := setupTestClient(t)
	op, err := client.ProfileOp()
	require.NoError(t, err)

	createTestProfile(t, op, "foo", "token", "secret")
	createTestProfile(t, op, "bar", "token", "secret")
	require.NoError(t, op.SetCurrentName("foo"))

	ctx := newTestContext(t, client)
	p := &useParameter{ProfileParameter: ProfileParameter{Name: "bar"}}
	require.NoError(t, useCommand.ValidateFunc(ctx, p))
	_, err = useFunc(ctx, p)
	require.NoError(t, err)

	current, err := op.GetCurrentName()
	require.NoError(t, err)
	require.Equal(t, "bar", current)
}

func TestListProfiles(t *testing.T) {
	_, client := setupTestClient(t)
	op, err := client.ProfileOp()
	require.NoError(t, err)

	createTestProfile(t, op, "foo", "token", "secret")
	createTestProfile(t, op, "bar", "token", "secret")
	require.NoError(t, op.SetCurrentName("bar"))

	ctx := newTestContext(t, client)
	_, err = listFunc(ctx, nil)
	require.NoError(t, err)

	out := ctx.io.(*testIO).out.String()
	require.Contains(t, out, "bar")
	require.Contains(t, out, "foo")
}

func TestShowProfile(t *testing.T) {
	_, client := setupTestClient(t)
	op, err := client.ProfileOp()
	require.NoError(t, err)

	createTestProfile(t, op, "foo", "token", "secret")

	ctx := newTestContext(t, client)
	p := &showParameter{ProfileParameter: ProfileParameter{Name: "foo"}}
	require.NoError(t, showCommand.ValidateFunc(ctx, p))
	_, err = showFunc(ctx, p)
	require.NoError(t, err)

	out := ctx.io.(*testIO).out.String()
	require.Contains(t, out, "token")
}

func TestPathProfile(t *testing.T) {
	dir, client := setupTestClient(t)
	op, err := client.ProfileOp()
	require.NoError(t, err)

	createTestProfile(t, op, "foo", "token", "secret")

	ctx := newTestContext(t, client)
	p := &pathParameter{ProfileParameter: ProfileParameter{Name: "foo"}}
	require.NoError(t, pathCommand.ValidateFunc(ctx, p))
	_, err = pathFunc(ctx, p)
	require.NoError(t, err)

	out := ctx.io.(*testIO).out.String()
	require.Contains(t, out, dir)
	require.Contains(t, out, "foo")
	require.Contains(t, out, "config.json")
}

func TestCurrentProfile(t *testing.T) {
	_, client := setupTestClient(t)
	op, err := client.ProfileOp()
	require.NoError(t, err)

	createTestProfile(t, op, "foo", "token", "secret")
	require.NoError(t, op.SetCurrentName("foo"))

	ctx := newTestContext(t, client)
	_, err = currentFunc(ctx, nil)
	require.NoError(t, err)

	out := ctx.io.(*testIO).out.String()
	require.Equal(t, "foo\n", out)
}

func TestEditProfile(t *testing.T) {
	t.Run("edit existing profile", func(t *testing.T) {
		_, client := setupTestClient(t)
		op, err := client.ProfileOp()
		require.NoError(t, err)
		createTestProfile(t, op, "foo", "old-token", "old-secret")

		ctx := newTestContext(t, client)
		p := &EditParameter{
			Name:              "foo",
			AccessToken:       "new-token",
			AccessTokenSecret: "new-secret",
			Zone:              "tk1a",
		}
		_, err = editProfile(ctx, p)
		require.NoError(t, err)

		read, err := op.Read("foo")
		require.NoError(t, err)
		require.Equal(t, "new-token", read.Attributes["AccessToken"])
		require.Equal(t, "tk1a", read.Attributes["Zone"])
	})

	t.Run("create new profile", func(t *testing.T) {
		_, client := setupTestClient(t)
		op, err := client.ProfileOp()
		require.NoError(t, err)

		ctx := newTestContext(t, client)
		p := &EditParameter{
			Name:              "new-profile",
			AccessToken:       "token",
			AccessTokenSecret: "secret",
			Zone:              "is1b",
		}
		_, err = editProfile(ctx, p)
		require.NoError(t, err)

		read, err := op.Read("new-profile")
		require.NoError(t, err)
		require.Equal(t, "token", read.Attributes["AccessToken"])
		require.Equal(t, "is1b", read.Attributes["Zone"])

		current, err := op.GetCurrentName()
		require.NoError(t, err)
		require.Equal(t, "default", current)
	})

	t.Run("with --use", func(t *testing.T) {
		_, client := setupTestClient(t)
		op, err := client.ProfileOp()
		require.NoError(t, err)
		createTestProfile(t, op, "bar", "token", "secret")
		createTestProfile(t, op, "baz", "token", "secret")
		require.NoError(t, op.SetCurrentName("bar"))

		ctx := newTestContext(t, client)
		p := &EditParameter{
			Name:              "baz",
			AccessToken:       "token",
			AccessTokenSecret: "secret",
			Zone:              "is1b",
			Use:               true,
		}
		_, err = editProfile(ctx, p)
		require.NoError(t, err)

		current, err := op.GetCurrentName()
		require.NoError(t, err)
		require.Equal(t, "baz", current)
	})
}
