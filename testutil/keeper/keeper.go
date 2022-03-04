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
	fundraisingkeeper "github.com/tendermint/fundraising/x/fundraising/keeper"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	monitoringckeeper "github.com/tendermint/spn/x/monitoringc/keeper"
	monitoringctypes "github.com/tendermint/spn/x/monitoringc/types"
	participationkeeper "github.com/tendermint/spn/x/participation/keeper"
	participationtypes "github.com/tendermint/spn/x/participation/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	rewardkeeper "github.com/tendermint/spn/x/reward/keeper"
	rewardtypes "github.com/tendermint/spn/x/reward/types"
)

var (
	// ExampleTimestamp is a timestamp used as the current time for the context of the keepers returned from the package
	ExampleTimestamp = time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC)

	// ExampleHeight is a block height used as the current block height for the context of test keeper
	ExampleHeight = int64(1111)
)

// TestKeepers holds all keepers used during keeper tests for all modules
type TestKeepers struct {
	CampaignKeeper           *campaignkeeper.Keeper
	LaunchKeeper             *launchkeeper.Keeper
	ProfileKeeper            *profilekeeper.Keeper
	RewardKeeper             *rewardkeeper.Keeper
	MonitoringConsumerKeeper *monitoringckeeper.Keeper
	BankKeeper               bankkeeper.Keeper
	IBCKeeper                *ibckeeper.Keeper
	FundraisingKeeper        *fundraisingkeeper.Keeper
	ParticipationKeeper      *participationkeeper.Keeper
}

// TestMsgServers holds all message servers used during keeper tests for all modules
type TestMsgServers struct {
	ProfileSrv       profiletypes.MsgServer
	LaunchSrv        launchtypes.MsgServer
	CampaignSrv      campaigntypes.MsgServer
	RewardSrv        rewardtypes.MsgServer
	MonitoringcSrv   monitoringctypes.MsgServer
	ParticipationSrv participationtypes.MsgServer
}

// NewTestSetup returns initialized instances of all the keepers and message servers of the modules
func NewTestSetup(t testing.TB) (sdk.Context, TestKeepers, TestMsgServers) {
	initializer := newInitializer()

	paramKeeper := initializer.Param()
	capabilityKeeper := initializer.Capability()
	authKeeper := initializer.Auth(paramKeeper)
	bankKeeper := initializer.Bank(paramKeeper, authKeeper)
	stakingkeeper := initializer.Staking(authKeeper, bankKeeper, paramKeeper)
	ibcKeeper := initializer.IBC(paramKeeper, stakingkeeper, *capabilityKeeper)
	fundraisingKeeper := initializer.Fundraising(paramKeeper, authKeeper, bankKeeper, make(map[string]bool))

	profileKeeper := initializer.Profile()
	launchKeeper := initializer.Launch(profileKeeper, paramKeeper)
	campaignKeeper := initializer.Campaign(launchKeeper, profileKeeper, bankKeeper, paramKeeper)
	rewardKeeper := initializer.Reward(bankKeeper, profileKeeper, launchKeeper, paramKeeper)
	participationKeeper := initializer.Participation(paramKeeper, fundraisingKeeper, stakingkeeper)
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
		Time:   ExampleTimestamp,
		Height: ExampleHeight,
	}, false, log.NewNopLogger())

	// Initialize params
	launchKeeper.SetParams(ctx, launchtypes.DefaultParams())
	rewardKeeper.SetParams(ctx, rewardtypes.DefaultParams())
	campaignKeeper.SetParams(ctx, campaigntypes.DefaultParams())
	fundraisingKeeper.SetParams(ctx, fundraisingtypes.DefaultParams())
	participationKeeper.SetParams(ctx, participationtypes.DefaultParams())
	setIBCDefaultParams(ctx, ibcKeeper)

	profileSrv := profilekeeper.NewMsgServerImpl(*profileKeeper)
	launchSrv := launchkeeper.NewMsgServerImpl(*launchKeeper)
	campaignSrv := campaignkeeper.NewMsgServerImpl(*campaignKeeper)
	rewardSrv := rewardkeeper.NewMsgServerImpl(*rewardKeeper)
	monitoringcSrv := monitoringckeeper.NewMsgServerImpl(*monitoringConsumerKeeper)
	participationSrv := participationkeeper.NewMsgServerImpl(*participationKeeper)

	return ctx, TestKeepers{
			CampaignKeeper:           campaignKeeper,
			LaunchKeeper:             launchKeeper,
			ProfileKeeper:            profileKeeper,
			RewardKeeper:             rewardKeeper,
			MonitoringConsumerKeeper: monitoringConsumerKeeper,
			BankKeeper:               bankKeeper,
			IBCKeeper:                ibcKeeper,
			FundraisingKeeper:        fundraisingKeeper,
			ParticipationKeeper:      participationKeeper,
		}, TestMsgServers{
			ProfileSrv:       profileSrv,
			LaunchSrv:        launchSrv,
			CampaignSrv:      campaignSrv,
			RewardSrv:        rewardSrv,
			MonitoringcSrv:   monitoringcSrv,
			ParticipationSrv: participationSrv,
		}
}

// MonitoringcWithIBCMocks returns a keeper of the monitoring consumer module for testing purpose with mocks for IBC keepers
func MonitoringcWithIBCMocks(
	t testing.TB,
	connectionMock []Connection,
	channelMock []Channel,
) (*monitoringckeeper.Keeper, sdk.Context) {
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
	monitoringConsumerKeeper.SetParams(ctx, monitoringctypes.DefaultParams())

	return monitoringConsumerKeeper, ctx
}

// setIBCDefaultParams set default params for IBC client and connection keepers
func setIBCDefaultParams(ctx sdk.Context, ibcKeeper *ibckeeper.Keeper) {
	ibcKeeper.ClientKeeper.SetParams(ctx, ibcclienttypes.DefaultParams())
	ibcKeeper.ConnectionKeeper.SetParams(ctx, ibcconnectiontypes.DefaultParams())
	ibcKeeper.ClientKeeper.SetNextClientSequence(ctx, 0)
}
