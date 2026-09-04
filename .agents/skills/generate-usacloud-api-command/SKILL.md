---
name: generate-usacloud-api-command
description: sacloud-sdk-go/api の API インターフェースを直接ラップする usacloud コマンドを、安全に設計・実装・検証する。SDK メソッドの棚卸し、パラメータ変換、登録、sakumock テストが必要な場合に使用する。
---

# usacloud API コマンド生成

`sacloud-sdk-go/api/<service>` の公開 `*API` インターフェースを、usacloud の
`*-api` リソースとして直接ラップするための手順です。既存 service package を利用
できる場合は、まず生成済みアダプターを優先し、この skill は使いません。

## 成果物と開始条件

実装前に対象 SDK のバージョン、サービス package、対象インターフェース、公開したい
メソッドに加え、サービスが解決する問題、典型的な利用場面、resource の役割と関係、
作成順、制約を確定します。SDK のすべての interface を機械的に公開してはいけません。
認証・内部用途・低レベル transport 用の interface は対象外です。

SDK の棚卸しを JSON として保存し、実装中とレビュー時に同じ入力を使います。
manifest はコミット対象ではない一時成果物です。session artifact 用の作業場所へ保存し、
`.agents/skills/` やソース tree には保存しません。

```sh
go run ./.agents/skills/generate-usacloud-api-command/scripts/sdk-interface-inventory \
  --package github.com/sacloud/sacloud-sdk-go/api/<service> \
  --interface <ServiceAPI> \
  --output "$SESSION_ARTIFACT_DIR/<service>-manifest.json"

go run ./.agents/skills/generate-usacloud-api-command/scripts/validate-implementation-coverage \
  --manifest "$SESSION_ARTIFACT_DIR/<service>-manifest.json" \
  --command-dir ./pkg/commands/<platform>/<service>-api \
  --strict
```

`--interface` は 1 interface のみに絞る任意指定です。省略時も interface を平坦化せず、
`interface`、そこから導出した `resource`、各 method を JSON に保持します。`--strict`
はこの identity ごとに `core.Command` と明示的な `Func` があることを、子 resource
directory まで再帰的に検証します。未公開メソッドは対応する interface の `excluded` に
理由を記録します。スクリプトは source を生成・変更しません。

## 情報源の制約

サービスの用途、resource の意味・関係、利用例、制約について信頼してよい情報源は
`manual.sakura.ad.jp` と、対象バージョンの `sacloud-sdk-go` 内ドキュメントだけです。
過去の記憶、想像、一般的な製品知識、検索結果の要約、第三者資料から説明を作っては
いけません。検索は許可された一次資料を見つけるためだけに使い、help の各事実は
`references/help-and-service-model.md` の根拠表で一次資料の原文に対応させます。
根拠がなければ説明を補完せず、実装を停止します。

## 実装フロー

1. `manual.sakura.ad.jp` または対象バージョンの `sacloud-sdk-go` 内ドキュメントだけを
   調査し、`references/help-and-service-model.md` に従って根拠表と利用者向けの
   サービスモデルを文章化する。
2. `references/sdk-selection.md` に従って API interface を選び、inventory を生成する。
3. メソッドを read/list/mutation/action に分類し、CLI の語彙と selector を決める。
   response の ID 型も確認し、通常の table/JSON/YAML 出力に加えて `--quiet` が内部 struct
   や union の表現ではなく、再利用可能な scalar ID を出力できることを設計に含める。
4. `assets/command.go.tmpl` を出発点に、resource、parameter、各 command を明示的に
   実装する。resource help は `assets/resource.go.tmpl` を出発点にする。テンプレートを
   自動生成物として扱わない。
5. scalar と単純な slice は typed flag、入れ子・complex slice・任意 JSON は JSON flag
   に写像し、更新の clear 意味を別フラグで保持する。
6. `ctx.SDKClient()` から `sacloud-sdk-go/common/saclient.ClientAPI` を遅延取得して SDK の `NewClient` と
   `NewXOp` を呼び、`core.Command.Func` で SDK interface を呼び出す。
7. 親 resource の help にサービスの目的と全体フロー、子 resource の help に役割、
   他 resource との関係、先に必要なもの、代表的な利用場面を記載する。
8. resource を該当する resources 集合に登録し、`make gen` を実行する（生成対象が
   ある場合）。登録後に help の順序、aliases、説明の理解可能性を確認する。
9. `references/testing.md` の sakumock テストと出力テストを追加する。SDK response を
   `output.NewIDOutput` に渡す `--quiet` 相当のテストを含め、manifest coverage を検証する。

## 変更停止条件

次のいずれかなら実装を止め、SDK 側または設計を先に確認します。

- 公開 API interface がなく、生成 client の endpoint 名から意味を推測するしかない。
- 許可された情報源に、サービスの用途、resource の意味、または help に記載しようと
  する事実の根拠がない。
- 公式資料を調査してもサービスの用途、resource 間の関係、作成順を利用者向けに説明
  できない。
- SDK request の必須 discriminator、union variant、または null/omit の意味が不明。
- SDK response の resource ID を selector と `--quiet` 用の安定した scalar ID に変換
  できない。特に union ID の discriminator と各 variant の意味が不明。
- selector の一意性、ゾーン、親リソースの関係を確定できない。
- `ctx.SDKClient()` が対象 SDK の `NewClient` と互換でなく、認証・base URL・retry を
  command 側で複製する必要がある。
- sakumock でリクエスト内容と API 呼び出しを再現できず、正常系だけのテストになる。

この場合、手書き HTTP、`ctx.Client()` の危険な型アサーション、SDK 内部型のコピーで
回避しないでください。`SDKClient()` は SDK 用 client を遅延生成します。legacy
`saclient-go` の `ctx.Saclient()` は型互換ではなく、iaas bill/coupon を含む既存 command
のために維持します。SDK API wrapper のために `Saclient()` を移行・置換せず、wrapper
からは必ず `SDKClient()` を使用します。

## 参照

- [SDK interface の選定と operation 分類](references/sdk-selection.md)
- [サービスモデルと help の設計](references/help-and-service-model.md)
- [CLI parameter と request の写像](references/parameter-mapping.md)
- [resource/command の構成、登録、停止条件](references/command-structure.md)
- [sakumock と coverage のテスト](references/testing.md)
- [resource help テンプレート](assets/resource.go.tmpl)
- [手書き用 command テンプレート](assets/command.go.tmpl)
- [manifest テンプレート](assets/coverage-manifest.json.tmpl)
