#!/bin/bash

set -e

DOCKER_IMAGE_NAME="usacloud-integration-test"
DOCKER_CONTAINER_NAME="usacloud-integration-test-container"

if [[ $(docker ps -a | grep $DOCKER_CONTAINER_NAME) != "" ]]; then
  docker rm -f $DOCKER_CONTAINER_NAME 2>/dev/null
fi

docker build -f scripts/Dockerfile.bats -t $DOCKER_IMAGE_NAME .

docker run --rm --name $DOCKER_CONTAINER_NAME \
       -e SAKURACLOUD_ACCESS_TOKEN \
       -e SAKURACLOUD_ACCESS_TOKEN_SECRET \
       -e SAKURACLOUD_ZONE \
       -e USACLOUD_PROFILE \
       $DOCKER_IMAGE_NAME

