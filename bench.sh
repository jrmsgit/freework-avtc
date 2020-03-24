#!/bin/sh
set -eu
ARGS=${@}
if test "X" = "X${ARGS}"; then
	ARGS='./cmd/gen/bench_test.go ./cmd/export/bench_test.go'
fi
echo -n "generate log data... "
./gen.sh
mkdir -p ./cmd/export/out
mv ./out/log.txt ./cmd/export/out
echo "done!"
for fn in $(ls ${ARGS}); do
	echo "--- ${fn}"
	go test -benchmem -bench=. ${fn}
done
exit 0
