# command 構成、登録、出力

## resource と command

package 名は resource 名を小文字・ハイフンなしにしたものを基本とします。予約語衝突は
既存の `iface`、`swytch` を踏襲します。`Resource` の category と command category /
order は既存 resource の同等操作に合わせます。未知の category は `AddCommand` が
拒否します。

SDK 直結 command は生成 service adapter では表せないため、各 command literal に
**明示的に** `Func: xxxFunc` を置きます。SDK operation は SDK の client/op constructor
で生成します。

```go
sdkClient, err := ctx.SDKClient()
if err != nil {
    return nil, fmt.Errorf("create SDK client: %w", err)
}
client, err := eventbus.NewClient(sdkClient)
if err != nil {
    return nil, fmt.Errorf("create EventBus client: %w", err)
}
op := eventbus.NewXOp(client)
result, err := op.<Method>(ctx, request)
```

`eventbus`/`X` は対象 SDK の package/interface 名へ置き換えます。共通 helper を置く
場合も、この `ctx.SDKClient()` → SDK `NewClient` → SDK `NewXOp` の流れを包むだけに
します。`SDKClient()` のエラーには生成段階を付けて返します。base URL、認証、retry の
ために client を独自生成してはいけません。legacy API 用の `ctx.Saclient()` は
`saclient-go` client のまま維持し、`sacloud-sdk-go/api/*` wrapper からは呼び出しません。

mutation（create/update/delete と状態を変える action）の parameter には
`cflag.ConfirmParameter` を埋め込みます。これにより非対話実行では
`--assumeyes` / `-y` が要求されます。read/list に確認を付けません。

## 表示と登録

list は `NoProgress: true` とし、resource の `defaultColumnDefs` または command 固有の
`ColumnDefs` を明示します。列の名前、順序、template は CLI 契約です。map の反復順や
JSON の偶然のフィールド順に依存せず、識別子（ID/Name）、親・zone、状態などを固定順
で選びます。詳細 response の全フィールドを table の列にしません。

resource を正しい集合へ登録します。

* IaaS: `pkg/commands/iaas/resources.go`
* その他のトップレベル resource: `pkg/resources.go`

新しい resource は必要な import とともに登録し、`make gen` を実行して生成対象を更新
します。`*_gen.go` と `Code generated` ファイルは手編集しません。追加後は `usacloud
<resource> -h` と command help で category/order/alias を確認します。

coverage validator には service root（例: `./pkg/commands/eventbusapi`）を渡します。
子 directory を再帰走査し、各 package の `core.Resource{Name: ...}` を取得して
`ProcessConfigurationAPI/process-configuration + delete` のように
interface/resource/command の三つ組で照合するため、別 resource の同名 `delete` で
欠落を満たすことはありません。
