#!/bin/sh
# Copyright 2017-2020 The Usacloud Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -e
sudo -k

echo "===== install usacloud by yum ====="

sudo sh <<'SCRIPT'
  set -x

  #import GPG key
  gpgkey_path=`mktemp`
  curl -fsSL -o $gpgkey_path https://releases.usacloud.jp/usacloud/repos/GPG-KEY-usacloud
  rpm --import $gpgkey_path
  rm $gpgkey_path

  cat >/etc/yum.repos.d/usacloud.repo <<'EOF';
[usacloud]
name=usacloud
baseurl=https://releases.usacloud.jp/usacloud/repos/centos/$basearch
gpgcheck=1
EOF

  yum install -y usacloud

SCRIPT

