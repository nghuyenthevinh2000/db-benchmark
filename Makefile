#!/usr/bin/make -f

GOFILES := $(shell ls -1 | grep .go | grep -v _test.go)

run-main-badgerdb:
	go run -ldflags '-X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb' -tags 'badgerdb' $(GOFILES)

run-main-pebbledb:
	go run -ldflags '-X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb' -tags 'pebbledb' $(GOFILES)

run-main-rocksdb:
	CGO_CFLAGS="-I/opt/homebrew/Cellar/rocksdb/7.8.3/include" CGO_LDFLAGS="-L/opt/homebrew/Cellar/rocksdb/7.8.3/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" go run -ldflags '-X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb' -tags 'rocksdb' $(GOFILES)

run-benchmark-badgerdb:
	go test -ldflags '-X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb' -tags 'badgerdb' -bench BenchmarkOrderedKeys -count=3 -benchmem github.com/nghuyenthevinh2000/db-benchmark

run-benchmark-pebbledb:
	go test -ldflags '-X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb' -tags 'pebbledb' -bench BenchmarkOrderedKeys -count=3 -benchmem github.com/nghuyenthevinh2000/db-benchmark

run-benchmark-rocksdb:
	bash rocksdb.sh