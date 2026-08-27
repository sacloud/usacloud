# OpenAPI / SDK からの半自動生成

## 目的

各サービスの API client のソースコードと OpenAPI 仕様を読み取り、全サービスの
低レベルコマンドを AI で生成・整備する。

- コマンド定義（リソース、アクション、フラグ）
- help テキスト
- Skill ファイル
- 出力カラム定義

これにより、新サービス追加時の工数を削減し、help text や Skill ファイルの網羅性・正確性を向上させる。

生成対象は `usacloud <service>-api <resource> <action>` である。AppRun や Object Storage
などに追加する `usacloud <service> <operation>` 形式の高レベルコマンドは、この機械的な
生成対象に含めない。

## 生成方針

### 低レベルコマンド

- 全サービスについて提供する。
- API client の service、resource、method、request/response 型を原則一対一に反映する。
- AI に SDK、OpenAPI、生成規約を与えて実装し、同じ入力から反復可能に生成できる
  状態を保つ。
- API の網羅性を優先し、高レベルコマンドの有無を理由に生成を省略しない。
- 生成後にコンパイル、contract test、help の検査を行う。

### 高レベルコマンド

高レベルな処理は、`iaas-service-go` と同様の service layer として
`sacloud-sdk-go` に実装する。複数 API の組み合わせ、read-modify-write、待機、
ローカルファイル操作など、具体的なユースケースを request/response として定義する。

CLI は service layer の薄いラッパーとし、service request の組み立て、入力検証の
表示、結果の出力を担当する。service のロジックを CLI に重複実装しない。

service layer と CLI wrapper は AI を利用して実装してよいが、API client の構造を
そのまま展開する自動生成物にはしない。AppRun の deploy や logs、Object Storage の
cp や sync のように、低レベル API の組み合わせを利用者へ要求することが大きな負担に
なる場合に限って追加する。

## API クライアントからの自動生成方法

`sacloud-sdk-go` に含まれる各サービスの API client の仕様を解釈し、低レベル CLI
を生成する。
その際の実装に統一感が出るように AGENTS.md を細かく定義する。

例えば以下のような事項が AGENTS.md に記載される。

### API クライアントからの命名導出

API クライアントの構造から usacloud のコマンド名を導出するルールは **例えば** 以下のようなものになるであろう。
このあたりは実際に実装をしていきながら詳細を詰めていく必要がある。

IaaS 以外のリソースについては Experimental 扱いで main line にマージしながら実装していき、すべて出揃った段階で再度ネーミングを調整していくような形が良いかもしれない。

#### サービス名

- API client のサービス名を基本とし、低レベル名前空間には `-api` を付ける。
- アンダースコアはハイフンに変換する（例: `object_storage` → `object-storage`）。
- 単複は原則として API クライアントのサービス名をそのまま用いる（例: `workflows`, `simplemq`）。

#### リソース名

- API クライアント内の API ファイル名（`*_api.go` や `*.go`）から導出する。
- ファイル名が `workflow.go`, `execution.go` ならリソース名は `workflow`, `execution`。
- アンダースコアはハイフンに変換する（例: `site_status.go` → `site-status`）。
- サービス名とリソース名が重複・類似しても、初期生成では省略しない
  （例: `workflows-api workflow`）。短縮形が必要な場合は、レビュー済みの明示的な
  エイリアスとして追加する。

#### アクション名

- API クライアントのメソッド名から導出する。

| API クライアントのメソッド | usacloud アクション |
|---|---|
| Create | create |
| List | list |
| Read | read |
| Update | update |
| Delete | delete |
| ListSuggest | list-suggest |
| Cancel | cancel |
| ListHistory | list-history |
| その他 | 必要に応じて定義 |

#### 子リソースの判定

- OpenAPI のパス階層と API クライアントの引数を照合して候補を抽出する。
- ID 引数があることだけでは子リソースと確定しない。曖昧な場合は生成設定に親リソースを明示し、人間のレビューなしに階層を決定しない。
- 子リソースは原則フラットに扱い、親 ID と子 ID の引数順を生成設定とテストで固定する。

### 低レベル update

`<service>-api ... update` は API client の UpdateRequest と PUT/PATCH セマンティクスを
そのまま公開する。CLI で read-modify-write や部分更新への変換を行わない。API が
全体置換の PUT のみを提供する場合は、低レベルコマンドも全体置換となることを help
へ明記する。

[更新モデル](../02-major-version-upgrade/update-model.md) に従う安全な部分更新は、
`sacloud-sdk-go` の service layer と、それを利用する高レベルコマンドに実装する。

## 品質保証

* 完全自動化は目指さず、人間が確認すること。
* 実際に動作確認を実施すること
    * 動作確認は本番環境に向けて行うこと
    * 本番環境での動作確認が難しいサービスについては sakumock での動作確認を行なっても良い
* 本番での動作確認の結果をログとして残しておき、マニュアルを生成するところまでを新規サービスサポートの一連のフローとしたい。
