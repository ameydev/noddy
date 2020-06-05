#!/bin/bash

docker build -t $TAG_NAME $CONTEXT_PATH/
# verify kube-context
# sleep 3000
# # docker info
# PATH=$1
# TAG=$2
# sleep 6000
# docker build $PATH -t $TAG
# # sh -c "docker $*"