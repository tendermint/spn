// Package keeper provides methods to initialize SDK keepers with local storage for test purposes
package keeper

import (
	"testing"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v7/modules/core/03-connection/types"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	"github.com/stretchr/testify/require"
	fundraisingkeeper "github.com/tendermint/fundraising/x/fundraising/keeper"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"

	claimkeeper "github.com/ignite/modules/x/claim/keeper"
	claimtypes "github.com/ignite/modules/x/claim/types"

	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/keeper/mocks"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	monitoringckeeper "github.com/tendermint/spn/x/monitoringc/keeper"
	monitoringctypes "github.com/tendermint/spn/x/monitoringc/types"
	monitoringpkeeper "github.com/tendermint/spn/x/monitoringp/keeper"
	participationkeeper "github.com/tendermint/spn/x/participation/keeper"
	participationtypes "github.com/tendermint/spn/x/participation/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	projectkeeper "github.com/tendermint/spn/x/project/keeper"
	projecttypes "github.com/tendermint/spn/x/project/types"
	rewardkeeper "github.com/tendermint/spn/x/reward/keeper"
	rewardtypes "github.com/tendermint/spn/x/reward/types"
)

var (
	// ExampleTimestamp is a timestamp used as the current time for the context of the keepers returned from the package
	ExampleTimestamp = time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC)

	// ExampleHeight is a block height used as the current block height for the context of test keeper
	ExampleHeight = int64(1111)
)

// HookMocks holds mocks for the module hooks
type HooksMocks struct {
	LaunchHooksMock *mocks.LaunchHooks
}

// TestKeepers holds all keepers used during keeper tests for all modules
type TestKeepers struct {
	T                        testing.TB
	ProjectKeeper            *projectkeeper.Keeper
	LaunchKeeper             *launchkeeper.Keeper
	ProfileKeeper            *profilekeeper.Keeper
	RewardKeeper             *rewardkeeper.Keeper
	MonitoringConsumerKeeper *monitoringckeeper.Keeper
	MonitoringProviderKeeper *monitoringpkeeper.Keeper
	AccountKeeper            authkeeper.AccountKeeper
	BankKeeper               bankkeeper.Keeper
	DistrKeeper              distrkeeper.Keeper
	IBCKeeper                *ibckeeper.Keeper
	StakingKeeper            *stakingkeeper.Keeper
	FundraisingKeeper        fundraisingkeeper.Keeper
	ParticipationKeeper      *participationkeeper.Keeper
	ClaimKeeper              *claimkeeper.Keeper
	HooksMocks               HooksMocks
}

// TestMsgServers holds all message servers used during keeper tests for all modules
type TestMsgServers struct {
	T                testing.TB
	ProfileSrv       profiletypes.MsgServer
	LaunchSrv        launchtypes.MsgServer
	ProjectSrv       projecttypes.MsgServer
	RewardSrv        rewardtypes.MsgServer
	MonitoringcSrv   monitoringctypes.MsgServer
	ParticipationSrv participationtypes.MsgServer
	ClaimSrv         claimtypes.MsgServer
}

// SetupOption represents an option that can be provided to NewTestSetup
type SetupOption func(*setupOptions)

// setupOptions represents the set of SetupOption
type setupOptions struct {
	LaunchHooksMock bool
}

// WithHooksMock sets a mock for the hooks in testing launch keeper
func WithLaunchHooksMock() func(*setupOptions) {
	return func(o *setupOptions) {
		o.LaunchHooksMock = true
	}
}

// NewTestSetup returns initialized instances of all the keepers and message servers of the modules
func NewTestSetup(t testing.TB, options ...SetupOption) (sdk.Context, TestKeepers, TestMsgServers) {
	// setup options
	var so setupOptions
	for _, option := range options {
		option(&so)
	}

	initializer := newInitializer()

	paramKeeper := initializer.Param()
	capabilityKeeper := initializer.Capability()
	authKeeper := initializer.Auth(paramKeeper)
	bankKeeper := initializer.Bank(paramKeeper, authKeeper)
	stakingKeeper := initializer.Staking(authKeeper, bankKeeper, paramKeeper)
	distrKeeper := initializer.Distribution(authKeeper, bankKeeper, stakingKeeper)
	upgradeKeeper := initializer.Upgrade()
	ibcKeeper := initializer.IBC(paramKeeper, stakingKeeper, *capabilityKeeper, upgradeKeeper)
	fundraisingKeeper := initializer.Fundraising(paramKeeper, authKeeper, bankKeeper, distrKeeper)
	profileKeeper := initializer.Profile()
	launchKeeper := initializer.Launch(profileKeeper, distrKeeper, paramKeeper)
	rewardKeeper := initializer.Reward(authKeeper, bankKeeper, profileKeeper, launchKeeper, paramKeeper)
	projectKeeper := initializer.Project(launchKeeper, profileKeeper, bankKeeper, distrKeeper, paramKeeper)
	participationKeeper := initializer.Participation(paramKeeper, fundraisingKeeper, stakingKeeper)
	launchKeeper.SetProjectKeeper(projectKeeper)
	monitoringConsumerKeeper := initializer.Monitoringc(
		*ibcKeeper,
		*capabilityKeeper,
		launchKeeper,
		rewardKeeper,
		paramKeeper,
		[]Connection{},
		[]Channel{},
	)
	launchKeeper.SetMonitoringcKeeper(monitoringConsumerKeeper)
	claimKeeper := initializer.Claim(paramKeeper, authKeeper, distrKeeper, bankKeeper)
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
	fundraisingParams := fundraisingtypes.DefaultParams()
	fundraisingParams.AuctionCreationFee = sdk.NewCoins()
	fundraisingKeeper.SetParams(ctx, fundraisingParams)
	participationKeeper.SetParams(ctx, participationtypes.DefaultParams())
	monitoringConsumerKeeper.SetParams(ctx, monitoringctypes.DefaultParams())
	claimKeeper.SetParams(ctx, claimtypes.DefaultParams())
	setIBCDefaultParams(ctx, ibcKeeper)

	// Set hooks
	var hooksMocks HooksMocks
	if so.LaunchHooksMock {
		launchHooksMock := mocks.NewLaunchHooks(t)
		launchKeeper = launchKeeper.SetHooks(launchHooksMock)
		hooksMocks.LaunchHooksMock = launchHooksMock
	}

	profileSrv := profilekeeper.NewMsgServerImpl(*profileKeeper)
	launchSrv := launchkeeper.NewMsgServerImpl(*launchKeeper)
	projectSrv := projectkeeper.NewMsgServerImpl(*projectKeeper)
	rewardSrv := rewardkeeper.NewMsgServerImpl(*rewardKeeper)
	monitoringcSrv := monitoringckeeper.NewMsgServerImpl(*monitoringConsumerKeeper)
	participationSrv := participationkeeper.NewMsgServerImpl(*participationKeeper)
	claimSrv := claimkeeper.NewMsgServerImpl(*claimKeeper)

	// set max shares - only set during app InitGenesis
	projectKeeper.SetTotalShares(ctx, spntypes.TotalShareNumber)

	return ctx, TestKeepers{
			T:                        t,
			ProjectKeeper:            projectKeeper,
			LaunchKeeper:             launchKeeper,
			ProfileKeeper:            profileKeeper,
			RewardKeeper:             rewardKeeper,
			MonitoringConsumerKeeper: monitoringConsumerKeeper,
			AccountKeeper:            authKeeper,
			BankKeeper:               bankKeeper,
			DistrKeeper:              distrKeeper,
			IBCKeeper:                ibcKeeper,
			StakingKeeper:            stakingKeeper,
			FundraisingKeeper:        fundraisingKeeper,
			ParticipationKeeper:      participationKeeper,
			ClaimKeeper:              claimKeeper,
			HooksMocks:               hooksMocks,
		}, TestMsgServers{
			T:                t,
			ProfileSrv:       profileSrv,
			LaunchSrv:        launchSrv,
			ProjectSrv:       projectSrv,
			RewardSrv:        rewardSrv,
			MonitoringcSrv:   monitoringcSrv,
			ParticipationSrv: participationSrv,
			ClaimSrv:         claimSrv,
		}
}

// NewTestSetupWithIBCMocks returns a keeper of the monitoring consumer module for testing purpose with mocks for IBC keepers
func NewTestSetupWithIBCMocks(
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
	fundraisingKeeper := initializer.Fundraising(paramKeeper, authKeeper, bankKeeper, distrKeeper)
	profileKeeper := initializer.Profile()
	launchKeeper := initializer.Launch(profileKeeper, distrKeeper, paramKeeper)
	rewardKeeper := initializer.Reward(authKeeper, bankKeeper, profileKeeper, launchKeeper, paramKeeper)
	projectKeeper := initializer.Project(launchKeeper, profileKeeper, bankKeeper, distrKeeper, paramKeeper)
	participationKeeper := initializer.Participation(paramKeeper, fundraisingKeeper, stakingKeeper)
	launchKeeper.SetProjectKeeper(projectKeeper)
	monitoringConsumerKeeper := initializer.Monitoringc(
		*ibcKeeper,
		*capabilityKeeper,
		launchKeeper,
		rewardKeeper,
		paramKeeper,
		connectionMock,
		channelMock,
	)
	launchKeeper.SetMonitoringcKeeper(monitoringConsumerKeeper)
	claimKeeper := initializer.Claim(paramKeeper, authKeeper, distrKeeper, bankKeeper)
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
	monitoringConsumerKeeper.SetParams(ctx, monitoringctypes.DefaultParams())
	claimKeeper.SetParams(ctx, claimtypes.DefaultParams())
	setIBCDefaultParams(ctx, ibcKeeper)

	profileSrv := profilekeeper.NewMsgServerImpl(*profileKeeper)
	launchSrv := launchkeeper.NewMsgServerImpl(*launchKeeper)
	projectSrv := projectkeeper.NewMsgServerImpl(*projectKeeper)
	rewardSrv := rewardkeeper.NewMsgServerImpl(*rewardKeeper)
	monitoringcSrv := monitoringckeeper.NewMsgServerImpl(*monitoringConsumerKeeper)
	participationSrv := participationkeeper.NewMsgServerImpl(*participationKeeper)

	// set max shares - only set during app InitGenesis
	projectKeeper.SetTotalShares(ctx, spntypes.TotalShareNumber)

	return ctx, TestKeepers{
			T:                        t,
			ProjectKeeper:            projectKeeper,
			LaunchKeeper:             launchKeeper,
			ProfileKeeper:            profileKeeper,
			RewardKeeper:             rewardKeeper,
			MonitoringConsumerKeeper: monitoringConsumerKeeper,
			AccountKeeper:            authKeeper,
			BankKeeper:               bankKeeper,
			IBCKeeper:                ibcKeeper,
			StakingKeeper:            stakingKeeper,
			FundraisingKeeper:        fundraisingKeeper,
			ParticipationKeeper:      participationKeeper,
			ClaimKeeper:              claimKeeper,
		}, TestMsgServers{
			T:                t,
			ProfileSrv:       profileSrv,
			LaunchSrv:        launchSrv,
			ProjectSrv:       projectSrv,
			RewardSrv:        rewardSrv,
			MonitoringcSrv:   monitoringcSrv,
			ParticipationSrv: participationSrv,
		}
}

// setIBCDefaultParams set default params for IBC client and connection keepers
func setIBCDefaultParams(ctx sdk.Context, ibcKeeper *ibckeeper.Keeper) {
	ibcKeeper.ClientKeeper.SetParams(ctx, ibcclienttypes.DefaultParams())
	ibcKeeper.ConnectionKeeper.SetParams(ctx, ibcconnectiontypes.DefaultParams())
	ibcKeeper.ClientKeeper.SetNextClientSequence(ctx, 0)
}
