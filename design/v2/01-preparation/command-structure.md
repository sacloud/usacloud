# コマンド体系の再整理

## 現状

現在の CLI は、サービス単位のトップレベルコマンドを持つ構成になっている。

```
usacloud [global options] <command> <sub-command> [options] [arguments] [flags]
```

トップレベルの Available Commands には以下が含まれる。

- `iaas` - SubCommands for IaaS
- `web-accelerator` - SubCommands for WebAccelerator
- `config`, `rest`, `completion` など

IaaS リソースには、以下の 2 通りの方法でアクセスできる。

```bash
usacloud iaas server list
usacloud server list
```

両者は同じ動作をする。つまり、`iaas` はサービスを明示するためのオプショナルなサブコマンドであり、IaaS リソースはルート直下にも登録されている。

> [!NOTE]
> ただし、トップレベルの help（`usacloud --help`）には `iaas` や `web-accelerator` のみが表示され、`server` などの個別リソースは表示されない。
> これは初学者や AI にとってのわかりにくさを産んでいる可能性がある。

## 目指すコマンド体系

### 基本方針

- IaaS リソースは `usacloud <resource> <action>` を第一級の推奨形とする
- `usacloud iaas <resource> <action>` も引き続き動作させる
- 全サービスに、API client の操作を網羅する低レベルコマンドを提供する
- AppRun や Object Storage など、複雑な複合操作が必要なサービスに限り
  高レベルコマンドを追加する

## 高レベルコマンドと低レベルコマンド

AWS CLI の `s3` と `s3api` の関係と同様に、コマンドを二層に分ける。

| 層 | 名前空間 | 目的 | 提供範囲 |
|---|---|---|---|
| 高レベル | `usacloud <service> ...` | service layer の複合操作を簡潔かつ安全に実行する | 必要なサービスのみ |
| 低レベル | `usacloud <service>-api <resource> <action>` | API client の各操作を直接利用する | 全サービス |

低レベルコマンドは API client の構造へ忠実であることと網羅性を優先し、AI による
生成を基本とする。高レベルコマンドは API の単純な別名ではなく、複数 API の
呼び出し、read-modify-write、待機、入力補完などをまとめる場合にのみ提供する。

高レベルコマンドが不要なサービスには `usacloud <service>` 名前空間を作らず、
`usacloud <service>-api` のみを提供する。すべてのサービスへ形式的に高レベルコマンドを
作ることは目標にしない。

高レベルな処理は `iaas-service-go` と同様の service layer として
`sacloud-sdk-go` に実装する。CLI 側は service の request/response をフラグと出力へ
接続する薄いラッパーにする。service layer と CLI wrapper の実装には AI を利用
できるが、API client からの機械的な生成対象にはしない。ユースケース、操作単位、
安全性、冪等性を設計し、人間がレビューする。

通常の利用者と AI エージェントには高レベルコマンドを優先して案内する。高レベルで
提供されない API 操作、API 固有機能、トラブルシューティングには低レベルコマンドを
使用する。高レベルコマンドが存在しても、対応する低レベル操作は省略しない。

既存 IaaS コマンドの `usacloud server ...` や `usacloud disk ...` は、`iaas-service-go` を
利用する service-layer wrapper として互換性のため維持する。一方、全サービス共通の
低レベル層として `usacloud iaas-api <resource> <action>` も提供する。既存 IaaS
コマンドの存在を理由に `iaas-api` の生成を省略しない。

### 例

```bash
# IaaS（プレフィックスなし）
usacloud server list
usacloud disk create
usacloud switch read

# IaaS(プレフィックスありでもこれまで同様に利用可能)
usacloud iaas server list
usacloud iaas disk create
usacloud iaas switch read

# 低レベルコマンド: 全サービスで提供する
usacloud iaas-api server list
usacloud web-accelerator-api site list
usacloud object-storage-api bucket list
usacloud iam-api user list
usacloud apprun-api application list
usacloud simplemq-api queue list

# 高レベルコマンド: 複合操作が必要なサービスに限り提供する
usacloud object-storage sync ./public s3://example-bucket/
usacloud object-storage cp ./artifact.zip s3://example-bucket/releases/
usacloud apprun deploy --name example --source .
usacloud apprun logs example --follow
```

## サービスの分類

対応候補となるサービスを大まかに分類すると以下のようになる。

### 既存対応済みリソース

- `iaas`: IaaS リソース
- `webaccel`: Web Accelerator

### 今後の追加候補（例）

- `object-storage`: オブジェクトストレージ
- `iam`: IAM
- `apigw`: API Gateway
- `apprun`: AppRun
- `simple-notification`: シンプル通知
- `simplemq`: SimpleMQ
- `workflows`: Workflows
- その他
  - addon, apprun-dedicated, cloudhsm, dedicated-storage
  - eventbus, kms, monitoring-suite, nosql
  - secretmanager, security-control, service-endpoint-gateway 等

## コマンド体系のパターン

### 低レベルコマンド

```
usacloud <service>-api <resource> <action> [args...] [flags]
```

API client の service、resource、method を原則として一対一に対応させる。

```bash
usacloud object-storage-api account read
usacloud object-storage-api bucket list
usacloud object-storage-api permission create
usacloud iam-api user list
usacloud iam-api service-principal list
```

サービス名と代表リソース名が重複しても省略しない。例えば Workflows は次の形を
正式な低レベルコマンドとする。

```bash
usacloud workflows-api workflow list
usacloud workflows-api execution list
usacloud workflows-api subscription read
```

### 高レベルコマンド

```
usacloud <service> <operation> [args...] [flags]
```

operation は API client のメソッド名から導出せず、利用者が達成したい操作から命名する。
実処理は `sacloud-sdk-go` の service layer が担い、CLI は service request の構築と
結果表示だけを担当する。

### 子リソースの扱い

workflows の `Execution` は `Workflow` の子リソース。Read には `workflowID` と `executionID` の両方が必要。

```bash
# 案 1: フラット（引数で親子を表現）
usacloud workflows-api execution read <workflow-id> <execution-id>

# 案 2: 親子階層を表現
usacloud workflows-api workflow execution read <workflow-id> <execution-id>
```

原則として、子リソースの操作が独立していれば案 1 のほうが素直な実装といえる。

初期生成では案 1 を採用する。親 ID、子 ID の順序を Usage とテストで固定し、
同じ名前の子リソースがサービス内で衝突する場合のみ案 2 を採用する。

### アクション名のマッピング

API client のメソッド名から低レベルコマンドのアクション名を決定する。

| API client メソッド | usacloud アクション |
|---|---|
| Create | create |
| List | list |
| Read | read |
| Update | update |
| Delete | delete |
| ListSuggest | list-suggest |
| Cancel | cancel |
| ListHistory | list-history |

## 互換性の扱い

### 原則

- 推奨形は `usacloud <resource> <action>`（既存 IaaS）とする。
- `usacloud iaas <resource> <action>` も引き続き動作させる。
- 新しい低レベル名前空間は API client のサービス名に `-api` を付け、
  必要に応じてハイフン区切りにする（例: `object-storage-api`）。
- 高レベル名前空間は `-api` を付けない（例: `object-storage`）。
- リソース名は API client の API ファイル名・パッケージ名を基本とし、snake_case は kebab-case に変換する。

## 内部パッケージ構成

CLI 上のコマンド体系と内部パッケージ構成を一致させる必要はない。
内部パッケージは整理のためサービス単位に分離する。

```
pkg/commands/
├── root/
├── config/
├── installskill/       # install-skill コマンド
├── iaas/               # IaaS サービス（CLI 上ではプレフィックスなし）
│   ├── server/
│   ├── disk/
│   └── ...
├── generated/          # API client から生成する低レベルコマンド
│   ├── objectstorage/
│   ├── apprun/
│   └── ...
└── operations/         # service layer の高レベルコマンドラッパー
    ├── objectstorage/
    ├── apprun/
    └── ...
```

## 移行計画

1. フェーズ 1: トップレベル help の改善
   - トップレベルの help（`usacloud --help`）に IaaS リソースを直接表示するか、IaaS カテゴリとして一覧表示する。
   - `usacloud server list` を第一級の推奨形として位置づける。

3. フェーズ 3: 新サービスの追加
   - 全サービスの低レベルコマンドを `usacloud <service>-api <resource> <action>` として生成する。
   - 複合操作が必要なサービスでは、`sacloud-sdk-go` に service layer を実装する。
   - service layer の薄い CLI wrapper を `usacloud <service> <operation>` として追加する。
   - AGENTS.md を整備し、AI が低レベルコマンドを反復可能に生成できるようにする。

## 実装上の注意

- IaaS リソースはルート直下にも登録し、トップレベル help から発見しやすくする。
- ルート直下の IaaS リソースは Hidden を解除し、`iaas` 配下の同じコマンドも維持する。ただし help のカスタムテンプレートで IaaS カテゴリにまとめ、同じコマンドを二重表示しない。
- それ以外のサービスは、低レベルコマンドを `<service>-api` に登録する。
- 高レベルコマンドがあるサービスだけ、`<service>` に operation を登録する。
- カテゴリ表示（Computing, Storage, Networking 等）は IaaS のカテゴリとして維持する。
- 新サービスは独立したカテゴリとして help に表示する。
- `iaas` サブコマンドは非推奨とせず、引き続き利用できるようにする。
- `config`、`rest`、`completion`、`version`、`update-self`、`install-skill`、`iaas` はルート予約語とする。`<service>` と `<service>-api` を含む名前の衝突を起動時またはテストで検出する。

## 関連ファイル

- `design/v2/01-preparation/openapi-sdk-generation.md`
- `design/v2/02-major-version-upgrade/skr-rename.md`
- `design/how_to_add_new_command.md`
