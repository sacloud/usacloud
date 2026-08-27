# コマンド体系の再整理

## 現状

現在の CLI は、サービス単位のトップレベルコマンドを持つ構成になっている。

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

> [!NOTE]
> ただし、トップレベルの help（`usacloud --help`）には `iaas` や `web-accelerator` のみが表示され、`server` などの個別リソースは表示されない。
> これは初学者や AI にとってのわかりにくさを産んでいる可能性がある。

## 目指すコマンド体系

### 基本方針

- IaaS リソースは `usacloud <resource> <action>` を第一級の推奨形とする
- `usacloud iaas <resource> <action>` も引き続き動作させる
- IaaS 以外のサービスはサービスプレフィックスを持つ: `skr <service> <resource> <action>`

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

# Web Accelerator(実装済み)
usacloud web-accelerator site list

# 将来追加されるサービスも同様に利用可能としていく
usacloud object-storage bucket list
usacloud iam user list
usacloud apigw service list
usacloud apprun application list
usacloud apprun-dedicated cluster list
usacloud simple-notification group list
usacloud simplemq queue list
usacloud workflows workflow list
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

各サービスの API クライアントの実際の構造を反映して、以下のようなパターンを想定する。

### 基本的なパターン

```
skr <service> <resource> <action> [args...] [flags]
```

### サービス名とリソース名の関係

#### サービス内に複数のリソースがあるケース

通常はこのようなケースとなる。このようなケースでは素直に API 設計をコマンド体系に落としていけば良い。

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

#### サービス名と代表リソース名が重複・類似するケース

workflows の場合、`workflow` がトップレベルリソースになるので短くかけるようにすることも考えてもいいかもしれない。
ただ、特定のコマンドだけショートカットがあるのも混乱をうむ可能性がある。

workflows の例:

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

初期実装は一貫性を優先して案 3 とする。案 1、案 2 の短縮形が必要な場合は、
正式なコマンドを置き換えず、衝突を確認した上でエイリアスとして追加する。

### 子リソースの扱い

workflows の `Execution` は `Workflow` の子リソース。Read には `workflowID` と `executionID` の両方が必要。

```bash
# 案 1: フラット（引数で親子を表現）
skr workflows execution read <workflow-id> <execution-id>

# 案 2: 親子階層を表現
skr workflows workflow execution read <workflow-id> <execution-id>
```

原則として、子リソースの操作が独立していれば案 1 のほうが素直な実装といえる。

初期生成では案 1 を採用する。親 ID、子 ID の順序を Usage とテストで固定し、
同じ名前の子リソースがサービス内で衝突する場合のみ案 2 を採用する。

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

- 推奨形は `usacloud <resource> <action>`（IaaS）、`usacloud <service> <resource> <action>`（IaaS 以外）。
- `usacloud iaas <resource> <action>` も引き続き動作させる。
- サービス名は API クライアントのサービス名を基本とし、必要に応じてハイフン区切りにする（例: `object-storage`）。
- リソース名は SDK の API ファイル名・パッケージ名を基本とし、snake_case は kebab-case に変換する。

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
├── webaccel/           # Web Accelerator サービス
│   └── webaccelerator/
└── ...                 # 将来のサービス
```

## 移行計画

1. フェーズ 1: トップレベル help の改善
   - トップレベルの help（`usacloud --help`）に IaaS リソースを直接表示するか、IaaS カテゴリとして一覧表示する。
   - `usacloud server list` を第一級の推奨形として位置づける。

3. フェーズ 3: 新サービスの追加
   - 新サービスは `pkg/commands/<service>/` 配下に追加。
   - コマンド体系は `skr <service> <resource> <action>` に統一。
   - AGENTS.md を整備し、AI が自動的に生成できるようにする。

## 実装上の注意

- IaaS リソースはルート直下にも登録し、トップレベル help から発見しやすくする。
- ルート直下の IaaS リソースは Hidden を解除し、`iaas` 配下の同じコマンドも維持する。ただし help のカスタムテンプレートで IaaS カテゴリにまとめ、同じコマンドを二重表示しない。
- それ以外のサービスはサービス名のサブコマンド下に登録する。
- カテゴリ表示（Computing, Storage, Networking 等）は IaaS のカテゴリとして維持する。
- 新サービスは独立したカテゴリとして help に表示する。
- `iaas` サブコマンドは非推奨とせず、引き続き利用できるようにする。
- `config`、`rest`、`completion`、`version`、`update-self`、`install-skill`、`iaas` はルート予約語とする。IaaS リソースやサービス名との衝突を起動時またはテストで検出する。

## 関連ファイル

- `design/v2/openapi-sdk-generation.md`
- `design/v2/skr-rename.md`
- `design/how_to_add_new_command.md`
