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

package cmd

import (
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/spf13/pflag"
)

type paramsAdapter struct {
	flags *pflag.FlagSet
}

func newParamsAdapter(flags *pflag.FlagSet) *paramsAdapter {
	return &paramsAdapter{flags: flags}
}

func (p *paramsAdapter) Changed(name string) bool {
	return p.flags.Changed(name)
}

func (p *paramsAdapter) Bool(name string) (bool, error) {
	return p.flags.GetBool(name)
}

func (p *paramsAdapter) String(name string) (string, error) {
	return p.flags.GetString(name)
}

func (p *paramsAdapter) StringSlice(name string) ([]string, error) {
	return p.flags.GetStringSlice(name)
}

func (p *paramsAdapter) Int(name string) (int, error) {
	return p.flags.GetInt(name)
}

func (p *paramsAdapter) IntSlice(name string) ([]int, error) {
	return p.flags.GetIntSlice(name)
}

func (p *paramsAdapter) Int64(name string) (int64, error) {
	return p.flags.GetInt64(name)
}

func (p *paramsAdapter) Int64Slice(name string) ([]int64, error) {
	return p.flags.GetInt64Slice(name)
}

func (p *paramsAdapter) ID(name string) (types.ID, error) {
	id, err := p.Int64(name)
	return types.ID(id), err
}

func (p *paramsAdapter) IDSlice(name string) ([]types.ID, error) {
	ids, err := p.Int64Slice(name)
	if err != nil {
		return []types.ID{}, err
	}
	out := make([]types.ID, len(ids))
	for i, id := range ids {
		out[i] = types.ID(id)
	}
	return out, nil
}
