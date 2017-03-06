#!/bin/sh
set -e
set -x

: "prepare deb build...\n"
    rm -rf repos/debian; mkdir -p repos/debian
    rm -rf package/deb-build ; mkdir -p package/deb-build
    cp rpmbuild/RPMS/x86_64/* package/deb-build/

: "building deb...\n"
	docker run --rm -v "$(PWD)/package/deb-build":/workdir sacloud/usacloud:deb-build

: "create apt repo...\n"
    cp package/deb-build/Release repos/debian/
    cp package/deb-build/Packages repos/debian/
    cp package/deb-build/*.deb repos/debian/
