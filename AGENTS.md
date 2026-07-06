# Agent Guidance for usacloud

本ドキュメントは、usacloud プロジェクトを扱う AI coding agent 向けのガイダンスです。

## プロジェクト概要

usacloud は、さくらインターネットの各種サービス（IaaS、PHY、オブジェクトストレージ、WebAccelerator など）を操作するための CLI ツールです。

- 言語はGoである。
- CLI フレームワークはspf13/cobra, spf13/pflagである。
- ビルドツールはGNU Makeである。
- コード生成は`tools/gen-commands`（Goのtext/templateベース）である。

## 主要ディレクトリ構成

```
.
├── design/                         # 設計文書
│   ├── core-architecture.md        # pkg/core アーキテクチャ
│   ├── core-command-lifecycle.md   # コマンド実行ライフサイクル
│   ├── core-resource-model.md      # リソースモデル
│   ├── code-generation-boundary.md # 生成コードと手書きコードの境界
│   ├── how_to_add_new_command.md   # コマンド追加ガイド
│   └── v1.md                       # v1.0.0 設計メモ
│
├── pkg/
│   ├── core/                       # コマンド実行フレームワーク（手書き・生成対象外）
│   ├── commands/                   # リソース・コマンド定義（主に手書き）
│   │   ├── iaas/                   # IaaS プラットフォーム
│   │   ├── webaccel/               # WebAccelerator プラットフォーム
│   │   ├── config/                 # config サブコマンド
│   │   ├── rest/                   # REST API サブコマンド
│   │   ├── root/                   # ルートコマンド
│   │   ├── completion/             # shell completion
│   │   ├── version/                # version コマンド
│   │   └── update-self/            # update-self コマンド
│   ├── services/                   # 生成されたサービス呼び出しコード
│   │   ├── iaas/
│   │   └── webaccel/
│   ├── cli/                        # CLI コンテキスト、IO、API クライアント
│   ├── cflag/                      # 共通パラメータ（フラグ）
│   ├── category/                   # リソース/コマンド/パラメータカテゴリ
│   ├── vdef/                       # 値定義（options タグで参照）
│   └── output/                     # 出力フォーマット処理
│
├── tools/                          # コード生成ツール
│   ├── gen-commands/               # コマンド・サービスコード生成
│   ├── clitag/                     # cli タグパーサー
│   └── utils/                      # 生成ユーティリティ
│
├── e2e/                            # E2E テスト
├── includes/                       # Makefile インクルード
└── GNUmakefile
```

## 設計文書へのリンク

コマンド追加・変更時は必ず以下を参照してください。

- [design/core-architecture.md](./design/core-architecture.md)
- [design/core-command-lifecycle.md](./design/core-command-lifecycle.md)
- [design/core-resource-model.md](./design/core-resource-model.md)
- [design/code-generation-boundary.md](./design/code-generation-boundary.md)
- [design/how_to_add_new_command.md](./design/how_to_add_new_command.md)

## よく使うコマンド

```sh
# ビルド
make

# コード生成
make gen

# 強制再生成（生成ファイルを全削除してから生成）
make gen-force

# テスト
make test

# フォーマット・lint
make fmt
make goimports
make lint
```

## コーディング上の注意

- 生成コード（`*_gen.go`）は原則として手で編集しない。テンプレートまたは定義元を修正して`make gen`を実行する。
- 新規リソース追加時は`pkg/commands/<platform>/<resource>/`配下に手書きファイルを配置し、`pkg/commands/<platform>/resources.go`に登録する。
- `pkg/core`はフレームワーク層である。リソース固有の知識を持ち込まない。
- Goのコメントは日本語/英語混在しており、主要な公開型にはGoDocスタイルのコメントを付けることを推奨する。

## 問い合わせ・参考

- リポジトリ: https://github.com/sacloud/usacloud。
- 既存リソースの実装が最も信頼できる参考例である。
