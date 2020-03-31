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

package commands

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/spf13/pflag"

	libsacloud "github.com/sacloud/libsacloud/v2"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/fake"
	"github.com/sacloud/libsacloud/v2/sacloud/trace"
	"github.com/sacloud/libsacloud/v2/utils/builder"
	"github.com/sacloud/libsacloud/v2/utils/setup"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/version"
)

type Context interface {
	Option() *CLIOptions
	Output() output.Output
	Client() sacloud.APICaller
	Zone() string
	IO() IO
	context.Context
}

type cliContext struct {
	parentCtx context.Context
	option    *CLIOptions
	output    output.Output
	cliIO     *cliIO

	client     sacloud.APICaller
	clientOnce sync.Once
}

func newCLIContext(globalFlags *pflag.FlagSet, formatter interface{}) (Context, error) {
	// TODO あとで変更する
	ctx := context.TODO()

	// TODO あとで切り替え処理を実装
	io := &cliIO{
		in:       os.Stdin,
		out:      os.Stdout,
		progress: os.Stderr,
		err:      os.Stderr,
	}

	option, err := initCLIOptions(globalFlags, io)
	if err != nil {
		return nil, err
	}

	return &cliContext{
		parentCtx: ctx,
		option:    option,
		output:    getOutputWriter(io, formatter),
		cliIO:     io,
	}, nil
}

func (c *cliContext) IO() IO {
	return c.cliIO
}

func (c *cliContext) Option() *CLIOptions {
	return c.option
}

func (c *cliContext) Output() output.Output {
	return c.output
}

func (c *cliContext) Client() sacloud.APICaller {
	c.clientOnce.Do(func() {
		o := c.Option()

		httpClient := http.DefaultClient
		httpClient.Timeout = time.Duration(o.HTTPRequestTimeout) * time.Second
		httpClient.Transport = &sacloud.RateLimitRoundTripper{RateLimitPerSec: o.HTTPRequestRateLimit}

		retryWaitMax := sacloud.APIDefaultRetryWaitMax
		retryWaitMin := sacloud.APIDefaultRetryWaitMin
		if o.RetryWaitMax > 0 {
			retryWaitMax = time.Duration(o.RetryWaitMax) * time.Second
		}
		if o.RetryWaitMin > 0 {
			retryWaitMin = time.Duration(o.RetryWaitMin) * time.Second
		}

		ua := fmt.Sprintf("Usacloud/ (+https://github.com/sacloud/usacloud) cli/v%s libsacloud/%s", version.Version, libsacloud.Version)

		caller := &sacloud.Client{
			AccessToken:       o.AccessToken,
			AccessTokenSecret: o.AccessTokenSecret,
			UserAgent:         ua,
			AcceptLanguage:    o.AcceptLanguage,
			RetryMax:          o.RetryMax,
			RetryWaitMax:      retryWaitMax,
			RetryWaitMin:      retryWaitMin,
			HTTPClient:        httpClient,
		}
		sacloud.DefaultStatePollingTimeout = 72 * time.Hour

		if o.TraceMode != "" {
			enableAPITrace := true
			enableHTTPTrace := true

			mode := strings.ToLower(o.TraceMode)
			switch mode {
			case "api":
				enableHTTPTrace = false
			case "http":
				enableAPITrace = false
			}

			if enableAPITrace {
				// note: exact once
				trace.AddClientFactoryHooks()
			}
			if enableHTTPTrace {
				caller.HTTPClient.Transport = &sacloud.TracingRoundTripper{
					Transport: caller.HTTPClient.Transport,
				}
			}
		}

		if o.FakeMode {
			if o.FakeStorePath != "" {
				fake.DataStore = fake.NewJSONFileStore(o.FakeStorePath)
			}
			// note: exact once
			fake.SwitchFactoryFuncToFake()

			defaultInterval := 10 * time.Millisecond

			// update default polling intervals: libsacloud/sacloud
			sacloud.DefaultStatePollingInterval = defaultInterval
			// update default polling intervals: libsacloud/utils/setup
			setup.DefaultDeleteWaitInterval = defaultInterval
			setup.DefaultProvisioningWaitInterval = defaultInterval
			setup.DefaultPollingInterval = defaultInterval
			// update default polling intervals: libsacloud/utils/builder
			builder.DefaultNICUpdateWaitDuration = defaultInterval
		}

		zones := o.Zones
		if len(zones) == 0 {
			zones = sacloud.SakuraCloudZones
		}
		if o.APIRootURL != "" {
			if strings.HasSuffix(o.APIRootURL, "/") {
				o.APIRootURL = strings.TrimRight(o.APIRootURL, "/")
			}
			sacloud.SakuraCloudAPIRoot = o.APIRootURL
		}
		c.client = caller
	})

	return c.client
}

func (c *cliContext) Zone() string {
	return c.Option().Zone
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

func getOutputWriter(io IO, rawFormatter interface{}) output.Output {
	if rawFormatter == nil {
		return nil
	}
	formatter, ok := rawFormatter.(output.Formatter)
	if !ok {
		return nil
	}

	out := io.Out()
	err := io.Err()

	if formatter.GetQuiet() {
		return output.NewIDOutput(out, err)
	}
	if formatter.GetFormat() != "" || formatter.GetFormatFile() != "" {
		return output.NewFreeOutput(out, err, formatter)
	}
	switch formatter.GetOutputType() {
	case "csv":
		return output.NewRowOutput(out, err, ',', formatter)
	case "tsv":
		return output.NewRowOutput(out, err, '\t', formatter)
	case "json":
		query := formatter.GetQuery()
		if query == "" {
			bQuery, _ := ioutil.ReadFile(formatter.GetQueryFile()) // nolint: err was already checked
			query = string(bQuery)
		}
		return output.NewJSONOutput(out, err, query)
	case "yaml":
		return output.NewYAMLOutput(out, err)
	default:
		return output.NewTableOutput(out, err, formatter)
	}
}
