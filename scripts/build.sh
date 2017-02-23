#!/bin/bash

set -e

OS="darwin linux windows"
ARCH="amd64 386"

rm -Rf bin/
mkdir bin/

for GOOS in $OS; do
    for GOARCH in $ARCH; do
        arch="$GOOS-$GOARCH"
        binary="usacloud"
        if [ "$GOOS" = "windows" ]; then
          binary="${binary}.exe"
        fi
        echo "Building $binary $arch"
        GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 \
            go build \
                -ldflags "-s -w  -X `go list ./version`.Revision=`git rev-parse --short HEAD 2>/dev/null`" \
                -o $binary \
                main.go
        zip -r "bin/usacloud_$arch" $binary
        rm -f $binary
    done
done
