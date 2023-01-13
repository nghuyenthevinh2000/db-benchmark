package simnode

import (
	"fmt"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/node"
	sm "github.com/tendermint/tendermint/state"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

func GetNode() {
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

func TendermintHandleGenesis() (dbm.DB, sm.State, *tmtypes.GenesisDoc, error) {
	db, err := GetDB()
	if err != nil {
		return nil, sm.State{}, nil, err
	}

	config := &config.Config{}
	config.Genesis = "genesis.json"
	genesisDocProvider := node.DefaultGenesisDocProviderFunc(config)

	state, genDoc, err := node.LoadStateFromDBOrGenesisDocProvider(db, genesisDocProvider)
	if err != nil {
		return nil, sm.State{}, nil, err
	}

	return db, state, genDoc, nil
}
