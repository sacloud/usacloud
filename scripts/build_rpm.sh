#!/bin/bash
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


SOURCE_DIR="package/rpm"
DESTINATION_DIR="package/rpm-build"
export GPG_PRIVATE_KEY="`cat usacloud_gpg_key`"

set -e
set -x

: "prepare rpm build..."
    rm -rf rpmbuild/
    mkdir -p rpmbuild/RPMS/{x86_64,noarch}
    rm -rf repos/centos
    mkdir -p repos/centos/{x86_64,noarch}
    rm -rf "${DESTINATION_DIR}"
    mkdir -p "${DESTINATION_DIR}/src"
    cp contrib/completion/bash/usacloud "${DESTINATION_DIR}/src/usacloud_bash_completion"
    cp "${SOURCE_DIR}/usacloud.spec" "${DESTINATION_DIR}/usacloud.spec"

: "building i386..."
    unzip -oq bin/usacloud_linux-386.zip -d bin/
	docker run --rm \
	    -v "$PWD":/workdir \
	    -v "$PWD/rpmbuild":/rpmbuild \
	    sacloud/rpm-build:latest \
	        --define "_sourcedir /workdir/package/rpm-build/src" \
	        --define "_builddir /workdir/bin" \
	        --define "_version ${CURRENT_VERSION}" \
	        --define "buildarch noarch" \
	        --target noarch \
	        -bb package/rpm-build/usacloud.spec

    # sign to rpm
	docker run --rm \
	    -v "$PWD/rpmbuild":/rpmbuild \
	    -e GPG_PRIVATE_KEY \
	    -e GPG_PASSPHRASE \
	    -e GPG_FINGERPRINT \
	    -e GPG_NAME \
	    --entrypoint /sign_to_rpm.sh \
	    --workdir /rpmbuild/RPMS/noarch \
	    sacloud/rpm-build:latest

: "building x86_64..."
    unzip -oq bin/usacloud_linux-amd64.zip -d bin/
	docker run --rm \
	    -v "$PWD":/workdir \
	    -v "$PWD/rpmbuild":/rpmbuild \
	    sacloud/rpm-build:latest \
	        --define "_sourcedir /workdir/package/rpm-build/src" \
	        --define "_builddir /workdir/bin" \
	        --define "_version ${CURRENT_VERSION}" \
	        --define "buildarch x86_64" \
	        --target x86_64 \
	        -bb package/rpm-build/usacloud.spec

    # sign to rpm
	docker run --rm \
	    -v "$PWD/rpmbuild":/rpmbuild \
	    -e GPG_PRIVATE_KEY \
	    -e GPG_PASSPHRASE \
	    -e GPG_FINGERPRINT \
	    -e GPG_NAME \
	    --entrypoint /sign_to_rpm.sh \
	    --workdir /rpmbuild/RPMS/x86_64 \
	    sacloud/rpm-build:latest

: "create yum repo..."
    cp -rf rpmbuild/RPMS/noarch/* repos/centos/noarch/
    cp -rf rpmbuild/RPMS/x86_64/* repos/centos/x86_64/
	docker run --rm \
	    -v "$PWD/repos/centos/noarch":/workdir \
	    --entrypoint createrepo \
	    sacloud/rpm-build:latest \
	        -v /workdir
	docker run --rm \
	    -v "$PWD/repos/centos/x86_64":/workdir \
	    --entrypoint createrepo \
	    sacloud/rpm-build:latest \
	        -v /workdir

