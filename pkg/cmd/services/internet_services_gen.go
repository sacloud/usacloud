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

package services

import (
	service "github.com/sacloud/iaas-service-go/internet"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/conv"
)

func init() {
	setDefaultServiceFunc("internet", "list-subnet",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.ListSubnetRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.ListSubnetWithContext(ctx, req)
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
	setDefaultListAllFunc("internet", "list-subnet",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
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
	setDefaultServiceFunc("internet", "add-subnet",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.AddSubnetRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.AddSubnetWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return []interface{}{res}, nil

		},
	)
	setDefaultListAllFunc("internet", "add-subnet",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
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
	setDefaultServiceFunc("internet", "delete-subnet",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.DeleteSubnetRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.DeleteSubnetWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	setDefaultListAllFunc("internet", "delete-subnet",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
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
	setDefaultServiceFunc("internet", "list",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

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
	setDefaultListAllFunc("internet", "list",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
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
	setDefaultServiceFunc("internet", "create",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.CreateRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.CreateWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return []interface{}{res}, nil

		},
	)
	setDefaultListAllFunc("internet", "create",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
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
	setDefaultServiceFunc("internet", "read",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.ReadRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.ReadWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return []interface{}{res}, nil

		},
	)
	setDefaultListAllFunc("internet", "read",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
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
	setDefaultServiceFunc("internet", "update",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.UpdateRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.UpdateWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return []interface{}{res}, nil

		},
	)
	setDefaultListAllFunc("internet", "update",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
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
	setDefaultServiceFunc("internet", "delete",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.DeleteRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.DeleteWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	setDefaultListAllFunc("internet", "delete",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
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
	setDefaultServiceFunc("internet", "enable-ipv6",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.EnableIPv6Request{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.EnableIPv6WithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return []interface{}{res}, nil

		},
	)
	setDefaultListAllFunc("internet", "enable-ipv6",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
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
	setDefaultServiceFunc("internet", "disable-ipv6",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.DisableIPv6Request{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.DisableIPv6WithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	setDefaultListAllFunc("internet", "disable-ipv6",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
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
	setDefaultServiceFunc("internet", "monitor-router",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.MonitorRouterRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.MonitorRouterWithContext(ctx, req)
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
	setDefaultListAllFunc("internet", "monitor-router",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
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
