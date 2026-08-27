# OpenAPI / SDK からの半自動生成

## 目的

sacloud-sdk-go のソースコードと OpenAPI 仕様を読み取り、skr の以下を半自動的に生成・整備する。

- コマンド定義（リソース、アクション、フラグ）
- help テキスト
- Skill ファイル
- 出力カラム定義

これにより、新サービス追加時の工数を削減し、help text や Skill ファイルの網羅性・正確性を向上させる。

## 情報源

### 1. sacloud-sdk-go

- Go の型定義（リクエスト/レスポンス struct）
- サービスメソッドのシグネチャ
- フィールドの型・タグ・コメント
- API ファイル名（リソース名の導出に使用）

### 2. OpenAPI 仕様

- エンドポイントと HTTP メソッド
- パス/クエリ/ボディパラメータ
- レスポンススキーマ
- 説明文（`description`, `summary`）

### 3. 既存の skr 定義

- `pkg/commands/iaas/<resource>/resource.go`（CLI 上では `iaas` プレフィックスなし/互換エイリアス）
- `pkg/commands/iaas/<resource>/<action>.go`
- `pkg/commands/<service>/<resource>/<action>.go`（IaaS 以外）
- `pkg/vdef/definitions.go` の値定義

## SDK からの命名導出

sacloud-sdk-go の構造から skr のコマンド名を導出するルール。

### サービス名

- SDK ディレクトリ名を基本とする。
- アンダースコアはハイフンに変換する（例: `object_storage` → `object-storage`）。
- 単複は SDK のディレクトリ名をそのまま用いる（例: `workflows`, `simplemq`）。

### リソース名

- SDK 内の API ファイル名（`*_api.go` や `*.go`）から導出する。
- ファイル名が `workflow.go`, `execution.go` ならリソース名は `workflow`, `execution`。
- アンダースコアはハイフンに変換する（例: `site_status.go` → `site-status`）。
- サービス名とリソース名が重複・類似する場合（例: `workflows/workflow`）は、各サービス導入時に省略 or 単数形サービス名を検討する。

### アクション名

- SDK のメソッド名から導出する。

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
| その他 | 必要に応じて定義 |

### 子リソースの判定

- メソッドの引数に親リソースの ID（例: `workflowID`）が含まれていれば子リソースとみなす。
- 子リソースは原則フラットに扱い、必要に応じて親子階層を検討する。

## 生成フロー

```
OpenAPI 定義
    ↓
[パーサー] API エンドポイント・パラメータ・レスポンス構造を抽出
    ↓
sacloud-sdk-go の Go 型定義とマッピング
    ↓
[ジェネレータ] skr 用の以下を生成
  - resource.go
  - <action>.go
  - zz_*_gen.go
  - help text 補完用データ
  - Skill ファイル用データ
    ↓
[AI 補完] 説明文・使用例・注意書きを生成
    ↓
[人間レビュー] 内容を確認・修正
    ↓
コミット
```

## 生成対象の優先度

1. **リソース・サブコマンド定義**（resource.go, zz_*_gen.go）
2. **フラグ定義**（struct tag からの変換）
3. **help text の下書き**（description, example）
4. **Skill ファイルの下書き**
5. **出力カラム定義**

## 人間レビューが必須な部分

完全自動化は目指さず、以下は人間が必ず確認する。

- 破壊的操作に関する注意書き
- ビジネスロジック・副作用のある処理
- 人間にとって自然な日本語/英語の説明文
- 実際の API 挙動と生成コードの整合性

## ツール構成

`tools/` 配下に新しい生成ツールを追加する。

```
tools/
├── clitag/           # 既存: struct tag パーサー
├── gen-commands/     # 新規: OpenAPI/SDK からコマンド定義を生成
└── gen-skill/        # 新規: Skill ファイルを生成
```

`make gen` で既存の生成と新しい生成を一括実行できるようにする。

## 段階的な導入

1. **フェーズ 1: 小さな実験**
   - 既存の 1 リソース（例: `server`）を対象に、OpenAPI/SDK から help text 下書きを生成してみる。
   - 品質を確認し、生成ロジックを調整する。

2. **フェーズ 2: 既存リソースへの展開**
   - 既存の IaaS リソースの help text を段階的に生成・置き換え。

3. **フェーズ 3: 新サービスへの適用**
   - 新サービス追加時に、生成ツールを使って一括でコマンド定義・Skill ファイルを作成。

## 成功指標

- 新サービス追加時のボイラープレート実装が大幅に減る。
- help text の網羅性が向上する。
- Skill ファイルの更新が生成ツールでカバーできる範囲が広がる。

## 関連ファイル

- `design/v2/agents.md`
- `design/v2/install-skill.md`
- `design/v2/help-text-improvement.md`
- `design/v2/command-structure.md`
- `design/how_to_add_new_command.md`
