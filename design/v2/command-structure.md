# コマンド体系の再整理

## 現状

現在の CLI（コマンド名は `skr` に一本化予定）は、サービス単位のトップレベルコマンドを持つ構成になっている。

```
skr [global options] <command> <sub-command> [options] [arguments] [flags]
```

トップレベルの Available Commands には以下が含まれる。

- `iaas` - SubCommands for IaaS
- `web-accelerator` - SubCommands for WebAccelerator
- `config`, `rest`, `completion` など

IaaS リソースには、以下の 2 通りの方法でアクセスできる。

```bash
skr iaas server list
skr server list
```

両者は同じ動作をする。つまり、`iaas` はサービスを明示するためのオプショナルなサブコマンドであり、IaaS リソースはルート直下にも登録されている。

ただし、トップレベルの help（`skr --help`）には `iaas` や `web-accelerator` のみが表示され、`server` などの個別リソースは表示されない。これが初学者や AI にとっての障害になっている。

## 問題

- `iaas` プレフィックスが冗長。短いコマンド名にした後も余計な階層は避けたい。
- IaaS はさくらのクラウドの中核サービスであり、デフォルトサービスとして扱うのが自然。
- 他サービス（Web Accelerator 等）との対称性がとれていない。

## 目指すコマンド体系

### 基本方針

- **IaaS リソースは `skr <resource> <action>` を第一級の推奨形とする**
- **`skr iaas <resource> <action>` も引き続き動作させる**
- **IaaS 以外のサービスはサービスプレフィックスを持つ**: `skr <service> <resource> <action>`

### 例

```bash
# IaaS（プレフィックスなし）
skr server list
skr disk create
skr switch read

# Web Accelerator
skr web-accelerator site list

# 将来追加されるサービス
skr object-storage bucket list
skr iam user list
skr apigw service list
skr apprun application list
skr apprun-dedicated cluster list
skr simple-notification group list
skr simplemq queue list
skr workflows workflow list
```

## なぜ `skr <resource> <action>` を推奨形とするか

- `skr` はもともと IaaS 中心の CLI として設計されている。
- 既存ユーザーにとって `skr server list` は馴染み深く、`iaas` プレフィックスは冗長に感じられる。
- IaaS 以外のサービスは相対的に少なく、プレフィックスを持つ形でも負担は小さい。
- コマンド名を短く保ちたい。
- トップレベルの help に IaaS リソースを直接表示することで、発見しやすくなる。

## サービスの分類

sacloud-sdk-go/api 配下のサービスを大まかに分類すると以下のようになる。

### 既存対応

- `iaas`: IaaS リソース（skr の主要リソース）
- `webaccel`: Web Accelerator

### 今後の追加候補（例）

- `object-storage`: オブジェクトストレージ
- `iam`: IAM
- `apigw`: API Gateway
- `apprun`: AppRun
- `simple-notification`: シンプル通知
- `simplemq`: SimpleMQ
- `workflows`: Workflows
- その他（addon, apprun-dedicated, cloudhsm, dedicated-storage, eventbus, kms, monitoring-suite, nosql, secretmanager, security-control, service-endpoint-gateway 等）

## コマンド体系のパターン

sacloud-sdk-go の実際の構造を反映して、以下のようなパターンを想定する。

### 1. 基本的なパターン

```
skr <service> <resource> <action> [args...] [flags]
```

### 2. サービス名とリソース名の関係

#### パターン A: サービス内に複数のリソースがある

object-storage の例:

```bash
skr object-storage account read
skr object-storage bucket list
skr object-storage permission create
skr object-storage site list
```

iam の例:

```bash
skr iam user list
skr iam group create
skr iam project read
skr iam service-principal list
```

#### パターン B: サービス名と代表リソース名が重複・類似する

workflows の例（案）:

```bash
# 案 1: サービス名を単数形にする
skr workflow list
skr workflow execution list
skr workflow subscription read

# 案 2: サービス名を複数形のまま、代表リソース名を省略する
skr workflows list
skr workflows execution list
skr workflows subscription read

# 案 3: 一貫性を重視して省略しない
skr workflows workflow list
skr workflows execution list
skr workflows subscription read
```

ここでは案を提示し、各サービス導入時に検討する。

### 3. 子リソースの扱い

workflows の `Execution` は `Workflow` の子リソース。Read には `workflowID` と `executionID` の両方が必要。

```bash
# 案 1: フラット（引数で親子を表現）
skr workflows execution read <workflow-id> <execution-id>

# 案 2: 親子階層を表現
skr workflows workflow execution read <workflow-id> <execution-id>
```

原則として、子リソースの操作が独立していれば案 1 でよい。親子関係が強く、かつ入れ子が深い場合は案 2 を検討する。

### 4. アクション名のマッピング

SDK のメソッド名から skr のサブコマンド名を決定する。

| SDK メソッド | skr アクション |
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

- 推奨形は `skr <resource> <action>`（IaaS）、`skr <service> <resource> <action>`（IaaS 以外）。
- `skr iaas <resource> <action>` も引き続き動作させる。
- サービス名は sacloud-sdk-go のディレクトリ名を基本とし、必要に応じてハイフン区切りにする（例: `object-storage`）。
- リソース名は SDK の API ファイル名・パッケージ名を基本とし、snake_case は kebab-case に変換する。

## 内部パッケージ構成

CLI 上のコマンド体系と内部パッケージ構成は必ずしも一致させる必要はないが、整理のためサービス単位に分離する。

```
pkg/commands/
├── root/
├── config/
├── installskill/       # install-skill コマンド
├── iaas/               # IaaS サービス（CLI 上ではプレフィックスなし）
│   ├── server/
│   ├── disk/
│   └── ...
├── webaccel/           # Web Accelerator サービス
│   └── webaccelerator/
└── ...                 # 将来のサービス
```

## 移行計画

1. **フェーズ 1: トップレベル help の改善**
   - トップレベルの help（`skr --help`）に IaaS リソースを直接表示するか、IaaS カテゴリとして一覧表示する。
   - `skr server list` を第一級の推奨形として位置づける。

2. **フェーズ 2: ドキュメント・Skill ファイルの更新**
   - help テキスト、Skill ファイル、ドキュメントの例を `skr server list` 主体にする。
   - `skr iaas server list` も動作することを明記。

3. **フェーズ 3: 新サービスの追加**
   - 新サービスは `pkg/commands/<service>/` 配下に追加。
   - コマンド体系は `skr <service> <resource> <action>` に統一。

## 実装上の注意

- IaaS リソースはルート直下にも登録し、トップレベル help から発見しやすくする。
- それ以外のサービスはサービス名のサブコマンド下に登録する。
- カテゴリ表示（Computing, Storage, Networking 等）は IaaS のカテゴリとして維持する。
- 新サービスは独立したカテゴリとして help に表示する。
- `iaas` サブコマンドは非推奨とせず、引き続き利用できるようにする。

## 関連ファイル

- `design/v2/openapi-sdk-generation.md`
- `design/v2/skr-rename.md`
- `design/how_to_add_new_command.md`
