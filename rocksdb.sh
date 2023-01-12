ROCKSDB_PATH=/opt/homebrew/Cellar/rocksdb/7.8.3
CGO_CFLAGS="-I$ROCKSDB_PATH/include -target x86_64-apple-darwin22.1.0"
CGO_LDFLAGS="-L$ROCKSDB_PATH/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -L/opt/homebrew/Cellar/snappy/1.1.9/lib -L/opt/homebrew/Cellar/lz4/1.9.4/lib -L/opt/homebrew/Cellar/zstd/1.5.2/lib"

CGO_CFLAGS=$CGO_CFLAGS CGO_LDFLAGS=$CGO_LDFLAGS go test -ldflags '-X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb' -tags 'rocksdb' -bench BenchmarkOrderedKeys -count=3 -benchmem github.com/nghuyenthevinh2000/db-benchmark
# go get github.com/cosmos/gorocksdb