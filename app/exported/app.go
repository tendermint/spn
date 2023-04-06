package exported

import (
	"io"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	appparams "github.com/ignite/modules/app/params"
)

type (
	// AppBuilder is a method that allows to build an app
	AppBuilder func(
		logger log.Logger,
		db dbm.DB,
		traceStore io.Writer,
		loadLatest bool,
		skipUpgradeHeights map[int64]bool,
		homePath string,
		invCheckPeriod uint,
		encodingConfig appparams.EncodingConfig,
		appOpts servertypes.AppOptions,
		baseAppOptions ...func(*baseapp.BaseApp),
	) App

	// App represents a Cosmos SDK application that can be run as a server and with an exportable state
	App interface {
		servertypes.Application
		ExportableApp
	}

	// ExportableApp represents an app with an exportable state
	ExportableApp interface {
		ExportAppStateAndValidators(
			forZeroHeight bool,
			jailAllowedAddrs []string,
			modulesToExport []string,
		) (servertypes.ExportedApp, error)
		LoadHeight(height int64) error
		Name() string
		LegacyAmino() *codec.LegacyAmino
		BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock
		EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock
		InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain
		SimulationManager() *module.SimulationManager
	}
)
