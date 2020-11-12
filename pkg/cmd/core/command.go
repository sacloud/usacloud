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

package core

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/sacloud/usacloud/pkg/validate"

	"github.com/sacloud/usacloud/pkg/cmd/cflag"

	"github.com/sacloud/usacloud/pkg/cmd/services"

	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/root"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/term"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/spf13/cobra"
)

type SelectorType int

const (
	SelectorTypeNone          = iota // セレクタなし(List系/自前実装系など)
	SelectorTypeRequireSingle        // ID or Name or Tagsを受け取る(複数ヒットNG)
	SelectorTypeRequireMulti         // ID or Name or Tagsを受け取る(複数ヒットOK)
)

type ValidateFunc func(ctx cli.Context, parameter interface{}) error

// Command コマンド定義、実行時のコンテキスト(設定保持)にもなる
type Command struct {
	Name      string   // コマンド名、ケバブケース(ハイフン区切り)で指定すること
	Aliases   []string // エイリアス
	Usage     string
	ArgsUsage string // Argumentsの説明、省略した場合はSelectorTypeの値に応じた内容が設定される // TODO 未実装

	Category string // カテゴリーのキー
	Order    int    // コマンドが属するリソース内での並び順

	// コマンド動作関連
	SelectorType   SelectorType
	NoProgress     bool // コマンド実行時のプログレス表示の有無
	NoConfirm      bool
	ConfirmMessage string

	// パラメータ関連
	ParameterCategories   []Category
	ParameterInitializer  func() interface{}
	ParameterVariableName string // コード生成用/省略可。 省略された場合は"コマンド名+Parameter"が利用される
	ServiceFuncAltName    string // デフォルトのlibsacloud service呼び出しコード生成用、空の場合はNameをCamelizeしたものが利用される

	// テーブル形式での出力対象列。省略した場合はIDとNameが出力される
	ColumnDefs []output.ColumnDef

	// TODO そのうちDeprecatedWarningとかも対応する
	ExperimentWarning string

	// 操作対象リソースの一覧取得用。通常はResourceに紐づけられたfuncを利用する
	ListAllFunc func(ctx cli.Context, parameter interface{}) ([]interface{}, error)

	// カスタムバリデーション用。空の場合usacloud/pkg/validate.Execが実行される
	ValidateFunc ValidateFunc

	// コマンドの実処理。設定してない場合はデフォルトのlibsacloud service呼び出しが行われる
	Func func(ctx cli.Context, parameter interface{}) ([]interface{}, error)

	// これらは実行時にセットされる
	resource         *Resource
	currentParameter interface{}
}

type CategorizedCommands struct {
	Category Category
	Commands []*Command
}

func (c *Command) ResourceName() string {
	if c.resource == nil {
		return ""
	}
	return c.resource.Name
}

func (c *Command) ValidateSchema() error {
	if c.resource == nil {
		return fmt.Errorf(`command "%s" has invalid schema: resource required`, c.Name)
	}
	if c.ParameterInitializer == nil {
		return fmt.Errorf(`command "%s %s" has invalid schema: ParameterInitializer required`, c.resource.Name, c.Name)
	}
	return nil
}

func (c *Command) CLICommand() *cobra.Command {
	c.currentParameter = c.ParameterInitializer()
	cmd := &cobra.Command{
		Use:          c.Name,
		Aliases:      c.Aliases,
		Short:        c.Usage,
		Long:         c.Usage,
		SilenceUsage: true,
		RunE:         c.Run,
	}

	if c, ok := c.currentParameter.(FlagInitializer); ok {
		c.SetupCobraCommandFlags(cmd)
	}
	return cmd
}

func (c *Command) confirmMessage() string {
	if c.ConfirmMessage != "" {
		return c.ConfirmMessage
	}

	return c.Name
}

func (c *Command) Run(cmd *cobra.Command, args []string) error {
	ctx, err := cli.NewCLIContext(c.resource.Name, c.Name, root.Command.PersistentFlags(), args, c.ColumnDefs, c.currentParameter)
	if err != nil {
		return err
	}

	c.completeParameterValue(cmd, ctx, c.currentParameter)

	if err := c.validateParameter(ctx, c.currentParameter); err != nil {
		return err
	}

	c.printCommandWarning(ctx)

	if cp, ok := c.currentParameter.(cflag.CommonParameterValueHolder); ok {
		// パラメータスケルトンの生成
		if cp.GenerateSkeletonFlagValue() {
			return generateSkeleton(ctx, c.currentParameter)
		}
		// --parameters/--parameter-fileフラグの処理
		if err := loadParameters(cp); err != nil {
			return err
		}
	}

	ids, err := c.expandIDsFromArgs(ctx, c.currentParameter, args)
	if err != nil {
		return err
	}

	// confirm
	if !c.NoConfirm {
		if cp, ok := c.currentParameter.(ConfirmParameterValueHandler); ok {
			if !cp.AssumeYesFlagValue() {
				if !term.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := util.ConfirmContinue(c.confirmMessage(), ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}
		}
	}

	// 各コマンド独自の処理を実行
	results, err := c.exec(ctx, ids)
	if err != nil {
		return err
	}
	return ctx.Output().Print(results)
}

func (c *Command) validateParameter(ctx cli.Context, parameter interface{}) error {
	validateFunc := c.ValidateFunc
	if validateFunc == nil {
		validateFunc = func(_ cli.Context, p interface{}) error {
			return validate.Exec(p)
		}
	}

	return validateFunc(ctx, parameter)
}

func (c *Command) completeParameterValue(cmd *cobra.Command, ctx cli.Context, parameter interface{}) {
	if p, ok := parameter.(FlagValueCleaner); ok {
		p.CleanupEmptyValue(cmd.Flags())
	}

	if zp, ok := parameter.(cflag.ZoneParameterValueHandler); ok {
		if zp.ZoneFlagValue() == "" {
			zp.SetZoneFlagValue(ctx.Zone())
		}
	}
}

func (c *Command) printCommandWarning(ctx cli.Context) {
	if c.resource.Warning != "" {
		c.printWarning(ctx.IO().Err(), ctx.Option().NoColor, c.resource.Warning)
	}

	if c.ExperimentWarning != "" {
		c.printWarning(ctx.IO().Err(), ctx.Option().NoColor, c.ExperimentWarning)
	}
}

func (c *Command) expandIDsFromArgs(ctx cli.Context, parameter interface{}, args []string) ([]types.ID, error) {
	if c.SelectorType == SelectorTypeNone {
		return nil, nil
	}

	listFn := c.ListAllFunc
	if listFn == nil {
		fn, ok := services.DefaultListAllFunc(c.ResourceName(), c.Name)
		if !ok {
			return nil, errors.New("ListAllFunc is not set")
		}
		listFn = fn
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("ID or Name or Tags arguments are required")
	}

	// 引数が1つ、かつtypes.IDへの変換が成功した場合は検索せずに返す
	if len(args) == 1 {
		id := types.StringID(args[0])
		if !id.IsEmpty() {
			return []types.ID{id}, nil
		}
	}

	resources, err := listFn(ctx, parameter)
	if err != nil {
		return nil, err
	}

	if c.SelectorType == SelectorTypeRequireSingle && len(resources) > 1 {
		return nil, fmt.Errorf("target resource not found: query=[%q]", args)
	}

	var ids []types.ID
	for _, r := range resources {
		for _, arg := range args {
			if v, ok := r.(accessor.ID); ok {
				// ID
				if v.GetID().String() == arg {
					ids = append(ids, v.GetID())
				}

				// Name
				if name, ok := r.(accessor.Name); ok {
					if name.GetName() == arg {
						ids = append(ids, v.GetID())
					}
				}

				// Tags
				if tags, ok := r.(accessor.Tags); ok {
					for _, tag := range tags.GetTags() {
						if tag == arg {
							ids = append(ids, v.GetID())
						}
					}
				}
			}
		}
	}
	return util.UniqIDs(ids), nil
}

func (c *Command) exec(ctx cli.Context, ids []types.ID) ([]interface{}, error) {
	if c.Func == nil {
		// use default func
		fn, ok := services.DefaultServiceFunc(c.ResourceName(), c.Name)
		if !ok {
			return nil, fmt.Errorf("default service func not found: resource:%s command:%s", c.ResourceName(), c.Name)
		}
		c.Func = fn
	}

	// プログレス表示が必要な場合はここでラップする
	if !c.NoProgress {
		fn := c.Func
		c.Func = func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			var results []interface{}
			err := ctx.ExecWithProgress(func() error {
				res, err := fn(ctx, parameter)
				if err != nil {
					return err
				}
				results = res
				return nil
			})
			return results, err
		}
	}

	if len(ids) == 0 {
		return c.Func(ctx, c.currentParameter)
	}
	return c.execParallel(ctx, ids)
}

func (c *Command) execParallel(ctx cli.Context, ids []types.ID) ([]interface{}, error) {
	var wg sync.WaitGroup
	var results []interface{}
	var errs []error
	for _, id := range ids {
		wg.Add(1)
		go func(ctx cli.Context, id types.ID) {
			p, err := c.parameterWithID(id)
			if err != nil {
				errs = append(errs, err)
			} else {
				res, err := c.Func(ctx, p)
				if err != nil {
					errs = append(errs, err)
				}
				results = append(results, res...)
			}
			wg.Done()
		}(ctx.WithID(id), id)
	}
	wg.Wait()
	return results, util.FlattenErrors(errs)
}

var mapconvDecoder = mapconv.Decoder{Config: &mapconv.DecoderConfig{TagName: "temp"}}

// parameterWithID 現在のパラメータ(c.currentParameter)を複製しidを設定して返す
func (c *Command) parameterWithID(id types.ID) (interface{}, error) {
	newParameter := c.ParameterInitializer()

	// mapconvDecoderを使うことで元のstructに付けられていたmapconvタグを無視する
	if err := mapconvDecoder.ConvertTo(c.currentParameter, newParameter); err != nil {
		return nil, err
	}
	if v, ok := newParameter.(IDParameterValueHandler); ok {
		v.SetIDFlagValue(id)
	}
	return newParameter, nil
}

func (c *Command) ParameterCategoryBy(key string) *Category {
	switch key {
	case "":
		return ParameterCategoryDefault
	case "output":
		return ParameterCategoryOutput
	case "input":
		return ParameterCategoryInput
	case "common":
		return ParameterCategoryCommon
	case "sort":
		return ParameterCategorySort
	case "limit-offset":
		return ParameterCategoryLimitOffset
	case "filter":
		return ParameterCategoryFilter
	default:
		if len(c.ParameterCategories) == 0 {
			return &Category{
				Key:         key,
				DisplayName: fmt.Sprintf("%s options", strings.Title(key)),
				Order:       1,
			}
		}

		for _, cat := range c.ParameterCategories {
			if cat.Key == key {
				return &cat
			}
		}
		return nil
	}
}
