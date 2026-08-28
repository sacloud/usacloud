# sakumock と manifest coverage のテスト

unit test は SDK operation の interface を sakumock で差し替えます。network を使わず、
各 command の `Func` が選択された ID、typed flag、JSON request、discriminator、
nullable clear を正しい SDK request にしたことを検査します。

最低限、次を table test にします。

1. scalar flag が request の対応フィールドへ変換される。
2. complex JSON の nested 値と derived discriminator が一致する。
3. 未指定、指定、`--clear-*` の三状態が異なる request になる。
4. single/multi selector が期待した resource 数で command を呼ぶ。
5. mutation が `--assumeyes` なしの非対話時に拒否される。
6. response が固定 `ColumnDefs` と期待する行数で出力される。

mock の期待値は request 全体、または対象の意味ある field を比較し、呼び出された
operation と context を検査します。SDK generated client を直接 HTTP mock するより、
`ctx.SDKClient()` から SDK の `NewClient` / `NewXOp` で得た interface を使います。SDK
API wrapper に legacy `ctx.Saclient()` を渡さず、既存 iaas bill/coupon の legacy client を
移行・置換しません。SDK interface が mock 不能なら、その interface を consumer 側で
狭めず SDK 側の mock 支援を追加します。

inventory/coverage script の unit test は hidden directory を `./...` が探索しないため、
明示的に実行します。

```sh
go test \
  ./.agents/skills/generate-usacloud-api-command/scripts/sdk-interface-inventory \
  ./.agents/skills/generate-usacloud-api-command/scripts/validate-implementation-coverage
```

実装後には対象 command package の test と、上記 script test を実行します。coverage
validator は service root 以下を再帰走査し、manifest の
`interface/resource/method/command` ごとに `core.Command{Name: ..., Func: ...}` を
機械検証します。resource identity は directory 名ではなく同一 package の
`core.Resource{Name: ...}` から取得します。`excluded` の inventoried method は coverage
から除外されます。mutation 確認、selector、mapping の意味は sakumock test とレビューで
検証します。
