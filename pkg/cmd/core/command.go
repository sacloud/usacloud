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
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/root"
	"github.com/sacloud/usacloud/pkg/cmd/services"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/printer"
	"github.com/sacloud/usacloud/pkg/term"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/usacloud/pkg/validate"
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
	ConfirmMessage string

	// パラメータ関連
	ParameterCategories  []Category
	ParameterInitializer func() interface{}
	ServiceFuncAltName   string // デフォルトのlibsacloud service呼び出しコード生成用、空の場合はNameをCamelizeしたものが利用される // TODO libsacloud側で対応すべき

	// テーブル形式での出力対象列。省略した場合はIDとNameが出力される
	ColumnDefs []output.ColumnDef

	// TODO そのうちDeprecatedWarningとかも対応する
	ExperimentWarning string

	// 操作対象リソースの一覧取得用。通常はResourceに紐づけられたfuncを利用する
	ListAllFunc func(ctx cli.Context, parameter interface{}) ([]interface{}, error)

	// 特殊な引数補完をしたい場合に設定する
	CustomCompletionFunc func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)

	// カスタムバリデーション用。空の場合usacloud/pkg/validate.Execが実行される
	ValidateFunc ValidateFunc

	// コマンドの実処理。設定してない場合はデフォルトのlibsacloud service呼び出しが行われる
	Func func(ctx cli.Context, parameter interface{}) ([]interface{}, error)

	resource         *Resource
	currentParameter interface{}
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
		Run: func(cmd *cobra.Command, args []string) {
			// コンテキスト構築
			ctx, needContinue, err := c.initCommandContext(cmd, args)
			if err != nil {
				// この段階ではctx.IO()が参照できないため標準エラーに出力する
				fmt.Fprintln(os.Stderr, err) // nolint
				return
			}
			if !needContinue {
				return
			}

			// エラー出力(可能ならカラーで出力)
			if err := c.Run(ctx, cmd, args); err != nil {
				out := ctx.IO().Err()
				(&printer.Printer{NoColor: ctx.Option().NoColor}).Fprint(out, color.New(color.FgHiRed), err)
				fmt.Fprintln(out, "") // nolint // エラーのあとは常に改行させる
				return
			}
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			// カスタムFuncが登録されていたらそちらを優先
			if c.CustomCompletionFunc != nil {
				return c.CustomCompletionFunc(cmd, args, toComplete)
			}

			if c.SelectorType == SelectorTypeNone {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}

			// コンテキスト構築
			ctx, needContinue, err := c.initCommandContext(cmd, args)
			if !needContinue || err != nil {
				return nil, cobra.ShellCompDirectiveError
			}

			return c.CompleteArgs(ctx, cmd, args, toComplete)
		},
	}

	if c, ok := c.currentParameter.(FlagInitializer); ok {
		c.SetupCobraCommandFlags(cmd)
	}

	cmd.InheritedFlags().SortFlags = false

	return cmd
}

func (c *Command) CompleteArgs(ctx cli.Context, cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// 各リソースのListAllFuncをコールし、toCompleteと前方一致するリソースを抽出
	zoned, err := c.collectResources(ctx)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	// ゾーンごとに分かれているため1つにまとめる
	var resources []interface{}
	for _, z := range zoned {
		resources = append(resources, z.resources...)
	}

	var results []string
	for _, r := range resources {
		results = append(results, c.collectCompletionValuesFromResource(r, toComplete)...)
	}

	return util.RemoveStringsFromSlice(util.UniqStrings(results), args), cobra.ShellCompDirectiveNoFileComp
}

func (c *Command) Run(ctx cli.Context, cmd *cobra.Command, args []string) error {
	if err := c.validateParameter(ctx, c.currentParameter); err != nil {
		return err
	}

	if customizer, ok := c.currentParameter.(ParameterCustomizer); ok {
		if err := customizer.Customize(ctx); err != nil {
			return err
		}
	}

	targets, err := c.expandResourceContextsFromArgs(ctx, args)
	if err != nil {
		return err
	}

	// confirm
	if ok, err := c.confirmContinue(ctx, targets); !ok || err != nil {
		return err
	}

	// 各コマンド独自の処理を実行
	results, err := c.exec(ctx, targets)
	if err != nil {
		return err
	}
	return ctx.Output().Print(results)
}

func (c *Command) resourceName() string {
	if c.resource == nil {
		return ""
	}
	return c.resource.Name
}

func (c *Command) completeParameterValue(cmd *cobra.Command, ctx cli.Context, parameter interface{}) {
	if p, ok := parameter.(FlagValueCleaner); ok {
		p.CleanupEmptyValue(cmd.Flags())
	}

	if !c.resource.IsGlobalResource {
		if zone := cflag.ZoneFlagValue(parameter); zone == "" {
			cflag.SetZoneFlagValue(parameter, ctx.Option().Zone)
		}
	}
}

func (c *Command) confirmMessage() string {
	if c.ConfirmMessage != "" {
		return c.ConfirmMessage
	}

	return c.Name
}

func (c *Command) collectCompletionValuesFromResource(resource interface{}, prefix string) []string {
	var results []string
	if v, ok := resource.(accessor.ID); ok {
		// ID
		if prefix == "" || strings.HasPrefix(v.GetID().String(), prefix) {
			results = append(results, v.GetID().String())
		}

		// Name(部分一致)
		if name, ok := resource.(accessor.Name); ok {
			if prefix == "" || strings.HasPrefix(name.GetName(), prefix) {
				results = append(results, name.GetName())
			}
		}

		// Tags
		if tags, ok := resource.(accessor.Tags); ok {
			for _, tag := range tags.GetTags() {
				if prefix == "" || strings.HasPrefix(tag, prefix) {
					results = append(results, tag)
				}
			}
		}
	}
	return results
}

func (c *Command) initCommandContext(cmd *cobra.Command, args []string) (cli.Context, bool, error) {
	ctx, err := cli.NewCLIContext(c.resource.Name, c.Name, root.Command.PersistentFlags(), args, c.ColumnDefs, c.currentParameter, c.resource.SkipLoadingProfile)
	if err != nil {
		return nil, false, err
	}

	c.printCommandWarning(ctx)

	// パラメータの補完処理(ポインタ型のクリアやコンテキストからのパラメータデフォルト値受け取りなど)
	c.completeParameterValue(cmd, ctx, c.currentParameter)

	// パラメータファイルの処理やスケルトン出力など
	needContinue, err := c.handleCommonParameters(ctx)
	return ctx, needContinue, err
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

func (c *Command) printCommandWarning(ctx cli.Context) {
	if c.resource.Warning != "" {
		c.printWarning(ctx.IO().Err(), ctx.Option().NoColor, c.resource.Warning)
	}

	if c.ExperimentWarning != "" {
		c.printWarning(ctx.IO().Err(), ctx.Option().NoColor, c.ExperimentWarning)
	}
}

func (c *Command) handleCommonParameters(ctx cli.Context) (bool, error) {
	if cp, ok := c.currentParameter.(cflag.InputParameterValueHolder); ok {
		// パラメータスケルトンの生成
		if cp.GenerateSkeletonFlagValue() {
			return false, generateSkeleton(ctx, c.currentParameter)
		}
		// --parameters/--parameter-fileフラグの処理
		if err := loadParameters(cp); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (c *Command) confirmContinue(ctx cli.Context, resources cli.ResourceContexts) (bool, error) {
	if cp, ok := c.currentParameter.(cflag.ConfirmParameterValueHandler); ok {
		if !cp.AssumeYesFlagValue() {
			if !term.IsTerminal() {
				return false, errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
			}
			result, err := util.ConfirmContinue(c.confirmMessage(), ctx.IO().In(), ctx.IO().Out(), resources.Strings()...)
			if err != nil || !result {
				return result, err
			}
		}
	}
	return true, nil
}

func (c *Command) allZoneResourceContext(ctx cli.Context) cli.ResourceContexts {
	if c.resource.IsGlobalResource {
		return cli.ResourceContexts{{Zone: ""}}
	}

	zone := cflag.ZoneFlagValue(c.currentParameter)
	if zone == "all" {
		results := cli.ResourceContexts{}
		for _, z := range ctx.Option().Zones {
			if z == "all" {
				continue
			}
			results.Append(cli.ResourceContext{Zone: z})
		}
		return results
	}
	return cli.ResourceContexts{{Zone: zone}}
}

func (c *Command) expandResourceContextsFromArgs(ctx cli.Context, args []string) (cli.ResourceContexts, error) {
	if c.SelectorType == SelectorTypeNone {
		return c.allZoneResourceContext(ctx), nil
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("ID or Name or Tags arguments are required")
	}

	collected, err := c.collectResources(ctx)
	if err != nil {
		return nil, err
	}

	results := cli.ResourceContexts{}
	for _, zonedResources := range collected {
		zone := zonedResources.zone
		resources := zonedResources.resources

		for _, r := range resources {
			if id, ok := c.extractMatchedResourceID(r, args); ok {
				results.Append(cli.ResourceContext{ID: id, Zone: zone})
			}
		}
	}

	if c.SelectorType == SelectorTypeRequireSingle && len(results) > 1 {
		return nil, fmt.Errorf("in this operation, only a single target is allowed: query=%q", args)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("target resource not found: query=%q", args)
	}
	return results, nil
}

func (c *Command) extractMatchedResourceID(r interface{}, idOrNameOrTag []string) (types.ID, bool) {
	if v, ok := r.(accessor.ID); ok {
		for _, cond := range idOrNameOrTag {
			// ID
			if v.GetID().String() == cond {
				return v.GetID(), true
			}

			// Name(部分一致)
			if name, ok := r.(accessor.Name); ok {
				if strings.Contains(name.GetName(), cond) {
					return v.GetID(), true
				}
			}

			// Tags
			if tags, ok := r.(accessor.Tags); ok {
				for _, tag := range tags.GetTags() {
					if tag == cond {
						return v.GetID(), true
					}
				}
			}
		}
	}
	return types.ID(0), false
}

func (c *Command) exec(ctx cli.Context, ids cli.ResourceContexts) (output.Contents, error) {
	if c.Func == nil {
		// use default func
		fn, ok := services.DefaultServiceFunc(c.resourceName(), c.Name)
		if !ok {
			return nil, fmt.Errorf("default service func not found: resource:%s command:%s", c.resourceName(), c.Name)
		}
		c.Func = fn
	}

	// プログレス表示が必要な場合はここでラップする
	if !c.NoProgress {
		fn := c.Func
		c.Func = func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			var results []interface{}
			progress := NewProgress(ctx)
			defer progress.Close()
			err := progress.Exec(func() error {
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

	return c.execParallel(ctx, ids)
}

func (c *Command) execParallel(ctx cli.Context, ids cli.ResourceContexts) (output.Contents, error) {
	var results output.Contents
	var errs []error

	type funcResult struct {
		results output.Contents
		err     error
	}
	resultCh := make(chan *funcResult)
	defer close(resultCh)

	var wg sync.WaitGroup
	wg.Add(len(ids))

	// 結果の受け取り
	go func() {
		for res := range resultCh {
			if res == nil {
				return
			}
			if res.err != nil {
				errs = append(errs, res.err)
			}
			if res.results != nil {
				results = append(results, res.results...)
			}
			wg.Done()
		}
	}()

	for _, rc := range ids {
		go func(ctx cli.Context) {
			p, err := c.parameterWithResourceContext(ctx)
			if err != nil {
				return
			}

			res, err := c.Func(ctx, p)
			if err != nil {
				resultCh <- &funcResult{err: err}
				return
			}

			var contents = output.Contents{}
			for _, r := range res {
				contents = append(contents, &output.Content{Zone: ctx.Zone(), ID: ctx.ID(), Value: r})
			}

			resultCh <- &funcResult{results: contents}
		}(ctx.WithResource(rc.ID, rc.Zone))
	}
	wg.Wait()

	results.Sort(ctx.Option().Zones)
	return results, util.FlattenErrors(errs)
}

func (c *Command) cloneCurrentParameter() (interface{}, error) {
	newParameter := c.ParameterInitializer()
	// mapconvDecoderを使うことで元のstructに付けられていたmapconvタグを無視する
	if err := cloneParameter(c.currentParameter, newParameter); err != nil {
		return nil, err
	}
	return newParameter, nil
}

// parameterWithID 現在のパラメータ(c.currentParameter)を複製しidを設定して返す
func (c *Command) parameterWithResourceContext(ctx cli.Context) (interface{}, error) {
	newParameter, err := c.cloneCurrentParameter()
	if err != nil {
		return nil, err
	}

	cflag.SetIDFlagValue(newParameter, ctx.ID())
	cflag.SetZoneFlagValue(newParameter, ctx.Zone())
	return newParameter, nil
}

func (c *Command) parameterWithZone(zone string) (interface{}, error) {
	newParameter, err := c.cloneCurrentParameter()
	if err != nil {
		return nil, err
	}
	cflag.SetZoneFlagValue(newParameter, zone)
	return newParameter, nil
}

func (c *Command) ParameterCategoryBy(key string) *Category {
	for _, cat := range ParameterCategories {
		if cat.Key == key {
			return cat
		}
	}

	if key == "" {
		key = c.resource.Name
		return &Category{
			Key:         key,
			DisplayName: fmt.Sprintf("%s-specific options", strings.Title(key)),
			Order:       100,
		}
	}

	for _, cat := range c.ParameterCategories {
		if cat.Key == key {
			return &cat
		}
	}
	return &Category{
		Key:         key,
		DisplayName: fmt.Sprintf("%s options", strings.Title(key)),
		Order:       200,
	}
}

type collectedResources struct {
	zone      string
	resources []interface{}
}

func (c *Command) collectResources(ctx cli.Context) ([]*collectedResources, error) {
	listFn := c.ListAllFunc
	if listFn == nil {
		fn, ok := services.DefaultListAllFunc(c.resourceName(), c.Name)
		if !ok {
			return nil, errors.New("ListAllFunc is not set")
		}
		listFn = fn
	}

	// 検索結果
	var results []*collectedResources
	var errs []error

	// 検索は各ゾーン単位で非同期なためchanで受け取る
	type funcResult struct {
		results *collectedResources
		err     error
	}
	resultCh := make(chan *funcResult)
	defer close(resultCh)

	allZones := c.allZoneResourceContext(ctx)

	var wg sync.WaitGroup
	wg.Add(len(allZones))

	// 非同期で実行されたAPIコールの結果受け取り
	go func() {
		for res := range resultCh {
			if res == nil {
				return
			}
			if res.err != nil {
				errs = append(errs, res.err)
			}
			if res.results != nil {
				results = append(results, res.results)
			}
			wg.Done()
		}
	}()

	// 各ゾーンごとに非同期にAPIコール実施
	for _, zoneCtx := range allZones {
		go func(zone string) {
			parameter, err := c.parameterWithZone(zone)
			if err != nil {
				resultCh <- &funcResult{err: err}
				return
			}

			resources, err := listFn(ctx, parameter)
			if err != nil {
				resultCh <- &funcResult{err: err}
				return
			}
			resultCh <- &funcResult{results: &collectedResources{zone: zone, resources: resources}}
		}(zoneCtx.Zone)
	}

	wg.Wait()
	if len(errs) > 0 {
		return nil, util.FlattenErrors(errs)
	}
	return results, nil
}
