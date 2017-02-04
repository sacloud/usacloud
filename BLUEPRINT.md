## CLI BLUEPRINT

### 概要

- [sacloud](https://github.com/sakura-internet/node-sacloud)ではない新CLI

- さくらのクラウドの最新機能に追随する

- golangを利用することでシングルバイナリ/クロスプラットフォーム化

- これまでバラバラに存在していた以下機能を取り込む
  - ウェブアクセラレータ関連
  - オブジェクトストレージ関連
  - アーカイブやISOイメージなどのFTPSでのアップロード/ダウンロード
  - リソース全削除


### コマンドのイメージ

ログイン(環境変数 or ~/.usacloud配下の設定ファイル or 都度オプション)

```bash
# configure(対話)

$ usacloud configure
> Sakura Cloud API Access Token [None]:
> Sakura Cloud API Access Token Secret [None]:
> Default Zone (is1a/is1b/tk1a/tk1v) [tk1a]:
> Default Output Format (table/json/csv/tsv) [table]: 

Written your settings to ~/.usacloud/configure

```

リソース操作形のコマンドフォーマット
```bash
$ usacloud ([グローバルオプション]) [リソース種別] [コマンド] ([オプション])

# 単体操作形で-nameが指定されていた場合、同じ名前のリソースが複数ある場合はエラーとする

# === コマンド例 ===
# サーバー作成
$ usacloud server create --name=server01
#
# サーバー詳細情報表示
$ usacloud server read --id=999999999999
$ # or
$ usacloud server read 999999999999 # IDは引数としても渡せる

# サーバーに接続されたディスク一覧(単体操作)
$ usacloud server disk-list --id=999999999999

```

### リソース

```bash
# サーバー
$ usacloud server
# ディスク系
$ usacloud disk
$ usacloud archive
$ usacloud iso-image
# ネットワーク
$ usacloud switch
$ usacloud packet-filter
$ usacloud bridge
# アプライアンス
$ usacloud loadbalancer
$ usacloud vpc-router
$ usacloud database
# サービスアイテム系
$ usacloud auto-backup
$ usacloud gslb
$ usacloud dns
$ usacloud simple-monitor
# 共通リソース
$ usacloud license
$ usacloud ssh-key
$ usacloud startup-script
$ usacloud icon
# 請求
$ usacloud bill
# 設備関連
$ usacloud facility # zone + region
# 商品関連
$ usacloud product # product-*
# ウェブアクセラレータ
$ usacloud webaccel 
# 請求関連
$ usacloud bill
# オブジェクトストレージ
$ usacloud object-storage 
```

## サーバー操作

```bash
# サーバー作成(全てデフォルト、サーバー名) ただしパスワードなどの警告は出すこと
$ server create server01

# その他オプションなど
$ server create 
# 1.サーバープラン
    --core=1 
    --memory=1 
# 2.ディスク
    --disk-mode=create          # or create / connect / diskless
    # disk-mode = createの場合
    --disk-plan=ssd             # ssd / hdd
    --disk-source-type=archive  # archive / my-archive / my-disk / blank
    --os=centos                 # centos / ubuntu / debian
    --disk-size=20              #
    --distant-from=disk-id      # 複数指定OK
    --use-disk-virtio
    # disk-mode = connectの場合
    --disk=disk-id
# 3.ネットワーク関係
    --network-mode=shared       # shared / switch / disconnect / none(cli固有)
    # network-modeがsharedの場合
    --use-nic-virtio
    --packet-filter=pf-id
    # network-modeがswitchの場合
    --switch=switck-id
# 4.ディスクの修正　(今はディスクの修正を「行わない」ことは指定できない)
    --password="password" 
    --hostname="hostname"
    # networkmodeがswitchの場合だけ
    --ipaddress="192.168.2.1"
    --netmask=28
    --default-gateway="192.168.2.1"
    --sshkey-mode=none          # none / id / generate / upload
    # sshkey-mode = none以外の場合
    --disable-password-auth
    # sshkey-mode=idの場合
    --sshkey=sshkey-id          
    # generateの場合
    --sshkey-name="name"
    --sshkey-pass-phrase="pass-pharase"
    --sshkey-description="desc"
    --sshkey-private-key-output="./keyname"
    # inputの場合(指定がない場合は~/.ssh/配下を使う-openssh互換)
    --sshkey-public-key="public-key-text"
    --sshkey-public-key-file="~/.ssh/id_rsa.pub"
# 5. シンプル監視(対応しない)
# 6. サーバーの情報
    --description="description"
    --tags="hoge" --tags="fuga"
    --icon=icon-id                  # アイコンのID
    --icon-file="path-to-icon-file" # アイコンファイル(アップロード)
# 7. その他
    --us-keyboard                   # USキーボードを有効化
    --boot-after-create             # 作成後すぐに起動
# 欄外:ISOイメージ
    --iso-image=image-id            # ISOイメージのID
    --iso-image-file="path/to/iso"  # ISOイメージファイル(アップロード)
```

createが対応しないこと

  - シンプル監視の同時作成:面倒だもん
  - ディスクの追加 : 後で別途コマンド発行してもらう
  - NICの追加 : 後で別途コマンド発行してもらう
  

操作一覧

```bash
# サーバーの操作一覧
$ server help

  list                 : List Servers(+find)
          
  create               : Create Server
  read                 : Show server detail
  update               : Update Server
  delete               : Delete Server
            
  start                :
  stop                 :
  reset                :
          
  iso-info             : Show ISO info
* iso-insert           : insert ISO
* iso-eject            : eject ISO

  disk-info            : List Disks which is connecting Server
* disk-attach          : 
* disk-detach          : 
  
  interface-info       : 
* interface-connect    : 
* interface-disconnect :
  
# 単体操作 = *
    
```

