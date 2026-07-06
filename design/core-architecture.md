# pkg/core アーキテクチャ

本ドキュメントでは、usacloud のコマンド実装の中核となる `pkg/core` パッケージの責務、位置づけ、および主要コンポーネント間の関係を説明します。

## 概要

`pkg/core` は、usacloud における「リソース」と「コマンド」を宣言的に定義し、それを Cobra/Pflag ベースの CLI として実行可能にするためのフレームワーク層です。

開発者は `pkg/core.Resource` および `pkg/core.Command` を用いてリソースとコマンドを定義するだけで、以下を自動的に得られます。

- Cobraサブコマンドの生成である。
- コマンドラインオプション（フラグ）の生成である。
- ゾーン・ID指定に基づく対象リソースの解決である。
- サービス関数（`iaas-service-go` / `webaccel-api-go` 等）の呼び出しである。
- テーブル/JSON/YAML形式での出力である。
- 進捗表示、確認ダイアログ、パラメータスケルトン出力などの横断的機能である。

`pkg/core` 自体は各 API クライアントの詳細を知りません。API 呼び出しは `pkg/services` 側のレジストリに登録された関数（通常はコード生成される）を介して行われます。

## レイヤ構成

usacloud は以下のレイヤで構成されます。

```
┌────────────────────────────────────────────────────────┐
│ CLI (cobra/pflag)                                      │  ユーザーが直接操作するコマンドライン
├────────────────────────────────────────────────────────┤
│ pkg/core                                               │  コマンド/リソース定義、実行フレームワーク
│   - Resource                                           │
│   - Command                                            │
│   - 実行ライフサイクル                                  │
├────────────────────────────────────────────────────────┤
│ pkg/services + pkg/commands/<platform>                 │  生成されたサービス呼び出しコード、フラグ定義
├────────────────────────────────────────────────────────┤
│ サービスレイヤ (iaas-service-go / webaccel-api-go 等)   │  実際の API クライアント、ビジネスロジック
└────────────────────────────────────────────────────────┘
```

## pkg/core の主要コンポーネント

### `Resource`

`core.Resource` は 1 つのリソース（例: `server`, `disk`, `simple-monitor`）を表します。

- リソース名、エイリアス、カテゴリ、プラットフォーム名（`iaas`/`phy`/`objectstorage`/`webaccel`）を保持する。
- `Command`をカテゴリ別に集約する。
- 子リソース（例: `server`配下の`interface`）を持つことができる。
- `CLICommand()`でCobraコマンドツリーを構築する。

詳細は [core-resource-model.md](./core-resource-model.md) を参照してください。

### `Command`

`core.Command` は 1 つのサブコマンド（例: `list`, `create`, `read`）を表します。

- コマンド名、エイリアス、カテゴリ、実行順序を保持する。
- パラメータstructの初期化関数（`ParameterInitializer`）を持つ。
- セレクタ種別（`SelectorType`）により、引数をID/Name/Tagsで指定可能かを制御する。
- `Func`が未設定の場合、サービスレジストリから自動的に実行関数を解決する。

詳細は [core-command-lifecycle.md](./core-command-lifecycle.md) を参照してください。

### `Resources`

`core.Resources` は `[]*Resource` のエイリアスです。

- カテゴリ別にリソースをグループ化する`CategorizedResources()`を提供する。
- `pkg/commands/<platform>/resources.go`などで全リソースの登録に使用される。

### `CategorizedResources` / `CategorizedCommands`

カテゴリ（`category.Category`）とリソース/コマンドの組み合わせを表します。

- ヘルプ表示時に「Basic Commands」「Operation Commands」などの見出しでグループ化するために利用される。

### `FlagSet`

Cobra/Pflag のフラグをグループ化し、ヘルプ表示をカテゴリ別に整えるための構造体です。

### ユーティリティ関数/構造体

| ファイル | 責務 |
| --- | --- |
| `progress.go` | 長時間実行コマンドの進捗表示 |
| `parameter_loader.go` | `--parameters`/`--parameter-file` による JSON パラメータの読み込み |
| `skeleton.go` | `--generate-skeleton` によるパラメータ雛形の JSON 出力 |
| `example.go` | `--example` によるパラメータ例の JSON 出力 |
| `usage.go` | カテゴリ別のヘルプ/フラグ表示テンプレート構築 |
| `labels.go` | ID/Name/Tags によるリソース識別と補完 |
| `mapconv.go` | パラメータのディープコピー（ゾーン・ID 付与時） |
| `warning.go` | リソース/コマンド単位の警告表示 |

## pkg/core と他パッケージの関係

### `pkg/commands/<platform>/<resource>`

各リソースの `Resource`/`Command` 定義を手書きで記述します。`init()` で `Resource.AddCommand(...)` を呼び出すことで、コマンドをリソースに紐づけます。

### `pkg/commands/<platform>/resources.go`

プラットフォーム内の全リソースを `core.Resources` として登録します。

### `pkg/services/<platform>/<resource>_services_gen.go`

コード生成により、`core.Command.Func` から呼ばれるサービス関数を `pkg/services/registry` に登録します。通常は手で編集しません。

### `pkg/commands/<platform>/<resource>/zz_*_gen.go`

コード生成により、各コマンドのパラメータ struct に対する Cobra フラグ定義が生成されます。通常は手で編集しません。

### `pkg/cli`

CLI 実行コンテキスト（`cli.Context`）を提供します。API クライアント、入出力、ゾーン/ID 情報、タイムアウト制御などを担います。

### `pkg/cflag`

共通パラメータ（`--output-type`, `--zone`, `--assumeyes` など）を提供する embeddable な struct 群です。

### `pkg/category`

リソース/コマンド/パラメータのカテゴリ定義を提供します。

## コード生成との関係

`pkg/core` はコード生成の対象ではありません。コード生成されるのは以下です。

- `pkg/commands/<platform>/<resource>/zz_*_gen.go` ... 各コマンドのフラグ定義である。
- `pkg/services/<platform>/*_services_gen.go` ... サービス関数のレジストリ登録である。
- `pkg/commands/<platform>/services_gen.go` ... サービス関数パッケージのimportである。

これらは `pkg/core` で定義された `Resource`/`Command` スキーマを読み取り、`tools/gen-commands` で生成されます。詳細は [code-generation-boundary.md](./code-generation-boundary.md) を参照してください。

## プラットフォーム対応

`Resource.PlatformName` により以下のプラットフォームを区別します。

| PlatformName | 対象サービス | クライアント型 |
| --- | --- | --- |
| `"iaas"` | IaaS (`iaas-service-go`) | `iaas.APICaller` |
| `"webaccel"` | ウェブアクセラレータ (`webaccel-api-go`) | `*webaccel.Client` |

`"phy"`、`"objectstorage"` は将来の拡張用に予約されている値ですが、現時点では未実装（実行時に panic するプレースホルダ）です。

空文字は `registry.platform()` で `"iaas"` に正規化されますが、`api_client.client()` では汎用 HTTP クライアント（`commonClient`）に分岐します。現状、空文字を明示的に使用するリソースは存在しません。

## まとめ

`pkg/core` は usacloud のコマンド実行エンジンであり、リソース・コマンドの「宣言」とその「実行」を分離しています。開発者は宣言に集中し、機械的なボイラープレート（フラグ定義やサービス呼び出し）はコード生成に任せることで、保守性と拡張性を両立しています。
