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
