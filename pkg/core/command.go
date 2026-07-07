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

package core

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/sacloud/usacloud/pkg/category"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/commands/root"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/printer"
	"github.com/sacloud/usacloud/pkg/services"
	"github.com/sacloud/usacloud/pkg/term"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/usacloud/pkg/validate"
	"github.com/spf13/cobra"
)

// SelectorType コマンドが引数として受け取るリソース選択子の種別
//
// ID/Name/Tags を使ったリソースの絞り込み方法を制御する
type SelectorType int

const (
	// SelectorTypeNone セレクタなし。
	// list や create のように、引数から操作対象リソースを特定しないコマンドで使用する。
	SelectorTypeNone = iota
	// SelectorTypeRequireSingle ID or Name or Tags を受け取り、単一のリソースに絞り込む。
	// 複数ヒットした場合はエラーとなる。read や update、delete などで使用する。
	SelectorTypeRequireSingle
	// SelectorTypeRequireMulti ID or Name or Tags を受け取り、複数のリソースを操作対象とする。
	// boot や shutdown のように複数リソースを一括操作するコマンドで使用する。
	SelectorTypeRequireMulti
)

// ValidateFunc コマンドパラメータのカスタムバリデーション関数
//
// 設定されていない場合は pkg/validate.Exec がデフォルトで使用される
type ValidateFunc func(ctx cli.Context, parameter interface{}) error

// Command 1 つのサブコマンドを表す定義。
//
// コマンド定義と同時に、実行時のコンテキスト（カレントパラメータの保持）としても機能する。
// ParameterInitializer は必須。Func が未設定の場合は pkg/services/registry から
// 自動生成されたサービス関数が解決される。
type Command struct {
	// Name コマンド名。ケバブケース（ハイフン区切り）で指定すること。
	Name string
	// Aliases コマンドのエイリアス。ls/find/select など。
	Aliases []string
	// Usage ヘルプに表示される短い説明文。
	Usage string
	// ArgsUsage Arguments の説明。省略した場合は SelectorType の値に応じた内容が自動設定される。
	ArgsUsage string

	// Category コマンドが属するカテゴリのキー。
	// basic/operation/power/monitor/other など、pkg/category/commands.go で定義される値を使用する。
	Category string
	// Order 同カテゴリ内での並び順。小さいほど先に表示される。
	Order int

	// コマンド動作関連
	// SelectorType 引数として受け取るリソース選択子の種別。
	SelectorType SelectorType
	// NoProgress true の場合、コマンド実行中のプログレス表示を行わない。
	NoProgress bool
	// ConfirmMessage 確認ダイアログで表示するメッセージ。未設定時は Name が使用される。
	ConfirmMessage string

	// パラメータ関連
	// ParameterCategories コマンド固有のパラメータカテゴリ。
	ParameterCategories []category.Category
	// ParameterInitializer コマンドパラメータを初期化する関数。必須。
	ParameterInitializer func() interface{}
	// ServiceFuncAltName 自動生成されるサービス関数呼び出しで、コマンド名以外のメソッド名を使う場合に指定する。
	// 空の場合は Name を Camelize したものが利用される。
	ServiceFuncAltName string

	// ColumnDefs テーブル形式での出力対象列。省略した場合は ID と Name が出力される。
	ColumnDefs []output.ColumnDef

	// ExperimentWarning 実験的機能として実行前に表示する警告メッセージ。
	ExperimentWarning string

	// ListAllFunc 操作対象リソースの一覧取得用関数。
	// 通常は Resource に紐づけられた自動生成関数を利用する。特殊な一覧取得が必要な場合に設定する。
	ListAllFunc func(ctx cli.Context, parameter interface{}) ([]interface{}, error)

	// CustomCompletionFunc 特殊な引数補完を行いたい場合に設定する関数。
	// 未設定の場合は ListAllFunc から取得したリソースの ID/Name/Tags で補完される。
	CustomCompletionFunc func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)

	// ValidateFunc カスタムバリデーション用。
	// 空の場合は usacloud/pkg/validate.Exec が実行される。
	ValidateFunc ValidateFunc

	// Func コマンドの実処理。
	// 設定してない場合はデフォルトの service（iaas-service-go / webaccel-api-go 等）呼び出しが行われる。
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

func (c *Command) argsUsage() string {
	// NOTE: cobra.Command#Useには"コマンド名 usage"のように最初のスペース以前がコマンド名となる値を設定しないと
	//       cobra.Command#Name()が正しい値を返さない仕様となっている。
	//       ref: https://github.com/spf13/cobra/blob/86f8bfd7fef868a174e1b606783bd7f5c82ddf8f/command.go#L1286-L1294
	if c.ArgsUsage != "" {
		usage := c.ArgsUsage
		if !strings.HasPrefix(usage, " ") {
			usage = " " + usage
		}
		return usage
	}

	if c.SelectorType == SelectorTypeNone {
		return ""
	}

	suffix := ""
	if c.SelectorType == SelectorTypeRequireMulti {
		suffix = "..."
	}
	return " { ID | NAME | TAG }" + suffix
}

func (c *Command) CLICommand() *cobra.Command {
	c.currentParameter = c.ParameterInitializer()
	cmd := &cobra.Command{
		Use:          c.Name + c.argsUsage(),
		Aliases:      c.Aliases,
		Short:        c.Usage,
		Long:         c.Usage,
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			// コンテキスト構築
			var err error
			ctx, cancel, needContinue, err := c.initCommandContext(cmd, args)
			if err != nil {
				// この段階ではctx.IO()が参照できないため標準エラーに出力する
				fmt.Fprintln(os.Stderr, err)
				return
			}
			defer cancel()

			if !needContinue {
				return
			}

			doneCh := make(chan error)
			defer close(doneCh)

			go func() {
				doneCh <- c.Run(ctx, cmd, args)
			}()

			select {
			case <-ctx.Done():
				err = fmt.Errorf("command[%s/%s] timed out: %s", c.resource.Name, c.Name, ctx.Err())
				break
			case e := <-doneCh:
				err = e
				break
			}
			if err != nil {
				out := ctx.IO().Err()
				fmt.Fprintln(out, "")
				(&printer.Printer{NoColor: ctx.Option().NoColor}).Fprint(out, color.New(color.FgHiRed), err)
				fmt.Fprintln(out, "")
				os.Exit(1)
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
			ctx, cancel, needContinue, err := c.initCommandContext(cmd, args)
			if !needContinue || err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			defer cancel()

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
	return c.resource.FullName()
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
	if labels := extractLabels(resource); labels != nil {
		// ID
		if prefix == "" || strings.HasPrefix(labels.Id, prefix) {
			results = append(results, labels.Id)
		}

		// Name
		if prefix == "" || strings.HasPrefix(labels.Name, prefix) {
			results = append(results, labels.Name)
		}

		// Tags
		for _, tag := range labels.Tags {
			if prefix == "" || strings.HasPrefix(tag, prefix) {
				results = append(results, tag)
			}
		}
	}
	return results
}

func (c *Command) initCommandContext(cmd *cobra.Command, args []string) (cli.Context, func(), bool, error) {
	ctx, cancel, err := cli.NewCLIContext(&cli.ContextParameter{
		PlatformName:       c.resource.PlatformName,
		ResourceName:       c.resource.Name,
		CommandName:        c.Name,
		GlobalFlags:        root.Command.PersistentFlags(),
		Args:               args,
		ColumnDefs:         c.ColumnDefs,
		Parameter:          c.currentParameter,
		SkipLoadingProfile: c.resource.SkipLoadingProfile,
	})
	if err != nil {
		return nil, nil, false, err
	}

	c.printCommandWarning(ctx)

	// パラメータの補完処理(ポインタ型のクリアやコンテキストからのパラメータデフォルト値受け取りなど)
	c.completeParameterValue(cmd, ctx, c.currentParameter)

	// パラメータファイルの処理やスケルトン出力など
	needContinue, err := c.handleCommonParameters(ctx, cmd)
	if needContinue {
		needContinue, err = c.handleExampleParameters(ctx)
	}
	return ctx, cancel, needContinue, err
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

func (c *Command) handleCommonParameters(ctx cli.Context, cmd *cobra.Command) (bool, error) {
	if cp, ok := c.currentParameter.(cflag.CommonParameterValueHolder); ok {
		// パラメータスケルトンの生成
		if cp.GenerateSkeletonFlagValue() {
			return false, generateSkeleton(ctx, c.currentParameter)
		}
		// --parameters/--parameter-fileフラグの処理
		if err := c.loadParameters(ctx, cmd, cp); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (c *Command) handleExampleParameters(ctx cli.Context) (bool, error) {
	if cp, ok := c.currentParameter.(cflag.ExampleParameterValueHolder); ok {
		if cp.ExampleFlagValue() {
			if eh, ok := c.currentParameter.(ExampleHolder); ok {
				return false, generateExampleParameters(ctx, eh)
			}
			// パラメータがExampleHolder未実装の場合はスケルトン出力
			return false, generateSkeleton(ctx, c.currentParameter)
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
			if id, ok := c.extractMatchedResourceID(ctx, r, args); ok {
				results.Append(cli.ResourceContext{ID: id, Zone: zone, Resource: r})
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

func (c *Command) extractMatchedResourceID(ctx cli.Context, r interface{}, idOrNameOrTag []string) (string, bool) {
	nameCompareFunc := c.selectNameComparingFuncByMode(ctx.Option().ArgumentMatchModeValue())
	if labels := extractLabels(r); labels != nil {
		for _, cond := range idOrNameOrTag {
			// ID
			if labels.Id == cond {
				return labels.Id, true
			}

			// Name(--argument-match-modeにしたがって比較)
			if nameCompareFunc(labels.Name, cond) {
				return labels.Id, true
			}

			// Tags
			for _, tag := range labels.Tags {
				if tag == cond {
					return labels.Id, true
				}
			}
		}
	}
	return "", false
}

func (c *Command) selectNameComparingFuncByMode(mode string) func(string, string) bool {
	switch mode {
	case "exact":
		return func(source, pattern string) bool {
			return source == pattern
		}
	default:
		return strings.Contains
	}
}

func (c *Command) exec(ctx cli.Context, ids cli.ResourceContexts) (output.Contents, error) {
	if c.Func == nil {
		// use default func
		fn, ok := services.DefaultServiceFunc(c.resource.PlatformName, c.resourceName(), c.Name)
		if !ok {
			return nil, fmt.Errorf("default service func not found: platform:%s resource:%s command:%s",
				c.resource.PlatformName, c.resourceName(), c.Name)
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
		}(ctx.WithResource(rc.ID, rc.Zone, rc.Resource))
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

func (c *Command) ParameterCategoryBy(key string) *category.Category {
	for _, cat := range category.ParameterCategories {
		if cat.Key == key {
			return cat
		}
	}

	if key == "" {
		key = c.resource.Name
		return &category.Category{
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
	return &category.Category{
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
		fn, ok := services.DefaultListAllFunc(c.resource.PlatformName, c.resourceName(), c.Name)
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
