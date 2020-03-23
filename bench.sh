#!/bin/sh
set -eu
ARGS=${@}
if test "X" = "X${ARGS}"; then
	ARGS='./cmd/gen/bench_test.go'
fi
for fn in $(ls ${ARGS}); do
	echo "--- ${fn}"
	go test -benchmem -bench=. ${ARGS}
done
exit 0
