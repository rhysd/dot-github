#! /bin/bash

set -e

gox -verbose
mkdir -p release
mv dot-github_* release/
cd release
for bin in `ls`; do
    mv "$bin" dot-github
    zip "${bin}.zip" dot-github
    rm dot-github
done
cd -
