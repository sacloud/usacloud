# `skr` コマンド名の導入

## 背景

`usacloud` は長いコマンド名であり、頻繁に打つには少し不便。
AI エージェントに指示する場合も、短いコマンド名の方がトークン効率が良い。

今後、CLI のコマンド名は `skr` とする。

## 方針

- CLI のコマンド名は `skr` のみとする。
- `usacloud` という名前のコマンドは提供しない。
- ただし、コマンド体系（`server list`, `iaas server list` 等）は維持する。
- 既存の `usacloud` ユーザーには、移行ガイドを提供する。

## 実装方法

### 1. バイナリ名

ビルド時のバイナリ名は `skr` とする。

```bash
go build -o skr .
```

### 2. Cobra コマンド名

```go
rootCmd := &cobra.Command{
    Use:   "skr",
    Short: "CLI to manage resources on SAKURA Cloud",
}
```

### 3. 互換性

`usacloud` というコマンド名は提供しない。

ただし、以下のコマンド体系は維持する。

```bash
# IaaS（プレフィックスなし）
skr server list

# IaaS（iaas プレフィックス付き）
skr iaas server list

# IaaS 以外
skr web-accelerator site list
```

## ドキュメント・Skill ファイル

Skill ファイルや help の例はすべて `skr` で統一する。

```markdown
# skr

さくらのクラウドを操作する CLI です。

## コマンド例

```bash
# サーバ一覧の取得
skr server list --zone=all

# IaaS プレフィックス付きでも可
skr iaas server list --zone=all
```
```

## 移行計画

1. **フェーズ 1: `skr` コマンドの実装**
   - バイナリ名を `skr` とする。
   - コマンド体系は現行の usacloud と同じ構造を維持。

2. **フェーズ 2: ドキュメント・Skill ファイルの更新**
   - すべての例を `skr` に統一。
   - `usacloud` から `skr` への移行ガイドを提供。

3. **フェーズ 3: リリース**
   - リリースアーカイブ名やバイナリ名を `skr` に統一。
   - Homebrew 等のインストーラーにも `skr` として反映。

## 影響範囲

- プロジェクト名 `usacloud`（リポジトリ名等）は維持してもよい。
- CLI 上のコマンド名は `skr` に一本化する。
- 既存スクリプトは `s/usacloud/skr/` で置き換える必要がある。

## 関連ファイル

- `design/v2/install-skill.md`
- `design/v2/command-structure.md`
