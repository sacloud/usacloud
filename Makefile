TEST?=$$(go list ./... | grep -v vendor)
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOGEN_FILES?=$$(go list ./... | grep -v vendor)
BIN_NAME?=usacloud

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

.PHONY: tools
tools: tools/bin/gen-command-funcs tools/bin/gen-input-models tools/bin/gen-cli-commands

tools/bin/gen-cli-commands: tools/gen-cli-commands/*.go
	go build -o $(CURDIR)/tools/bin/gen-cli-commands $(CURDIR)/tools/gen-cli-commands/*.go

tools/bin/gen-command-funcs: tools/gen-command-funcs/*.go
	go build -o $(CURDIR)/tools/bin/gen-command-funcs $(CURDIR)/tools/gen-command-funcs/*.go

tools/bin/gen-input-models: tools/gen-input-models/*.go
	go build -o $(CURDIR)/tools/bin/gen-input-models $(CURDIR)/tools/gen-input-models/*.go

.PHONY: gen
gen: tools command/*_gen.go

.PHONY: gen-force
gen-force: clean-all tools
	go generate $(GOGEN_FILES)

command/*_gen.go: define/*.go tools/gen-cli-commands/*.go tools/gen-command-funcs/*.go tools/gen-input-models/*.go
	go generate $(GOGEN_FILES)

.PHONY: build
build: clean gen vet
	go build -ldflags "-s -w -X `go list ./version`.Revision=`git rev-parse --short HEAD 2>/dev/null`" -o $(CURDIR)/bin/$(BIN_NAME) $(CURDIR)/main.go

.PHONY: build-x
build-x: clean gen vet
	sh -c "'$(CURDIR)/scripts/build.sh' '$(BIN_NAME)'"

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

.PHONY: docker-run
docker-run:
	sh -c "$(CURDIR)/scripts/build_docker_image.sh" ; \
	sh -c "$(CURDIR)/scripts/run_on_docker.sh"

.PHONY: docker-test
docker-test:
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'test'"

.PHONY: docker-build
docker-build: clean
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'build-x'"


