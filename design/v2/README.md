# Usacloud v2 Design

Usacloud（CLI 名は `skr`）の次の段階として、AI を第一級の利用者とし、IaaS 以外の各種サービスにも対応しつつ拡張性を持たせるための設計・実装計画をまとめる。

## 背景

- CLI の利用者が人間だけでなく、AI コーディングエージェント（Claude Code, GitHub Copilot など）へと広がっている。
- さくらのクラウドのサービスが IaaS 以外にも増えており（Web Accelerator / SimpleMQ / AppRun 等）、将来さらに増えることが予想される。
- 既存の help テキストは、引数やフラグにどのような値を入れたらよいかが分かりにくい箇所が多い。
- コマンド名 `usacloud` が長く、今後 CLI 名は `skr` とする。
- これらを効率的に進めるため、AGENTS.md の整備や各サービスの API クライアント / OpenAPI からの半自動生成を活用したい。

## Goal

- AI コーディングエージェントが `usacloud` を使いこなせる状態にする。
    - AI が使いやすければ、人間にとっても理解しやすい状態になる。
- IaaS リソースで `usacloud <resource> <action>` を第一級の推奨形としつつ、IaaS 以外のサービス（Web Accelerator, Object Storage, IAM, API Gateway 等）への対応を拡張しやすいコマンド体系にする。
- 既存の help テキストを改善し、`usacloud` の引数やフラグへ指定する値を明確にする。
- 新サービス・新リソース追加時の工数を半自動化で削減する。
- CLI 名を `skr` に一本化し、コマンド体系の互換性を維持する。

## Non-Goal

- MCP サーバーの実装（現段階ではメンテナンスコストに見合わない)
- JSON Schema 出力機能などはつけない。
- Structured Error の実装。

## 計画一覧 / Index

| # | 計画 | ファイル | 概要 |
|---|---|---|---|
| 1 | [AGENTS.md 整備](agents.md) | `agents.md` | AI エージェントがプロジェクトを理解し、help text や Skill を生成しやすくするためのガイドライン整備。 |
| 2 | [コマンド体系の再整理](command-structure.md) | `command-structure.md` | 今後のコマンド体系を整理する。 |
| 3 | [Help text 改善](help-text-improvement.md) | `help-text-improvement.md` | 引数・フラグに入れる値が明確になるよう、例や選択肢を help に追加。 |
| 4 | [OpenAPI / SDK からの半自動生成](openapi-sdk-generation.md) | `openapi-sdk-generation.md` | 各サービスの API クライアントと OpenAPI 仕様からコマンド定義・help text・Skill ファイルを半自動生成する仕組み。 |
| 5 | [`install-skill` コマンド](install-skill.md) | `install-skill.md` | `skr install-skill` で `.claude/skills/skr/SKILL.md` を生成。 |
| 6 | [`skr` コマンド名への改名](skr-rename.md) | `skr-rename.md` | CLI 名を `skr` に一本化し、コマンド体系の互換性を維持する。 |

- 計画の各ステップは少なからず重複する内容をはらんでいる。
- Note: install-skill はある程度 help text が改善されてからでないと、メリットが薄い。
- Note: Help text の改善は iaas がメインとなるが、iaas の API ドキュメンテーション改善が先。

## サブコマンドの拡張方式について

`usacloud apprun` などのサブリソースの実装についての議論。

* grpc で通信する
* `git` のようにサブコマンドは exec する形にする

などの方法も取りうるが、実際にはそこまでの複雑なアーキテクチャにする必然性は薄い(本当に?)。

シングルバイナリで進めて、それで無理そうならアーキテクチャを変更するという形が良さそう。
