#
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
#
VETARGS         ?=-all
GOFMT_FILES     ?=$$(find . -name '*.go' | grep -v vendor)
GOGEN_FILES     ?=$$(go list ./... | grep -v vendor)
BIN_NAME        ?=usacloud
GO_FILES        ?=$(shell find . -name '*.go')
AUTHOR          ?="The Usacloud Authors"
COPYRIGHT_YEAR  ?="2017-2020"
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
	GO111MODULE=off go get golang.org/x/tools/cmd/goimports
	GO111MODULE=off go get golang.org/x/tools/cmd/stringer
	GO111MODULE=off go get github.com/sacloud/addlicense
	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/v1.23.8/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.23.8

.PHONY: gen _gen
gen: _gen set-license fmt goimports

_gen:
	go generate ./...

.PHONY: gen-force
gen-force: clean-all gen


.PHONY: build build-x build-darwin build-windows build-linux
build: bin/usacloud

bin/usacloud: $(GO_FILES)
	OS="`go env GOOS`" ARCH="`go env GOARCH`" ARCHIVE= BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"


build-x: build-darwin build-windows build-linux build-bsd build-wasm

build-darwin: bin/usacloud_darwin-amd64.zip

build-windows: bin/usacloud_windows-386.zip bin/usacloud_windows-amd64.zip

build-linux: bin/usacloud_linux-386.zip bin/usacloud_linux-amd64.zip bin/usacloud_linux-arm.zip

build-bsd: bin/usacloud_freebsd-386.zip bin/usacloud_freebsd-amd64.zip

build-wasm: bin/usacloud_js-wasm.zip

bin/usacloud_darwin-amd64.zip:
	OS="darwin"  ARCH="amd64"     ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

bin/usacloud_windows-386.zip:
	OS="windows" ARCH="386"     ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

bin/usacloud_windows-amd64.zip:
	OS="windows" ARCH="amd64"     ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

bin/usacloud_linux-386.zip:
	OS="linux"   ARCH="386" ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

bin/usacloud_linux-amd64.zip:
	OS="linux"   ARCH="amd64" ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

bin/usacloud_linux-arm.zip:
	OS="linux"   ARCH="arm" ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

bin/usacloud_freebsd-386.zip:
	OS="freebsd"   ARCH="386" ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

bin/usacloud_freebsd-amd64.zip:
	OS="freebsd"   ARCH="amd64" ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

bin/usacloud_js-wasm.zip:
	OS="js"   ARCH="wasm" ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

.PHONY: rpm deb
rpm: build-linux
	CURRENT_VERSION="$(CURRENT_VERSION)" sh -c "'$(CURDIR)/scripts/build_rpm.sh'"

deb: rpm
	CURRENT_VERSION="$(CURRENT_VERSION)" sh -c "'$(CURDIR)/scripts/build_apt.sh'"

.PHONY: test
test: 
	go test $(TESTARGS) -v ./...

.PHONY: integration-test
integration-test: bin/usacloud
	test/integration/run-bats.sh test/integration/bats ;

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: goimports
goimports:
	goimports -l -w pkg/ tools/

.PHONY: fmt
fmt:
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

.PHONY: build-docs serve-docs lint-docs
build-docs:
	sh -c "'$(CURDIR)/scripts/build_docs.sh'"

serve-docs:
	sh -c "'$(CURDIR)/scripts/serve_docs.sh'"

lint-docs:
	sh -c "'$(CURDIR)/scripts/lint_docs.sh'"

.PHONY: docker-run docker-test docker-build docker-rpm
docker-run:
	sh -c "$(CURDIR)/scripts/build_docker_image.sh" ; \
	sh -c "$(CURDIR)/scripts/run_on_docker.sh"

docker-test:
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'test'"

docker-integration-test:
	sh -c "'$(CURDIR)/scripts/run_integration_test.sh'"

docker-build: clean
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'build-x'"

docker-rpm: clean
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'rpm'"

set-license:
	@addlicense -c $(AUTHOR) -y $(COPYRIGHT_YEAR) $(COPYRIGHT_FILES)


build-completion-test-image:
	GOOS=linux GOARCH=amd64 go build -o usacloud-linux main.go
	docker build -t usacloud-bash-completion -f scripts/completion-dev.dockerfile .
	rm -f usacloud-linux

run-completion-test: build-completion-test-image
	docker run -it --rm usacloud-bash-completion
