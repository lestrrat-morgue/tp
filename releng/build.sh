#!/bin/bash

set -e

TPDIR=$(cd $(dirname $0)/..; pwd -P)
id=$(echo $(date) $$| shasum | awk '{print $1}')
docker run --rm \
    --name tp-build-$id \
    -v $TPDIR:/work/src/github.com/lestrrat/tp/ \
    -e RESULTSDIR=/work/artifacts \
    tp-docker \
    ./build-tp.sh