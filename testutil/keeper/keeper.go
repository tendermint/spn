package keeper

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	ibctransfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	"github.com/stretchr/testify/require"
	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	monitoringcmodulekeeper "github.com/tendermint/spn/x/monitoringc/keeper"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	// ExampleTimestamp is a timestamp used as the current time for the context of the keepers returned from the package
	ExampleTimestamp = time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC)

	moduleAccountPerms = map[string][]string{
		authtypes.FeeCollectorName:  nil,
		minttypes.ModuleName:        {authtypes.Minter},
		ibctransfertypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		campaigntypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
	}
)

// AllKeepers returns initialized instances of all the keepers of the module
func AllKeepers(t testing.TB) (
	*campaignkeeper.Keeper,
	*launchkeeper.Keeper,
	*profilekeeper.Keeper,
	*monitoringcmodulekeeper.Keeper,
	bankkeeper.Keeper,
	sdk.Context,
) {
	initializer := newInitializer()

	paramKeeper := initializer.Param()
	capabilityKeeper := initializer.Capability()

	authKeeper := initializer.Auth(paramKeeper)
	bankKeeper := initializer.Bank(paramKeeper, authKeeper)

	stakingkeeper := initializer.Staking(authKeeper, bankKeeper, paramKeeper)
	ibcKeeper := initializer.IBC(paramKeeper, stakingkeeper, *capabilityKeeper)

	profileKeeper := initializer.Profile()
	launchKeeper := initializer.Launch(profileKeeper, paramKeeper)
	campaignKeeper := initializer.Campaign(launchKeeper, profileKeeper, bankKeeper)
	launchKeeper.SetCampaignKeeper(campaignKeeper)
	monitoringConsumerKeeper := initializer.Monitoringc(*ibcKeeper, *capabilityKeeper, launchKeeper, paramKeeper)
	require.NoError(t, initializer.StateStore.LoadLatestVersion())

	// Create a context using a custom timestamp
	ctx := sdk.NewContext(initializer.StateStore, tmproto.Header{
		Time: ExampleTimestamp,
	}, false, log.NewNopLogger())

	// Initialize params
	launchKeeper.SetParams(ctx, launchtypes.DefaultParams())

	return campaignKeeper, launchKeeper, profileKeeper, monitoringConsumerKeeper, bankKeeper, ctx
}

// Profile returns a keeper of the profile module for testing purpose
func Profile(t testing.TB) (*profilekeeper.Keeper, sdk.Context) {
	initializer := newInitializer()

	keeper := initializer.Profile()
	require.NoError(t, initializer.StateStore.LoadLatestVersion())

	return keeper, sdk.NewContext(initializer.StateStore, tmproto.Header{}, false, log.NewNopLogger())
}

// Launch returns a keeper of the launch module for testing purpose
func Launch(t testing.TB) (*launchkeeper.Keeper, sdk.Context) {
	initializer := newInitializer()

	paramKeeper := initializer.Param()
	profileKeeper := initializer.Profile()
	launchKeeper := initializer.Launch(profileKeeper, paramKeeper)
	require.NoError(t, initializer.StateStore.LoadLatestVersion())

	// Create a context using a custom timestamp
	ctx := sdk.NewContext(initializer.StateStore, tmproto.Header{
		Time: ExampleTimestamp,
	}, false, log.NewNopLogger())

	// Initialize params
	launchKeeper.SetParams(ctx, launchtypes.DefaultParams())

	return launchKeeper, ctx
}

// Campaign returns a keeper of the campaign module for testing purpose
func Campaign(t testing.TB) (*campaignkeeper.Keeper, sdk.Context) {
	campaignKeeper, _, _, _, _, ctx := AllKeepers(t) // nolint
	return campaignKeeper, ctx
}

// Monitoringc returns a keeper of the monitoring consumer module for testing purpose
func Monitoringc(t testing.TB) (*monitoringcmodulekeeper.Keeper, sdk.Context) {
	_, _, _, monitoringcKeeper, _, ctx := AllKeepers(t) // nolint
	return monitoringcKeeper, ctx
}
