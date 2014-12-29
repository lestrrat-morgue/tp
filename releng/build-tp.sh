#!/bin/bash

DIR=/work/src/github.com/lestrrat/tp
pushd $DIR
goxc \
    -tasks "xc archive" \
    -bc "linux windows darwin" \
    -wd $DIR \
    -resources-include "README*,Changes" \
    -d /work/artifacts