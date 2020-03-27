package params

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// Input 入力値を保持/参照するためのインターフェース
type Input interface {
	Changed(name string) bool
	Bool(name string) (bool, error)
	String(name string) (string, error)
	StringSlice(name string) ([]string, error)
	Int(name string) (int, error)
	IntSlice(name string) ([]int, error)
	Int64(name string) (int64, error)
	Int64Slice(name string) ([]int64, error)
	ID(name string) (types.ID, error)
	IDSlice(name string) ([]types.ID, error)
}
