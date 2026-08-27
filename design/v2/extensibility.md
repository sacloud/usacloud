# 拡張性の土台

## 目的

- IaaS 以外のサービス（Web Accelerator, Object Storage, IAM, Workflows 等、sacloud-sdk-go/api 配下）を追加しやすくする。
- 将来の新サービス追加時に、コア部分の変更を最小限に抑える。
- コード生成と連携して、新サービスの追加を容易にする。

## 目標とする設計

### 1. コアの抽象化

`pkg/core` において、以下を interface として定義する。

- `Resource`：リソース定義
- `Command`：サブコマンド定義
- `Parameter`：パラメータ
- `OutputFormatter`：出力形式

各サービスはこれらの interface を満たす実装を提供するだけで、CLI に登録できる。

### 2. サービス登録の標準化

```go
// サービスからリソース一覧を取得
var Services = []core.Service{
    iaas.Service,        // CLI 上ではプレフィックスなし/iaas 互換エイリアス
    webaccel.Service,
    // objectstorage.Service,
    // iam.Service,
    // apigw.Service,
    // apprun.Service,
    // apprunDedicated.Service,
    // simpleNotification.Service,
    // simplemq.Service,
    // workflows.Service,
    // その他 sacloud-sdk-go/api 配下のサービス
}
```

各サービスは自分のリソースとコマンドを自己登録する。

### 3. 出力フォーマッタの分離

```
pkg/output/
├── table.go
├── json.go
├── yaml.go
└── interface.go
```

出力形式をサービスに依存しない形で切り替えられるようにする。

## コード生成との連携

OpenAPI/SDK からの半自動生成は、拡張性の土台の上で動作する。

- 生成ツールが `core.Resource` や `core.Command` を満たすコードを出力。
- サービス登録ファイルに新サービスを 1 行追加するだけで有効化。

## 段階的な導入

1. **フェーズ 1: 現在の core パッケージを整理**
   - 既存の `pkg/core` を見直し、サービス横断で使える抽象を明確にする。

2. **フェーズ 2: コマンド体系の再整理**
   - IaaS リソースを `pkg/commands/iaas/` に集約しつつ、CLI 上では `skr <resource> <action>` を推奨形とする。
   - Web Accelerator を `pkg/commands/webaccel/` として統合。

3. **フェーズ 3: 新サービス追加の雛形作成**
   - 新サービス追加時のテンプレート（resource.go, command.go, テスト）を提供。
   - 生成ツールとの連携。

## 関連ファイル

- `design/v2/command-structure.md`
- `design/v2/openapi-sdk-generation.md`
- `design/how_to_add_new_command.md`
