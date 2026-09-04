# Usacloud エージェント向け指示

## プロジェクト概要

`usacloud` はさくらのクラウド公式 CLI です。`main.go` は `pkg.Run()` に処理を委譲し、そこで Cobra コマンドを登録してルートコマンドを実行します。

- `pkg/commands/root` は永続的な設定フラグと Cobra のルートコマンドを定義する。
- `pkg/resources.go` はトップレベルのリソースを登録する。
- IaaS リソースは `pkg/commands/iaas/resources.go` にまとめられている。
- これらは `usacloud iaas` の配下に追加され、後方互換性のため、非表示のルートレベルコマンドとしても公開される。
- `pkg/commands/<platform>/<resource>` 配下の各リソースパッケージは `core.Resource` を定義し、`init()` から `core.Command` の値を登録する。パラメータ構造体は、共有の `cflag` 型とリソース固有のフィールドを組み合わせる。
- `pkg/core` は共通のコマンドライフサイクルを制御する。
- 設定とプロファイルの読み込み、クライアントと出力の初期化、パラメータの検証をする。
- また、ゾーンをまたぐセレクターの展開、破壊的操作の確認、コマンドの呼び出し、結果の出力も行う。
- 独自の `core.Command.Func` を持たないコマンドは、`pkg/services/<platform>/*_services_gen.go` に生成されたアダプターを使用し、対応するサービスライブラリを呼び出す。コマンドパッケージ内の生成された `zz_*_gen.go` ファイルは、パラメータ構造体のタグから Cobra フラグを作成する。
- `tools/gen-commands` は登録済みのリソースとコマンドを読み取り、これらのアダプターとフラグファイルを生成する。ファイル名が `_gen.go` で終わるファイルや `Code generated` と記載されたファイルは手動で編集せず、リソースやコマンドの定義、またはジェネレーターテンプレートを変更してから再生成する。

## コマンド

Go `1.25.8`（`go.mod` で指定）を使用します。

| 目的 | コマンド |
| --- | --- |
| 開発ツールのインストール | `make tools` |
| CLI のビルド | `make build` |
| コマンド／サービスアダプター、ライセンス、フォーマットの再生成 | `make gen` |
| すべてのユニットテストの実行 | `make test` |
| 1 パッケージのテストの実行 | `go test ./pkg/config -run '^TestFillDefaults_ZonesEmpty$' -v` |
| コマンドパッケージ内の 1 テストケースの実行 | `go test ./pkg/commands/iaas/disk -run '^TestCreate_ConvertToServiceRequest$' -v` |
| End-to-End テストの実行 | `make e2e-test` |
| 1 件の End-to-End テストの実行 | `make e2e-test TESTARGS='-run ^TestE2E_minimum$'` |
| Go の lint とフォーマット | `make lint-go` |
| テキストの lint | `make lint-text` |
| GitHub Actions ワークフローの lint | `make lint-action` |

`make test` は race detector を有効にして `go test ./...` を実行します。End-to-End テストには `e2e` ビルドタグがあり、最初にローカルの CLI をインストールします。完了まで最大 240 分かかる場合があります。

## コマンド実装の規約

- CLI フラグとサービスリクエストへのマッピングは、パラメータ構造体で宣言的に定義する。埋め込まれた共有パラメータには通常 `cli:",squash"` と `mapconv:",squash"` の両方を使用し、サービスリクエストへ渡してはならないフィールドには `mapconv:"-"` を使用する。
- 標準的な検証には `validate` タグを使用する。フラグに `options=...`、フィルタリング、キー／値変換が必要な場合は、`pkg/vdef/definitions.go` に名前付きの値マッピングを追加する。複数のフィールドに依存する検証は、コマンド固有のバリデーターに記述する。
- API キー、パスワード、トークンなど、コマンド固有の認証情報やシークレットを通常のフラグ値として受け取らない。シェル履歴やプロセス一覧への露出を避けるため、権限を制限したファイル、標準入力、または公式にサポートされた環境変数から読み取る。シークレットは設定、ログ、エラー、コマンド出力へ含めず、パラメータ構造体からリクエストへ自動変換されないようにする。
- リソース ID、名前の一部、またはタグを引数として受け取るコマンドには、`SelectorTypeRequireSingle` または `SelectorTypeRequireMulti` を設定する。
- リソースの検索とゾーンごとの実行は core が処理する。セレクターのロジックを重複して実装せず、`--zone=all` を考慮して `ctx.WithResource` を使用する。
- コマンドの `Category` と `Order` は共有カテゴリーと整合させる。`Resource.AddCommand` は未知のカテゴリーを拒否し、これらの値を使ってヘルプ出力を並べ替える。
- 新しい IaaS リソースは `pkg/commands/iaas/resources.go` に、IaaS 以外の新しいトップレベルリソースは `pkg/resources.go` に追加する。ジェネレーターがアダプターを作成できるよう、`ServiceType` には対応するサービスライブラリの型を設定する。
- `sacloud-sdk-go/api/<service>` を直接ラップする `*-api` コマンドを追加するときは、`.agents/skills/generate-usacloud-api-command/` の skill を使用する。
- `<service>-api` コマンドは `.agents/skills/generate-usacloud-api-command/SKILL.md` に従い、各コマンドを `sacloud-sdk-go/api/<service>` の公開 API メソッドと 1:1 で対応させる。
    - 1 コマンドから別の API メソッドを暗黙に追加実行したり、複数メソッドを組み合わせたワークフローにしたりしてはならない。
    - 複数メソッドを組み合わせたワークフローが必要な場合は `<service>` という名前の高レベルサブコマンドを別途実装する。
- サービスメソッドがコマンドと一致する場合は、生成されるデフォルトのサービス呼び出しを優先する。名前が一致しない場合は `ServiceFuncAltName` を指定し、サービスアダプターで表現できない動作に限って独自の `Func` を使用する。
- リソースを変更するコマンドには `cflag.ConfirmParameter` を埋め込む。これにより、非対話的な呼び出しでは `--assumeyes`／`-y` が必須になる。list コマンドでは慣例として `NoProgress: true` を設定し、`ColumnDefs` を明示的に定義する。

## リポジトリの規約

- すべてのコミットで DCO への同意が必須である。`git commit --signoff` を使用し、コミットメッセージに `Signed-off-by` トレーラーを追加する。
- 新しい Go ファイルでは、Apache ライセンスヘッダーとリポジトリの著作権年のパターンを維持する。`make gen` は `set-license`、`gofmt`、`gosimports` を実行する。
- `golangci-lint` は意図的に生成ファイルを除外し、`--fix` を付けて実行される。コミット前に、適用されたフォーマット変更を確認する。
- ユニットテストには標準の `testing` パッケージと `testify/require` を使用する。End-to-End テストは `e2e` ビルドタグ付きで `e2e/` 配下に分離されている。
