#!/bin/bash

set -e

TPDIR=$(cd $(dirname $0)/..; pwd -P)

if [ -z "$TP_VERSION" ]; then
    echo "TP_VERSION must be specified"
    exit 1
fi

if [ -z "$GITHUB_TOKEN_FILE" ]; then
    GITHUB_TOKEN_FILE=github_token
fi

if [ ! -e "$GITHUB_TOKEN_FILE" ]; then
    echo "GITHUB_TOKEN_FILE does not exist"
    exit 1
fi

docker run --rm \
    -v $TPDIR:/work/src/github.com/lestrrat/tp/ \
    -e TP_VERSION=$TP_VERSION \
    -e GITHUB_USERNAME=lestrrat \
    -e GITHUB_TOKEN=`cat $GITHUB_TOKEN_FILE` \
    tp-docker \
    /release-tp.sh

