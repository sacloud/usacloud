# SDK インターフェースの選定と分類

## 選定

対象は `sacloud-sdk-go/api/<service>` package の exported な `XxxAPI` interface です。
`NewXxxOp` が返す interface と、その `XxxOp` 実装を対にして確認します。client、
transport、認証、pager、`apis/v*` の生成型を直接 CLI に露出しません。

inventory は `go list -json` で解決した package の実ファイルを読み、`go/parser` /
`go/ast` および `go/types` で interface と署名を取得します。そのため SDK 更新後にも
import path と型表記を手で転記しません。

`--interface` を省略すると、すべての `*API` interface を inventory します。ただし
`List` や `Create` という同名 method を 1 配列へ混在させません。
schema v2 の `interfaces[]` は、各要素に `interface` と単数 CLI resource 名、
`methods[]` を保持します。単数 CLI resource 名には、interface 名から `API` を除いて
kebab-case にした合意済みの名前を使います。たとえば `ProcessConfigurationAPI` には
`process-configuration` を使います。略語や SDK 名と CLI 名が異なる
場合は、inventory 後に `resource` を合意済みの CLI 名へ修正します。resource は package
directory 名ではなく、`core.Resource{Name: ...}` と coverage の照合キーの一部です。

```sh
go run ./.agents/skills/generate-usacloud-api-command/scripts/sdk-interface-inventory \
  --package github.com/sacloud/sacloud-sdk-go/api/<service> --output -
```

候補ごとに、次を manifest の `excluded` へ理由付きで残します。

* CLI の resource 操作として意味があるか（認証・helper は除外）。
* request/response を安定した CLI 契約に変換できるか。
* `ctx.SDKClient()` から得た `sacloud-sdk-go` client を対象 SDK の `NewClient` に
  渡せるか。
* selector、ゾーン、親 ID が明確か。

`ctx.SDKClient()` は `sacloud-sdk-go/common/saclient.ClientAPI` を遅延生成して返します。
legacy `saclient-go` の `ctx.Saclient()` はこの型と互換ではありません。iaas bill/coupon
などが引き続き `Saclient()` を必要とするため、SDK API wrapper の追加時に legacy client を
移行・置換せず、wrapper は `SDKClient()` だけを使います。

## operation の分類

inventory は名称から `list` (`List`/`Find`)、`read` (`Read`/`Get`)、`create`、
`update`、`delete`、それ以外を `action` と暫定分類します。これはレビューの起点で
あり、API のセマンティクスを置換しません。

| 分類 | 通常の CLI 名 | selector | 確認 |
| --- | --- | --- | --- |
| list | `list` | なし | 不要。`NoProgress: true` |
| read | `read` | 1 件なら `RequireSingle` | 不要 |
| create | `create` | 親のみ必要なら親 selector | 必須 |
| update | `update` | `RequireSingle` | 必須 |
| delete | `delete` | `RequireMulti` が基本 | 必須 |
| action | SDK 意味を表す kebab-case | 対象数に合わせる | 状態変更なら必須 |

`Read` は CLI では `read`、`Get` も通常 `read` とします。複数取得は `list` に統一
し、SDK 名をそのまま command 名にしません。既存 resource と同じ語彙を優先します。
`Start`, `Stop`, `RotateKey` のような action は曖昧な `update` に畳み込まず、
`start`、`stop`、`rotate-key` として公開します。

完全一致の CRUD 名だけを `list` / `read` / `create` / `update` / `delete` に揃えます。
`UpdateSecret` のように suffix を持つ method は mutation と分類されますが、suggestion は
`update-secret` として固有性を保ちます。manifest の `command` は推測値を必ずレビューして
更新します。1 API メソッドを公開
しない場合も `methods` から削除せず、その interface の `excluded` に
`{ "method": "...", "reason": "..." }` を書きます。validator は reason が空、または
inventoried されていない method の除外を拒否し、正当な除外は coverage と件数から除きます。
