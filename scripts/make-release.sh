#! /bin/bash

set -e

gox -verbose
mkdir -p release
mv dot-github_* release/
cd release
for bin in `ls`; do
    if [[ "$bin" == *windows* ]]; then
        command="dot-github.exe"
    else
        command="dot-github"
    fi
    mv "$bin" "$command"
    zip "${bin}.zip" "$command"
    rm "$command"
done
cd -
