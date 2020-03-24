#!/bin/sh
set -eu
go build -mod=vendor -o ./build/export -i ./cmd/export
echo "export ./out/log.txt data"
./build/export
exit 0
