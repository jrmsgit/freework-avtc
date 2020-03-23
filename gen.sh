#!/bin/sh
set -eu
go run ./cmd/gen/main.go >out/log.txt
echo "out/log.txt file created"
exit 0
