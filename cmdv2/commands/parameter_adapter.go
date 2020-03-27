package commands

import (
	"strconv"
	"strings"

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
	f := p.flags.Lookup(name)
	if f == nil {
		return []int64{}, nil
	}
	val := strings.Trim(f.Value.String(), "[]")
	if len(val) == 0 {
		return []int64{}, nil
	}
	ss := strings.Split(val, ",")
	out := make([]int64, len(ss))
	for i, d := range ss {
		var err error
		out[i], err = strconv.ParseInt(d, 10, 64)
		if err != nil {
			return nil, err
		}

	}
	return out, nil
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
