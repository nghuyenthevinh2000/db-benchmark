#!/usr/bin/make -f

GOFILES := $(shell ls -1 | grep .go | grep -v _test.go)

run-main-badgerdb:
	rm -rf testdb*
	go run -ldflags '-X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb' -tags 'badgerdb' $(GOFILES)

run-main-pebbledb:
	rm -rf testdb*
	go run -ldflags '-X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb' -tags 'pebbledb' $(GOFILES)

run-benchmark-badgerdb:
	go test -ldflags '-X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb' -tags 'badgerdb' -bench BenchmarkOrderedKeys -count=3 -benchmem github.com/nghuyenthevinh2000/db-benchmark

run-benchmark-pebbledb:
	go test -ldflags '-X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb' -tags 'pebbledb' -bench BenchmarkOrderedKeys -count=3 -benchmem github.com/nghuyenthevinh2000/db-benchmark