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
- Terraform などの宣言的管理に馴染まない一時的な運用操作を、安全で意図の明確な
  action command として提供する。
- 新サービス・新リソース追加時の工数を半自動化で削減する。
- CLI 名を `skr` に一本化し、コマンド体系の互換性を維持する。

## Non-Goal

- MCP サーバーの実装（現段階ではメンテナンスコストに見合わない)
- JSON Schema 出力機能などはつけない。
- Structured Error の実装。

## 実施フェーズ

ディレクトリを実施タイミングで分ける。

```text
design/v2/
├── README.md
├── 01-preparation/             # メジャーバージョン更新前に順次実施可能
│   ├── agents.md
│   ├── command-structure.md
│   ├── help-text-improvement.md
│   ├── openapi-sdk-generation.md
│   └── install-skill.md
└── 02-major-version-upgrade/   # 同じリリースで一括して実施
    ├── skr-rename.md
    └── update-model.md
```

### Phase 1: 先行整備

互換性を維持したまま順次実施できる。Phase 2 の完了を待つ必要はない。

| 計画 | ファイル | 概要 |
|---|---|---|
| [AGENTS.md 整備](01-preparation/agents.md) | `01-preparation/agents.md` | AI エージェントがプロジェクトを理解し、help text や Skill を生成しやすくするためのガイドライン整備。 |
| [コマンド体系の再整理](01-preparation/command-structure.md) | `01-preparation/command-structure.md` | 今後のコマンド体系を整理する。 |
| [Help text 改善](01-preparation/help-text-improvement.md) | `01-preparation/help-text-improvement.md` | 引数・フラグに入れる値が明確になるよう、例や選択肢を help に追加。 |
| [OpenAPI / SDK からの半自動生成](01-preparation/openapi-sdk-generation.md) | `01-preparation/openapi-sdk-generation.md` | 各サービスの API クライアントと OpenAPI 仕様からコマンド定義・help text・Skill ファイルを半自動生成する仕組み。 |
| [`install-skill` コマンド](01-preparation/install-skill.md) | `01-preparation/install-skill.md` | `skr install-skill` で `.claude/skills/skr/SKILL.md` を生成。 |

`install-skill` は help text の改善後に実施する。IaaS の help text 改善は、IaaS API
ドキュメントの改善を前提とする。計画間に重複はあるが、Phase 1 内では準備が整った
ものから進めてよい。

### Phase 2: メジャーバージョン更新

次の 2 項目は別々にリリースせず、`usacloud` から `skr` へ切り替える同じ
メジャーバージョン更新で一括して実施する。

| 計画 | ファイル | 概要 |
|---|---|---|
| [`skr` コマンド名への改名](02-major-version-upgrade/skr-rename.md) | `02-major-version-upgrade/skr-rename.md` | CLI 名を `skr` に一本化し、移行ガイドを提供する。 |
| [更新モデル](02-major-version-upgrade/update-model.md) | `02-major-version-upgrade/update-model.md` | 非互換なフラグ変更を含め、部分更新、値の削除、list 操作、nested field、PUT API の read-modify-write を統一する。 |

更新モデルは曖昧な既存フラグを互換性の対象外とするため、先行してリリースしない。
`skr` への改名を互換性境界とし、新しいコマンド名と更新セマンティクスを同時に
導入する。

## サブコマンドの拡張方式について

`usacloud apprun` などのサブリソースの実装についての議論。

* grpc で通信する
* `git` のようにサブコマンドは exec する形にする

などの方法も取りうるが、実際にはそこまでの複雑なアーキテクチャにする必然性は薄い(本当に?)。

シングルバイナリで進めて、それで無理そうならアーキテクチャを変更するという形が良さそう。
