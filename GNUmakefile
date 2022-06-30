#
# Copyright 2017-2022 The sacloud/usacloud Authors
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
#====================
AUTHOR         ?= The sacloud/usacloud Authors
COPYRIGHT_YEAR ?= 2017-2022

BIN            ?= usacloud
BUILD_LDFLAGS   ?= "-s -w -X github.com/sacloud/usacloud/pkg/version.Revision=`git rev-parse --short HEAD`"

include includes/go/common.mk
include includes/go/single.mk
#====================

default: gen $(DEFAULT_GOALS)
tools: dev-tools

.SUFFIXES:
.SUFFIXES: .go

export GO111MODULE=on
export GOPROXY=https://proxy.golang.org

.PHONY: clean-all
clean-all:
	rm -Rf bin/* ; rm -Rf tools/bin/* ; find . -name "*_gen.go" -type f | xargs rm -rf

.PHONY: gen _gen
gen: _gen set-license fmt goimports

_gen:
	go generate ./...

.PHONY: gen-force
gen-force: clean-all gen

.PHONY: e2e-test
e2e-test: install
	@echo "[INFO] When you run e2e-test for the first time, run 'make tools' first."
	(cd e2e; go test $(TESTARGS) -v -tags=e2e -timeout 240m ./...)