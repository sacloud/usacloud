# API command 実装支援スクリプト

このディレクトリには、`sacloud-sdk-go/api/<service>` の API interface を
usacloud の `*-api` コマンドとして実装するときに使う検査ツールがあります。
ソースコードを生成・変更するツールではありません。

## 全体の流れ

```text
SDK interface
    |
    | sdk-interface-inventory
    v
一時 manifest
    |
    | resource 名、command 名、excluded をレビュー
    v
usacloud コマンドを実装
    |
    | validate-implementation-coverage
    v
API メソッドの実装漏れを検査
    |
    | sakumock を使った unit test
    v
引数変換と実際の SDK 呼び出しを検査
```

manifest は実装中とレビュー時に使う一時成果物です。リポジトリにはコミットせず、
session artifact 用のディレクトリなどに保存してください。

## `sdk-interface-inventory`

指定した Go package の公開 API interface を `go list`、`go/parser`、`go/types` で
解析し、schema v2 の JSON manifest を出力します。

```sh
go run ./.agents/skills/generate-usacloud-api-command/scripts/sdk-interface-inventory \
  --package github.com/sacloud/sacloud-sdk-go/api/<service> \
  --interface <ServiceAPI> \
  --output "$SESSION_ARTIFACT_DIR/<service>-manifest.json"
```

`--interface` を省略すると、package 内の公開 `*API` interface をすべて対象にします。
`--output -` を指定すると標準出力へ書き出します。

manifest には、interface ごとに次の情報が記録されます。

- interface 名と、そこから推定した単数形の kebab-case resource 名
- 各メソッドの名前、引数型、戻り値型
- メソッド名から暫定分類した operation
- operation とメソッド名から推定した command 名

operation は `List`/`Find`、`Read`/`Get`、`Create`、`Update`、`Delete` などの
接頭辞から機械的に分類され、それ以外は `action` になります。この分類、resource 名、
command 名は設計上の確定値ではないため、生成後に必ずレビューしてください。

公開しないメソッドも `methods` から削除せず、対応する interface の `excluded` に
メソッド名と理由を追記します。manifest の形は
[`../assets/coverage-manifest.json.tmpl`](../assets/coverage-manifest.json.tmpl)も
参照してください。

`testdata/sdkfixture/api.go` は解析対象となる模擬 SDK package、
`testdata/interfaces.golden.json` は期待する出力を固定した golden file です。

## `validate-implementation-coverage`

レビュー済みmanifestと、実装したcommand packageを照合します。

```sh
go run ./.agents/skills/generate-usacloud-api-command/scripts/validate-implementation-coverage \
  --manifest "$SESSION_ARTIFACT_DIR/<service>-manifest.json" \
  --command-dir ./pkg/commands/<platform>/<service>-api \
  --strict
```

`--command-dir` 以下のpackageを再帰的に型解析し、除外されていない各SDKメソッドに
ついて次を検査します。

- manifest の `resource` と一致する `core.Resource{Name: ...}` がある
- manifest の `command` と一致する `core.Command{Name: ...}` がある
- その `core.Command` に明示的な非 `nil` の `Func` がある
- `excluded` の各項目が実在するメソッドを指し、空でない理由を持つ
- `--strict` 指定時、各メソッドのcommand名が空でない

resource identity はディレクトリ名ではなく、同じpackageにある
`core.Resource{Name: ...}` から取得します。

`testdata/commands/` は、実装済みcommand、`Func` のないcommand、複数resourceなどを
再現するfixtureです。`testdata/manifest.json` はfixtureと照合する入力例です。

## 保証しないこと

coverage validator が保証するのは、manifest と
`core.Resource` / `core.Command` の構造上の対応関係までです。次の内容は検査しません。

- `Func` がmanifestに対応するSDKメソッドを実際に呼んでいるか
- CLI flag が正しいSDK request fieldへ変換されるか
- nullable、clear、discriminatorなどの意味が保たれているか
- selector、mutation確認、出力列が適切か

これらは、SDK operationをsakumockで差し替えたcommand packageのunit testで検証します。
詳細は[`../references/testing.md`](../references/testing.md)を参照してください。

## スクリプト自体のテスト

このディレクトリはhidden directory配下にあるため、通常の`go test ./...`では
探索されません。次のようにpackageを明示して実行します。

```sh
go test \
  ./.agents/skills/generate-usacloud-api-command/scripts/sdk-interface-inventory \
  ./.agents/skills/generate-usacloud-api-command/scripts/validate-implementation-coverage
```
