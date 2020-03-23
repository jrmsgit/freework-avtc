#!/bin/sh
set -eu
mkdir -p ./out
go build -mod=vendor -o ./build/gen -i ./cmd/gen
./build/gen >./out/log.txt
echo "./out/log.txt file created"
exit 0
