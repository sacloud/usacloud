FROM golang:1.12
MAINTAINER Kazumichi Yamamoto <yamamoto.febc@gmail.com>

RUN  apt-get update && apt-get -y install \
        bash \
        git  \
        make \
        zip  \
        bats  \
        jq \
        genisoimage \
      && apt-get clean \
      && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*

ADD . /go/src/github.com/sacloud/usacloud
WORKDIR /go/src/github.com/sacloud/usacloud


