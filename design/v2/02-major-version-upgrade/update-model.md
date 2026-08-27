# 更新モデル

## 目的

`skr <resource> update` の挙動を API ごとの差異から切り離し、利用者から見て一貫した
部分更新として提供する。

更新コマンドは、次の原則を満たさなければならない。

- 指定していないフィールドは変更しない。
- 値の設定と削除を区別できる。
- リストの要素追加、要素削除、全体置換を区別できる。
- ネストしたフィールドでも、各操作の意味を `--help` だけで判断できる。
- API が PUT のみを提供する場合も、CLI の表面上は同じ部分更新として扱う。

以下では、CLI が受け取る変更内容を **update plan**、API に送るリクエストを
**update request** と呼ぶ。update plan は値だけでなく、フィールドごとの操作と
「指定されていない」という状態を保持する。

## 現状の課題

現行実装は、scalar については pointer と `CleanupEmptyValue` により「フラグ未指定」
と「空値の指定」を区別している。一方、次の点はコマンドごとに異なる。

- `--tags` や `--records` が全体置換なのか追加なのか、フラグ名から判断できない。
- `--network-interfaces` など、ネストした値を JSON でのみ受け取るフラグがある。
- VPC ルーターのように、一部の nested leaf は専用フラグ、一部は JSON になっている。
- API の update request へ直接変換するため、PUT で未指定値を保持する責務が不明確。
- help に型、JSON schema、更新範囲、削除方法が表示されない。

この設計では JSON の利用自体を廃止せず、操作の意味がフラグ名と help から分かる
形に統一する。

### DNS update の例

DNS は現状の問題がまとまって現れている代表例である。

- CLI は `--name` を表示するが、利用している service の `UpdateRequest` に Name は
  なく、実際には更新対象にならない。
- `--records` は JSON を読み込み、空の `DNSRecords` に append した後で update
  request へ渡す。既存レコードへ追加する処理ではないため、非空の指定は実質的に
  全体置換となるが、フラグ名と help からは判断できない。
- service 層は現在値を Read してから Update するが、Records は pointer ではなく
  `omitempty` の list である。この表現では「未指定」と「空 list への置換」を
  update request 上で区別できず、全レコード削除を一貫して指定できない。
- DNS レコードの各フィールドや削除時の照合条件が help に表示されない。

`skr dns update` では、少なくとも次のインターフェースに変更する。

```console
# Add one record without changing existing records
skr dns update example.com \
  --records-add='[{"Name":"www","Type":"A","RData":"192.0.2.1","TTL":300}]'

# Remove records identified by Name, Type, and RData
skr dns update example.com \
  --records-remove='[{"Name":"old","Type":"A","RData":"192.0.2.2"}]'

# Replace every record
skr dns update example.com --records-replace=@records.json

# Remove every record
skr dns update example.com --clear-records
```

DNS record の identity は `Name,Type,RData` とし、TTL は identity に含めない。
同じ identity のレコードへ異なる TTL を add した場合は、既存要素の TTL を更新する。
同じ update plan で remove と add を行う場合は remove、add の順に適用する。

DNS の update plan には、API と service が実際に更新できるフィールドだけを公開する。
Name のように update request に存在しないフィールドは表示しない。

Records の add/remove は現在値を必要とするため、既存の service 層の
read-modify-write を利用する。ただし空 list を「未指定」として扱う現行
`UpdateRequest` では clear/replace-to-empty を表現できないため、optional な
collection operation を受け取れるよう service adapter を変更する。

## 一時的な運用操作

更新可能な値の中には、Terraform などの宣言的管理へ含めるより、障害対応や
メンテナンス時に一時的に切り替えるほうが自然なものがある。例えばロードバランサの
特定 backend の有効・無効、リソースの接続・切断などである。

これらを nested object/list 全体の JSON 更新だけで操作させると、次の問題がある。

- 現在の設定を取得し、対象要素だけを書き換え、完全な JSON を再送する必要がある。
- unrelated な backend や health check を誤って削除または上書きしやすい。
- `Enabled: false` が「一時停止」なのか構成全体の置換なのか、コマンドから分からない。
- Terraform へ一時的な状態を記述すると、通常の構成変更と運用上の override が混ざる。

頻繁に行う一時的な操作には、汎用の `update` フラグだけでなく、意図を表す専用
action command を提供する。

### Load Balancer backend の例

ロードバランサでは Virtual IP を `VirtualIPAddress,Port`、その配下の backend を
`IPAddress,Port` で識別し、次のようなコマンドを提供する。

```console
# Temporarily stop sending traffic to one backend
skr load-balancer backend disable example-lb \
  --virtual-ip-address=192.0.2.10 \
  --virtual-port=443 \
  --ip-address=192.0.2.101 \
  --port=443

# Resume traffic to the backend
skr load-balancer backend enable example-lb \
  --virtual-ip-address=192.0.2.10 \
  --virtual-port=443 \
  --ip-address=192.0.2.101 \
  --port=443
```

`enable` と `disable` は冪等にする。既に目的の状態であれば成功し、対象の Virtual IP
または backend が存在しなければ、全体を作り直さず not found エラーを返す。

API が Load Balancer 設定全体の PUT しか提供しない場合、実装は次のようにする。

1. Load Balancer の現在値を read する。
2. 複合 identity で対象 Virtual IP と backend を一意に特定する。
3. 対象 backend の Enabled だけを変更する。
4. SettingsHash など API の競合検出値を維持して設定全体を apply する。
5. 適用完了を待つ。ただし `--no-wait` が指定された場合を除く。

この処理は update plan の nested leaf 更新と同じ applier を再利用し、専用 command
から list 全体を再実装しない。

### 専用 action command にする基準

次をすべて満たす操作は、`update` のみで提供せず専用 action command を設ける。

- 障害対応、メンテナンス、切り戻しなどで繰り返し実行される。
- 対象を安定した identity で一意に指定できる。
- 操作名が利用者の意図を明確に表せる。
- 他のフィールドを維持したまま冪等に実行できる。

専用 action は API が専用 endpoint を持つことを条件としない。CLI が安全な
read-modify-write を実装できればよい。逆に identity や更新後の整合性を保証できない
場合は、便利さのために曖昧な action を追加しない。

Terraform など外部の desired state 管理下にあるリソースへ action を実行すると、
次回の reconcile/apply で元の値へ戻る可能性がある。これはエラーではないが、
該当 action の help に明記する。CLI は Terraform の state を書き換えず、運用上の
override と宣言的な構成変更を混同しない。

## 基本セマンティクス

### 未指定、設定、削除

各フィールドは最低でも次の 3 状態を区別する。

| 状態 | CLI の例 | 意味 |
|---|---|---|
| 未指定 | フラグなし | 現在値を維持する |
| 設定 | `--description="new value"` | 指定した値に変更する |
| 削除 | `--clear-description` | 値を削除し、API 上の未設定値にする |

空文字列、`0`、`false`、空リストは有効な値であり、未指定として扱わない。
削除は空値の設定で代用せず、`--clear-<field>` で明示する。API が削除を表現できない
フィールドには clear フラグを生成しない。

`--<field>` と `--clear-<field>` は同時に指定できず、指定された場合は API を
呼び出す前にエラーにする。

boolean は通常の値と同様に扱い、`--enabled=true` と `--enabled=false` の両方を
設定操作とする。引数を省略した `--enabled` を許容する場合も、その意味は `true`
への設定であり、未指定とは区別する。

### 更新可能なフィールド

更新用 SDK 型に存在するという理由だけでフラグを生成しない。各フィールドには
以下のメタデータを持たせ、生成時に更新方法を決定する。

| メタデータ | 内容 |
|---|---|
| path | API リクエスト内のフィールドパス |
| type | scalar、object、list、map |
| operations | set、clear、add、remove、replace のうち利用可能な操作 |
| identity | list 要素の同一性を判定するキー |
| transport | PATCH、PUT、専用 API のいずれで反映するか |
| readable | read-modify-write 用に現在値を読み戻せるか |

OpenAPI や SDK から確定できない情報、特に削除可能性と list 要素の identity は、
生成設定で明示する。推測で破壊的な操作を生成しない。

## 型ごとの操作

### scalar

文字列、数値、boolean、ID、enum などは次の形式にする。

```console
--name VALUE
--enabled=true
--clear-icon-id
```

- `--<field>`: 値を設定する。
- `--clear-<field>`: API が許す場合に値を削除する。
- enum の候補、値の単位、clear の結果は help に記載する。

### object

頻繁に操作するネストした object は、更新可能な leaf ごとにフラグを公開する。
フラグ名は object のパスをハイフンで連結する。

```console
--health-check-protocol=https
--health-check-path=/ready
--clear-health-check-path
```

この例で `--health-check-path` だけを指定した場合、`protocol` を含む同じ object の
他フィールドは変更しない。

object 全体の置換が必要な場合は、明示的に `--<field>-replace` を使用する。

```console
--health-check-replace='{"Protocol":"https","Path":"/ready"}'
--clear-health-check
```

`--<field>-replace` は、省略された子フィールドを現在値から補完しない。指定した
JSON object で全体を置き換える。leaf の更新フラグと object の replace/clear は
同時に指定できない。

すべてのネストした型を leaf フラグへ展開する必要はない。低頻度または将来拡張が
多い object は replace のみでもよいが、その場合もフラグ名と help で全体置換で
あることを明示する。単なる `--health-check` のように部分更新か置換か判別できない
名前は新規に導入しない。

### list

list は次の操作を明示的に分離する。

```console
--tags-add=production
--tags-remove=staging
--tags-replace=production,web
--clear-tags
```

| 操作 | 意味 |
|---|---|
| `--<field>-add` | 現在の list に要素を追加する |
| `--<field>-remove` | 現在の list から一致する要素を削除する |
| `--<field>-replace` | list 全体を指定値で置き換える |
| `--clear-<field>` | list を空にする、または nullable list 自体を削除する |

`replace` または `clear` は、`add`、`remove` および互いに排他的とする。`add` と
`remove` は同時指定でき、その場合は remove、add の順に適用する。同じ要素を
remove と add の両方へ指定した場合、最終的に要素は存在する。

scalar list は値そのものを identity とする。重複を許さない list では、既存要素の
add は冪等に扱う。順序に意味がある list では add は末尾への追加とし、help に
順序の意味を記載する。

object list は JSON object または JSON array を受け取る。

```console
--records-add='[{"Name":"www","Type":"A","RData":"192.0.2.1"}]'
--records-remove='[{"Name":"www","Type":"A","RData":"192.0.2.1"}]'
--records-replace=@records.json
```

object list の remove は、リソースごとに定義した identity フィールドで照合する。
上の DNS レコードで identity を `Name,Type,RData` と定義した場合、それらすべてが
一致する要素を削除する。identity を安全に定義できない list には remove を提供せず、
replace のみを提供する。削除対象が存在しない場合はエラーにせず、冪等な成功とする。

`--clear-<field>` が「空 list」と「null」のどちらを送るかは API のデータモデルに
従い、その結果を help に明記する。両方を区別する必要がある API では
`--<field>-replace='[]'` を空 list、`--clear-<field>` を null とする。

### map

map は key 単位の更新と全体置換を分ける。

```console
--labels-set=environment=production
--labels-remove=temporary
--labels-replace=@labels.json
--clear-labels
```

`set` と `remove` は同時指定でき、remove、set の順に適用する。`replace` と `clear`
は他の map 操作と排他的とする。

## JSON 入力

JSON はコマンドラインに直接指定できるほか、`@path` でファイルを指定できるように
統一する。標準入力は `@-` とする。

```console
--records-replace=@records.json
--records-add=@-
```

update コマンドの `--parameters` は **update plan の JSON 表現**を受け取る。
トップレベルおよびネストした object では、キーの省略は「変更しない」、`null` は
「削除」、その他の値は「設定」を意味する。list の値は全体置換を意味する。

```json
{
  "Description": null,
  "HealthCheck": {
    "Path": "/ready"
  },
  "Tags": ["production", "web"]
}
```

この例では Description を削除し、HealthCheck.Path のみを設定し、Tags 全体を
置き換える。他のフィールドは変更しない。

JSON 内でも list の add/remove が必要な場合は、操作 object を使用する。

```json
{
  "Records": {
    "$remove": [
      {"Name": "old", "Type": "A", "RData": "192.0.2.1"}
    ],
    "$add": [
      {"Name": "www", "Type": "A", "RData": "192.0.2.2"}
    ]
  }
}
```

`$replace`、`$clear` も対応する CLI フラグと同じ排他規則に従う。通常の list 値は
`{"$replace": [...]}` の短縮形とする。

`--parameters` と個別フラグは併用できる。まず `--parameters` の update plan を
読み込む。同じ field path を個別フラグでも指定した場合は、その field path に
対する JSON 側の操作をすべて破棄し、個別フラグ側の操作を採用する。その後、
個別フラグ内の排他関係を検証する。例えば JSON の `Tags.$replace` と
`--tags-add` を併用した場合は、replace を破棄して add のみを行う。

JSON の未知のキーは typo による意図しない未更新を防ぐためエラーにする。

## API への反映

### PATCH または部分更新 API

API が省略フィールドを保持する PATCH、または同等の update API を提供する場合は、
update plan をその API のリクエストへ変換する。CLI の未指定状態を失わないよう、
生成する Go 型は pointer または明示的な optional 型を使用する。

### PUT

API が全体置換の PUT しか提供しない場合、CLI 内部で次の read-modify-write を行う。

1. 対象リソースを read する。
2. read 結果から、API が受け付ける更新可能なモデルを構築する。
3. update plan をそのモデルへ適用する。
4. 完成したモデルを検証する。
5. PUT を実行する。

read 結果をそのまま PUT しない。read-only フィールドを除外し、更新可能な
フィールドだけを明示的に投影する。write-only など read で復元できない必須値が
ある API は、自動的な read-modify-write の対象にせず、専用実装または必須フラグを
用意する。

API が ETag、世代番号、revision などの条件付き更新を提供する場合は必ず利用し、
read 後に他の更新が入った場合は競合エラーを返す。条件付き更新がない API では
競合を完全には防げないため、read と PUT を連続して実行し、可能なら API 固有の
revision を比較する。この制約は該当コマンドの help に記載する。

複数リソースを一度に更新する場合、各対象について独立して read-modify-write を
行う。一部だけ成功した場合は、成功・失敗した対象を明示し、ロールバック済みで
あるかのように扱わない。

## help の要件

update コマンドの help には、冒頭で次の共通説明を表示する。

```text
Updates only fields explicitly specified. Unspecified fields are preserved.
Use --clear-<field> to remove a value. List flags ending in -add, -remove,
and -replace add elements, remove elements, and replace the entire list.
```

各フラグの説明には、少なくとも以下を含める。

- 操作種別: set、clear、add、remove、replace
- 値の型と形式。JSON の場合は object/array の schema または具体例
- enum の全候補
- list remove に使用する identity
- clear が null と空値のどちらになるか
- ネストしたフィールドのうち、一緒に変更される範囲
- read-modify-write を行う場合はその旨と競合検出の有無

例:

```text
Resource options:
      --description string          Set Description. Empty string is allowed.
      --clear-description           Remove Description (set it to null).
      --tags-add strings            Add tags; existing tags are preserved.
      --tags-remove strings         Remove tags by exact match.
      --tags-replace strings        Replace the entire tag list.
      --clear-tags                  Replace Tags with an empty list.
      --records-add string          Add DNS records from a JSON array.
      --records-remove string       Remove records matching Name, Type, and RData.
      --records-replace string      Replace all DNS records from a JSON array.
```

`--generate-skeleton` は update plan の skeleton を出力する。ただし、すべてのキーを
埋めた JSON は全項目更新に見えるため、デフォルト出力は空 object とフィールドごとの
コメントを生成できる別形式に改めるか、`--help` と `--example` で更新例を示す。
少なくとも、生成した skeleton をそのまま実行すると全フィールドを変更する形式には
しない。

## エラーと安全性

- update plan に更新操作が一つもなければエラーにする。
- 同一フィールドへの排他的な操作は API 呼び出し前にエラーにする。
- read-only フィールド、未知の JSON キー、未対応の操作はエラーにする。
- 複数対象の更新では、対象ごとの結果を識別できるエラーを返す。
- dry-run を実装する場合は、API リクエストではなく update plan 適用後の差分を
  表示する。secret や write-only 値は表示しない。

## 実装モデル

現在の `CleanupEmptyValue` と pointer フィールドだけでは、未指定、null、空値に加え、
list の各操作を一貫して表現できない。v2 ではフラグから直接 SDK の UpdateRequest を
組み立てず、中間の update plan を導入する。

概念上は次のように scalar/object と list/map で型を分ける。実際の型安全な
field path と値型はコード生成する。

```go
type Change[T any] struct {
    Set   *T
    Clear bool
}

type CollectionChange[T any, K any] struct {
    Add     []T
    Remove  []K
    Replace *[]T
    Clear   bool
}
```

`Change` は Set と Clear を排他的にする。`CollectionChange` は Add と Remove の
併用を許し、Replace と Clear は他のすべての操作と排他的にする。map の Remove は
key、object list の Remove は identity 型 `K` を用いる。

処理は次の層に分離する。

1. flag/JSON parser: 入力を update plan に変換する。
2. validator: 利用可能な操作、排他関係、値、identity を検証する。
3. applier: 現在値が必要な操作をモデルへ適用する。
4. transport adapter: PATCH、PUT、専用 API のリクエストへ変換する。

API ごとの例外は parser ではなく、生成メタデータまたは transport adapter に置く。

## 移行方針

`usacloud` から `skr` への改名を更新インターフェースの互換性境界とする。現行の
更新コマンドは操作の意味が曖昧で、DNS など実用上必要な操作を表現できないため、
曖昧な旧フラグの互換 alias は原則として設けない。

1. `--name` のように意味が明確な scalar フラグは維持する。
2. 現在 list 全体を置換する `--tags`、`--records` などは
   `--tags-replace`、`--records-replace` へ変更する。
3. object 全体を受け取る既存 JSON フラグは `--<field>-replace` へ変更する。
4. 利用頻度が高い nested leaf に専用フラグを追加する。
5. 各 update API を PATCH、PUT、専用更新 API に分類し、PUT には
   read-modify-write adapter を追加する。
6. 全 update コマンドについて「フラグなし」「空値」「clear」「list の各操作」
   「ネストした leaf のみ」の contract test を生成する。

`skr` で旧フラグを受け付けず、unknown flag と代替フラグを含む migration guide
によって移行を案内する。例外的に旧フラグを残す場合は、単に新フラグへ alias
するだけで意味が完全に一致するときに限る。例えば旧 `--records` は追加とも置換とも
読めるため alias にしない。

## 対象外

- API 自体が提供しない更新や削除を擬似的に実現すること。
- identity を定義できない list の要素単位削除。
- 条件付き更新機構がない API に対する完全な競合防止。
- JSON Patch の任意 path や index 操作。API モデルの変更に弱く、help から
  更新可能範囲を発見しにくいため、公開インターフェースには採用しない。
