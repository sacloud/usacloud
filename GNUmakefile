#
# Copyright 2017-2021 The Usacloud Authors
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
#
GO_FILES        ?=$(shell find . -name '*.go')
AUTHOR          ?="The Usacloud Authors"
COPYRIGHT_YEAR  ?="2017-2021"
COPYRIGHT_FILES ?=$$(find . \( -name "*.dockerfile" -or -name "*.go" -or -name "*.sh" -or -name "*.pl" -or -name "*.bats" -or -name "*.bash" \) -print | grep -v "/vendor/")
BUILD_LDFLAGS   ?= "-s -w -X github.com/sacloud/usacloud/pkg/version.Revision=`git rev-parse --short HEAD`"

export GO111MODULE=on
export GOPROXY=https://proxy.golang.org

.PHONY: default
default: gen lint test build

.PHONY: run
run:
	go run $(CURDIR)/main.go $(ARGS)

.PHONY: clean
clean:
	rm -Rf bin/*

.PHONY: clean-all
clean-all:
	rm -Rf bin/* ; rm -Rf tools/bin/* ; find . -name "*_gen.go" -type f | xargs rm -rf

.PHONY: tools
tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/sacloud/addlicense@latest
	go install github.com/client9/misspell/cmd/misspell@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/v1.38.0/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.38.0

.PHONY: gen _gen
gen: _gen set-license fmt goimports

_gen:
	go generate ./...

.PHONY: gen-force
gen-force: clean-all gen


.PHONY: build
build: bin/usacloud

bin/usacloud: $(GO_FILES)
	GOOS=$${OS:-"`go env GOOS`"} GOARCH=$${ARCH:-"`go env GOARCH`"} CGO_ENABLED=0 go build -ldflags=$(BUILD_LDFLAGS) -o bin/usacloud main.go

.PHONY: shasum
shasum:
	(cd bin/; shasum -a 256 * > usacloud_SHA256SUMS)

.PHONY: test
test: 
	go test $(TESTARGS) -v ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: goimports
goimports:
	goimports -l -w pkg/ tools/

.PHONY: fmt
fmt:
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

set-license:
	@addlicense -c $(AUTHOR) -y $(COPYRIGHT_YEAR) $(COPYRIGHT_FILES)

.SUFFIXES:
.SUFFIXES: .go