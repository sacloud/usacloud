# コマンド実行ライフサイクル

本ドキュメントでは、`pkg/core.Command` が Cobra のコマンドとして呼び出されてから、サービス関数の実行・結果の出力に至るまでのライフサイクルを段階的に説明します。

## 全体の流れ

```
Cobra.Run()
  │
  ├─ 1. CLIContext の構築 (initCommandContext)
  │     ├─ 設定ファイル/プロファイルの読み込み
  │     ├─ API クライアントの初期化
  │     ├─ 出力 Writer の構築
  │     ├─ パラメータの補完（ゾーンなど）
  │     ├─ --generate-skeleton / --example の処理
  │     └─ --parameters / --parameter-file の読み込み
  │
  ├─ 2. パラメータのバリデーション (validateParameter)
  │
  ├─ 3. パラメータのカスタマイズ (ParameterCustomizer.Customize)
  │
  ├─ 4. 対象リソースの解決 (expandResourceContextsFromArgs)
  │     ├─ SelectorType に応じた引数チェック
  │     ├─ 全ゾーン or 指定ゾーンでのリソース一覧取得
  │     └─ ID / Name / Tags によるマッチング
  │
  ├─ 5. 確認ダイアログ (confirmContinue)
  │
  ├─ 6. コマンド実行 (exec / execParallel)
  │     ├─ サービス関数の解決
  │     ├─ プログレス表示のラップ
  │     └─ ゾーン/リソースごとの並列実行
  │
  └─ 7. 結果の出力 (ctx.Output().Print)
```

## 1. CLIContext の構築

`Command.CLICommand()` が返す `cobra.Command` の `Run` から、`Command.initCommandContext()` が呼び出されます。

主な処理:

- `cli.NewCLIContext()` を呼び、グローバルフラグ/プロファイル/API クライアント/入出力を初期化する。
- リソースが global でない場合、`completeParameterValue()` で `zone` パラメータにデフォルト値を設定する。
- `handleCommonParameters()` で `--generate-skeleton` / `--parameters` / `--parameter-file` を処理する。
- `handleExampleParameters()` で `--example` を処理する。
- スケルトン/例出力時は `needContinue=false` を返し、以降の処理をスキップする。

タイムアウトは設定ファイルの `ProcessTimeout` に基づき、`context.WithTimeout()` で管理されます。

## 2. パラメータのバリデーション

`Command.validateParameter()` が実行されます。

- `Command.ValidateFunc`が明示的に設定されていればそれを使用する。
- 未設定の場合は`pkg/validate.Exec`を呼び出し、structタグ（`validate:"required"` など）に基づいて検証する。
- バリデーションエラー時はエラーメッセージを表示して終了する。

## 3. パラメータのカスタマイズ

パラメータが `ParameterCustomizer` インターフェースを実装していれば、`Customize(ctx)` が呼び出されます。

これは、タグベースのバリデーションやフラグ設定だけでは表現できない、コマンド固有の前処理をするための拡張ポイントです。例: 複数フィールド間の依存関係の調整、デフォルト値の動的決定など。

## 4. 対象リソースの解決

`Command.expandResourceContextsFromArgs()` が引数から操作対象のリソースを解決します。

### SelectorType

`SelectorType` は引数の扱いを決めます。

| 値 | 意味 | 例 |
| --- | --- | --- |
| `SelectorTypeNone` | 引数を受け取らない | `list`, `create` |
| `SelectorTypeRequireSingle` | ID/Name/Tags で 1 件を指定 | `read`, `update`, `delete` |
| `SelectorTypeRequireMulti` | ID/Name/Tags で複数件を指定 | `boot`, `shutdown` |

### ゾーンの扱い

- グローバルリソース（`Resource.IsGlobalResource == true`）はゾーンを持たない。
- 通常のリソースは`--zone`または設定ファイルのデフォルトゾーンが適用される。
- `--zone=all` が指定された場合、全ゾーンを対象にする。

### リソースのマッチング

`Command.collectResources()` は `ListAllFunc`（未設定時は `pkg/services/registry` から解決）を用いて対象ゾーンのリソース一覧を取得します。取得後、`extractMatchedResourceID()` で引数と ID/Name/Tags を比較し、一致するリソースを `cli.ResourceContexts` として返します。

Name の比較は `--argument-match-mode` に従います。デフォルトは部分一致（`strings.Contains`）、`exact` を指定すると完全一致となります。

## 5. 確認ダイアログ

パラメータが `cflag.ConfirmParameterValueHandler` を実装しており、かつ `--assumeyes(-y)` が指定されていない場合、`confirmContinue()` が確認ダイアログを表示します。

ターミナル以外ではダイアログを表示できないため、`--assumeyes` の指定を求めるエラーとなります。

## 6. コマンド実行

### サービス関数の解決

`Command.exec()` は以下の順で実行関数を決定します。

1. `Command.Func`が直接設定されていればそれを使用する。
2. 未設定の場合、`pkg/services.DefaultServiceFunc()`からプラットフォーム名・リソース名・コマンド名で検索する。

`Func` は `func(ctx cli.Context, parameter interface{}) ([]interface{}, error)` のシグネチャを持ちます。

### プログレス表示

`Command.NoProgress == false` の場合、`Func` を `Progress` でラップし、実行中に進捗メッセージを表示します。

### 並列実行

`Command.execParallel()` は `cli.ResourceContexts` の各要素に対して別 goroutine で `Func` を実行します。

- 各goroutineは`ctx.WithResource(id, zone, resource)`でコンテキストを複製する。
- パラメータも`cloneCurrentParameter()`で複製し、対象のID/ゾーンを設定する。
- 全goroutineの結果を収集し、ゾーン順にソートして返す。
- 一部でエラーが発生しても他の処理は継続し、最後にエラーをまとめて返す。

## 7. 結果の出力

`Command.Run()` の最後に `ctx.Output().Print(results)` を呼び出し、結果を出力します。

出力形式は `cli.Context` 構築時に決定されます。

- `--quiet` → IDのみ出力である。
- `--format` → Goテンプレート出力である。
- `--query` → JSONPath/JQクエリ付きJSON出力である。
- `--output-type json/yaml` → JSON/YAML出力である。
- それ以外 → テーブル出力（`ColumnDefs`に基づく）である。

## タイムアウトとキャンセル

`CLICommand()` の `Run` では、コマンド実行を別 goroutine で行い、`ctx.Done()` または完了を待ちます。

- `ctx.Done()`が先に発火した場合、「command[resource/command]timed out」エラーとして処理する。
- エラー発生時は赤文字で標準エラーに出力し、`os.Exit(1)`で終了する。

## 補完

`CLICommand()` は `ValidArgsFunction` を設定し、シェル補完を提供します。

- `CustomCompletionFunc`が設定されていればそれを優先する。
- それ以外は`collectResources()`で取得したリソースのID/Name/Tagsをもとに前方一致で補完候補を返す。

## まとめ

`pkg/core.Command` のライフサイクルは「コンテキスト構築 → バリデーション → 対象解決 → 実行 → 出力」という一連の流れを固定しつつ、各所でインターフェースによる拡張ポイントを提供しています。これにより、基本的な CRUD コマンドは宣言的に定義するだけで実現でき、特殊な処理が必要なコマンドのみが `Func` や `ParameterCustomizer` を実装すればよい構成となっています。
