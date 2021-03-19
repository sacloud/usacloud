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

FROM golang:1.16 AS builder
MAINTAINER Usacloud Authors <sacloud.users@gmail.com>

RUN  apt-get update && apt-get -y install \
        bash \
        git  \
        make \
        zip  \
        bzr  \
      && apt-get clean \
      && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*

ADD . /go/src/github.com/sacloud/usacloud
WORKDIR /go/src/github.com/sacloud/usacloud
ENV CGO_ENABLED 0
RUN make tools build
# ======

FROM alpine:3.12
MAINTAINER Usacloud Authors <sacloud.users@gmail.com>

RUN apk add --no-cache --update ca-certificates
COPY --from=builder /go/src/github.com/sacloud/usacloud/bin/usacloud /usr/bin/

ENTRYPOINT ["/usr/bin/usacloud"]