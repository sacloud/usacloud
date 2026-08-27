# `install-skill` コマンド

## 目的

AI コーディングエージェントが `usacloud` を使いこなせるよう、Claude Code などが自動的に読み込む Skill ファイルを生成・配置する。

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

`--path` は生成先の基準ディレクトリを指定する。この例では
`./docs/.claude/skills/skr/SKILL.md` に生成する。指定しない場合はカレント
ディレクトリを基準とする。

生成先にファイルが存在する場合はエラーにし、`--force` が指定された場合のみ
skr が管理するファイルを上書きする。途中で失敗して既存ファイルを壊さないよう、
一時ディレクトリへの生成後に置き換える。

## 生成される Skill ファイルの構成

生成先: `.claude/skills/skr/SKILL.md`

含める内容：

- usacloud とはなにか?
- コマンド体系
   - IaaS: `usacloud <resource> <action> [options]`（推奨形）
   - IaaS 互換: `usacloud iaas <resource> <action> [options]`
   - IaaS 以外: `usacloud <service> <resource> <action> [options]`
- 共通オプション
   - ゾーン指定は `--zone` で行う。`--zone=all` で全ゾーン。
   - 破壊的操作は確認を求められる。非対話的には `--assumeyes`（`-y`）。
   - 出力形式は `--output-type=json`（短縮形は `-o json`、エイリアスは `--out=json`）で機械可読にする。
- サブコマンドについては別ファイルで help の出力結果を入れておく
- マニュアル/サービス概要ページや料金プランページへのリンクなどもあるとベター

## 実装方針

### Skill ファイルのテンプレート

プロジェクト内に `includes/skill/SKILL.md` のようなテンプレートを置き、コマンド実行時にコピーする。
自動生成した方がよいものは自動生成する。

### 実装時期

まず usacloud コマンドを使い易く整理した上で、やったほうが良い。
