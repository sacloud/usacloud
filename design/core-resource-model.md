# リソースモデル

本ドキュメントでは、`pkg/core` が提供するリソース・コマンド・カテゴリのモデルと、それらの関係を説明します。

## Resource

`core.Resource` は usacloud CLI における 1 つのリソース（コマンド群の namespace）を表します。

```go
type Resource struct {
    Name               string
    Aliases            []string
    Usage              string
    DefaultCommandName string
    Category           category.Category
    Warning            string
    IsGlobalResource   bool
    PlatformName       string
    ServiceType        reflect.Type
    SkipLoadingProfile bool

    resources []*Resource // 子リソース
    parent    *Resource

    categorizedCommands []*CategorizedCommands
}
```

### 主要フィールド

| フィールド | 説明 |
| --- | --- |
| `Name` | リソース名。ケバブケース（例: `simple-monitor`）で指定します。 |
| `Aliases` | リソースのエイリアス（例: `simplemonitor`）。 |
| `Usage` | ヘルプに表示される短い説明文。 |
| `DefaultCommandName` | リソース名だけが指定された際に自動実行するコマンド名（例: `list`）。空の場合はヘルプを表示します。 |
| `Category` | リソースが属するカテゴリ。ヘルプ表示のグループ分けに使用します。 |
| `Warning` | コマンド実行前に表示する警告メッセージ。 |
| `IsGlobalResource` | true の場合、ゾーンを持たないリソースとして扱います。 |
| `PlatformName` | `"iaas"`, `"phy"`, `"objectstorage"`, `"webaccel"` のいずれか。空の場合は `"iaas"` として扱われます。 |
| `ServiceType` | 対応する service（`iaas-service-go` / `webaccel-api-go` 等）の型情報。コード生成に使用します。 |
| `SkipLoadingProfile` | true の場合、プロファイルの読み込みをスキップします。 |

### リソース階層

`Resource` は `parent` と `resources` により親子関係を持てます。

- `AddChild(resource *Resource)`で子リソースを追加する。
- 子リソースは親の`IsGlobalResource`と`PlatformName`を継承する。
- 子リソース名は`parent.FullName() + "_" + Name`として扱われる。

例: `server` リソース配下の `interface` 子リソースは、`server_interface` としてサービスレジストリやコマンド名で参照されます。

### コマンドの登録

`Resource.AddCommand(command *Command)` でコマンドを登録します。

- コマンドの`Category`は`Resource`に存在するカテゴリキーのいずれかでなければならない。
- 同一カテゴリ内で同名のコマンドを登録しようとするとfatalとなる。
- 登録時にコマンドの`resource`フィールドが紐づけられる。
- 各カテゴリ内では`Order`の昇順でソートされる。

### Cobra コマンドの構築

`Resource.CLICommand()` は `*cobra.Command` を返します。

- 自身のNameをUseに持つコマンドを生成する。
- 登録済みの`Command`をサブコマンドとして追加する。
- 子リソースも再帰的にサブコマンドとして追加する。
- `DefaultCommandName`が設定されている場合、引数なしで該当コマンドを実行する。
- カテゴリ別のヘルプ表示を構築する。

## Command

`core.Command` は 1 つのサブコマンドを表します。実行時にはコンテキストとしても機能します。

```go
type Command struct {
    Name      string
    Aliases   []string
    Usage     string
    ArgsUsage string

    Category string
    Order    int

    SelectorType   SelectorType
    NoProgress     bool
    ConfirmMessage string

    ParameterCategories  []category.Category
    ParameterInitializer func() interface{}
    ServiceFuncAltName   string

    ColumnDefs []output.ColumnDef

    ExperimentWarning string

    ListAllFunc          func(ctx cli.Context, parameter interface{}) ([]interface{}, error)
    CustomCompletionFunc func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)
    ValidateFunc         ValidateFunc
    Func                 func(ctx cli.Context, parameter interface{}) ([]interface{}, error)

    resource         *Resource
    currentParameter interface{}
}
```

### 主要フィールド

| フィールド | 説明 |
| --- | --- |
| `Name` | コマンド名。ケバブケース（例: `list`, `monitor-response-time`）。 |
| `Aliases` | コマンドのエイリアス（例: `ls`, `find`, `select`）。 |
| `Usage` | ヘルプに表示される短い説明文。 |
| `ArgsUsage` | 引数の説明。未設定時は `SelectorType` の値に応じた内容が自動設定されます。 |
| `Category` | コマンドが属するカテゴリのキー。basic/operation/power/monitor/other など、pkg/category/commands.go で定義される値を使用します。 |
| `Order` | 同カテゴリ内での並び順。小さいほど先に表示されます。 |
| `SelectorType` | 引数として受け取るリソース選択子の種別。 |
| `NoProgress` | true の場合、コマンド実行中のプログレス表示はしません。 |
| `ConfirmMessage` | 確認ダイアログで表示するメッセージ。未設定時は `Name` が使用されます。 |
| `ParameterCategories` | コマンド固有のパラメータカテゴリ。 |
| `ParameterInitializer` | コマンドパラメータを初期化する関数。必須。 |
| `ServiceFuncAltName` | 自動生成されるサービス関数呼び出しで、コマンド名以外のメソッド名を使う場合に指定します。空の場合は `Name` を Camelize したものが利用されます。 |
| `ColumnDefs` | テーブル形式での出力対象列。省略した場合は ID と Name が出力されます。 |
| `ExperimentWarning` | 実験的機能として実行前に表示する警告メッセージ。 |
| `ListAllFunc` | 操作対象リソースの一覧取得用関数。通常は `Resource` に紐づけられた自動生成関数を利用します。特殊な一覧取得が必要な場合に設定します。 |
| `CustomCompletionFunc` | 特殊な引数補完が必要な場合に設定する関数。未設定の場合は `ListAllFunc` から取得したリソースの ID/Name/Tags で補完されます。 |
| `ValidateFunc` | カスタムバリデーション用。空の場合は `usacloud/pkg/validate.Exec` が実行されます。 |
| `Func` | コマンドの実処理。設定してない場合はデフォルトの service（`iaas-service-go` / `webaccel-api-go` 等）呼び出しが行われます。 |

### パラメータ struct

`ParameterInitializer` が返す struct は通常以下を含みます。

- `cflag.CommonParameter` ... `--parameters`, `--generate-skeleton`, `--example`である。
- `cflag.OutputParameter` ... `--output-type`, `--quiet`, `--format`, `--query`である。
- `cflag.ConfirmParameter` ... `--assumeyes`（変更系コマンドで必要）である。
- `cflag.ZoneParameter` ... `--zone`（ゾーン指定が必要なコマンドで必要）である。
- リソース/コマンド固有のフィールドである。

これらの struct は `cli` タグによりフラグ名/カテゴリ/エイリアス等が制御され、`mapconv` タグにより service（`iaas-service-go` / `webaccel-api-go` 等）のリクエスト型への変換が制御されます。

## Category

`category.Category` はリソース/コマンド/パラメータをグループ化するための識別子です。

```go
type Category struct {
    Key         string
    DisplayName string
    Order       int
}
```

### コマンドカテゴリ

`pkg/category/commands.go` で定義されます。各リソースはその中から使用するカテゴリを選び、`Resource` 定義時に列挙します。

主要な例:

- `basic` ... 一覧/参照/作成/更新/削除などの基本CRUDである。
- `operation` ... 電源操作やリソース固有の操作である。
- `power` ... 電源系操作である。
- `monitor` ... 監視系操作である。
- `other` ... その他である。

### リソースカテゴリ

`pkg/commands/<platform>/category` 配下でプラットフォーム固有に定義されます。ヘルプ表示時にリソースをグループ化するために使用されます。

### パラメータカテゴリ

`pkg/category/parameters.go` で定義されます。`cli` タグの `category` 属性に対応し、ヘルプのフラグ表示をグループ化します。

## Resources

`core.Resources` は `[]*Resource` のエイリアスで、プラットフォーム全体のリソース登録に使用されます。

```go
type Resources []*Resource
```

`CategorizedResources(categories []category.Category)` メソッドにより、指定されたカテゴリ順にリソースをグループ化した `[]*CategorizedResources` を取得できます。

## CategorizedResources / CategorizedCommands

カテゴリと、それに紐づくリソース/コマンドのリストを保持する構造体です。

```go
type CategorizedResources struct {
    Category  category.Category
    Resources []*Resource
}

type CategorizedCommands struct {
    Category category.Category
    Commands []*Command
}
```

これらは `usage.go` でヘルプテンプレートを構築する際に使用されます。

## Labels

`core.Labels` は usacloud でリソースを識別するための最小限の情報です。

```go
type Labels struct {
    Id   string
    Name string
    Tags []string
}
```

リソース解決やシェル補完では、`LabelsExtractors` に登録された extractor 関数を使い、サービスから返された値を `Labels` に変換して抽出します。

## まとめ

`pkg/core` のリソースモデルは、リソース・コマンド・カテゴリという 3 つの軸で CLI を構成します。

- `Resource`がリソースのnamespaceとコマンド群の集約を担う。
- `Command`が個別の操作とそのパラメータ、実行ロジックを担う。
- `Category`がリソース/コマンド/パラメータの表示グループ化を担う。

これらを宣言的に組み合わせることで、usacloud は大量のリソース・コマンドを統一的かつ拡張可能に管理しています。
