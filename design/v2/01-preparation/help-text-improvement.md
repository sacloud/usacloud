# Help text 改善

## 現状の問題

`skr <resource> <action> --help` を見ても、以下のような情報が不足している箇所が多い。

- 引数にどのような値を入れたらよいか分からない。
- フラグに入れる値の形式や選択肢が不明確。
- コマンドの使用例がない、または少ない。
- skr 固有の慣習（ゾーン指定方法など）が説明されていない。

これは人間だけでなく、AI コーディングエージェントにとっても大きな障害になる。

## 改善方針

コマンドが出力する help text は英語で統一する。対象にはコマンドの概要、引数、
フラグ、エラーメッセージ、使用例を含める。設計文書自体は日本語でもよいが、
コード例に示す help 出力と Cobra の `Use`、`Short`、`Long`、`Example` は英語にする。

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
  --output-type string
```

改善後:
```
Flags:
  --zone string         Target zone (e.g. is1a, is1b, tk1a, all)
  --output-type string  Output format (table, json, yaml)
```

### 3. 使用例を追加

Cobra の `Example` フィールドを充実させる。

```go
var shutdownCommand = &core.Command{
    Name: "shutdown",
    Example: strings.TrimSpace(`
# Shut down a server by its exact name
skr server shutdown example-server --argument-match-mode=exact --zone=is1a

# Shut down a server by ID
skr server shutdown 123456789012 --zone=is1a
`),
}
```

### 4. enum 値の列挙

選択肢のあるフラグは、help に列挙する。

```
Flags:
  --output-type string  Output format (table, json, yaml; default: table)
```

`options=...` タグで定義済みの値を help テキストに自動挿入できると理想。

### 5. update の操作を明示

update コマンドでは、フラグが値の設定、削除、list への追加・削除、全体置換の
どれを行うか明記する。未指定フィールドは変更しないこと、ネストした object の
どの範囲が更新されるかも help に表示する。

詳細は [更新モデル](../02-major-version-upgrade/update-model.md) を参照。

## 実装方法

* これら基準を満たすような help message を生成するように規定した AGENTS.md を用意する
* 原則として sdk から生成されたものについては、上記の情報が生成できる程度に openapi 定義が充実していることが前提となる
    * 足りない場合にはまず openapi 定義を更新する
