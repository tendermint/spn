package keeper

import (
	"testing"

	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"

	spntypes "github.com/tendermint/spn/pkg/types"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	monitoringptypes "github.com/tendermint/spn/x/monitoringp/types"
	participationkeeper "github.com/tendermint/spn/x/participation/keeper"
	participationtypes "github.com/tendermint/spn/x/participation/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	projectkeeper "github.com/tendermint/spn/x/project/keeper"
	projecttypes "github.com/tendermint/spn/x/project/types"
	rewardkeeper "github.com/tendermint/spn/x/reward/keeper"
	rewardtypes "github.com/tendermint/spn/x/reward/types"
)

// NewTestSetupWithMonitoringp returns a test keepers struct and servers struct with the monitoring provider module
func NewTestSetupWithMonitoringp(t testing.TB) (sdk.Context, TestKeepers, TestMsgServers) {
	return NewTestSetupWithIBCMocksMonitoringp(t, []Connection{}, []Channel{})
}

// NewTestSetupWithIBCMocksMonitoringp returns a keeper of the monitoring provider module for testing purpose with mocks for IBC keepers
func NewTestSetupWithIBCMocksMonitoringp(
	t testing.TB,
	connectionMock []Connection,
	channelMock []Channel,
) (sdk.Context, TestKeepers, TestMsgServers) {
	initializer := newInitializer()

	paramKeeper := initializer.Param()
	capabilityKeeper := initializer.Capability()
	authKeeper := initializer.Auth(paramKeeper)
	bankKeeper := initializer.Bank(paramKeeper, authKeeper)
	stakingKeeper := initializer.Staking(authKeeper, bankKeeper, paramKeeper)
	distrKeeper := initializer.Distribution(authKeeper, bankKeeper, stakingKeeper)
	upgradeKeeper := initializer.Upgrade()
	ibcKeeper := initializer.IBC(paramKeeper, stakingKeeper, *capabilityKeeper, upgradeKeeper)
	monitoringProviderKeeper := initializer.Monitoringp(
		stakingKeeper,
		*ibcKeeper,
		*capabilityKeeper,
		paramKeeper,
		connectionMock,
		channelMock,
	)
	fundraisingKeeper := initializer.Fundraising(paramKeeper, authKeeper, bankKeeper, distrKeeper)
	profileKeeper := initializer.Profile()
	launchKeeper := initializer.Launch(profileKeeper, distrKeeper, paramKeeper)
	rewardKeeper := initializer.Reward(authKeeper, bankKeeper, profileKeeper, launchKeeper, paramKeeper)
	projectKeeper := initializer.Project(launchKeeper, profileKeeper, bankKeeper, distrKeeper, paramKeeper)
	participationKeeper := initializer.Participation(paramKeeper, fundraisingKeeper, stakingKeeper)
	launchKeeper.SetProjectKeeper(projectKeeper)

	require.NoError(t, initializer.StateStore.LoadLatestVersion())

	// Create a context using a custom timestamp
	ctx := sdk.NewContext(initializer.StateStore, tmproto.Header{
		Time:   ExampleTimestamp,
		Height: ExampleHeight,
	}, false, log.NewNopLogger())

	// Initialize community pool
	distrKeeper.SetFeePool(ctx, distrtypes.InitialFeePool())

	// Initialize params
	distrKeeper.SetParams(ctx, distrtypes.DefaultParams())
	stakingKeeper.SetParams(ctx, stakingtypes.DefaultParams())
	launchKeeper.SetParams(ctx, launchtypes.DefaultParams())
	rewardKeeper.SetParams(ctx, rewardtypes.DefaultParams())
	projectKeeper.SetParams(ctx, projecttypes.DefaultParams())
	fundraisingKeeper.SetParams(ctx, fundraisingtypes.DefaultParams())
	participationKeeper.SetParams(ctx, participationtypes.DefaultParams())
	monitoringProviderKeeper.SetParams(ctx, monitoringptypes.DefaultParams())
	setIBCDefaultParams(ctx, ibcKeeper)

	profileSrv := profilekeeper.NewMsgServerImpl(*profileKeeper)
	launchSrv := launchkeeper.NewMsgServerImpl(*launchKeeper)
	projectSrv := projectkeeper.NewMsgServerImpl(*projectKeeper)
	rewardSrv := rewardkeeper.NewMsgServerImpl(*rewardKeeper)
	participationSrv := participationkeeper.NewMsgServerImpl(*participationKeeper)

	// set max shares - only set during app InitGenesis
	projectKeeper.SetTotalShares(ctx, spntypes.TotalShareNumber)

	return ctx, TestKeepers{
			T:                        t,
			ProjectKeeper:            projectKeeper,
			LaunchKeeper:             launchKeeper,
			ProfileKeeper:            profileKeeper,
			RewardKeeper:             rewardKeeper,
			MonitoringProviderKeeper: monitoringProviderKeeper,
			BankKeeper:               bankKeeper,
			IBCKeeper:                ibcKeeper,
			StakingKeeper:            stakingKeeper,
			FundraisingKeeper:        fundraisingKeeper,
			ParticipationKeeper:      participationKeeper,
		}, TestMsgServers{
			T:                t,
			ProfileSrv:       profileSrv,
			LaunchSrv:        launchSrv,
			ProjectSrv:       projectSrv,
			RewardSrv:        rewardSrv,
			ParticipationSrv: participationSrv,
		}
}
