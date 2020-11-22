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

package cli

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sacloud/libsacloud/v2"
	"github.com/sacloud/libsacloud/v2/helper/api"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/config"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/validate"
	"github.com/sacloud/usacloud/pkg/version"
	"github.com/spf13/pflag"
)

type Context interface {
	Option() *config.Config
	Output() output.Output
	Client() sacloud.APICaller
	IO() IO
	context.Context

	Args() []string

	ResourceName() string
	CommandName() string

	// WithResource 特定のリソース向け作業をする際に呼ばれる。
	// リソースのIDとゾーンを保持した新しいコンテキストを返す
	// 新しいコンテキストの親コンテキストには現在のコンテキストが設定される
	WithResource(id types.ID, zone string) Context
	ID() types.ID
	Zone() string
}

type cliContext struct {
	parentCtx context.Context
	option    *config.Config
	output    output.Output
	cliIO     IO
	args      []string

	resourceName string
	commandName  string
	resource     ResourceContext
}

func NewCLIContext(resourceName, commandName string, globalFlags *pflag.FlagSet, args []string, columnDefs []output.ColumnDef, parameter interface{}) (Context, error) {
	// TODO あとでグローバルなタイムアウトなどを実装する
	ctx := context.TODO()

	io := newIO()

	option, err := config.LoadConfigValue(globalFlags, io.Err())
	if err != nil {
		return nil, err
	}

	// initialize validator with contextual values
	validate.InitializeValidator(option.Zones)

	return &cliContext{
		parentCtx:    ctx,
		option:       option,
		output:       getOutputWriter(io, columnDefs, parameter),
		resourceName: resourceName,
		commandName:  commandName,
		cliIO:        io,
		args:         args,
	}, nil
}

func (c *cliContext) IO() IO {
	return c.cliIO
}

func (c *cliContext) Option() *config.Config {
	return c.option
}

func (c *cliContext) Output() output.Output {
	return c.output
}

func (c *cliContext) ResourceName() string {
	return c.resourceName
}

func (c *cliContext) CommandName() string {
	return c.commandName
}

func (c *cliContext) ID() types.ID {
	return c.resource.ID
}

func (c *cliContext) Zone() string {
	return c.resource.Zone
}

func (c *cliContext) WithResource(id types.ID, zone string) Context {
	return &cliContext{
		parentCtx:    c,
		option:       c.option,
		output:       c.output,
		cliIO:        c.cliIO,
		args:         c.args,
		resourceName: c.resourceName,
		commandName:  c.commandName,
		resource:     ResourceContext{ID: id, Zone: zone},
	}
}

func (c *cliContext) Client() sacloud.APICaller {
	o := c.Option()
	if o.FakeMode {
		// libsacloud fakeドライバはlogパッケージにシステムログを出すがusacloudからは利用しないため出力を抑制する
		log.SetOutput(ioutil.Discard)
	}

	return api.NewCaller(&api.CallerOptions{
		AccessToken:          o.AccessToken,
		AccessTokenSecret:    o.AccessTokenSecret,
		APIRootURL:           o.APIRootURL,
		DefaultZone:          o.DefaultZone,
		AcceptLanguage:       o.AcceptLanguage,
		HTTPClient:           http.DefaultClient,
		HTTPRequestTimeout:   o.HTTPRequestTimeout,
		HTTPRequestRateLimit: o.HTTPRequestRateLimit,
		RetryMax:             o.RetryMax,
		RetryWaitMax:         o.RetryWaitMax,
		RetryWaitMin:         o.RetryWaitMin,
		UserAgent:            fmt.Sprintf("Usacloud/v%s (+https://github.com/sacloud/usacloud) libsacloud/%s", version.Version, libsacloud.Version),
		TraceAPI:             o.EnableAPITrace(),
		TraceHTTP:            o.EnableHTTPTrace(),
		FakeMode:             o.FakeMode,
		FakeStorePath:        o.FakeStorePath,
	})
}

func (c *cliContext) Deadline() (deadline time.Time, ok bool) {
	return c.parentCtx.Deadline()
}

func (c *cliContext) Done() <-chan struct{} {
	return c.parentCtx.Done()
}

func (c *cliContext) Err() error {
	return c.parentCtx.Err()
}

func (c *cliContext) Value(key interface{}) interface{} {
	return c.parentCtx.Value(key)
}

func (c *cliContext) Args() []string {
	return c.args
}

func getOutputWriter(io IO, columnDefs []output.ColumnDef, rawOptions interface{}) output.Output {
	if rawOptions == nil {
		return output.NewDiscardOutput()
	}
	options, ok := rawOptions.(output.Option)
	if !ok {
		return output.NewDiscardOutput()
	}

	out := io.Out()
	err := io.Err()

	if options.QuietFlagValue() {
		return output.NewIDOutput(out, err)
	}
	if options.FormatFlagValue() != "" || options.FormatFileFlagValue() != "" {
		return output.NewFreeOutput(out, err, options)
	}
	switch options.OutputTypeFlagValue() {
	case "json":
		query := options.QueryFlagValue()
		if query == "" {
			bQuery, _ := ioutil.ReadFile(options.QueryFileFlagValue()) // nolint: err was already checked
			query = string(bQuery)
		}
		return output.NewJSONOutput(out, err, query)
	case "yaml":
		return output.NewYAMLOutput(out, err)
	default:
		return output.NewTableOutput(out, err, columnDefs)
	}
}
