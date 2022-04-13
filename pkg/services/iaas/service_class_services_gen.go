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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands'; DO NOT EDIT

package generated_services

import (
	"github.com/sacloud/iaas-api-go"
	service "github.com/sacloud/iaas-service-go/serviceclass"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/conv"
	"github.com/sacloud/usacloud/pkg/services/registry"
)

func init() {
	registry.SetDefaultServiceFunc("iaas", "service-class", "list",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))

			req := &service.FindRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.FindWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil

		},
	)
	registry.SetDefaultListAllFunc("iaas", "service-class", "list",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))
			req := &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}

			type requester interface {
				FindRequest() *service.FindRequest
			}
			if v, ok := parameter.(requester); ok {
				req = v.FindRequest()
			}

			res, err := svc.FindWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
}
