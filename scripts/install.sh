#!/bin/bash
# Copyright 2017-2022 The Usacloud Authors
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


install_by_brew() {
  echo "===== install usacloud by homebrew ====="
  set -x
  brew tap sacloud/usacloud
  brew install usacloud
}

install_by_dnf() {
  echo "===== install usacloud by dnf ====="
  sudo sh <<'SCRIPT'
    set -x
    dnf install -y curl zip
SCRIPT
  install_with_curl
}

install_by_yum() {
  echo "===== install usacloud by yum ====="
  sudo sh <<'SCRIPT'
    set -x
    yum install -y curl zip
SCRIPT
  install_with_curl
}

install_by_apt() {
  echo "===== install usacloud by apt ====="
  sudo sh <<'SCRIPT'
    set -x
    apt-get update -qq
    apt-get install -y curl apt-transport-https zip
SCRIPT
  install_with_curl
}

install_with_curl() {
  sudo sh <<'SCRIPT'
    curl -LO https://github.com/sacloud/usacloud/releases/latest/download/usacloud_linux-amd64.zip
    unzip -o usacloud_linux-amd64.zip && rm usacloud_linux-amd64.zip
    chmod +x usacloud
    mv usacloud /usr/local/bin/
SCRIPT
}

### main
set -e
sudo -k

if [ "$(uname)" == 'Darwin' ]; then
  OS='Mac'
  if type brew >/dev/null 2>&1; then
    install_by_brew
    exit 0
  else
    echo "To install usacloud, you need 'brew' command"
    exit 1
  fi
elif [ "$(expr substr $(uname -s) 1 5)" == 'Linux' ]; then
  OS='Linux'
  if type dnf >/dev/null 2>&1; then
    install_by_dnf
    exit 0
  elif type yum > /dev/null 2>&1; then
    install_by_yum
    exit 0
  elif type apt > /dev/null 2>&1; then
    install_by_apt
    exit 0
  elif type brew > /dev/null 2>&1; then
    install_by_brew
    exit 0
  fi

  echo "To install usacloud, you need 'apt' or 'yum'/'dnf' or 'brew' command"
  exit 1

else
  echo "Your platform ($(uname -a)) is not supported."
  exit 1
fi

exit 0