#!/bin/bash
set -ex

source scripts/common.sh

docker build -t stevemcquaid/$PACKAGE_NAME:latest .

# docker run --rm -v "$PWD/src":/usr/src/teflon -w /usr/src/teflon stevemcquaid/$PACKAGE_NAME:latest go build -v
