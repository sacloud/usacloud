# 統合テスト

usacloudの統合テストには[bats](https://github.com/sstephenson/bats)を利用しています。

## 実行

統合テストの実行には以下の環境変数の設定が必要です。

- `SAKURACLOUD_ACCESS_TOKEN`
- `SAKURACLOUD_ACCESS_TOKEN_SECRET`
- `SAKURACLOUD_ZONE`

また、Docker上ではなくローカルマシン上でテストを実行する場合は以下のコマンドが必要です。

- `jq`
- `mkisofs` or `genisoimage` or `hdiutil`

```bash
# ローカルマシン上で実行する場合(bats/jqが必要)
$ make integration-test
# 個別のテストだけ実行したい場合(ディレクトリ単位、またはファイル単位でも可)
$ test/integration/run_bats.sh test/integration/bats/対象ディレクトリor対象ファイル名

# Docker上で実行する場合(bats/jq不要)
$ make docker-integration-test    
```
    

# テストの書き方

### リソース作成時の命名規則

`run_bats.sh`では開始時/終了時にリソースのクリーンアップ(削除)処理を行います。  
対象は`usacloud-integration-test`タグが付与されたリソース、または名前に`usacloud-integration-test`を含むリソースです。

テストの中でリソースの作成を行う場合はこの命名規則に従ってください。  
これらは以下環境変数から参照可能です。(helpers.bashにて定義されています)  

- `TEST_TARGET_NAME`
- `TEST_TARGET_TAG`

また、クリーンアップ処理をスキップしたい場合は`SKIP_CLEANUP`環境変数に`1`を設定してください。

```bash
$ SKIP_CLEANUP=1 make integration-test
```

### 1) ディレクトリの作成

統合テストは`bats`配下の各ディレクトリ内にある`*.bat`ファイルに実装されています。  
`bats`配下のディレクトリはテストのカテゴリーごとに分けられています。  
テストが属するカテゴリーに対応するディレクトリが存在しない場合は新たに作成してください。  

### 2) テスト用スクリプトの作成

`bats`配下の任意のディレクトリに`*.bat`ファイルを作成します。

batファイルでは以下のように記載することでテスト用ヘルパー関数が利用可能です。

    load ${BASE_TEST_DIR}/helpers.bash
   
usacloudコマンドの実行にはヘルパー関数`usacloud_run`(標準エラー出力なし)、または`usacloud_run_with_stderr`(標準エラー出力あり)を利用してください。  
(usacloud/bin配下のusacloudバイナリが利用されます) 


### テストで利用できるライブラリ

現在は以下のライブラリをテスト内で利用可能です。

- `jq`
- `genisoimage`(利用する場合は`helpers.bash`で定義されている`MK_ISO_CMD`環境変数経由で利用してください)

ライブラリを追加したい場合は[scripts/bats.dockerfile](/scripts/bats.dockerfile)を編集してインストールするようにしてください。

### その他注意点

統合テストはmacOSとLinux上で行えるように改行コードの扱いなどで環境依存しないように注意してください。  
また、CI環境での統合テスト実施のために、テスト内の各コマンドは非対話モードにて実行してください。(usacloudの場合`--assumeyes`オプションをつけるなど)
