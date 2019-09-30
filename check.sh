#!/bin/bash

set -e

for f in *.go; do
    result="$(go run $f 2>&1)"
    diff -u <(echo "$result") <(cat "$f")
done