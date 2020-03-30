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
TEST            ?=$$(go list ./... | grep -v vendor)
VETARGS         ?=-all
GOFMT_FILES     ?=$$(find . -name '*.go' | grep -v vendor)
GOGEN_FILES     ?=$$(go list ./... | grep -v vendor)
BIN_NAME        ?=usacloud
CURRENT_VERSION := $(shell git log --merges --oneline | perl -ne 'if(m/^.+Merge pull request \#[0-9]+ from .+\/bump-version-([0-9\.]+)/){print $$1;exit}')
GO_FILES        ?=$(shell find . -name '*.go')
AUTHOR          ?="The Usacloud Authors"
COPYRIGHT_YEAR  ?="2017-2020"
COPYRIGHT_FILES ?=$$(find . \( -name "*.dockerfile" -or -name "*.go" -or -name "*.sh" -or -name "*.pl" -or -name "*.bats" -or -name "*.bash" \) -print | grep -v "/vendor/")

export GO111MODULE=on
export GOPROXY=https://proxy.golang.org

BUILD_LDFLAGS = "-s -w \
	  -X github.com/sacloud/usacloud/version.Revision=`git rev-parse --short HEAD` \
	  -X github.com/sacloud/usacloud/version.Version=$(CURRENT_VERSION)"

.PHONY: default
default: test build

.PHONY: run
run:
	go run $(CURDIR)/main.go $(ARGS)

.PHONY: run-v1
run-v1:
	go run $(CURDIR)/cmdv2/main.go $(ARGS)

.PHONY: clean
clean:
	rm -Rf bin/*

.PHONY: clean-all
clean-all:
	rm -Rf bin/* ; rm -Rf tools/bin/* ; rm -f command/*_gen.go; \
	rm -f command/cli/*_gen.go \
	rm -f command/funcs/*_gen.go \
	rm -f command/params/*_gen.go \
	rm -f cmdv2/commands/*_gen.go \
	rm -f cmdv2/params/*_gen.go \


.PHONY: tools
tools:
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports
	GO111MODULE=off go get -u github.com/motemen/gobump/cmd/gobump
	GO111MODULE=off go get -u golang.org/x/lint/golint
	GO111MODULE=off go get github.com/sacloud/addlicense

.PHONY: gen
gen: command/cli/*_gen.go command/funcs/*_gen.go command/params/*_gen.go

.PHONY: gen-force
gen-force: clean-all _gen-force set-license fmt goimports
_gen-force: 
	go generate $(GOGEN_FILES);

command/*_gen.go: define/*.go tools/gen-cli-commands/*.go tools/gen-command-funcs/*.go tools/gen-input-models/*.go
	go generate $(GOGEN_FILES); gofmt -s -l -w $(GOFMT_FILES)

.PHONY: build build-x build-darwin build-windows build-linux
build: bin/usacloud

bin/usacloud: $(GO_FILES)
	OS="`go env GOOS`" ARCH="`go env GOARCH`" ARCHIVE= BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

build-x: build-darwin build-windows build-linux build-bsd

build-darwin: bin/usacloud_darwin-amd64.zip

build-windows: bin/usacloud_windows-386.zip bin/usacloud_windows-amd64.zip

build-linux: bin/usacloud_linux-386.zip bin/usacloud_linux-amd64.zip bin/usacloud_linux-arm.zip

build-bsd: bin/usacloud_freebsd-386.zip bin/usacloud_freebsd-amd64.zip

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

.PHONY: rpm deb
rpm: build-linux
	CURRENT_VERSION="$(CURRENT_VERSION)" sh -c "'$(CURDIR)/scripts/build_rpm.sh'"

deb: rpm
	CURRENT_VERSION="$(CURRENT_VERSION)" sh -c "'$(CURDIR)/scripts/build_apt.sh'"

.PHONY: test
test: 
	go test $(TEST) $(TESTARGS) -v -timeout=30m -parallel=4 ;

.PHONY: integration-test
integration-test: bin/usacloud
	test/integration/run-bats.sh test/integration/bats ;

.PHONY: lint
lint: golint
	gometalinter --vendor --skip=vendor/ --disable-all --enable vet --enable goimports --deadline=5m ./...
	@echo

.PHONY: golint
golint: goimports
	for pkg in $$(go list ./... | grep -v /vendor/ ) ; do \
        test -z "$$(golint $$pkg | grep -v '_gen.go' | grep -v '_string.go' | grep -v 'should have comment' | grep -v 'func ServerMonitorCpu' | grep -v 'func ServerSsh' | grep -v 'DatabaseMonitorCpu' | grep -v "func MobileGatewayDnsUpdate" | tee /dev/stderr)" || RES=1; \
    done ;exit $$RES

.PHONY: goimports
goimports:
	find . -name '*.go' | grep -v vendor | xargs goimports -l -w

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

.PHONY: bump-patch bump-minor bump-major version
bump-patch:
	gobump patch -w

bump-minor:
	gobump minor -w

bump-major:
	gobump major -w

version:
	gobump show

git-tag:
	git tag v`gobump show -r`

set-license:
	@addlicense -c $(AUTHOR) -y $(COPYRIGHT_YEAR) $(COPYRIGHT_FILES)
