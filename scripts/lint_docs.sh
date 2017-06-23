#!/bin/bash

set -e

DOCKER_IMAGE_NAME="sacloud/usacloud:textlint"

docker run -ti --rm \
       -v $PWD/build_docs:/workdir \
       $DOCKER_IMAGE_NAME .
