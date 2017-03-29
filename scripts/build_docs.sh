#!/bin/bash

set -e

DOCKER_IMAGE_NAME="usacloud-docs"
DOCKER_CONTAINER_NAME="usacloud-docs-container"

if [[ $(docker ps -a | grep $DOCKER_CONTAINER_NAME) != "" ]]; then
  docker rm -f $DOCKER_CONTAINER_NAME 2>/dev/null
fi

docker build -t $DOCKER_IMAGE_NAME -f scripts/Dockerfile.docs .

docker run --name $DOCKER_CONTAINER_NAME \
       -w /go/src/github.com/sacloud/usacloud/build_docs \
       $DOCKER_IMAGE_NAME

rm -rf docs/
docker cp $DOCKER_CONTAINER_NAME:/go/src/github.com/sacloud/usacloud/build_docs/site docs
docker rm -f $DOCKER_CONTAINER_NAME 2>/dev/null
