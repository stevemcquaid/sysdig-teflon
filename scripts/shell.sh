#!/bin/bash
set -ex

source scripts/common.sh

docker run -it  -p 80:80 -v $PWD:/src --rm stevemcquaid/$PACKAGE_NAME:latest /bin/bash
