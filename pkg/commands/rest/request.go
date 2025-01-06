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

package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hokaccha/go-prettyjson"
	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/query"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/usacloud/pkg/validate"
)

var requestCommand = &core.Command{
	Name:       "request",
	Category:   "basic",
	Order:      20,
	NoProgress: true,

	ParameterInitializer: func() interface{} {
		return newRequestParameter()
	},

	Func: requestFunc,

	ValidateFunc: validateRequest,
}

type requestParameter struct {
	Zone        string `validate:"omitempty,zone"` // 通常のzoneと扱いが異なるためcflag.ZoneParameterを利用しない
	Method      string `cli:",short=X,options=rest_method" validate:"required,rest_method"`
	Data        string `cli:",short=d" validate:"omitempty,file|json"`
	Query       string `cli:",category=output,desc=Query for JSON output" validate:"omitempty" json:"-"`
	QueryDriver string `cli:",category=output,desc=Name of the driver that handles queries to JSON output options: [jmespath/jq]" json:"-" validate:"omitempty,oneof=jmespath jq"`
}

func newRequestParameter() *requestParameter {
	return &requestParameter{
		Method: "get",
	}
}

func init() {
	Resource.AddCommand(requestCommand)
}

func (p *requestParameter) queryValue() (string, error) {
	if p.Query != "" {
		return util.StringFromPathOrContent(p.Query)
	}
	return "", nil
}

func (p *requestParameter) queryDriverValue(ctx cli.Context) string {
	driver := ctx.Option().DefaultQueryDriver
	if p.QueryDriver != "" {
		driver = p.QueryDriver
	}
	if driver == "" {
		driver = query.DriverJMESPath
	}
	return driver
}

func validateRequest(ctx cli.Context, parameter interface{}) error {
	p, ok := parameter.(*requestParameter)
	if !ok {
		return fmt.Errorf("invalid parameter: %v", parameter)
	}

	if p.Zone == "" {
		p.Zone = ctx.Option().Zone
	}

	if err := validate.Exec(p); err != nil {
		return err
	}

	if p.Zone == "all" {
		// restコマンドはallゾーンに対応していないため個別に弾く
		return validate.NewValidationError(
			validate.NewFlagError("--zone", "all is not supported by rest command"),
		)
	}

	if len(ctx.Args()) != 1 {
		return validate.NewValidationError(
			validate.NewFlagError("arguments", fmt.Sprintf("accepts only 1 arg, received %d", len(ctx.Args()))),
		)
	}

	url := ctx.Args()[0]
	if !strings.HasPrefix(url, "https://") {
		if p.Zone == "" {
			return validate.NewValidationError(
				validate.NewFlagError("--zone", "required if there is no prefix 'https://' in the URL"),
			)
		}
	}
	return nil
}

func requestFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*requestParameter)
	if !ok {
		return nil, fmt.Errorf("invalid parameter: %v", parameter)
	}

	url := ctx.Args()[0]

	if !strings.HasPrefix(url, "https://") {
		if !strings.HasPrefix(url, "/") {
			url = "/" + url
		}
		url = fmt.Sprintf("%s/%s/api/cloud/1.1%s", iaas.SakuraCloudAPIRoot, p.Zone, url)
	}

	var body io.Reader
	if p.Data != "" {
		data, err := util.BytesFromPathOrContent(p.Data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, strings.ToUpper(p.Method), url, body)
	if err != nil {
		return nil, err
	}

	resp, err := ctx.Client().(client.HttpRequestDoer).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	results, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(results) > 0 {
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

		var temp interface{}
		if err := json.Unmarshal(results, &temp); err != nil {
			return nil, err
		}
		if p.Query != "" {
			driver := p.queryDriverValue(ctx)
			q, err := p.queryValue()
			if err != nil {
				return nil, err
			}
			return nil, query.Executor(driver)(temp, q, printFn)
		}
		return nil, printFn(temp)
	}
	return nil, nil
}
