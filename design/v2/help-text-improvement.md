# Help text 改善

## 現状の問題

`skr <resource> <action> --help` を見ても、以下のような情報が不足している箇所が多い。

- 引数にどのような値を入れたらよいか分からない。
- フラグに入れる値の形式や選択肢が不明確。
- コマンドの使用例がない、または少ない。
- skr 固有の慣習（ゾーン指定方法など）が説明されていない。

これは人間だけでなく、AI コーディングエージェントにとっても大きな障害になる。

## 改善方針

### 1. 引数の説明を具体化

現在:
```
Usage:
  skr server shutdown [flags]
```

改善後:
```
Usage:
  skr server shutdown <name or id or tag> [flags]
```

### 2. フラグの説明に具体例を追加

現在:
```
Flags:
  --zone string
  --output string
```

改善後:
```
Flags:
  --zone string    対象ゾーン (例: is1a, is1b, tk1a, all)
  --output string  出力形式 (json, yaml, table)
```

### 3. 使用例を追加

Cobra の `Example` フィールドを充実させる。

```go
var shutdownCommand = &core.Command{
    Name: "shutdown",
    Example: strings.TrimSpace(`
# 名前に "example-" を含むサーバを停止
skr server shutdown "example-" --zone=is1a

# ID を指定して停止
skr server shutdown 123456789012 --zone=is1a
`),
}
```

### 4. enum 値の列挙

選択肢のあるフラグは、help に列挙する。

```
Flags:
  --output string  出力形式 (json, yaml, table; default: table)
```

`options=...` タグで定義済みの値を help テキストに自動挿入できると理想。

### 5. 入門ガイドの追加

```
skr help getting-started
skr help examples
```

のような統一的なガイドを提供する。

## 実装方法

### 段階的アプローチ

1. よく使われるコマンド（`server list`, `server shutdown`, `disk create` 等）から改善を始める。
2. コード生成部（`tools/clitag` 等）を修正し、struct tag から help テキストの補足を生成しやすくする。
3. OpenAPI / SDK からの半自動生成と連携し、説明文や例を補完する。

### 自動生成との連携

sacloud-sdk-go のコメントや OpenAPI の `description` を元に、help テキストの下書きを生成する。
生成後は人間がレビュー・調整してコミットする。

## 成功指標

- `skr <resource> <action> --help` を見ただけで、引数とフラグに入れる値が想像できる。
- AI エージェントが help を読んで、初見のコマンドでも正しく実行できる。

## 関連ファイル

- `design/v2/openapi-sdk-generation.md`
- `design/how_to_add_new_command.md`
