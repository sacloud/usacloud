// Copyright 2016-2021 The Libsacloud Authors
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

package database

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) ListParameter(req *ListParameterRequest) ([]*Parameter, error) {
	return s.ListParameterWithContext(context.Background(), req)
}

func (s *Service) ListParameterWithContext(ctx context.Context, req *ListParameterRequest) ([]*Parameter, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	client := sacloud.NewDatabaseOp(s.caller)
	parameters, err := client.GetParameter(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}
	var results []*Parameter
	for _, p := range parameters.MetaInfo {
		var setting interface{}
		for k, v := range parameters.Settings {
			if p.Name == k {
				setting = v
				break
			}
		}
		results = append(results, &Parameter{
			Key:   p.Label,
			Value: setting,
			Meta:  p,
		})
	}
	return results, nil
}
