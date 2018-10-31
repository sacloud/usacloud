TEST?=$$(go list ./... | grep -v vendor)
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOGEN_FILES?=$$(go list ./... | grep -v vendor)
BIN_NAME?=usacloud
CURRENT_VERSION = $$(gobump show -r version/)
NEXT_VERSION?= $$(git symbolic-ref --short HEAD | sed -n -E 's/^bump-version-([0-9]+\.[0-9]+\.[0-9]+)$$/\1/p')
GO_FILES?=$(shell find . -name '*.go')
export GO111MODULE=on

BUILD_LDFLAGS = "-s -w \
	  -X github.com/sacloud/usacloud/version.Revision=`git rev-parse --short HEAD` \
	  -X github.com/sacloud/usacloud/version.Version=$(CURRENT_VERSION)"

.PHONY: default
default: test build

.PHONY: run
run:
	go run $(CURDIR)/main.go $(ARGS)

.PHONY: clean
clean:
	rm -Rf bin/*

.PHONY: clean-all
clean-all:
	rm -Rf bin/* ; rm -Rf tools/bin/* ; rm -f command/*_gen.go; \
	rm -f command/cli/*_gen.go \
	rm -f command/completion/*_gen.go \
	rm -f command/funcs/*_gen.go \
	rm -f command/params/*_gen.go \
	rm -f contrib/completion/bash/usacloud


.PHONY: tools
tools:
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports
	GO111MODULE=off go get -u github.com/motemen/gobump/cmd/gobump
	GO111MODULE=off go get -u golang.org/x/lint/golint

contrib/completion/bash/usacloud: define/*.go
	go run tools/gen-bash-completion/main.go

.PHONY: gen
gen: command/cli/*_gen.go command/completion/*_gen.go command/funcs/*_gen.go command/params/*_gen.go

.PHONY: gen-force
gen-force: clean-all contrib/completion/bash/usacloud
	go generate $(GOGEN_FILES); gofmt -s -l -w $(GOFMT_FILES); goimports -l -w $(GOFMT_FILES)

command/*_gen.go: define/*.go tools/gen-cli-commands/*.go tools/gen-command-funcs/*.go tools/gen-input-models/*.go
	go generate $(GOGEN_FILES); gofmt -s -l -w $(GOFMT_FILES); goimports -l -w $(GOFMT_FILES)

.PHONY: build build-x build-darwin build-windows build-linux
build: bin/usacloud

bin/usacloud: contrib/completion/bash/usacloud $(GO_FILES)
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
	goimports -l -w $(GOFMT_FILES)

.PHONY: fmt
fmt:
	gofmt -s -l -w $(GOFMT_FILES)

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

# -----------------------------------------------
# for release
# -----------------------------------------------
.PHONY: bump-patch bump-minor bump-major
bump-patch:
	gobump patch -w version/

bump-minor:
	gobump minor -w version/

bump-major:
	gobump major -w version/

.PHONY: current-version next-version
current-version:
	@echo $(CURRENT_VERSION)

next-version:
	@echo $(NEXT_VERSION)

update-authors:
	@scripts/generate-authors.sh

.PHONY: create-release-pr
create-release-pr: build-docs update-authors
	$(eval CURRENT := $(shell echo $(CURRENT_VERSION)))
	gobump set "$(NEXT_VERSION)" -w version/ && \
	docker run --rm \
        -e APP_NAME=usacloud \
        -e REPO_NAME=sacloud/usacloud \
        -e ENABLE_RPM=1 \
        -e ENABLE_DEB=1 \
        -e ENABLE_PR=1 \
        -e RELEASE_FROM="$(CURRENT)" \
        -e RELEASE_TO="$(NEXT_VERSION)" \
        -e GITHUB_TOKEN \
        -v $(PWD):/workdir \
        sacloud/generate-changelog:latest

git-tag:
	git tag v$(CURRENT_VERSION)
