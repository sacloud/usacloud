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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands'; DO NOT EDIT

package generated_services

import (
	"github.com/sacloud/iaas-api-go"
	service "github.com/sacloud/iaas-service-go/disk"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/conv"
	"github.com/sacloud/usacloud/pkg/services/registry"
)

func init() {
	registry.SetDefaultServiceFunc("iaas", "disk", "connect-to-server",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))

			req := &service.ConnectToServerRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.ConnectToServerWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	registry.SetDefaultListAllFunc("iaas", "disk", "connect-to-server",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))
			req := &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()}

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
	registry.SetDefaultServiceFunc("iaas", "disk", "edit",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))

			req := &service.EditRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.EditWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	registry.SetDefaultListAllFunc("iaas", "disk", "edit",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))
			req := &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()}

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
	registry.SetDefaultServiceFunc("iaas", "disk", "disconnect-from-server",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))

			req := &service.DisconnectFromServerRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.DisconnectFromServerWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	registry.SetDefaultListAllFunc("iaas", "disk", "disconnect-from-server",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))
			req := &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()}

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
	registry.SetDefaultServiceFunc("iaas", "disk", "resize-partition",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))

			req := &service.ResizePartitionRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.ResizePartitionWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	registry.SetDefaultListAllFunc("iaas", "disk", "resize-partition",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))
			req := &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()}

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
	registry.SetDefaultServiceFunc("iaas", "disk", "list",
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
	registry.SetDefaultListAllFunc("iaas", "disk", "list",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))
			req := &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()}

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
	registry.SetDefaultServiceFunc("iaas", "disk", "create",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))

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
	registry.SetDefaultListAllFunc("iaas", "disk", "create",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))
			req := &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()}

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
	registry.SetDefaultServiceFunc("iaas", "disk", "read",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))

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
	registry.SetDefaultListAllFunc("iaas", "disk", "read",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))
			req := &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()}

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
	registry.SetDefaultServiceFunc("iaas", "disk", "update",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))

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
	registry.SetDefaultListAllFunc("iaas", "disk", "update",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))
			req := &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()}

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
	registry.SetDefaultServiceFunc("iaas", "disk", "delete",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))

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
	registry.SetDefaultListAllFunc("iaas", "disk", "delete",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))
			req := &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()}

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
	registry.SetDefaultServiceFunc("iaas", "disk", "monitor-disk",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))

			req := &service.MonitorDiskRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.MonitorDiskWithContext(ctx, req)
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
	registry.SetDefaultListAllFunc("iaas", "disk", "monitor-disk",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))
			req := &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()}

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
	registry.SetDefaultServiceFunc("iaas", "disk", "wait-until-ready",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))

			req := &service.WaitReadyRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			err := svc.WaitReadyWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return nil, nil

		},
	)
	registry.SetDefaultListAllFunc("iaas", "disk", "wait-until-ready",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client().(iaas.APICaller))
			req := &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()}

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
