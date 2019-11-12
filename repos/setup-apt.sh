#!/bin/sh
# Copyright 2017-2019 The Usacloud Authors
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
