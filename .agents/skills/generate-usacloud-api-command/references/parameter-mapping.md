# parameter と SDK request の写像

## 基本原則

CLI parameter は SDK request をそのまま露出する DTO ではありません。利用者が安全に
指定できる typed な入力を保持し、`Func` 内で SDK request を組み立てます。共有
parameter は `cli:",squash"` と `mapconv:"-"` を付け、SDK に渡しません。

| SDK フィールド | CLI | 変換規則 |
| --- | --- | --- |
| string, bool, int, duration 相当、enum | typed flag | enum は `options=` と `validate=` を付ける |
| ID / name / tag が対象 | selector | resource selector を使い、自前検索しない |
| `[]string`、`[]int`、`[]enum` のような単純 slice | typed の繰り返し flag | 各要素を validate し SDK の typed slice に変換する |
| map、入れ子 struct、union、入れ子要素を持つ complex slice | JSON flag | JSON document を decode し、型と必須項目を検証する |
| API 応答専用、server computed | 露出しない | request にコピーしない |

単純な scalar と complex JSON を混在させても、同じ意味の値を二重に受け取らないように
します。`--config` の JSON が `--name` と競合するなら、validate で排他にします。
JSON は `encoding/json` で decode し、unknown field の扱いは SDK 契約に合わせて
明示します。秘密値をエラー、出力、fixture に含めません。

complex JSON flag は型を知っている実装者だけが使える状態にしてはいけません。
flag の短い説明から command help を参照させます。`core.Command.LongUsage` には
JSON object の必須 field、任意 field、enum 値、時刻・単位などの形式を記載します。
`Example` には、そのまま実行できる inline JSON と JSON file の例を載せます。
親 request envelope や自動導出する
discriminator は例へ含めません。resource 一覧に表示される `Usage` は一行に保ちます。

## discriminator と nullable

resource/interface から一意に決まる SDK の discriminator は、ユーザーに flag や JSON
として入力させず、`Func` が導出します。例えば `EventBusAPI` 用 command は対応する
request の kind を固定して設定します。discriminator を生 JSON に任せると variant と
payload が食い違います。複数 variant を公開する場合にだけ利用者は業務上の variant を
選び、その選択から `Func` が discriminator を設定します。variant に必要な field、
使用不能な field は validator で検査します。

optional pointer / nullable フィールドには三状態があります。

1. flag 未指定: request から omit（既存値を維持）。
2. 値を指定: 新値を設定。
3. 明示的 clear: `--clear-<field>` を指定して API が期待する null/空値を送る。

`--clear-<field>` と値 flag は排他にし、clear が API の null と空文字列・空配列の
どれに対応するかをテストで固定します。Go の zero value だけではこの区別を失うため、
pointer、optional wrapper、または入力の「変更されたか」の情報を使用します。

## selector と context

resource ID、部分一致名、タグを位置引数で受ける command には、単一対象なら
`core.SelectorTypeRequireSingle`、複数削除などは
`core.SelectorTypeRequireMulti` を設定します。`Func` で検索・ゾーン展開を再実装せず、
core が解決した resource を `ctx.WithResource` から取得します。親 ID を SDK の path
parameter に使う場合も、親子の selector 範囲と `--zone=all` の可否を先に決めます。
