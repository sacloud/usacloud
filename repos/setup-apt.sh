#!/bin/sh

set -e
sudo -k

echo "===== install usacloud by apt ====="
sudo sh <<'SCRIPT'
  set -x
  apt-get update -qq
  apt-get install -y curl apt-transport-https
  echo "deb https://releases.usacloud.jp/usacloud/repos/debian /" > /etc/apt/sources.list.d/usacloud.list
  curl -fsS https://releases.usacloud.jp/usacloud/repos/GPG-KEY-usacloud | apt-key add -
  apt-get update -qq
  apt-get install -y usacloud
SCRIPT
