package main

import (
	"encoding/json"
	"os"

	"github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/node"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	sm "github.com/tendermint/tendermint/state"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/std"
	storemulti "github.com/cosmos/cosmos-sdk/store/rootmulti"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

//============Tendermint============

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

//============Cosmos============

var (
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
	)

	Modules = []string{
		authtypes.ModuleName,
	}

	Keys = sdk.NewKVStoreKeys(
		authtypes.StoreKey,
	)

	Encoding = GetEncodingConfig()
)

func CosmosHandleGenesis(db dbm.DB, genDoc *tmtypes.GenesisDoc) (*storemulti.Store, map[string]*types.KVStoreKey, error) {
	// get new db if there is no prior db
	if db == nil {
		var err error
		db, err = GetDB()
		if err != nil {
			return nil, nil, err
		}
	}

	// create new Store
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	store := storemulti.NewStore(db, logger)

	// mount store of each keeper
	for _, key := range Keys {
		store.MountStoreWithDB(key, types.StoreTypeIAVL, nil)
	}

	// load store to create new empty iavl store
	if err := store.LoadLatestVersion(); err != nil {
		return nil, nil, err
	}

	// genesis state handling
	var genesisState GenesisState
	if err := json.Unmarshal(genDoc.AppState, &genesisState); err != nil {
		panic(err)
	}

	// get a basic module manager and init genesis
	ctx := sdk.NewContext(store, tmproto.Header{}, false, logger)

	mm := NewModuleManager()
	for _, name := range Modules {
		mm.Modules[name].InitGenesis(ctx, Encoding.Marshaler, genesisState[name])
	}

	return store, Keys, nil
}

func NewModuleManager() *module.Manager {
	appCodec := Encoding.Marshaler

	accountKeeper := authkeeper.NewAccountKeeper(
		appCodec, Keys[authtypes.StoreKey], paramtypes.Subspace{}, authtypes.ProtoBaseAccount, nil,
	)

	mm := module.NewManager(
		auth.NewAppModule(appCodec, accountKeeper, nil),
	)

	mm.SetOrderInitGenesis(Modules...)

	return mm
}

func GetEncodingConfig() EncodingConfig {
	encodingConfig := MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	return encodingConfig
}
