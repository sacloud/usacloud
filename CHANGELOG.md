# Changelog

## 0.6.2 (2017-12-04)

* ロードバランサVIPの重複確認にIPアドレスとポート番号の組み合わせを利用 #261 (yamamoto-febc)


## 0.6.1 (2017-12-04)

* AppVeyorでのCI #255 (yamamoto-febc)
* AUTHORS出力処理の追加 #256 (yamamoto-febc)
* ロードバランサでのVIP操作時にSIGSEGVが発生する問題の修正 #258 (yamamoto-febc)


## 0.6.0 (2017-11-24)

* デフォルト出力形式の設定オプション #249 (yamamoto-febc)
* TravisCIからの通知設定追加 #250 (yamamoto-febc)
* 外部連携用の資格情報更新 #251 (yamamoto-febc)


## 0.5.0 (2017-11-16)

* 配布サイト用に静的コンテンツのビルドを実行 #244 (yamamoto-febc)
* リリースサイトのHTTPS対応 #245 (yamamoto-febc)
* 専有ホスト対応 #246 (yamamoto-febc)


## 0.4.0 (2017-11-09)

* VPCルータへの機能追加(DHCPでのDNSサーバ配布/NICごとのファイアウォール)  #242 (yamamoto-febc)


## 0.3.1 (2017-11-02)

* Fix typo on messages #239 (ariarijp)
* VPC APIのレスポンス処理修正 #240 (yamamoto-febc)


## 0.3.0 (2017-10-26)

* アーカイブアップロード時のout of memory対応 #225, #227 (higebu)
* VNCスナップショット機能の追加 #226 (misodengaku)
* Pleskパブリックアーカイブを除去 #230 (yamamoto-febc)
* ostypeパラメータによるCentOS6パブリックアーカイブの指定 #231 (yamamoto-febc)
* シンプル監視でのSSLサーバ証明書 有効期限監視 #232 (yamamoto-febc)
* configディレクトリが存在しない場合のconfig listコマンドエラー修正 #233 (yamamoto-febc)
* 統合テスト(初期実装) #235 (yamamoto-febc)
* FTPSアップロードでレスポンス226が応答されない問題の修正 #236 (yamamoto-febc)
* IPv4/IPv6関連コマンドの追加 #237 (yamamoto-febc)


## 0.2.2 (2017-10-01)

* NFSアプライアンス プラン追加 #220 (yamamoto-febc)
* ドキュメント更新 #221 (yamamoto-febc)


## 0.2.1 (2017-09-26)

* フラグ名変更 enabled -> disabled #215 (yamamoto-febc)
* 対象ゾーンとAPIのルートURL設定用グローバルオプション追加 #216 (yamamoto-febc)
* スタートアップスクリプトへClass属性追加 #218 (yamamoto-febc)


## 0.2.0 (2017-09-07)

* サーバ起動APIでのエラー時リトライ #204 (yamamoto-febc)
* NFSアプライアンス #205 (yamamoto-febc)
* DNSレコードのインデックスを特定しやすくする #206 (yamamoto-febc)
* quietモードでのIndex列の表示 #207 (yamamoto-febc)


## 0.1.1 (2017-08-16)

* 0.1.1リリース #201 (yamamoto-febc)


## 0.1.0 (2017-08-15)

* VPCルータでのログ出力機能追加 #179 (yamamoto-febc)
* データベースでのログ出力機能追加 #181 (yamamoto-febc)
* データベースでのバックアップ機能追加 #183 (yamamoto-febc)
* データベースでのモニター機能追加 #185 (yamamoto-febc)
* モニター機能デフォルトキー変更 #187 (yamamoto-febc)
* 複数のAPIキーの切り替え機能 #188 (yamamoto-febc)
* summaryコマンドの追加 #190 (yamamoto-febc)
* interface-driver項目の追加 #192 (yamamoto-febc)
* 請求CSVでのデフォルトターゲット指定 #194 (yamamoto-febc)
* パッケージレイアウトのリファクタリング #195 (yamamoto-febc)
* summaryコマンドヘルプ修正 #197 (yamamoto-febc)
* データベースのログ名称変更 #198 (yamamoto-febc)


## 0.0.13 (2017-07-11)

* textlint導入 #159 (yamamoto-febc)
* Chocolateyでのインストール方法追記 #160 (yamamoto-febc)
* textlintルール追加 #161 (yamamoto-febc)
* メンテナンス情報取得コマンド追加 #162 (yamamoto-febc)
* パラメータテンプレートでのフィールド名をハイフン区切りに変更 #168 (yamamoto-febc)
* --selectorオプション追加 #171 (yamamoto-febc)
* server ssh-execの複数ホストでのコマンド実行対応 #173 (yamamoto-febc)
* server scpコマンドでのサーバ名での対象指定 #174 (yamamoto-febc)


## 0.0.12 (2017-06-21)

* Windows2016 SQLServer Standard(RDS+Office)パブリックアーカイブ追加 #155 (yamamoto-febc)
* サーバ操作関連のドキュメント追加 #156 (yamamoto-febc)
* SiteGuardパブリックアーカイブ除去 #157 (yamamoto-febc)


## 0.0.11 (2017-06-15)

* linuxbrew対応 #138 (yamamoto-febc)
* --format-fileオプションの追加 #139 (yamamoto-febc)
* ビルド時にGo1.8(latest)を利用 #140 (yamamoto-febc)
* Windows上でのgo generate対応 #144 (yamamoto-febc)
* go generateで生成されるファイルの整理 #145 (yamamoto-febc)
* コピー元アーカイブ/ディスクID検索パラメータ追加 #146 (yamamoto-febc)
* コマンドパラメータのテンプレート対応 #149 (yamamoto-febc)
* パラメータファイルのスケルトン出力機能 #150 (yamamoto-febc)
* 環境変数からのオプション設定処理修正 #153 (yamamoto-febc)


## 0.0.10 (2017-06-12)

* VNCコマンドへサーバ起動待ち用オプション追加 #130 (yamamoto-febc)
* タグによる検索機能 #132 (yamamoto-febc)
* lsコマンドの出力を他コマンドで利用 #133 (yamamoto-febc)
* readコマンドでの操作対象を単一リソースのみに制限 #135 (yamamoto-febc)
* list or lsコマンドに起動状態列を追加 #136 (yamamoto-febc)


## 0.0.9 (2017-05-22)

* パブリックアーカイブ(`rancheros`,`plesk`)追加 #96 (yamamoto-febc)
* RancherOSでのデフォルトSSHユーザー名設定 #97 (yamamoto-febc)
* オブジェクトストレージでの環境変数読み込み順修正 #98 (yamamoto-febc)
* コードクリーンアップ、日本語化など  #100 , #101 (yamamoto-febc)
* ブランクディスク/ディスク接続処理の実装 #104 (yamamoto-febc)
* リソースへのデフォルトコマンド追加 #105 (yamamoto-febc)
* 請求情報の出力修正 #107 (yamamoto-febc)
* テーブル形式出力でのヘッダ書式修正 #110 (yamamoto-febc)
* auth-statusコマンドの追加 #111 (yamamoto-febc)
* --without-diskオプションの導入 #116 (yamamoto-febc)
* プログレス表示関数の置き換え #117 (yamamoto-febc)
* コマンドでの1文字エイリアスを除去 #118 (yamamoto-febc)
* アプライアンス(ロードバランサ/データベース/VPCルータ) #114, #119, #121 (yamamoto-febc)
* ヘルプ表示でのカテゴリ/ソート対応 #126, #128 (yamamoto-febc)


## 0.0.8 (2017-04-19)

* Add basic_usage guide #91 (yamamoto-febc)
* Add help text when no command is found #92 (yamamoto-febc)
* Update sacloud API #93 (yamamoto-febc)


## 0.0.7 (2017-04-09)

* Add gh-pages using mkdocs #83 (yamamoto-febc)
* Update docs #84 (yamamoto-febc)
* Fix SSHKeyID param #86 (yamamoto-febc)
* Fix displaying server IP Address #87 (yamamoto-febc)
* Refactoring ServerBuild #88 (yamamoto-febc)
* Add vnc-send subcommand #89 (yamamoto-febc)


## 0.0.6 (2017-03-28)

* Support homebrew #80 (yamamoto-febc)


## 0.0.5 (2017-03-27)

* Fix getSSHDefaultUserName func #74 (yamamoto-febc)
* Allow multiple target for vnc subcommand #76 (yamamoto-febc)
* Display IPAddress #77 (yamamoto-febc)


## 0.0.4 (2017-03-23)

* Fix build windows server #69 (yamamoto-febc)
* Bugfix - don't use ioutil.Discard for stdout #70 (yamamoto-febc)
* Windows Server 2016 + SQLServer(web/standard) #71 (yamamoto-febc)


## 0.0.3 (2017-03-22)

* Use colorable stdout #65 (yamamoto-febc)


## 0.0.2 (2017-03-22)

* Monitoring APIs #51 (yamamoto-febc)
* Output options #53 (yamamoto-febc)
* Progress output writer #55 (yamamoto-febc)
* Allow multiple ID or Name args #56 (yamamoto-febc)
* Open VNC client #59 (yamamoto-febc)
* Add config command #62 (yamamoto-febc)


## 0.0.1 (2017-03-14)

  * Set SkipAuth flag to object-storage commands #41 (yamamoto-febc)
  * Add confirmation to dangerous operations #44 (yamamoto-febc)
  * Build the Deploy pipeline #30,#31,#32,#34,#36,#37,#38,#39,#40,#43,#45 (yamamoto-febc)

## 0.0.0 (2017-03-09)

* First release (yamamoto-febc)
