# Usacloud v2 Design

Usacloud（CLI 名は `skr`）の次の段階として、AI を第一級の利用者とし、IaaS 以外の各種サービスにも対応しつつ拡張性を持たせるための設計・実装計画をまとめる。

## 背景 / Background

- CLI の利用者が人間だけでなく、AI コーディングエージェント（Claude Code, GitHub Copilot など）へと広がっている。
- さくらのクラウドのサービスが IaaS 以外にも増えており（Web Accelerator / SimpleMQ / AppRun 等）、将来さらに増えることが予想される。
- 既存の help テキストは、引数やフラグにどのような値を入れたらよいかが分かりにくい箇所が多い。
- コマンド名 `usacloud` が長く、今後 CLI 名は `skr` とする。
- これらを効率的に進めるため、AGENTS.md の整備や sacloud-sdk-go / OpenAPI からの半自動生成を活用したい。

## Goal

- AI コーディングエージェントが `skr` を使いこなせる状態にする。
    - AI が使い易い状態はこれすなわち人間にとっても理解しやすい状態になるものである。
- IaaS リソースで `skr <resource> <action>` を第一級の推奨形としつつ、IaaS 以外のサービス（Web Accelerator, Object Storage, IAM, API Gateway 等、sacloud-sdk-go/api 配下）への対応を拡張しやすいコマンド体系にする。
- 既存の help テキストを改善し、`skr` コマンドで引数やフラグに入れる値が明確になるようにする。
- 新サービス・新リソース追加時の工数を半自動化で削減する。
- CLI 名を `skr` に一本化し、コマンド体系の互換性を維持する。

## Non-Goal

- MCP サーバーの実装（現段階ではメンテナンスコストに見合わない。skill の生成と `--help` の充実が最近のトレンドである）。
- JSON Schema 出力機能などはつけない
- Structured Error の実装
- `--output=json` 自体の大きな変更（現状で動作しているため）。
- サブコマンドの完全な自動生成（人間によるレビューは必須とするし、使い勝手向上のためのコマンド追加なども行いたい）。

## 計画一覧 / Index

| # | 計画 | ファイル | 概要 |
|---|---|---|---|
| 1 | [AGENTS.md 整備](agents.md) | `agents.md` | AI エージェントがプロジェクトを理解し、help text や Skill を生成しやすくするためのガイドライン整備。 |
| 2 | [`install-skill` コマンド](install-skill.md) | `install-skill.md` | `skr install-skill` で `.claude/skills/skr/SKILL.md` を生成。 |
| 3 | [Help text 改善](help-text-improvement.md) | `help-text-improvement.md` | 引数・フラグに入れる値が明確になるよう、例や選択肢を help に追加。 |
| 4 | [コマンド体系の再整理](command-structure.md) | `command-structure.md` | IaaS を `skr <resource> <action>` の推奨形とし、`skr iaas ...` も維持。IaaS 以外は `skr <service> <resource> <action>` 体系で追加する。 |
| 5 | [OpenAPI / SDK からの半自動生成](openapi-sdk-generation.md) | `openapi-sdk-generation.md` | sacloud-sdk-go と OpenAPI 仕様からコマンド定義・help text・Skill ファイルを半自動生成する仕組み。 |
| 6 | [拡張性の土台](extensibility.md) | `extensibility.md` | コアの抽象化とコード生成により、将来のサービス追加を容易にする。 |
| 7 | [`skr` コマンド名の追加](skr-rename.md) | `skr-rename.md` | CLI 名を `skr` に一本化し、コマンド体系の互換性を維持する。 |

## トラッキング

- 本計画に関連する Issue や Pull Request は随時追記する。
