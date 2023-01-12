# db benchmark

this repository is used to perform actions on store directly

CGO_CFLAGS="-I/opt/homebrew/Cellar/rocksdb/7.8.3/include" CGO_LDFLAGS="-L/opt/homebrew/Cellar/rocksdb/7.8.3/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" go install -tags rocksdb ./...