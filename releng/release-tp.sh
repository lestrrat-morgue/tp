#!/bin/bash

set -e

if [ -z "$GITHUB_TOKEN" ]; then
    echo "GITHUB_TOKEN environment variable must be set"
    exit 1
fi

if [ -z "$GITHUB_USERNAME" ]; then
    echo "GITHUB_USERNAME environment variable must be set"
    exit 1
fi

if [ -z "$TP_VERSION" ]; then
    echo "TP_VERSION environment variable must be set"
    exit 1
fi

# Change directory to the project because that makes
# things much easier
cd /work/src/github.com/lestrrat/tp

/build-tp.sh
ghr --debug -u "$GITHUB_USERNAME" $TP_VERSION /work/artifacts/snapshot