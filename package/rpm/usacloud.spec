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
