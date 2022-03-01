// Package keeper provides methods to initialize SDK keepers with local storage for test purposes
package keeper

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	ibcclienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v2/modules/core/03-connection/types"
	ibckeeper "github.com/cosmos/ibc-go/v2/modules/core/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	monitoringcmodulekeeper "github.com/tendermint/spn/x/monitoringc/keeper"
	monitoringcmoduletypes "github.com/tendermint/spn/x/monitoringc/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	rewardkeeper "github.com/tendermint/spn/x/reward/keeper"
	rewardtypes "github.com/tendermint/spn/x/reward/types"
)

var (
	// ExampleTimestamp is a timestamp used as the current time for the context of the keepers returned from the package
	ExampleTimestamp = time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC)
)

type TestKeepers struct {
	CampaignKeeper           *campaignkeeper.Keeper
	LaunchKeeper             *launchkeeper.Keeper
	ProfileKeeper            *profilekeeper.Keeper
	RewardKeeper             *rewardkeeper.Keeper
	MonitoringConsumerKeeper *monitoringcmodulekeeper.Keeper
	BankKeeper               bankkeeper.Keeper
	IBCKeeper                *ibckeeper.Keeper
}

// NewTestKeepers returns initialized instances of all the keepers of the module
func NewTestKeepers(t testing.TB) (sdk.Context, TestKeepers) {
	initializer := newInitializer()

	paramKeeper := initializer.Param()
	capabilityKeeper := initializer.Capability()
	authKeeper := initializer.Auth(paramKeeper)
	bankKeeper := initializer.Bank(paramKeeper, authKeeper)
	stakingkeeper := initializer.Staking(authKeeper, bankKeeper, paramKeeper)
	ibcKeeper := initializer.IBC(paramKeeper, stakingkeeper, *capabilityKeeper)

	profileKeeper := initializer.Profile()
	launchKeeper := initializer.Launch(profileKeeper, paramKeeper)
	campaignKeeper := initializer.Campaign(launchKeeper, profileKeeper, bankKeeper, paramKeeper)
	rewardKeeper := initializer.Reward(bankKeeper, profileKeeper, launchKeeper, paramKeeper)
	launchKeeper.SetCampaignKeeper(campaignKeeper)
	monitoringConsumerKeeper := initializer.Monitoringc(
		*ibcKeeper,
		*capabilityKeeper,
		launchKeeper,
		rewardKeeper,
		paramKeeper,
		[]Connection{},
		[]Channel{},
	)
	require.NoError(t, initializer.StateStore.LoadLatestVersion())

	// Create a context using a custom timestamp
	ctx := sdk.NewContext(initializer.StateStore, tmproto.Header{
		Time: ExampleTimestamp,
	}, false, log.NewNopLogger())

	// Initialize params
	launchKeeper.SetParams(ctx, launchtypes.DefaultParams())
	rewardKeeper.SetParams(ctx, rewardtypes.DefaultParams())
	campaignKeeper.SetParams(ctx, campaigntypes.DefaultParams())
	setIBCDefaultParams(ctx, ibcKeeper)

	return ctx, TestKeepers{
		CampaignKeeper:           campaignKeeper,
		LaunchKeeper:             launchKeeper,
		ProfileKeeper:            profileKeeper,
		RewardKeeper:             rewardKeeper,
		MonitoringConsumerKeeper: monitoringConsumerKeeper,
		BankKeeper:               bankKeeper,
		IBCKeeper:                ibcKeeper,
	}

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
	ctx, tk := NewTestKeepers(t)
	return tk.CampaignKeeper, ctx
}

// Reward returns a keeper of the reward module for testing purpose
func Reward(t testing.TB) (*rewardkeeper.Keeper, sdk.Context) {
	ctx, tk := NewTestKeepers(t)

	// Initialize params
	tk.RewardKeeper.SetParams(ctx, rewardtypes.DefaultParams())

	return tk.RewardKeeper, ctx
}

// Monitoringc returns a keeper of the monitoring consumer module for testing purpose
func Monitoringc(t testing.TB) (*monitoringcmodulekeeper.Keeper, sdk.Context) {
	ctx, tk := NewTestKeepers(t)
	tk.MonitoringConsumerKeeper.SetParams(ctx, monitoringcmoduletypes.DefaultParams())

	return tk.MonitoringConsumerKeeper, ctx
}

// MonitoringcWithIBCMocks returns a keeper of the monitoring consumer module for testing purpose with mocks for IBC keepers
func MonitoringcWithIBCMocks(
	t testing.TB,
	connectionMock []Connection,
	channelMock []Channel,
) (*monitoringcmodulekeeper.Keeper, sdk.Context) {
	initializer := newInitializer()

	paramKeeper := initializer.Param()
	capabilityKeeper := initializer.Capability()
	authKeeper := initializer.Auth(paramKeeper)
	bankKeeper := initializer.Bank(paramKeeper, authKeeper)
	stakingkeeper := initializer.Staking(authKeeper, bankKeeper, paramKeeper)
	ibcKeeper := initializer.IBC(paramKeeper, stakingkeeper, *capabilityKeeper)

	profileKeeper := initializer.Profile()
	launchKeeper := initializer.Launch(profileKeeper, paramKeeper)
	campaignKeeper := initializer.Campaign(launchKeeper, profileKeeper, bankKeeper, paramKeeper)
	rewardKeeper := initializer.Reward(bankKeeper, profileKeeper, launchKeeper, paramKeeper)
	launchKeeper.SetCampaignKeeper(campaignKeeper)
	monitoringConsumerKeeper := initializer.Monitoringc(
		*ibcKeeper,
		*capabilityKeeper,
		launchKeeper,
		rewardKeeper,
		paramKeeper,
		connectionMock,
		channelMock,
	)
	require.NoError(t, initializer.StateStore.LoadLatestVersion())

	// Create a context using a custom timestamp
	ctx := sdk.NewContext(initializer.StateStore, tmproto.Header{
		Time: ExampleTimestamp,
	}, false, log.NewNopLogger())

	// Initialize params
	launchKeeper.SetParams(ctx, launchtypes.DefaultParams())
	monitoringConsumerKeeper.SetParams(ctx, monitoringcmoduletypes.DefaultParams())

	return monitoringConsumerKeeper, ctx
}

// setIBCDefaultParams set default params for IBC client and connection keepers
func setIBCDefaultParams(ctx sdk.Context, ibcKeeper *ibckeeper.Keeper) {
	ibcKeeper.ClientKeeper.SetParams(ctx, ibcclienttypes.DefaultParams())
	ibcKeeper.ConnectionKeeper.SetParams(ctx, ibcconnectiontypes.DefaultParams())
	ibcKeeper.ClientKeeper.SetNextClientSequence(ctx, 0)
}
