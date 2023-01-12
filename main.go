package main

import (
	"fmt"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func main() {
	db, _, genDoc, err := TendermintHandleGenesis()
	if err != nil {
		panic(err)
	}

	store, keys, _ := CosmosHandleGenesis(db, genDoc)
	authstore := store.GetKVStore(keys[authtypes.StoreKey])
	iterator := authstore.Iterator(nil, nil)
	for ; iterator.Valid(); iterator.Next() {
		fmt.Printf("auth value = %v \n", iterator.Value())
	}
	iterator.Close()
}
