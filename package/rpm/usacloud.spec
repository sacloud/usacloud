# sudo yum -y install rpmdevtools go && rpmdev-setuptree
# rpmbuild -ba ~/rpmbuild/SPECS/usacloud.spec

%define _binaries_in_noarch_packages_terminate_build 0

Summary: CLI client of the SakuraCloud
Name:    usacloud
Version: %{_version}
Release: 1
BuildArch: %{buildarch}
License: Apache-2.0
Group:   SakuraCloud
URL:     https://github.com/sacloud/usacloud

Source0:   %{_sourcedir}/usacloud_bash_completion
BuildRoot: %{_tmppath}/%{name}-%{version}-%{release}-root

%description
CLI client of the SakuraCloud

%prep

%build

%install
%{__rm} -rf %{buildroot}
%{__install} -Dp -m0755 %{_builddir}/%{name}  %{buildroot}%{_bindir}/%{name}
%{__mkdir} -p %{buildroot}%{_sysconfdir}/bash_completion.d
%{__install} -m 644 -T %{SOURCE0} %{buildroot}%{_sysconfdir}/bash_completion.d/usacloud


%clean
%{__rm} -rf %{buildroot}

%post

%files
%defattr(-,root,root)
%{_bindir}/%{name}
%{_sysconfdir}/bash_completion.d/usacloud

%changelog
* Thu Feb 27 2020 <sacloud.users@gmail.com> - 0.32.1-1
- Fix converting ID when parsing arguments (by yamamoto-febc)

* Thu Feb 27 2020 <sacloud.users@gmail.com> - 0.32.0-1
- Update copyright (by yamamoto-febc)
- Add v1 design idea (by yamamoto-febc)
- Introduce sacloud.ID type for avoiding json marshaling error (by yamamoto-febc)

* Wed Dec 18 2019 <sacloud.users@gmail.com> - 0.31.1-1
- エンハンスドロードバランサのスキーマ変更 (by yamamoto-febc)

* Wed Nov 20 2019 <sacloud.users@gmail.com> - 0.31.0-1
- os_typeに指定可能な値の追加 (by yamamoto-febc)

* Tue Nov 19 2019 <sacloud.users@gmail.com> - 0.30.0-1
- ライセンスヘッダの追加 (by yamamoto-febc)
- コピーライト表記の修正 (by yamamoto-febc)
- VPCルータ 4000Mbpsプラン (by yamamoto-febc)
- PTRレコード (by yamamoto-febc)
- エンハンスドロードバランサでのTCPサポート (by yamamoto-febc)

* Mon Nov 11 2019 <sacloud.users@gmail.com> - 0.29.1-1
- 専有ホストのプラン検索パラメータ (by yamamoto-febc)
- k3OS利用時のデフォルトユーザー名検出処理の修正 (by yamamoto-febc)

* Wed Oct 30 2019 <sacloud.users@gmail.com> - 0.29.0-1
- ウェブアクセラレータでの証明書登録 (by yamamoto-febc)

* Tue Oct 29 2019 <sacloud.users@gmail.com> - 0.28.0-1
- Go v1.13 / libsacloud v1.28.1 (by yamamoto-febc)
- Aliasレコード (by yamamoto-febc)
- Use go module mirror (by yamamoto-febc)
- シンプル監視での再通知間隔の設定 (by yamamoto-febc)
- vendoring (by yamamoto-febc)
- Use alpine 3.10 (by yamamoto-febc)
- ディスクの修正APIでのBackgroundパラメータの利用 (by yamamoto-febc)
- パーティションリサイズを行うサブコマンドを追加 (by yamamoto-febc)

* Fri Aug 30 2019 <sacloud.users@gmail.com> - 0.27.1-1
- 0.27.0のリリースビルドエラーによる再リリース (by yamamoto-febc)

* Fri Aug 30 2019 <sacloud.users@gmail.com> - 0.27.0-1
- エンハンスドロードバランサ更新 (by yamamoto-febc)

* Thu Jul 04 2019 <sacloud.users@gmail.com> - 0.26.0-1
- エンハンスドロードバランサでのカスタムレスポンスヘッダ設定 (by yamamoto-febc)

* Fri Jun 28 2019 <sacloud.users@gmail.com> - 0.25.1-1
- Use libsacloud v1.25.1 - fix database apis (by yamamoto-febc)

* Thu Jun 27 2019 <sacloud.users@gmail.com> - 0.25.0-1
- libsacloud v1.25.0 - k3OS (by yamamoto-febc)
- エンハンスドロードバランサでのセッション維持機能 (by yamamoto-febc)

* Wed Jun 26 2019 <sacloud.users@gmail.com> - 0.24.0-1
- エンハンスドロードバランサの機能拡充 (by yamamoto-febc)

* Fri Jun 14 2019 <sacloud.users@gmail.com> - 0.23.1-1
- CAA Recordの作成 (by yamamoto-febc)

* Thu May 16 2019 <sacloud.users@gmail.com> - 0.23.0-1
- コア専有プラン (by yamamoto-febc)
- Add test codes (by yamamoto-febc)

* Thu Apr 18 2019 <sacloud.users@gmail.com> - 0.22.0-1
- APIリクエスト時の*http.Clientタイムアウト指定機能 (by yamamoto-febc)
- スイッチ+ルータでの5000Mbpsプラン (by yamamoto-febc)

* Wed Mar 27 2019 <sacloud.users@gmail.com> - 0.21.0-1
- エンハンスドLBでのプラン変更 (by yamamoto-febc)
- NFSへのSSDプラン追加 (by yamamoto-febc)

* Tue Mar 12 2019 <sacloud.users@gmail.com> - 0.20.1-1
- 実験的機能への警告表示 (by yamamoto-febc)

* Thu Mar 07 2019 <sacloud.users@gmail.com> - 0.20.0-1
- ProxyLB (by yamamoto-febc)
- summaryコマンドへProxyLB追加 (by yamamoto-febc)

* Wed Feb 20 2019 <sacloud.users@gmail.com> - 0.19.0-1
- summaryコマンドへNFS/SIM/MobileGateway追加 (by yamamoto-febc)
- libsacloud to v1.14.0 (by yamamoto-febc)
- アクティビティAPIでのソート順修正 (by yamamoto-febc)
- YAML出力オプションの追加 (by yamamoto-febc)
- --query-fileオプションの追加 (by yamamoto-febc)
- DNSレコードのバルク更新コマンド追加 (by yamamoto-febc)

* Fri Jan 25 2019 <sacloud.users@gmail.com> - 0.18.0-1
- Usacloud Sandboxへのリンク追加 (by yamamoto-febc)
- クーポン情報取得コマンド (by yamamoto-febc)
- Use sacloud/go-jmespath (by yamamoto-febc)
- --output-typeオプションに一文字エイリアス(-o)を追加 (by yamamoto-febc)

* Thu Dec 20 2018 <sacloud.users@gmail.com> - 0.17.0-1
- シンプル監視での死活監視ステータス (by yamamoto-febc)

* Wed Dec 19 2018 <sacloud.users@gmail.com> - 0.16.4-1
- Windows Server 2019対応 (by yamamoto-febc)

* Fri Dec 14 2018 <sacloud.users@gmail.com> - 0.16.3-1
- ロードバランサでのVIP削除エラー対応 (by yamamoto-febc)

* Thu Dec 06 2018 <sacloud.users@gmail.com> - 0.16.2-1
- SIMマルチキャリア機能での綴り誤り修正 (by yamamoto-febc)

* Thu Dec 06 2018 <sacloud.users@gmail.com> - 0.16.1-1
- リリースビルド: エラー修正 (by yamamoto-febc)

* Thu Dec 06 2018 <sacloud.users@gmail.com> - 0.16.0-1
- use go 1.11 modules (by yamamoto-febc)
- SIMのマルチキャリア対応 (by yamamoto-febc)

* Fri Oct 12 2018 <sacloud.users@gmail.com> - 0.15.0-1
- サーバ/ディスクの新プラン対応 (by yamamoto-febc)
- トラフィックコントロール機能 (by yamamoto-febc)
- データベースサブコマンド追加(clone/replica) (by yamamoto-febc)

* Fri Sep 14 2018 <sacloud.users@gmail.com> - 0.14.3-1
- CentOS6向けタグの更新 (by yamamoto-febc)
- パブリックアーカイブ追加/ストレージ分離オプションのリクエストレイアウト対応 (by yamamoto-febc)
- DNSゾーン登録時のパラメータにデフォルト値を明示  (by yamamoto-febc)

* Tue Aug 28 2018 <sacloud.users@gmail.com> - 0.14.2-1
- Change DNS record max count from 300 to 1000 (by hnakamur)

* Tue Aug 28 2018 <sacloud.users@gmail.com> - 0.14.1-1
- ディスクの再インストール時のパラメータ修正 (by yamamoto-febc)

* Fri Aug 24 2018 <sacloud.users@gmail.com> - 0.14.0-1
- シンプル監視でのBasic認証 (by yamamoto-febc)
- ロードバランサーVIPでの説明欄 (by yamamoto-febc)

* Thu Aug 09 2018 <sacloud.users@gmail.com> - 0.13.2-1
- VPCルータでのサイト間VPNステータス確認用コマンド (by yamamoto-febc)

* Tue Jul 19 2018 <sacloud.users@gmail.com> - 0.13.1-1
- SQL Server 2017パブリックアーカイブ対応 (by yamamoto-febc)

* Thu Jul 05 2018 <sacloud.users@gmail.com> - 0.13.0-1
- VPCルータでのインターネット接続 有効/無効 設定 (by yamamoto-febc)

* Tue Jul 03 2018 <sacloud.users@gmail.com> - 0.12.0-1
- シンプル監視でのSNIオプション追加 (by yamamoto-febc)
- モバイルゲートウェイでのルーティング設定 (by yamamoto-febc)

* Tue Jun 05 2018 <sacloud.users@gmail.com> - 0.11.1-1
- データベースアプライアンスのデフォルトバージョン更新 (by yamamoto-febc)

* Mon Apr 09 2018 <sacloud.users@gmail.com> - 0.11.0-1
- Add --class option to startup-script (by blp1526)
- Update copyright (by yamamoto-febc)
- Fix startup-script to startup-script-ids (by blp1526)
- ディスク修正のテスト追加 (by yamamoto-febc)
- FreeBSD版のビルド/リリース (by yamamoto-febc)
- データベースアプライアンス 500GB/1TBプラン対応 (by yamamoto-febc)
- 開発環境の改善 (by yamamoto-febc)
- --queryオプションでのJMESPath指定 (by yamamoto-febc)
- マルチパートアップロードの修正 (by yamamoto-febc)
- バージョン情報にOS/ARCHを表示 (by yamamoto-febc)
- リリース時のドキュメント生成を自動化 (by yamamoto-febc)
- セキュアモバイル(SIM/モバイルゲートウェイ)対応 (by yamamoto-febc)

* Wed Feb 14 2018 <sacloud.users@gmail.com> - 0.10.1-1
- タイムアウトオプションの追加 (by yamamoto-febc)
- プログレス表示崩れ対応 (by yamamoto-febc)

* Tue Feb 13 2018 <sacloud.users@gmail.com> - 0.10.0-1
- server build時のバリデーション誤り修正 (by yamamoto-febc)
- インストールスクリプトの複数OS対応 (by yamamoto-febc)
- Windows系OSの作成時パラメータ例誤り修正 (by yamamoto-febc)
- auth-statusでのレスポンスハンドリング修正 (by yamamoto-febc)
- Sophos UTM対応 (by yamamoto-febc)
- VNC接続時の確認ダイアログ除去 (by yamamoto-febc)
- リモートデスクトップ接続(RDP)コマンドの追加 (by yamamoto-febc)
- リモートデスクトップ接続コマンド名変更 (by yamamoto-febc)
- グローバルオプションの追加 (by yamamoto-febc)
- database create時のヘルプテキスト改善 (by yamamoto-febc)
- no-colorオプションの追加 (by yamamoto-febc)

* Wed Jan 24 2018 <sacloud.users@gmail.com> - 0.9.1-1
- MariaDB作成時のAPIパラメータ更新 (by yamamoto-febc)

* Fri Dec 22 2017 <sacloud.users@gmail.com> - 0.9.0-1
- ウェブアクセラレータ コマンド追加 (by yamamoto-febc)
- ソース生成ツール修正 (by yamamoto-febc)

* Tue Dec 19 2017 <sacloud.users@gmail.com> - 0.8.0-1
- ストレージ情報の追加 (by yamamoto-febc)

* Tue Dec 05 2017 <sacloud.users@gmail.com> - 0.7.0-1
- 標準入出力の取り扱い改善 (by yamamoto-febc)
- オブジェクトストレージへのマルチパートアップロード (by yamamoto-febc)
- TravisCI上でのCI/CD改善 (by yamamoto-febc)
- selfコマンドの追加 (by yamamoto-febc)

* Tue Dec 05 2017 <sacloud.users@gmail.com> - 0.6.3-1
- summaryコマンドでの件数集計でTotalフィールドを利用 (by yamamoto-febc)

* Mon Dec 04 2017 <sacloud.users@gmail.com> - 0.6.2-1
- ロードバランサVIPの重複確認にIPアドレスとポート番号の組み合わせを利用 (by yamamoto-febc)

* Mon Dec 04 2017 <sacloud.users@gmail.com> - 0.6.1-1
- AppVeyorでのCI (by yamamoto-febc)
- AUTHORS出力処理の追加 (by yamamoto-febc)
- ロードバランサでのVIP操作時にSIGSEGVが発生する問題の修正 (by yamamoto-febc)

* Fri Nov 24 2017 <sacloud.users@gmail.com> - 0.6.0-1
- デフォルト出力形式の設定オプション (by yamamoto-febc)
- TravisCIからの通知設定追加 (by yamamoto-febc)
- 外部連携用の資格情報更新 (by yamamoto-febc)

* Thu Nov 16 2017 <yamamoto.febc@gmail.com> - 0.5.0-1
- 配布サイト用に静的コンテンツのビルドを実行 (by yamamoto-febc)
- リリースサイトのHTTPS対応 (by yamamoto-febc)
- 専有ホスト対応 (by yamamoto-febc)

* Thu Nov 09 2017 <yamamoto.febc@gmail.com> - 0.4.0-1
- VPCルータへの機能追加(DHCPでのDNSサーバ配布/NICごとのファイアウォール)  (by yamamoto-febc)

* Thu Nov 02 2017 <yamamoto.febc@gmail.com> - 0.3.1-1
- Fix typo on messages (by ariarijp)
- VPC APIのレスポンス処理修正 (by yamamoto-febc)

* Thu Oct 26 2017 <yamamoto.febc@gmail.com> - 0.3.0-1
- アーカイブアップロード時のout of memory対応 (by higebu)
- VNCスナップショット機能の追加 (by misodengaku)
- Pleskパブリックアーカイブを除去 (by yamamoto-febc)
- ostypeパラメータによるCentOS6パブリックアーカイブの指定 (by yamamoto-febc)
- シンプル監視でのSSLサーバ証明書 有効期限監視 (by yamamoto-febc)
- configディレクトリが存在しない場合のconfig listコマンドエラー修正 (by yamamoto-febc)
- 統合テスト(初期実装) (by yamamoto-febc)
- FTPSアップロードでレスポンス226が応答されない問題の修正 (by yamamoto-febc)
- IPv4/IPv6関連コマンドの追加 (by yamamoto-febc)

* Sun Oct 01 2017 <yamamoto.febc@gmail.com> - 0.2.2-1
- NFSアプライアンス プラン追加 (by yamamoto-febc)
- ドキュメント更新 (by yamamoto-febc)

* Tue Sep 26 2017 <yamamoto.febc@gmail.com> - 0.2.1-1
- フラグ名変更 enabled -> disabled (by yamamoto-febc)
- 対象ゾーンとAPIのルートURL設定用グローバルオプション追加 (by yamamoto-febc)
- スタートアップスクリプトへClass属性追加 (by yamamoto-febc)

* Thu Sep 07 2017 <yamamoto.febc@gmail.com> - 0.2.0-1
- サーバ起動APIでのエラー時リトライ (by yamamoto-febc)
- NFSアプライアンス (by yamamoto-febc)
- DNSレコードのインデックスを特定しやすくする (by yamamoto-febc)
- quietモードでのIndex列の表示 (by yamamoto-febc)

* Wed Aug 16 2017 <yamamoto.febc@gmail.com> - 0.1.1-1
- 0.1.1リリース (by yamamoto-febc)

* Tue Aug 15 2017 <yamamoto.febc@gmail.com> - 0.1.0-1
- VPCルータでのログ出力機能追加 (by yamamoto-febc)
- データベースでのログ出力機能追加 (by yamamoto-febc)
- データベースでのバックアップ機能追加 (by yamamoto-febc)
- データベースでのモニター機能追加 (by yamamoto-febc)
- モニター機能デフォルトキー変更 (by yamamoto-febc)
- 複数のAPIキーの切り替え機能 (by yamamoto-febc)
- summaryコマンドの追加 (by yamamoto-febc)
- interface-driver項目の追加 (by yamamoto-febc)
- 請求CSVでのデフォルトターゲット指定 (by yamamoto-febc)
- パッケージレイアウトのリファクタリング (by yamamoto-febc)
- summaryコマンドヘルプ修正 (by yamamoto-febc)
- データベースのログ名称変更 (by yamamoto-febc)

* Tue Jul 11 2017 <yamamoto.febc@gmail.com> - 0.0.13-1
- textlint導入 (by yamamoto-febc)
- Chocolateyでのインストール方法追記 (by yamamoto-febc)
- textlintルール追加 (by yamamoto-febc)
- メンテナンス情報取得コマンド追加 (by yamamoto-febc)
- パラメータテンプレートでのフィールド名をハイフン区切りに変更 (by yamamoto-febc)
- --selectorオプション追加 (by yamamoto-febc)
- server ssh-execの複数ホストでのコマンド実行対応 (by yamamoto-febc)
- server scpコマンドでのサーバ名での対象指定 (by yamamoto-febc)

* Wed Jun 21 2017 <yamamoto.febc@gmail.com> - 0.0.12-1
- Windows2016 SQLServer Standard(RDS+Office)パブリックアーカイブ追加 (by yamamoto-febc)
- サーバ操作関連のドキュメント追加 (by yamamoto-febc)
- SiteGuardパブリックアーカイブ除去 (by yamamoto-febc)

* Thu Jun 15 2017 <yamamoto.febc@gmail.com> - 0.0.11-1
- linuxbrew対応 (by yamamoto-febc)
- --format-fileオプションの追加 (by yamamoto-febc)
- ビルド時にGo1.8(latest)を利用 (by yamamoto-febc)
- Windows上でのgo generate対応 (by yamamoto-febc)
- go generateで生成されるファイルの整理 (by yamamoto-febc)
- コピー元アーカイブ/ディスクID検索パラメータ追加 (by yamamoto-febc)
- コマンドパラメータのテンプレート対応 (by yamamoto-febc)
- パラメータファイルのスケルトン出力機能 (by yamamoto-febc)
- 環境変数からのオプション設定処理修正 (by yamamoto-febc)

* Mon Jun 12 2017 <yamamoto.febc@gmail.com> - 0.0.10-1
- VNCコマンドへサーバ起動待ち用オプション追加 (by yamamoto-febc)
- タグによる検索機能 (by yamamoto-febc)
- lsコマンドの出力を他コマンドで利用 (by yamamoto-febc)
- readコマンドでの操作対象を単一リソースのみに制限 (by yamamoto-febc)
- list or lsコマンドに起動状態列を追加 (by yamamoto-febc)

* Mon May 22 2017 <yamamoto.febc@gmail.com> - 0.0.9-1
- パブリックアーカイブ(RancherOS,Plesk)追加 (by yamamoto-febc)
- RancherOSでのデフォルトSSHユーザー名設定 (by yamamoto-febc)
- オブジェクトストレージでの環境変数読み込み順修正 (by yamamoto-febc)
- コードのクリーンアップ (by yamamoto-febc)
- README.mdの日本語化 to Japanese (by yamamoto-febc)
- ブランクディスク/ディスク接続処理 (by yamamoto-febc)
- リソースへのデフォルトコマンド追加 (by yamamoto-febc)
- 請求情報の出力修正 (by yamamoto-febc)
- テーブル形式出力でのヘッダ書式修正 (by yamamoto-febc)
- auth-statusコマンドの追加 (by yamamoto-febc)
- ロードバランサ実装 (by yamamoto-febc)
- --with-diskオプションの変更 (by yamamoto-febc)
- プログレス表示関数の置き換え (by yamamoto-febc)
- コマンドでの1文字エイリアスを除去 (by yamamoto-febc)
- データベース追加 (by yamamoto-febc)
- VPCRouter追加 (by yamamoto-febc)
- ヘルプ表示でのカテゴリ/ソート対応 (by yamamoto-febc)
- トップレベルコマンドでのリソース並び順修正 (by yamamoto-febc)

* Wed Apr 19 2017 <yamamoto.febc@gmail.com> - 0.0.8-1
- Add basic_usage guide (by yamamoto-febc)
- Add help text when no command is found (by yamamoto-febc)
- Update sacloud API (by yamamoto-febc)

* Sun Apr 09 2017 <yamamoto.febc@gmail.com> - 0.0.7-1
- Add gh-pages using mkdocs (by yamamoto-febc)
- Update docs (by yamamoto-febc)
- Fix SSHKeyID param (by yamamoto-febc)
- Fix displaying server IP Address (by yamamoto-febc)
- Refactoring ServerBuild (by yamamoto-febc)
- Add vnc-send subcommand (by yamamoto-febc)

* Tue Mar 28 2017 <yamamoto.febc@gmail.com> - 0.0.6-1
- Support homebrew (by yamamoto-febc)

* Mon Mar 27 2017 <yamamoto.febc@gmail.com> - 0.0.5-1
- Fix getSSHDefaultUserName func (by yamamoto-febc)
- Allow multiple target for vnc subcommand (by yamamoto-febc)
- Display IPAddress (by yamamoto-febc)

* Thu Mar 23 2017 <yamamoto.febc@gmail.com> - 0.0.4-1
- Fix build windows server (by yamamoto-febc)
- Bugfix - don't use ioutil.Discard for stdout (by yamamoto-febc)
- Windows Server 2016 + SQLServer(web/standard) (by yamamoto-febc)

* Wed Mar 22 2017 <yamamoto.febc@gmail.com> - 0.0.3-1
- Use colorable stdout (by yamamoto-febc)

* Wed Mar 22 2017 <yamamoto.febc@gmail.com> - 0.0.2-1
- Monitoring APIs (by yamamoto-febc)
- Output options (by yamamoto-febc)
- Progress output writer (by yamamoto-febc)
- Allow multiple ID or Name args (by yamamoto-febc)
- Open VNC client (by yamamoto-febc)
- Add config command (by yamamoto-febc)

* Tue Mar 14 2017 <yamamoto.febc@gmail.com> - 0.0.1-1
- Set SkipAuth flag to object-storage commands (by yamamoto-febc)
- Add confirmation to dangerous operations (by yamamoto-febc)
- Build the Deploy pipeline (by yamamoto-febc)

* Fri Mar 3 2017 <yamamoto.febc@gmail.com> - 0.0.0
- Initial commit
