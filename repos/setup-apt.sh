#!/bin/sh

set -e
sudo -k

sudo sh <<'SCRIPT'
  set -x
  echo "deb https://releases.usacloud.jp/usacloud/repos/debian /" > /etc/apt/sources.list.d/usacloud.list
  curl -fsS https://releases.usacloud.jp/usacloud/repos/GPG-KEY-usacloud | apt-key add -
  apt-get update -qq

  apt-get install -y usacloud
SCRIPT
