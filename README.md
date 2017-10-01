# Usacloud

![usacloud_logo_h.png](build_docs/docs/images/usacloud_logo_h.png)

[`usacloud`](https://github.com/sacloud/usacloud)は[さくらのクラウド](http://cloud.sakura.ad.jp/index.html)用のCLIクライアントです。  

[![Build Status](https://travis-ci.org/sacloud/usacloud.svg?branch=master)](https://travis-ci.org/sacloud/usacloud)  

## 主な特徴

- さくらのクラウドの最新機能に追随。請求情報やオブジェクトストレージ、ウェブアクセラレータなども対応済み
- クロスプラットフォーム(Windows/macOS/Linux)サポート。ARMでも動作可能。
- Go言語で実装されたシングルバイナリ、インストールはバイナリをコピーするだけ(yum/apt/brewもサポート)
- SSH/SCP/SFTP/VNCなどをバイナリ単体でサポート

## インストール

### macOS(`homebrew`) / Linux(`linuxbrew`)

    brew tap sacloud/usacloud; brew install usacloud

### RHEL / CentOS

    curl -fsSL http://releases.usacloud.jp/usacloud/repos/setup-yum.sh | sh

### Ubuntu / debian / bash on Windows(Ubuntu)

    curl -fsSL http://releases.usacloud.jp/usacloud/repos/setup-apt.sh | sh

### Windows(`chocolatey`)

    choco install usacloud

> chocolateyの[usacloudパッケージ](https://chocolatey.org/packages/usacloud)は @223n さんによってメンテナンスされています。

### Windows / その他の場合

以下のリンクからバイナリーファイルをダウンロードして展開し、任意のフォルダー内に配置してください。  
(PATHを通しておくと便利です)

- Windows 64bit版 : [http://releases.usacloud.jp/usacloud/repos/windows/usacloud_windows-amd64.zip](http://releases.usacloud.jp/usacloud/repos/windows/usacloud_windows-amd64.zip)
- Windows 32bit版 : [http://releases.usacloud.jp/usacloud/repos/windows/usacloud_windows-386.zip](http://releases.usacloud.jp/usacloud/repos/windows/usacloud_windows-386.zip)
- その他の場合 : [https://github.com/sacloud/usacloud/releases/latest/](https://github.com/sacloud/usacloud/releases/latest/)

`bash_completion`が利用できる場合は、以下のコマンドで`usacloud`用の`bash_completion`を導入することが出来ます。

```bash
curl -s -L http://releases.usacloud.jp/usacloud/contrib/completion/bash/usacloud >> ~/.bashrc
```

> ※bash_completionを有効化するには上記コマンドを実行後に再ログインしてください。

## 初期設定

`usacloud config` コマンドを用いてAPIキーを設定しておきます。 
(APIキーの設定は`~/.usacloud_config`にファイルとして保存されます。)

```bash
    $ usacloud config

    Setting SakuraCloud API Token => 
    	Enter token: [ENTER YOUR_API_TOKEN]

    Setting SakuraCloud API Secret => 
    	Enter secret: [ENTER YOUR_API_SECRET]
    	
    Setting SakuraCloud Zone => 
    	Enter zone[is1a/is1b/tk1a/tk1v](default:tk1a): [ENTER ZONE]
   
    Written your settings to ~/.usacloud_config
```

APIキーの設定は`usacloud config --show`コマンドで確認可能です。  

```bash
   $ usacloud config show
   
   token  = [YOUR_API_TOKEN]
   secret = [YOUR_API_SECRET]
   zone   = [YOUR_ZONE]
   
```

Note: APIキーは環境変数を用いて設定することも可能です。  

```bash
   $ export SAKURACLOUD_ACCESS_TOKEN=[YOUR_API_TOKEN]
   $ export SAKURACLOUD_ACCESS_TOKEN_SECRET=[YOUR_API_SECRET]
   $ export SAKURACLOUD_ZONE=tk1v
```
  
詳細は[Usacloudドキュメント](https://sacloud.github.io/usacloud/)を参照してください。  
   
### 使い方

[Usacloudドキュメント:基本的な使い方](https://sacloud.github.io/usacloud/basic_usage/)を参照してください。

```bash
NAME:
   usacloud - CLI client for SakuraCloud

USAGE:
   usacloud [global options] resource command [command options] [arguments...]

VERSION:
   NN.NN.NN, build xxxxxx

COMMANDS:
   config, profile                  A manage command of APIKey settings
   auth-status                      A manage commands of AuthStatus
   server                           A manage commands of Server
   archive                          A manage commands of Archive
   auto-backup                      A manage commands of AutoBackup
   disk                             A manage commands of Disk
   iso-image                        A manage commands of ISOImage
   bridge                           A manage commands of Bridge
   interface                        A manage commands of Interface
   internet                         A manage commands of Internet
   packet-filter                    A manage commands of PacketFilter
   switch                           A manage commands of Switch
   database                         A manage commands of Database
   load-balancer                    A manage commands of LoadBalancer
   nfs                              A manage commands of NFS
   vpc-router                       A manage commands of VPCRouter
   dns                              A manage commands of DNS
   gslb                             A manage commands of GSLB
   simple-monitor                   A manage commands of SimpleMonitor
   icon                             A manage commands of Icon
   license                          A manage commands of License
   ssh-key                          A manage commands of SSHKey
   startup-script, note             A manage commands of StartupScript
   bill                             A manage commands of Bill
   object-storage, ojs              A manage commands of ObjectStorage
   web-accel                        A manage commands of WebAccel
   price, public-price              A manage commands of Price
   product-disk, disk-plan          A manage commands of ProductDisk
   product-internet, internet-plan  A manage commands of ProductInternet
   product-license, license-info    A manage commands of ProductLicense
   product-server, server-plan      A manage commands of ProductServer
   region                           A manage commands of Region
   zone                             A manage commands of Zone
   summary                          Show summary of resource usage

GLOBAL OPTIONS:
   --token value                    API Token of SakuraCloud (default: none) [$SAKURACLOUD_ACCESS_TOKEN]
   --secret value                   API Secret of SakuraCloud (default: none) [$SAKURACLOUD_ACCESS_TOKEN_SECRET]
   --zone value                     Target zone of SakuraCloud (default: tk1a) [$SAKURACLOUD_ZONE]
   --config value, --profile value  Config(Profile) name [$USACLOUD_PROFILE]
   --help, -h                       show help (default: false)
   --version, -v                    print the version (default: false)

COPYRIGHT:
   Copyright (C) 2017 Kazumichi Yamamoto.
```   


#### Examples: 一覧表示/検索

```bash
    # 全件表示
    $ usacloud switch ls 
   
    # 名称に"example"を含むものを一覧表示
    $ usacloud switch ls --name example
    
    # ソート条件指定(名称での昇順、IDでの降順)
    $ usacloud switch ls --sort Name --sort -ID
    
    # Limit/Offset指定(最大5件、2件目から表示)
    $ usacloud switch ls --max 5 --from 2
    
```   

#### Examples: CRUD操作

```bash
    # 作成(Create)
    $ usacloud switch create --name "Example" --desc "description" --tags "Tag1" --tags "Tag2"
     
    # 詳細表示(Read)
    $ usacloud switch read <ID または 名称>
   
    # 更新(Update)
    $ usacloud switch update --name "Example-update" <ID または 名称>
    
    # 削除(Delete)
    $ usacloud switch delete <ID または 名称> 
    
```  

#### Examples: サーバ作成

```bash
    # CentOSインストール済みのサーバを構築
    $ usacloud server build \
             --name server01 \               # サーバ名
             --os-type centos \              # OS種別(パブリックアーカイブを指定)
             --hostname server01 \           # ホスト名
             --password "$YOUR_PASSWORD" \   # 管理者パスワード
             --ssh-key-mode generate \       # SSH公開鍵(クラウド上で生成する)
             --disable-pw-auth               # SSH接続時のパスワード、チャレンジ/レスポンス認証の無効化

    # generated private-key is saved to ~/.ssh/sacloud_pkey_[ServerID]
```   

#### Examples: サーバ操作(電源周り)

```bash
    # 起動
    $ usacloud server boot <ID または 名称> 
    
    # シャットダウン(graceful)
    $ usacloud server shutdown <ID または 名称> 
    
    # シャットダウン(force)
    $ usacloud server shutdown-force <ID または 名称> 
    
    # リセット(hard)
    $ usacloud server reset <ID または 名称> 
    
```

#### Examples: サーバへの接続(SSH/SCP/VNC)

```bash
    # サーバへのSSH接続
    # デフォルトでは~/.ssh/sacloud_pkey_[サーバID]ファイルが存在すれば秘密鍵として利用する(-iオプションで明示も可)
    $ usacloud server ssh <ID または 名称> 
    
    # サーバにSSH接続し、任意のコマンドを実行(Windowsコマンドプロンプトからでも実行可能)
    $ usacloud server ssh-exec <ID または 名称>  cat /etc/passwd
    
    # SCPでのアップロード/ダウンロード
    $ usacloud server scp local-file.txt [ID または 名称]:/home/ubuntu/remote-file.txt # ローカルからリモートへ
    $ usacloud server scp [ID または 名称]:/home/ubuntu/remote-file.txt local-file.txt # リモートからローカルへ
    $ usacloud server scp -r local-dir [ID または 名称]:/home/ubuntu/remote-dir        # ディレクトリに対して再帰的に処理
   
    # OSのデフォルトVNCクライアントを用いてサーバへVNC接続
    # (Windowsの場合、.vnc拡張子に適切なVNCクライアントを関連付けしておく必要あり)
    $ usacloud server vnc <ID または 名称> 
```

#### Examples: FTPSでのアップロード/ダウンロード(アーカイブ/ISOイメージ)

```bash
    # ISOイメージのアップロード
    $ usacloud iso-image create --name example --iso-file example.iso
    
    # アーカイブのダウンロード(マイカーカイブのみダウンロード可能)
    $ usacloud archive download --file-destination example.img <ID または 名称> 
    
```

#### Examples: 請求関連

```bash
    # 請求情報一覧
    $ usacloud bill list
    
    # 請求CSVの出力
    $ usacloud bill csv [BillID]

``` 

#### Examples: オブジェクトストレージの操作

```bash
    # オブジェクトストレージ用にAPIキー設定(バケットごと)
    $ export SACLOUD_OJS_ACCESS_KEY_ID="[YOUR_BUCKET_ACCESS_KEY]"
    $ export SACLOUD_OJS_SECRET_ACCESS_KEY="[YOUR_BUCKET_SECRET_KEY]"
    
    # オブジェクト一覧表示
    $ usacloud object-storage ls 
    $ usacloud object-storage ls dir1/dir2

    # オブジェクトのダウンロード
    $ usacloud object-storage get remote.txt           # 標準出力へ
    $ usacloud object-storage get remote.txt local.txt # ローカルファイルへ
    $ usacloud object-storage get -r remote/ local/    # ディレクトリを再帰的に処理

    # オブジェクトのアップロード
    $ usacloud object-storage put local.txt remote.txt 
    $ usacloud object-storage put local.txt dir1/dir2/remote.txt
    $ usacloud object-storage put -r local/ remote/    # ディレクトリを再帰的に処理
    
    # オブジェクトの削除
    $ usacloud object-storage del remote.txt
    $ usacloud object-storage del -r remote/           # ディレクトリを再帰的に処理
```   

#### Examples: ウェブアクセラレータ

```bash
    # ウェブアクセラレータ上のキャッシュを削除
    $ usacloud web-accel purge https://example.com https://foobar.com

```   

#### Examples: 出力の定義

```bash
    # テーブル形式(デフォルト)
    $ usacloud switch ls

    # JSON形式
    $ usacloud switch ls --output-type json

    # CSV/TSV形式
    $ usacloud switch ls --output-type csv # or tsv

    # 出力する列を指定(CSV/TSV形式での出力時に指定可能)
    $ usacloud switch ls --output-type tsv --col ID --col Name
   
    # IDまたはキーのみ出力
    $ usacloud swtich ls -q # or --quiet
    
    # golangのテンプレートを用いてカスタム出力
    $ usacloud switch ls --format "ID is '{{.ID}}', Name is '{{.Name}}'"
    ID is '123456789012', Name is 'example'
```

#### Examples: 出力の定義(モニタリングツール用)

```bash

    # target resource ID = 123456789012

    # for munin
    $ usacloud internet monitor --format "target.value {{.In}}" 123456789012
    
    # for zabbix_sender(zabbix_hostname=router01 , item_key=packet.in)
    $ usacloud internet monitor --format "router01 packet.in {{.UnixTime}} {{.In}}" 123456789012 \
         | zabbix_sender -z your.zabbix.hostname -p 10051 -T -i -
    
    # for sensu/mackerel
    $ OUTPUT_FORMAT=`echo -e "{{.Key}}\t{{.In}}\t{{.UnixTime}}"`
    $ usacloud internet monitor --format "$OUTPUT_FORMAT" 123456789012
```

### Examples: 複数のAPIキーの利用(プロファイル機能)

```bash
    # 一覧
    $ usacloud config list

    # プロファイルの作成(対話形式)
    $ usacloud config edit your-profile-name1

    # プロファイルの作成(非対話形式)
    $ usacloud config edit --zone "is1a" --token "token" --secret "secret" your-profile-name1
    
    # プロファイル内容の表示
    $ usacloud config show your-profile-name1

    #現在選択中のプロファイル名表示
    $ usacloud config current

    #プロファイル切り替え
    $ usacloud config use your-profile-name1

    #プロファイルの削除
    $ usacloud config delete your-profile-name1
```

## 開発

#### ビルド

    $ make build

#### ビルド(Docker上でのビルド/クロスプラットフォーム向けビルド)
    
    $ make docker-build
        
#### テスト
    
    $ make test
        
#### 各コマンドのソース生成

    $ make gen
    $ # or
    $ make gen-force
    
#### 新しいリソース/コマンドの追加

`define`配下に定義ファイルを作成し`make gen-force`コマンドでソース生成してください。  

#### ドキュメント

ドキュメントはGithub Pagesを利用しています。(masterブランチの`docs`ディレクトリ配下)  
静的ファイルの生成は`mkdocs`コマンドで行なっています。  
ドキュメントの追加や修正は`build_docs`ディレクトリ以下のファイルの追加/修正を行なった上で`mkdocs`コマンドでファイル生成してコミットしてください。

    # ドキュメントのプレビュー用サーバー起動(http://localhost/でプレビュー可能)
    make serve-docs
    
    # ドキュメントの検証(textlint)
    make lint-docs
    
    # build_docs配下のファイルからドキュメント生成(docsディレクトリ再生成)
    make build-docs

## License

 `usacloud` Copyright (C) 2017 Kazumichi Yamamoto.

  This project is published under [Apache 2.0 License](LICENSE.txt).
  
## Author

  * Kazumichi Yamamoto ([@yamamoto-febc](https://github.com/yamamoto-febc))
