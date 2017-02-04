TEST?=$$(go list ./... | grep -v vendor)
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOGEN_FILES?=$$(go list ./... | grep -v vendor)
BIN_NAME?=usacloud

default: test vet

run:
	go run $(CURDIR)/main.go $(ARGS)

clean:
	rm -Rf bin/*

clean-all:
	rm -Rf bin/* ; rm -Rf tools/bin/* ; rm -f command/*_gen.go

tools: tools/bin/gen-command-funcs tools/bin/gen-input-models tools/bin/gen-cli-commands

tools/bin/gen-cli-commands: tools/gen-cli-commands/*.go
	go build -o $(CURDIR)/tools/bin/gen-cli-commands $(CURDIR)/tools/gen-cli-commands/*.go

tools/bin/gen-command-funcs: tools/gen-command-funcs/*.go
	go build -o $(CURDIR)/tools/bin/gen-command-funcs $(CURDIR)/tools/gen-command-funcs/*.go

tools/bin/gen-input-models: tools/gen-input-models/*.go
	go build -o $(CURDIR)/tools/bin/gen-input-models $(CURDIR)/tools/gen-input-models/*.go

gen: tools command/*_gen.go

gen-force: clean-all tools
	go generate $(GOGEN_FILES)

command/*_gen.go: define/*.go tools/gen-cli-commands/*.go tools/gen-command-funcs/*.go tools/gen-input-models/*.go
	go generate $(GOGEN_FILES)

build: clean gen vet
	go build -ldflags "-s -w -X `go list ./version`.Revision=`git rev-parse --short HEAD 2>/dev/null`" -o $(CURDIR)/bin/$(BIN_NAME) $(CURDIR)/main.go

build-x: clean vet
	sh -c "'$(CURDIR)/scripts/build.sh' '$(BIN_NAME)'"

test: vet
	go test $(TEST) $(TESTARGS) -v -timeout=30m -parallel=4 ;

vet: fmt gen
	@echo "go tool vet $(VETARGS) ."
	@go tool vet $(VETARGS) $$(ls -d */ | grep -v vendor) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

golint: fmt
	golint ./...

fmt:
	gofmt -s -l -w $(GOFMT_FILES)

docker-run: 
	sh -c "'$(CURDIR)/scripts/build_docker_image.sh' '$(BIN_NAME)'" ; \
	sh -c "'$(CURDIR)/scripts/run_on_docker.sh' '$(BIN_NAME)'"

docker-daemon:
	sh -c "'$(CURDIR)/scripts/build_docker_image.sh' '$(BIN_NAME)'" ; \
	sh -c "'$(CURDIR)/scripts/run_on_docker_daemon.sh' '$(BIN_NAME)'"

docker-logs:
	docker logs -f $(BIN_NAME)

docker-rm:
	docker rm -f $(BIN_NAME)

docker-test: 
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'test'"

docker-build: clean 
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'build-x'"


.PHONY: default build run clean test vet fmt golint
