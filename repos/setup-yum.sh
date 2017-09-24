#!/bin/sh

set -e
sudo -k

sudo sh <<'SCRIPT'
  set -x

  #import GPG key
  gpgkey_path=`mktemp`
  curl -fsS -o $gpgkey_path http://releases.usacloud.jp/usacloud/repos/GPG-KEY-usacloud
  rpm --import $gpgkey_path
  rm $gpgkey_path

  cat >/etc/yum.repos.d/usacloud.repo <<'EOF';
[usacloud]
name=usacloud
baseurl=http://releases.usacloud.jp/usacloud/repos/centos/$basearch
gpgcheck=1
EOF

  yum install -y usacloud

SCRIPT

