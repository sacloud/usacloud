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

package rest

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/core"
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
	Zone   string `validate:"omitempty,zone"` // 通常のzoneと扱いが異なるためcflag.ZoneParameterを利用しない
	Method string `cli:",short=X,options=rest_method" validate:"required,rest_method"`
	Data   string `cli:",short=d" validate:"omitempty,file|json"`
}

func newRequestParameter() *requestParameter {
	return &requestParameter{
		Method: "get",
	}
}

func init() {
	Resource.AddCommand(requestCommand)
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
		url = fmt.Sprintf("%s/%s/api/cloud/1.1%s", sacloud.SakuraCloudAPIRoot, p.Zone, url)
	}

	var body interface{}
	if p.Data != "" {
		data, err := util.BytesFromPathOrContent(p.Data)
		if err != nil {
			return nil, err
		}
		b := make(map[string]interface{})
		if err := json.Unmarshal(data, &b); err != nil {
			return nil, err
		}
		body = b
	}

	results, err := ctx.Client().Do(ctx, strings.ToUpper(p.Method), url, body)
	if err != nil {
		return nil, err
	}

	if len(results) > 0 {
		temp := make(map[string]interface{})
		if err := json.Unmarshal(results, &temp); err != nil {
			return nil, err
		}
		formattedJSON, err := json.MarshalIndent(temp, "", "    ")
		if err != nil {
			return nil, err
		}

		_, err = fmt.Fprintln(ctx.IO().Out(), string(formattedJSON))
		return nil, err
	}
	return nil, nil
}
