package app_test

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stretchr/testify/require"
	spnsim "github.com/tendermint/spn/testutil/simapp"
)

func BenchmarkSimulation(b *testing.B) {
	config, db, dir, logger, _, err := simapp.SetupSimulation("goleveldb-app-sim", "Simulation")
	require.NoError(b, err, "simulation setup failed")

	defer func() {
		db.Close()
		err = os.RemoveAll(dir)
		if err != nil {
			b.Fatal(err)
		}
	}()

	spnApp := spnsim.New(db, dir, logger)

	// Run randomized simulations
	_, simParams, simErr := simulation.SimulateFromSeed(
		b,
		os.Stdout,
		spnApp.GetBaseApp(),
		simapp.AppStateFn(spnApp.AppCodec(), spnApp.SimulationManager()),
		simulationtypes.RandomAccounts,
		simapp.SimulationOperations(spnApp, spnApp.AppCodec(), config),
		spnApp.ModuleAccountAddrs(),
		config,
		spnApp.AppCodec(),
	)

	// export state and simParams before the simulation error is checked
	if err = simapp.CheckExportSimulation(spnApp, config, simParams); err != nil {
		b.Fatal(err)
	}

	if simErr != nil {
		b.Fatal(simErr)
	}

	if config.Commit {
		simapp.PrintStats(db)
	}
}
