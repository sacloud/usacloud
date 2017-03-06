CURRENT_VERSION?=0.0.1alpha.11
TEST?=$$(go list ./... | grep -v vendor)
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOGEN_FILES?=$$(go list ./... | grep -v vendor)
BIN_NAME?=usacloud

BUILD_LDFLAGS = "-s -w \
	  -X github.com/sacloud/usacloud/version.Revision=`git rev-parse --short HEAD` \
	  -X github.com/sacloud/usacloud/version.Version=$(CURRENT_VERSION)"

.PHONY: default
default: test vet

.PHONY: run
run:
	go run $(CURDIR)/main.go $(ARGS)

.PHONY: clean
clean:
	rm -Rf bin/*

.PHONY: clean-all
clean-all:
	rm -Rf bin/* ; rm -Rf tools/bin/* ; rm -f command/*_gen.go

.PHONY: deps
deps:
	go get -u github.com/kardianos/govendor


.PHONY: tools
tools: tools/bin/*

tools/bin/gen-cli-commands: tools/gen-cli-commands/*.go
	go build -o $(CURDIR)/tools/bin/gen-cli-commands $(CURDIR)/tools/gen-cli-commands/*.go

tools/bin/gen-command-funcs: tools/gen-command-funcs/*.go
	go build -o $(CURDIR)/tools/bin/gen-command-funcs $(CURDIR)/tools/gen-command-funcs/*.go

tools/bin/gen-input-models: tools/gen-input-models/*.go
	go build -o $(CURDIR)/tools/bin/gen-input-models $(CURDIR)/tools/gen-input-models/*.go

tools/bin/gen-command-completion: tools/gen-command-completion/*.go
	go build -o $(CURDIR)/tools/bin/gen-command-completion $(CURDIR)/tools/gen-command-completion/*.go


.PHONY: gen-bash-completion
gen-bash-completion: gen tools/bin/gen-bash-completion
	tools/bin/gen-bash-completion

tools/bin/gen-bash-completion: tools/gen-bash-completion/*.go
	go build -o $(CURDIR)/tools/bin/gen-bash-completion $(CURDIR)/tools/gen-bash-completion/*.go

contrib/completion/bash/usacloud: define/*.go gen-bash-completion

.PHONY: gen
gen: tools command/*_gen.go

.PHONY: gen-force
gen-force: clean-all tools
	go generate $(GOGEN_FILES); gofmt -s -l -w $(GOFMT_FILES)

command/*_gen.go: define/*.go tools/gen-cli-commands/*.go tools/gen-command-funcs/*.go tools/gen-input-models/*.go
	go generate $(GOGEN_FILES); gofmt -s -l -w $(GOFMT_FILES)

.PHONY: build build-x build-darwin build-windows build-linux

build: clean gen vet contrib/completion/bash/usacloud
	OS="`go env GOOS`" ARCH="`go env GOARCH`" ARCHIVE= BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

build-x: build-darwin build-windows build-linux

build-darwin: clean gen vet contrib/completion/bash/usacloud
	OS="darwin"  ARCH="amd64"     ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

build-windows: clean gen vet contrib/completion/bash/usacloud
	OS="windows" ARCH="386 amd64"     ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

build-linux: clean gen vet contrib/completion/bash/usacloud
	OS="linux"   ARCH="amd64 arm" ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) sh -c "'$(CURDIR)/scripts/build.sh'"

.PHONY: rpm deb
rpm: build-linux
	CURRENT_VERSION="$(CURRENT_VERSION)" sh -c "'$(CURDIR)/scripts/build_rpm.sh'"

deb: rpm
	CURRENT_VERSION="$(CURRENT_VERSION)" sh -c "'$(CURDIR)/scripts/build_apt.sh'"


.PHONY: test
test: vet
	go test $(TEST) $(TESTARGS) -v -timeout=30m -parallel=4 ;

.PHONY: vet
vet: fmt gen
	@echo "go tool vet $(VETARGS) ."
	@go tool vet $(VETARGS) $$(ls -d */ | grep -v vendor) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

.PHONY: golint
golint: fmt
	golint ./...

.PHONY: fmt
fmt:
	gofmt -s -l -w $(GOFMT_FILES)

.PHONY: docker-run docker-test docker-build docker-rpm
docker-run:
	sh -c "$(CURDIR)/scripts/build_docker_image.sh" ; \
	sh -c "$(CURDIR)/scripts/run_on_docker.sh"

docker-test:
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'test'"

docker-build: clean
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'build-x'"

docker-rpm: clean
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'rpm'"

