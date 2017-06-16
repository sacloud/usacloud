# サーバの作成

---

## 概要

サーバの作成は`server build`コマンドで行います。

指定するオプションにより、以下の5つのタイプのサーバが作成可能です。  
タイプにより指定可能なオプションが異なります。

- 1) パブリックアーカイブ(非Windows系)を利用してディスクを作成  
- 2) パブリックアーカイブ(Windows系)を利用してディスクを作成  
- 3) 既存のディスクを接続/マイアーカイブを利用してディスクを作成  
- 4) ブランクディスクを作成(ISOイメージなどからOSをインストール)  
- 5) ディスクレス(ISOイメージやネットワークからブート)

---

#### オプションのカテゴリ

指定できるオプションは以下のようなカテゴリに分かれています。 

- `サーバプラン`: コア数/メモリサイズ
- `ISOイメージ`: ISOイメージの挿入
- `ネットワーク`: 接続する上流ネットワークの指定やパケットフィルタの指定
- `ディスク`: 接続/作成するディスクや、コピー元のアーカイブの指定など
- `ディスクの修正`: パスワードやホスト名など、ディスクの修正方法についての指定([**注1](#notice1))
    - `OS基本設定`: ホスト名やパスワード、SSH接続時のパスワード/チャレンジレスポンス認証の無効化など
    - `ネットワーク`: IPアドレスやネットワークマスク、デフォルトゲートウェイの設定
    - `スタートアップスクリプト`: スタートアップスクリプト関連の設定
    - `SSH公開鍵`: SSH公開鍵の生成/登録/利用などの設定
- `サーバ基本設定`: サーバ名や説明、タグなど
- `その他`: キーボードのUS配列有効化やサーバ作成後の自動起動の抑制など

各カテゴリの指定可能なオプションの一覧は[コマンドリファレンス:`server build`](commands/server/#build)を参照してください。

---

#### `ディスクの修正`カテゴリのオプションについて

`ディスクの修正`カテゴリのオプションについては、作成/接続するディスクや、コピー元アーカイブにより指定可能な項目が異なります。 

**1) パブリックアーカイブ(非Windows系)を利用してディスクを作成する場合**

`ディスクの修正`カテゴリの全てのオプションが利用可能です。

**2) パブリックアーカイブ(Windows系)を利用してディスクを作成する場合**

`ディスクの修正`カテゴリのうち、以下の項目のみ利用可能です。  

- パスワード
- IPアドレス(共有セグメント以外に接続する場合)
- ネットワークマスク(共有セグメント以外に接続する場合)
- デフォルトゲートウェイ(共有セグメント以外に接続する場合)

**3) 既存のディスクやマイアーカイブを接続する場合**  
**4) ブランクディスクを接続し、ISOイメージなどからOSをインストールする場合**  
**5) ディスクレスな場合**  

`ディスクの修正`カテゴリのオプションは指定できません。

---


# 各構成でのオプション指定例

---

## 最小構成

- 共有セグメントに接続
- OSはCentOSパブリックアーカイブを利用

```console
$ usacloud server build \
        --os-type centos \
        --hostname [put-your-hostname] \
        --password [put-your-password] \
        --name [put-your-server-name] 
```

---

## スイッチ/スイッチ+ルータに接続

- 接続するスイッチのIDを指定
- NIC関連の設定を指定(IPアドレス/ネットワークマスク長/デフォルトルート)

```console
$ usacloud server build \
        --os-type centos \
        --hostname [put-your-hostname] \
        --password [put-your-password] \
        --name [put-your-server-name] \
        --network-mode switch \
        --switch-id [put-your-switchID] \
        --ipaddress [put-your-IPAddress] \
        --nw-masklen [put-your-network-masklen] \
        --default-route [put-your-default-route]
```



--

## スタートアップスクリプトの利用

#### 登録済みのスタートアップスクリプトを利用(IDを指定、複数指定OK)

```console
$ usacloud server build \
        --os-type centos \
        --hostname [put-your-hostname] \
        --password [put-your-password] \
        --name [put-your-server-name] \
        --startup-script-ids [put-your-script-id] \
        --startup-script-ids [put-your-script-id]
```

#### スタートアップスクリプトの内容を記述した文字列を指定(複数指定OK)

```console
$ usacloud server build \
        --os-type centos \
        --hostname [put-your-hostname] \
        --password [put-your-password] \
        --name [put-your-server-name] \
        --startup-scripts "`cat your-script-file`" \
        --startup-scripts "`cat your-script-file`"
```

## SSH公開鍵の利用

#### 登録済みのSSH公開鍵を利用(IDを指定、複数指定OK)

```console
$ usacloud server build \
        --os-type centos \
        --hostname [put-your-hostname] \
        --password [put-your-password] \
        --name [put-your-server-name] \
        --ssh-key-mode id
        --ssh-key-ids [put-your-key-id]
```

#### SSH公開鍵をアップロード(複数指定OK)

公開鍵を記述した文字列を直接指定、またはファイル名での指定が可能です。  

```console
$ usacloud server build \
        --os-type centos \
        --hostname [put-your-hostname] \
        --password [put-your-password] \
        --name [put-your-server-name] \
        --ssh-key-mode upload
        --ssh-key-public-keys "`cat your-publickey-file`" \
        --ssh-key-public-key-files your-publickey-file \
```

#### SSHキーペアを生成

キーペアをさくらのクラウド上で生成し利用します。  
生成した秘密鍵はデフォルトで`~/.ssh/sacloud_pkey_[サーバID]`に保存されます。  
保存先は`--ssh-key-private-key-output`オプションで変更可能です。

```console
$ usacloud server build \
        --os-type centos \
        --hostname [put-your-hostname] \
        --password [put-your-password] \
        --name [put-your-server-name] \
        --ssh-key-mode generate
        --ssh-key-name [put-your-keyname]
        --ssh-key-pass-phrase [put-your-pass-phrase]
        --ssh-key-description [put-your-key-description]
```

## Windows系OSを利用

#### 共有セグメントに接続する場合

```console
$ usacloud server build \
        --os-type windows2016-rds \
        --password [put-your-password] \
        --name [put-your-server-name] \
```

#### スイッチ/スイッチ+ルータに接続する場合

```console
$ usacloud server build \
        --os-type windows2016-rds \
        --password [put-your-password] \
        --name [put-your-server-name] \
        --network-mode switch \
        --switch-id [put-your-switchID] \
        --ipaddress [put-your-IPAddress] \
        --nw-masklen [put-your-network-masklen] \
        --default-route [put-your-default-route]
```

## アーカイブIDを指定

```console
$ usacloud server build \
        --hostname [put-your-hostname] \
        --password [put-your-password] \
        --name [put-your-server-name] \
        --source-archive-id [put-your-archive-id]
```

Note: 指定するアーカイブによってはディスクの修正関連のオプションが使用できません。

## コピー元ディスクIDを指定

```console
$ usacloud server build \
        --hostname [put-your-hostname] \
        --password [put-your-password] \
        --name [put-your-server-name] \
        --source-disk-id [put-your-archive-id]
```

Note: 指定するディスクによってはディスクの修正関連のオプションが使用できません。

## 既存のディスクを接続

```console
$ usacloud server build \
        --disk-mode connect \
        --disk-id [put-your-disk-id] \
        --name [put-your-server-name] 
```

Note: `server build`では複数のディスクを持つサーバを作成できません。  
複数のディスクを使用したい場合、サーバ作成後に`server disk-connect`を実行してください。  

## ブランクディスク

`--disk-mode`が`create`(デフォルト値)、かつ`--source-archive-id`と`--source-disk-id`が指定されていない場合はブランクディスクとなります。  
ISOイメージを挿入しておく場合は`--iso-image-id`オプションを併用してください。  

```console
$ usacloud server build \
        --iso-image-is [put-your-iso-id]
        --name [put-your-server-name] 
```


## ディスクレス

ISOイメージを挿入しておく場合は`--iso-image-id`オプションを併用してください。  

```console
$ usacloud server build \
        --disk-mode diskless\
        --name [put-your-server-name] 
```

