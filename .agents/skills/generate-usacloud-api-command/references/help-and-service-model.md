# サービスモデルと help の設計

API method を 1:1 で command に写すだけでは、利用者は何を作るべきか判断できません。

用途と意味の根拠として利用できる情報源は、`manual.sakura.ad.jp` のドキュメントと、
対象バージョンの `sacloud-sdk-go` repository/module 内にある README、Go doc、
example だけです。検索結果の要約、第三者の記事、Terraform provider の説明、過去の
記憶、一般的な製品知識は根拠にしてはいけません。SDK の request/response 型は field、
required、enum など機械的な入力契約の根拠に限り、用途や意味を型名から推測しません。

実装前に許可された情報源を読み、次をサービスモデルとして整理します。

1. サービスが解決する問題と、利用者が選ぶ典型的な場面
2. 各 resource が表す概念（「何を」「いつ」「何を契機に」など）
3. resource 間の参照関係、必須の作成順、削除順
4. 代替関係と併用関係（schedule と trigger のどちらを選ぶか、併用できるか）
5. 実行保証、最小間隔、scope など、誤用につながる重要な制約
6. 初回利用者が最短で動かすための操作順

名称や field から用途を推測しません。各説明について、説明文、出典 URL または module
内 path、根拠となる原文を一時的な根拠表へ記録します。根拠表は manifest と同様に
session artifact へ置き、コミットしません。原文から直接確認できない説明は help に
含めません。許可された資料間で意味が矛盾する、または関係を説明できない場合は実装を
止めて確認します。

| 説明 | 出典 | 根拠となる原文 |
| --- | --- | --- |
| `<help に記載する一文>` | `<manual URL または SDK 内 path@version>` | `<該当箇所の引用>` |

## help の階層

短い `Usage` は command 一覧で意味が伝わる動詞句にし、`SubCommands for ...` のような
構造の言い換えにしません。長い説明は `LongUsage`、実行例は `Example` に分けます。
各 command にも一行の `Usage` を設定し、子 resource の command 一覧に空欄を残しません。

親 `*-api` resource の help には次を含めます。

- サービスの目的と、使うべき代表的な場面
- resource を作る順序と参照関係
- 重要な実行上の制約
- 各子 resource の `create --help` へ進む例
- 公式マニュアルの URL

子 resource の名前と一行の役割は Cobra が生成する `Available Commands` に表示される
ため、`LongUsage` に同じ対応表を重複させません。`LongUsage` はサービスの概念、作成順、
制約を文章で補い、command 一覧は `Usage` を唯一の一行説明として使います。

子 resource の help には次を含めます。

- その resource が「何を」決めるものか
- 単独で効果を持つか、別 resource から参照されるか
- 前提として先に作る resource
- 主な入力概念と代表的な利用場面
- `create --help` と `list` の例

`core.Resource.Usage` は一行、`LongUsage` は複数段落にします。command の complex JSON
については `parameter-mapping.md` に従い、さらに field と入力例を command help に
記載します。

## 検証

実バイナリで親とすべての子 resource の `--help` を表示し、SDK を知らない利用者が
次の問いに回答できることをテストします。

- このサービスはいつ使うのか。
- 各 resource は何を表すのか。
- 最初に何を作るのか。
- 次にどの command の help を見ればよいか。

help test は `Short` と `Long` が異なること、resource 間の役割・作成順を示す主要語、
実行例が存在することを固定します。さらに、help の各事実記述が根拠表の許可された
情報源に対応することをレビューします。
