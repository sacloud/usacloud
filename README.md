# Usacloud

![usacloud_logo_h.png](usacloud_logo_h.png)

[`usacloud`](https://github.com/sacloud/usacloud)は[さくらのクラウド](http://cloud.sakura.ad.jp/index.html)用の公認CLIクライアントです。  

![Test Status](https://github.com/sacloud/usacloud/workflows/Tests/badge.svg)
[![Slack](https://slack.usacloud.jp/badge.svg)](https://slack.usacloud.jp/)
[![License](https://img.shields.io/github/license/sacloud/usacloud)](LICENSE.txt)
[![Version](https://img.shields.io/github/v/tag/sacloud/usacloud)](https://github.com/sacloud/usacloud/releases/latest)
![Downloads](https://img.shields.io/github/downloads/sacloud/usacloud/total)
[![Documents](https://img.shields.io/badge/docs-Documents%20%20for%20Usacloud-green)](https://docs.usacloud.jp/usacloud)

## Installation / インストール

[Documents: https://docs.usacloud.jp/usacloud/installation/start_guide](https://docs.usacloud.jp/usacloud/installation/start_guide)

### Quick Start

- [GitHub Releases](https://github.com/sacloud/usacloud/releases/latest)から自身のプラットフォーム向けのファイルをダウンロード&展開
- [さくらのクラウド ドキュメント: APIキーの新規作成・編集](https://manual.sakura.ad.jp/cloud/api/apikey.html#id3) を参照してAPIキーを作成
- `usacloud profile`コマンドでAPIキーを設定

## Usage / 基本的な使い方

コマンドは以下の書式で指定します。

    usacloud <リソース> <サブコマンド> [オプション] [対象リソースのID or 名前(部分一致) or タグ]

リソースやサブコマンド、オプションは`usacloud -h`、`usacloud <リソース名> -h`、または`usacloud <リソース名> <サブコマンド> -h`で確認できます。

#### コマンドの例

```bash
# 全ゾーンのサーバ一覧を取得
$ usacloud server list --zone=all

# 石狩第1ゾーンで名前に"example-"を含むサーバをすべてシャットダウン(オプションの位置は引数の後ろでもOK)
$ usacloud server shutdown "example-" --zone=is1a
```

### その他の使い方

Usacloud ドキュメントを参照してください。
> [Usacloud ドキュメント](https://docs.usacloud.jp/usacloud)

## License

 `usacloud` Copyright (C) 2017-2020 The Usacloud Authors.

  This project is published under [Apache 2.0 License](LICENSE.txt).
