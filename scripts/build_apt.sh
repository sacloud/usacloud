#!/bin/bash
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


USACLOUD_VERSION=$(grep -o -e "[0-9]\+.[0-9]\+.[0-9]\+-[0-9]" package/deb/debian/changelog | head -1 | sed 's/-.*$//')
export GPG_PRIVATE_KEY="`cat usacloud_gpg_key`"

set -e
set -x


: "prepare deb build..."
    rm -rf repos/debian
    mkdir -p repos/debian
    rm -rf package/deb-build
    mkdir -p package/deb-build
    cp -r package/deb package/deb-build/

    unzip -oq bin/usacloud_linux-amd64.zip -d bin/
    mv bin/usacloud package/deb-build/deb/debian/usacloud.bin
    cp contrib/completion/bash/usacloud package/deb-build/deb/debian/usacloud_bash_completion
    cp package/dummy-empty.tar.gz package/deb-build/usacloud_${USACLOUD_VERSION}.orig.tar.gz


: "building deb..."
	docker run --rm -v "$PWD/package/deb-build":/workdir sacloud/deb-build:latest
    # sign to Release file
	docker run --rm \
	    -v "$PWD/package/deb-build":/workdir \
	    -e GPG_PRIVATE_KEY \
	    -e GPG_PASSPHRASE \
	    -e GPG_FINGERPRINT \
	    -e GPG_NAME \
	    --entrypoint /sign_to_deb.sh \
	    sacloud/rpm-build

: "create apt repo..."
    cp package/deb-build/Release repos/debian/
    cp package/deb-build/Release.gpg repos/debian/
    cp package/deb-build/Packages repos/debian/
    cp package/deb-build/*.deb repos/debian/
