# AGENTS.md 整備

## 目的

AI コーディングエージェント（Claude Code, GitHub Copilot, OpenCode 等）が `skr`（旧称 usacloud）プロジェクトを短時間で理解し、適切に変更・生成できる状態にする。

特に以下を目指す。

- プロジェクトの構成、ビルド・テスト方法、命名規則を説明する。
- help text や Skill ファイルの生成時に、AI が品質を担保できる指針を示す。
- 新規コントリビューターよりも、継続的に作業する AI エージェントを想定した内容にする。

## 配置場所

プロジェクトルートの `AGENTS.md` を整備する。

既存の `README.md` は人間向けのクイックスタートに留め、`AGENTS.md` には AI 向けの詳細なコンテキストを集約する。

## 含めるべき内容

### 1. プロジェクト概要

- `skr`（旧称 usacloud）とは何か（さくらのクラウド公式 CLI）。
- リポジトリの主要ディレクトリ構成。
- Go バージョン、主要な依存ライブラリ（sacloud-sdk-go, libsacloud, cobra 等）。

### 2. ビルド・テスト・開発フロー

```bash
make          # コード生成 + ビルド
make test     # テスト実行
make gen      # コード生成
make lint     # lint 実行
```

- コード生成の仕組み（`tools/` 配下）についての概要。
- 生成ファイル（`zz_*_gen.go`）は手動で編集しない原則。

### 3. コマンド体系と命名規則

#### コマンド体系

- 現在: `skr iaas <resource> <action>`（IaaS）、`skr <service> <resource> <action>`（IaaS 以外）
- 目指す: `skr <resource> <action>`（IaaS）、`skr <service> <resource> <action>`（IaaS 以外）
- `skr iaas <resource> <action>` は維持

#### 命名規則

- サービス名: sacloud-sdk-go のディレクトリ名を基本とし、アンダースコアはハイフンに変換（例: `object_storage` → `object-storage`）。
- リソース名: SDK 内の API ファイル名から導出し、アンダースコアはハイフンに変換（例: `site_status.go` → `site-status`）。
- アクション名: SDK のメソッド名から導出（Create → create, List → list, Read → read 等）。
- フラグのカテゴリ分け（`cli:,category=...`）。

### 4. コード生成の仕組み

- `tools/clitag/` などで struct tag から CLI 定義を生成していること。
- 新リソース追加時の流れ（`design/how_to_add_new_command.md` への参照）。
- 今後 OpenAPI / SDK からの半自動生成を導入する予定があること。

### 5. AI 向けの作業指針

- help text を追加・修正するときの注意点（具体例を含める、enum 値を列挙する等）。
- Skill ファイル生成時のトーンや含めるべき情報。
- 破壊的操作に関する注意書きの書き方。
- テストを追加・変更する際の方針。

### 6. ブランチ運用・コミット方針

- main ブランチへの直接 push は行わない。
- 作業用ブランチを切って Pull Request を作成する。
- DCO（`Signed-off-by:`）対応。

## 更新方針

- プロジェクト構造やビルド手順が変わったら必ず更新する。
- 新サービス追加時のパターンが確立したら、該当セクションを追加する。
- AI が作業して気づいた点は、都度 `AGENTS.md` に還元していく。

## 関連ファイル

- `README.md`
- `design/how_to_add_new_command.md`
- `design/v2/openapi-sdk-generation.md`
