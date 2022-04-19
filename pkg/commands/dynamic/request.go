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

package dynamic

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hokaccha/go-prettyjson"
	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/phy-api-go"
	phyServices "github.com/sacloud/phy-service-go"
	"github.com/sacloud/services/dispatcher"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
)

var requestCommand = &core.Command{
	Name:       "request",
	Category:   "basic",
	Order:      20,
	NoProgress: true,

	ParameterInitializer: func() interface{} {
		return newRequestParameter()
	},

	Func: run,
}

type requestParameter struct {
	Data string `cli:",short=d" validate:"omitempty,file|json"`
}

func newRequestParameter() *requestParameter {
	return &requestParameter{}
}

func init() {
	Resource.AddCommand(requestCommand)
}

func run(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	param := parameter.(*requestParameter)

	rawArgs := ctx.Args()
	if len(rawArgs) < 2 {
		return nil, fmt.Errorf("invalid args: %s", rawArgs)
	}

	platform := rawArgs[0]
	if err := initService(ctx, platform); err != nil {
		return nil, err
	}

	var data interface{}
	if param.Data != "" {
		var v map[string]interface{}
		if err := json.Unmarshal([]byte(param.Data), &v); err != nil {
			return nil, err
		}
		data = v
	}

	result, err := dispatcher.Dispatch(rawArgs, data)
	if err != nil {
		return nil, err
	}

	if result != nil {
		printFn := func(v interface{}) error {
			formatter := prettyjson.NewFormatter()
			formatter.DisabledColor = ctx.Option().NoColor
			formatter.Indent = 4

			data, err := formatter.Marshal(v)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(ctx.IO().Out(), string(data))
			return err
		}

		return nil, printFn(result)
	}
	return nil, nil
}

func initService(ctx cli.Context, platform string) error {
	o := ctx.Option()
	switch platform {
	case "phy":
		services := phyServices.Services(&phy.Client{
			Options: &client.Options{
				AccessToken:          o.AccessToken,
				AccessTokenSecret:    o.AccessTokenSecret,
				AcceptLanguage:       o.AcceptLanguage,
				HttpClient:           http.DefaultClient,
				HttpRequestTimeout:   o.HTTPRequestTimeout,
				HttpRequestRateLimit: o.HTTPRequestRateLimit,
				RetryMax:             o.RetryMax,
				RetryWaitMax:         o.RetryWaitMax,
				RetryWaitMin:         o.RetryWaitMin,
				UserAgent:            cli.UserAgent,
				Trace:                o.EnableHTTPTrace(),
			},
		})
		dispatcher.Register("phy", services...)
		return nil
	}
	return fmt.Errorf("unsupported platform: %s", platform)
}
