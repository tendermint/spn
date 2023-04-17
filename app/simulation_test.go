package app_test

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/app"
	"github.com/tendermint/spn/cmd"
)

func init() {
	simcli.GetSimulatorFlags()
}

type StoreKeysPrefixes struct {
	A        storetypes.StoreKey
	B        storetypes.StoreKey
	Prefixes [][]byte
}

// fauxMerkleModeOpt returns a BaseApp option to use a dbStoreAdapter instead of
// an IAVLStore for faster simulation speed.
func fauxMerkleModeOpt(bapp *baseapp.BaseApp) {
	bapp.SetFauxMerkleMode()
}

// go test ./app -v -benchmem -run=^$ -bench ^BenchmarkSimulation -Commit=true -cpuprofile cpu.out
func BenchmarkSimulation(b *testing.B) {
	simcli.FlagSeedValue = 10
	simcli.FlagVerboseValue = true

	config := simcli.NewConfigFromFlags()
	config.ChainID = app.DefaultChainID

	db, dir, logger, _, err := simtestutil.SetupSimulation(
		config,
		"leveldb-app-sim",
		"Simulation",
		simcli.FlagVerboseValue,
		simcli.FlagEnabledValue,
	)
	require.NoError(b, err, "simulation setup failed")

	b.Cleanup(func() {
		require.NoError(b, db.Close())
		require.NoError(b, os.RemoveAll(dir))
	})

	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = app.DefaultNodeHome
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue

	cmdApp := app.New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		0,
		cmd.MakeEncodingConfig(app.ModuleBasics),
		appOptions,
		baseapp.SetChainID(app.DefaultChainID),
	)

	bApp, ok := cmdApp.(*app.App)
	require.True(b, ok)

	// Run randomized simulations
	_, simParams, simErr := simulation.SimulateFromSeed(
		b,
		os.Stdout,
		bApp.GetBaseApp(),
		simtestutil.AppStateFn(
			bApp.AppCodec(),
			bApp.SimulationManager(),
			app.NewDefaultGenesisState(bApp.AppCodec()),
		),
		simtypes.RandomAccounts,
		simtestutil.SimulationOperations(bApp, bApp.AppCodec(), config),
		bApp.BlockedModuleAccountAddrs(),
		config,
		bApp.AppCodec(),
	)

	// export state and simParams before the simulation error is checked
	err = simtestutil.CheckExportSimulation(bApp, config, simParams)
	require.NoError(b, err)
	require.NoError(b, simErr)

	if config.Commit {
		simtestutil.PrintStats(db)
	}
}

func TestAppImportExport(t *testing.T) {
	config := simcli.NewConfigFromFlags()
	config.ChainID = app.DefaultChainID

	db, dir, logger, skip, err := simtestutil.SetupSimulation(
		config,
		"leveldb-app-sim",
		"Simulation",
		simcli.FlagVerboseValue,
		simcli.FlagEnabledValue,
	)
	if skip {
		t.Skip("skipping application import/export simulation")
	}
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		require.NoError(t, db.Close())
		require.NoError(t, os.RemoveAll(dir))
	}()

	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = app.DefaultNodeHome
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue

	cmdApp := app.New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		0,
		cmd.MakeEncodingConfig(app.ModuleBasics),
		appOptions,
		fauxMerkleModeOpt,
		baseapp.SetChainID(app.DefaultChainID),
	)
	bApp, ok := cmdApp.(*app.App)
	require.True(t, ok)

	// run randomized simulation
	_, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		bApp.GetBaseApp(),
		simtestutil.AppStateFn(
			bApp.AppCodec(),
			bApp.SimulationManager(),
			app.NewDefaultGenesisState(bApp.AppCodec()),
		),
		simtypes.RandomAccounts,
		simtestutil.SimulationOperations(bApp, bApp.AppCodec(), config),
		bApp.BlockedModuleAccountAddrs(),
		config,
		bApp.AppCodec(),
	)
	require.NoError(t, simErr)

	// export state and simParams before the simulation error is checked
	err = simtestutil.CheckExportSimulation(bApp, config, simParams)
	require.NoError(t, err)

	if config.Commit {
		simtestutil.PrintStats(db)
	}

	fmt.Printf("exporting genesis...\n")

	exported, err := bApp.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	newDB, newDir, _, _, err := simtestutil.SetupSimulation(
		config,
		"leveldb-app-sim-2",
		"Simulation-2",
		simcli.FlagVerboseValue,
		simcli.FlagEnabledValue,
	)
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		require.NoError(t, newDB.Close())
		require.NoError(t, os.RemoveAll(newDir))
	}()

	cmdNewApp := app.New(
		log.NewNopLogger(),
		newDB,
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		0,
		cmd.MakeEncodingConfig(app.ModuleBasics),
		appOptions,
		fauxMerkleModeOpt,
		baseapp.SetChainID(app.DefaultChainID),
	)

	newApp, ok := cmdNewApp.(*app.App)
	require.True(t, ok)

	var genesisState app.GenesisState
	err = json.Unmarshal(exported.AppState, &genesisState)
	require.NoError(t, err)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("%v", r)
			if !strings.Contains(err, "validator set is empty after InitGenesis") {
				panic(r)
			}
			logger.Info("Skipping simulation as all validators have been unbonded")
			logger.Info("err", err, "stacktrace", string(debug.Stack()))
		}
	}()

	ctxA := bApp.NewContext(true, tmproto.Header{Height: bApp.LastBlockHeight()})
	ctxB := newApp.NewContext(true, tmproto.Header{Height: bApp.LastBlockHeight()})
	newApp.InitGenesis(ctxB, bApp.AppCodec(), genesisState)
	newApp.StoreConsensusParams(ctxB, exported.ConsensusParams)

	fmt.Printf("comparing stores...\n")

	storeKeysPrefixes := []StoreKeysPrefixes{
		{bApp.GetKey(authtypes.StoreKey), newApp.GetKey(authtypes.StoreKey), [][]byte{}},
		{
			bApp.GetKey(stakingtypes.StoreKey), newApp.GetKey(stakingtypes.StoreKey),
			[][]byte{
				stakingtypes.UnbondingQueueKey, stakingtypes.RedelegationQueueKey, stakingtypes.ValidatorQueueKey,
				stakingtypes.HistoricalInfoKey, stakingtypes.UnbondingIDKey, stakingtypes.UnbondingIndexKey, stakingtypes.UnbondingTypeKey, stakingtypes.ValidatorUpdatesKey,
			},
		}, // ordering may change but it doesn't matter
		{bApp.GetKey(slashingtypes.StoreKey), newApp.GetKey(slashingtypes.StoreKey), [][]byte{}},
		{bApp.GetKey(minttypes.StoreKey), newApp.GetKey(minttypes.StoreKey), [][]byte{}},
		{bApp.GetKey(distrtypes.StoreKey), newApp.GetKey(distrtypes.StoreKey), [][]byte{}},
		{bApp.GetKey(banktypes.StoreKey), newApp.GetKey(banktypes.StoreKey), [][]byte{banktypes.BalancesPrefix}},
		{bApp.GetKey(paramstypes.StoreKey), newApp.GetKey(paramstypes.StoreKey), [][]byte{}},
		{bApp.GetKey(govtypes.StoreKey), newApp.GetKey(govtypes.StoreKey), [][]byte{}},
		{bApp.GetKey(evidencetypes.StoreKey), newApp.GetKey(evidencetypes.StoreKey), [][]byte{}},
		{bApp.GetKey(capabilitytypes.StoreKey), newApp.GetKey(capabilitytypes.StoreKey), [][]byte{}},
		{bApp.GetKey(authzkeeper.StoreKey), newApp.GetKey(authzkeeper.StoreKey), [][]byte{authzkeeper.GrantKey, authzkeeper.GrantQueuePrefix}},
	}

	for _, skp := range storeKeysPrefixes {
		storeA := ctxA.KVStore(skp.A)
		storeB := ctxB.KVStore(skp.B)

		failedKVAs, failedKVBs := sdk.DiffKVStores(storeA, storeB, skp.Prefixes)
		require.Equal(t, len(failedKVAs), len(failedKVBs), "unequal sets of key-values to compare")

		fmt.Printf("compared %d different key/value pairs between %s and %s\n", len(failedKVAs), skp.A, skp.B)
		require.Equal(t, 0, len(failedKVAs), simtestutil.GetSimulationLog(skp.A.Name(), bApp.SimulationManager().StoreDecoders, failedKVAs, failedKVBs))
	}
}

func TestAppSimulationAfterImport(t *testing.T) {
	config := simcli.NewConfigFromFlags()
	config.ChainID = app.DefaultChainID

	db, dir, logger, skip, err := simtestutil.SetupSimulation(
		config,
		"leveldb-app-sim",
		"Simulation",
		simcli.FlagVerboseValue,
		simcli.FlagEnabledValue,
	)
	if skip {
		t.Skip("skipping application simulation after import")
	}
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		require.NoError(t, db.Close())
		require.NoError(t, os.RemoveAll(dir))
	}()

	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = app.DefaultNodeHome
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue

	cmdApp := app.New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		0,
		cmd.MakeEncodingConfig(app.ModuleBasics),
		appOptions,
		fauxMerkleModeOpt,
		baseapp.SetChainID(app.DefaultChainID),
	)
	bApp, ok := cmdApp.(*app.App)
	require.True(t, ok)

	// run randomized simulation
	stopEarly, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		bApp.BaseApp,
		simtestutil.AppStateFn(
			bApp.AppCodec(),
			bApp.SimulationManager(),
			app.NewDefaultGenesisState(bApp.AppCodec()),
		),
		simtypes.RandomAccounts, // replace with own random account function if using keys other than secp256k1
		simtestutil.SimulationOperations(bApp, bApp.AppCodec(), config),
		bApp.BlockedModuleAccountAddrs(),
		config,
		bApp.AppCodec(),
	)
	require.NoError(t, simErr)

	// export state and simParams before the simulation error is checked
	err = simtestutil.CheckExportSimulation(bApp, config, simParams)
	require.NoError(t, err)

	if config.Commit {
		simtestutil.PrintStats(db)
	}

	if stopEarly {
		fmt.Println("can't export or import a zero-validator genesis, exiting test...")
		return
	}

	fmt.Printf("exporting genesis...\n")

	exported, err := bApp.ExportAppStateAndValidators(true, []string{}, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	newDB, newDir, _, _, err := simtestutil.SetupSimulation(
		config,
		"leveldb-app-sim-2",
		"Simulation-2",
		simcli.FlagVerboseValue,
		simcli.FlagEnabledValue,
	)
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		require.NoError(t, newDB.Close())
		require.NoError(t, os.RemoveAll(newDir))
	}()

	cmdNewApp := app.New(
		log.NewNopLogger(),
		newDB,
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		0,
		cmd.MakeEncodingConfig(app.ModuleBasics),
		appOptions,
		fauxMerkleModeOpt,
		baseapp.SetChainID(app.DefaultChainID),
	)
	newApp, ok := cmdNewApp.(*app.App)
	require.True(t, ok)

	newApp.InitChain(abci.RequestInitChain{
		ChainId:       app.DefaultChainID,
		AppStateBytes: exported.AppState,
	})

	_, _, err = simulation.SimulateFromSeed(
		t,
		os.Stdout,
		newApp.BaseApp,
		simtestutil.AppStateFn(
			bApp.AppCodec(),
			bApp.SimulationManager(),
			app.NewDefaultGenesisState(bApp.AppCodec()),
		),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		simtestutil.SimulationOperations(newApp, newApp.AppCodec(), config),
		newApp.BlockedModuleAccountAddrs(),
		config,
		bApp.AppCodec(),
	)
	require.NoError(t, err)
}

func TestAppStateDeterminism(t *testing.T) {
	if !simcli.FlagEnabledValue {
		t.Skip("skipping application simulation")
	}

	config := simcli.NewConfigFromFlags()
	config.InitialBlockHeight = 1
	config.ExportParamsPath = ""
	config.OnOperation = true
	config.AllInvariants = true
	config.ChainID = app.DefaultChainID

	var (
		r                    = rand.New(rand.NewSource(time.Now().Unix()))
		numSeeds             = 3
		numTimesToRunPerSeed = 5
		appHashList          = make([]json.RawMessage, numTimesToRunPerSeed)
	)

	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = app.DefaultNodeHome
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue

	for i := 0; i < numSeeds; i++ {
		config.Seed = r.Int63()

		for j := 0; j < numTimesToRunPerSeed; j++ {
			var logger log.Logger
			if simcli.FlagVerboseValue {
				logger = log.TestingLogger()
			} else {
				logger = log.NewNopLogger()
			}

			var (
				chainID = fmt.Sprintf("chain-id-%d-%d", i, j)
				db      = dbm.NewMemDB()
				cmdApp  = app.New(
					logger,
					db,
					nil,
					true,
					map[int64]bool{},
					app.DefaultNodeHome,
					simcli.FlagPeriodValue,
					cmd.MakeEncodingConfig(app.ModuleBasics),
					appOptions,
					fauxMerkleModeOpt,
					baseapp.SetChainID(chainID),
				)
			)
			config.ChainID = chainID

			bApp, ok := cmdApp.(*app.App)
			require.True(t, ok)

			fmt.Printf(
				"running non-determinism simulation; seed %d: %d/%d, attempt: %d/%d\n",
				config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
			)

			_, _, err := simulation.SimulateFromSeed(
				t,
				os.Stdout,
				bApp.BaseApp,
				simtestutil.AppStateFn(
					bApp.AppCodec(),
					bApp.SimulationManager(),
					app.NewDefaultGenesisState(bApp.AppCodec()),
				),
				simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
				simtestutil.SimulationOperations(bApp, bApp.AppCodec(), config),
				bApp.BlockedModuleAccountAddrs(),
				config,
				bApp.AppCodec(),
			)
			require.NoError(t, err)

			if config.Commit {
				simtestutil.PrintStats(db)
			}

			appHash := bApp.LastCommitID().Hash
			appHashList[j] = appHash

			if j != 0 {
				require.Equal(
					t, string(appHashList[0]), string(appHashList[j]),
					"non-determinism in seed %d: %d/%d, attempt: %d/%d\n", config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
				)
			}
		}
	}
}
