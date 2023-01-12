package main

import (
	"fmt"
	"testing"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

var kvStore storetypes.CommitKVStore

func init() {
	var err error
	kvStore, err = GetIAVLKVStore()
	if err != nil {
		panic(err.Error())
	}
}

var table = []struct {
	keysNum      uint64
	dataByteSize uint64
}{
	{
		keysNum:      uint64(1000),
		dataByteSize: uint64(8),
	},
}

func BenchmarkOrderedKeys(b *testing.B) {
	// set 1000 keys (8bytes) of value (8bytes) to kv store
	for _, v := range table {
		b.Run(fmt.Sprintf("running ordered keysNum = %v with dataByteSize = %v bytes", v.keysNum, v.dataByteSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// writing action
				PopulateStoreWithOrderedKeys(v.keysNum, v.dataByteSize, kvStore)

				// commit store to disk
				kvStore.Commit()

				// reading action
				iter := kvStore.Iterator(nil, nil)
				for ; iter.Valid(); iter.Next() {
					iter.Value()
				}
				iter.Close()
			}
		})
	}
}

func BenchmarkUnOrderedKeys(b *testing.B) {
	// set 1000 keys (8bytes) of value (8bytes) to kv store
	for _, v := range table {
		b.Run(fmt.Sprintf("running unordered keysNum = %v with dataByteSize = %v bytes", v.keysNum, v.dataByteSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// writing action
				PopulateStoreWithUnorderedKeys(v.keysNum, v.dataByteSize, kvStore)

				// commit store to disk
				kvStore.Commit()

				// reading action
				iter := kvStore.Iterator(nil, nil)
				for ; iter.Valid(); iter.Next() {
					iter.Value()
				}
				iter.Close()
			}
		})
	}
}
