# usacloud v1 でのコマンド実装方法

本ドキュメントでは、usacloud に新しいリソース・コマンドを追加する手順を説明します。

`pkg/core` フレームワークの概要やコマンド実行ライフサイクルについては、以下も併せて参照してください。

- [core-architecture.md](./core-architecture.md)
- [core-command-lifecycle.md](./core-command-lifecycle.md)
- [core-resource-model.md](./core-resource-model.md)
- [code-generation-boundary.md](./code-generation-boundary.md)

## 全体の流れ

1. `pkg/commands/<platform>/<リソース名>`配下にパッケージを作成する。
2. パッケージ内にリソース定義`resource.go`を作成する。
3. パッケージ内にコマンド定義（例: `list.go`, `create.go`）を作成する。
4. 必要に応じてカラム定義`columns.go`やラベル抽出`labels.go`を作成する。
5. `pkg/commands/<platform>/resources.go`に作成したリソース定義を追加する。
6. `make`または`make gen`を実行する。

## 各手順について

### 1. `pkg/commands/<platform>/<リソース名>` 配下にパッケージを作成

追加したいリソースが `simple-monitor` で、プラットフォームが `iaas` の場合、`pkg/commands/iaas/simplemonitor` ディレクトリを作成します。

#### 命名規則

- リソース名を小文字にしたものである。
- 単語の区切りはなし。ハイフンやアンダーバーを用いない。
- 以下2つはGoの予約語とぶつからないように特殊なルールを適用する。
  - NIC（インターフェース）リソース → `iface`である。
  - スイッチリソース → `swytch`である。

### 2. リソース定義 `resource.go` を作成

以下のように実装します。詳細は `core.Resource` や既存リソースの実装を参照してください。

```go
package simplemonitor

import (
    "reflect"

    "github.com/sacloud/iaas-service-go/simplemonitor"
    "github.com/sacloud/usacloud/pkg/commands/iaas/category"
    "github.com/sacloud/usacloud/pkg/core"
)

var Resource = &core.Resource{
    PlatformName:     "iaas",
    Name:             "simple-monitor",
    Aliases:          []string{"simplemonitor"},
    ServiceType:      reflect.TypeOf(&simplemonitor.Service{}),
    Category:         category.ResourceCategoryCommonServiceItem,
    IsGlobalResource: true,
}
```

#### 主要フィールドの選び方

| フィールド | 設定指針 |
| --- | --- |
| `PlatformName` | 対象プラットフォーム。現状は `"iaas"` または `"webaccel"` が使用される。`"phy"` / `"objectstorage"` は将来拡張用に予約されている。 |
| `Name` | ケバブケースのリソース名。コマンドパスにそのまま使用されます。 |
| `Aliases` | リソースの別名。後方互換性維持などに使用します。 |
| `ServiceType` | 対応する service（`iaas-service-go` / `webaccel-api-go` 等）の型。コード生成でメソッドを検索する際に使用します。 |
| `Category` | リソースカテゴリ。ヘルプ表示のグループ分けに使用します。 |
| `IsGlobalResource` | ゾーンを持たないリソースの場合は true。 |
| `DefaultCommandName` | リソース名だけで実行したいデフォルトコマンドがあれば設定（例: `"list"`）。 |

### 3. コマンド定義を作成

各コマンドは `pkg/commands/<platform>/<リソース名>/<コマンド名>.go` に定義します。

#### 一覧系コマンドの例

```go
var listCommand = &core.Command{
    Name:               "list",
    Aliases:            []string{"ls", "find", "select"},
    Category:           "basic",
    Order:              10,
    ServiceFuncAltName: "Find",
    NoProgress:         true,

    ColumnDefs: defaultColumnDefs,

    ParameterInitializer: func() interface{} {
        return newListParameter()
    },
}

type listParameter struct {
    cflag.CommonParameter      `cli:",squash" mapconv:"-"`
    cflag.OutputParameter      `cli:",squash" mapconv:"-"`
    cflag.LimitOffsetParameter `cli:",squash" mapconv:",squash"`

    cflag.FilterByNamesParameter `cli:",squash" mapconv:",omitempty,squash"`
    cflag.FilterByTagsParameter  `cli:",squash" mapconv:",omitempty,squash"`
}

func newListParameter() *listParameter {
    return &listParameter{}
}

func init() {
    Resource.AddCommand(listCommand)
}
```

#### 作成系コマンドの例

```go
var createCommand = &core.Command{
    Name:     "create",
    Category: "basic",
    Order:    20,

    ColumnDefs: defaultColumnDefs,

    ParameterInitializer: func() interface{} {
        return newCreateParameter()
    },
}

type createParameter struct {
    cflag.CommonParameter  `cli:",squash" mapconv:"-"`
    cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
    cflag.OutputParameter  `cli:",squash" mapconv:"-"`

    Target                string `validate:"required"`
    cflag.DescParameter   `cli:",squash" mapconv:",squash"`
    cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
    cflag.IconIDParameter `cli:",squash" mapconv:",squash"`
}

func newCreateParameter() *createParameter {
    return &createParameter{}
}

func init() {
    Resource.AddCommand(createCommand)
}
```

### 4. カラム定義の追加

#### カラム定義 `columns.go`

テーブル出力時の列を定義します。

```go
var defaultColumnDefs = []output.ColumnDef{
    {Name: "Zone"},
    {Name: "ID"},
    {Name: "Name"},
    {Name: "Scope", Template: "{{ scope_to_key .Scope }}"},
}
```

#### ラベル抽出（プラットフォーム単位の汎用実装）

ID/Name/Tags による引数解決やシェル補完は、`pkg/commands/<platform>/labels.go` でプラットフォーム単位に汎用実装されています。例えば `iaas` リソースでは `accessor.ID` / `accessor.Name` / `accessor.Tags` インターフェースを使った extractor が、既に `pkg/commands/iaas/labels.go` に登録されています。新しい `iaas` リソースを追加しても、リソースごとに専用の `labels.go` を作成する必要はありません。

### 5. リソースをプラットフォームのリソース一覧に追加

`pkg/commands/<platform>/resources.go` にリソースを追加します。

```go
var Resources = core.Resources{
    // ...
    simplemonitor.Resource,
    // ...
}
```

### 6. コード生成

```sh
make
# または
make gen
```

コード生成により以下が生成されます。

- `pkg/commands/<platform>/<リソース名>/zz_*_gen.go` ... 各コマンドのフラグ定義である。
- `pkg/services/<platform>/<リソース名>_services_gen.go` ... サービス関数のレジストリ登録である。
- `pkg/commands/<platform>/services_gen.go` ... サービス関数パッケージのimportである。

生成後、ビルド/テストを実行して問題がないか確認します。

```sh
make build
make test
```

生成したコードが原因でビルドできなくなったり `make` が失敗する場合は、`make clean-all` で生成コードを一旦全削除してから再生成すると解消することがあります。

## カスタマイズが必要なケース

### カスタムバリデーション

タグベースのバリデーションではカバーできない複雑な検証が必要な場合、`Command.ValidateFunc` を設定します。

```go
var createCommand = &core.Command{
    Name:         "create",
    ValidateFunc: validateCreateParameter,
    // ...
}

func validateCreateParameter(ctx cli.Context, parameter interface{}) error {
    p := parameter.(*createParameter)
    if p.Target == "" {
        return errors.New("target is required")
    }
    return validate.Exec(p)
}
```

### カスタム実行処理

自動生成されるサービス呼び出しでは実現できない処理が必要な場合、`Command.Func` を設定します。

```go
var customCommand = &core.Command{
    Name: "custom",
    Func: func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
        // カスタム処理
        return []interface{}{result}, nil
    },
    // ...
}
```

`Func` を設定すると、対応するサービス関数は生成されません。

### カスタムシェル補完

`Command.CustomCompletionFunc` を設定すると、標準の ID/Name/Tags 補完の代わりに独自の補完を提供できます。

## 値定義の追加

タグで `options=xxx` を指定する場合、`pkg/vdef/definitions.go` に対応する定義を追加する必要があります。

```go
var definitions = map[string][]*definition{
    "scope": {
        {key: types.Scopes.User.String(), value: types.Scopes.User},
        {key: types.Scopes.Shared.String(), value: types.Scopes.Shared},
    },
}
```

## まとめ

新しいコマンドを追加する主な作業は「リソース定義」「コマンド定義」「リソース登録」の 3 つです。フラグ定義やサービス呼び出しの大部分はコード生成に任せることができます。特殊な処理が必要な場合のみ、`ValidateFunc`/`Func`/`CustomCompletionFunc` などの拡張ポイントを活用してください。
