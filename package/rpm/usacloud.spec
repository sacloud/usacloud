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
