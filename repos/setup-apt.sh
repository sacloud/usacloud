#!/bin/sh

set -e
sudo -k

sudo sh <<'SCRIPT'
  set -x
  echo "deb https://usacloud.b.sakurastorage.jp/repos/debian /" > /etc/apt/sources.list.d/usacloud.list
  # curl -fsS https://usacloud.b.sakurastorage.jp/repos/GPG-KEY-usacloud | apt-key add -
  apt-get update -qq

  apt-get install -y --allow-unauthenticated usacloud
SCRIPT
