# `install-skill` コマンド

## 目的

AI コーディングエージェントが `skr` を使いこなせるよう、Claude Code などが自動的に読み込む Skill ファイルを生成・配置する。

Playwright CLI や Confluence CLI と同様に、`install-skill` サブコマンドで `.claude/skills/skr/SKILL.md` を生成する。

## コマンド設計

```bash
# カレントディレクトリに Skill ファイルを生成
skr install-skill

# 上書き生成
skr install-skill --force

# 任意のディレクトリ配下に生成
skr install-skill --path ./docs
```

## 生成される Skill ファイルの構成

生成先: `.claude/skills/skr/SKILL.md`

含める内容：

1. **ツールの概要**
   - `skr` とは何か。
   - 対応サービス（IaaS, Web Accelerator, Object Storage, IAM 等、sacloud-sdk-go/api 配下のサービス）。

2. **コマンド体系**
   - IaaS: `skr <resource> <action> [options]`（推奨形）
   - IaaS 互換: `skr iaas <resource> <action> [options]`
   - IaaS 以外: `skr <service> <resource> <action> [options]`

3. **重要なルール**
   - ゾーン指定は `--zone` で行う。`--zone=all` で全ゾーン。
   - 破壊的操作は確認を求められる。非対話的には `--yes`。
   - 出力形式は `--output=json` で機械可読に。
   - 安全のため、AI は初回実行時に `--dry-run` を検討する。

4. **よく使う操作の例**
   ```markdown
   # サーバ一覧の取得
   skr server list --zone=all --output=json

   # 名前でサーバを特定して停止
   skr server shutdown "example-" --zone=is1a

   # リソースの作成
   skr disk create --name example-disk --zone=is1a
   ```

5. **サービス別のガイドへの参照**
   - `.claude/skills/skr/iaas.md`
   - `.claude/skills/skr/web-accelerator.md`
   - 今後追加されるサービスごとに生成。

## 実装方針

### コマンドの追加場所

`pkg/commands/` 配下に新規パッケージ `installskill`（または `skill`）を作成する。

ただし、`install-skill` はエンドユーザー向けのコマンドなので、`pkg/commands/root/command.go` から直接登録する形がシンプル。

### Skill ファイルのテンプレート

プロジェクト内に `includes/skill/SKILL.md` のようなテンプレートを置き、コマンド実行時にコピーする。

テンプレートは将来的にコード生成で一部を自動生成できるようにしておく。

### サービス別 Skill ファイル

将来的には、サービスごとに `SKILL.md` を分割し、`install-skill` で一括生成する。

   ```
   .claude/skills/skr/
   ├── SKILL.md              # 共通ガイド
   ├── iaas.md               # IaaS 固有
   ├── web-accelerator.md    # Web Accelerator 固有
   ├── object-storage.md     # Object Storage 固有
   ├── iam.md                # IAM 固有
   ├── apigw.md              # API Gateway 固有
   ├── apprun.md             # AppRun 固有
   ├── apprun-dedicated.md   # AppRun Dedicated 固有
   ├── simple-notification.md# Simple Notification 固有
   ├── simplemq.md           # SimpleMQ 固有
   ├── workflows.md          # Workflows 固有
   └── ...
   ```

## 将来の拡張

- OpenAPI / SDK からの半自動生成と連携し、新サービス追加時に Skill ファイルも自動更新する。

## 関連ファイル

- `design/v2/agents.md`
- `design/v2/openapi-sdk-generation.md`
- `design/v2/skr-rename.md`
