package launch_test

import (
	"math/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	"github.com/tendermint/spn/x/launch"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (
	*keeper.Keeper,
	*profilekeeper.Keeper,
	*campaignkeeper.Keeper,
	types.MsgServer,
	profiletypes.MsgServer,
	campaigntypes.MsgServer,
	sdk.Context,
) {
	campaignKeeper, launchLKeeper, profileKeeper, _, ctx := testkeeper.AllKeepers(t)
	return launchLKeeper,
		profileKeeper,
		campaignKeeper,
		keeper.NewMsgServerImpl(*launchLKeeper),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		campaignkeeper.NewMsgServerImpl(*campaignKeeper),
		ctx
}

func TestFindAccount(t *testing.T) {
	var (
		r    = rand.New(rand.NewSource(1))
		accs = simulation.RandomAccounts(r, 5)
	)
	tests := []struct {
		name    string
		address string
		want    simulation.Account
		wantErr bool
	}{
		{
			name:    "invalid address",
			address: "invalid_address",
			wantErr: true,
		},
		{
			name:    "first account",
			address: accs[0].Address.String(),
			want:    accs[0],
		},
		{
			name:    "last account",
			address: accs[len(accs)-1].Address.String(),
			want:    accs[len(accs)-1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := launch.FindAccount(accs, tt.address)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFindChain(t *testing.T) {
	var (
		k, _, _, _, profileSrv, _, sdkCtx = setupMsgServer(t)

		ctx            = sdk.WrapSDKContext(sdkCtx)
		msgCreateCoord = sample.MsgCreateCoordinator(sample.Address())
	)

	// Create coordinator
	res, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoord)
	require.NoError(t, err)

	t.Run("chain without coordinator", func(t *testing.T) {
		k.AppendChain(sdkCtx, types.Chain{
			CoordinatorID:   1000,
			LaunchTriggered: true,
		})
		_, found := launch.FindChain(sdkCtx, *k, true)
		require.False(t, found)
	})

	t.Run("not launch triggered chain", func(t *testing.T) {
		k.AppendChain(sdkCtx, types.Chain{
			CoordinatorID:   res.CoordinatorId,
			LaunchTriggered: false,
		})
		_, found := launch.FindChain(sdkCtx, *k, true)
		require.False(t, found)
		got, found := launch.FindChain(sdkCtx, *k, false)
		require.True(t, found)
		require.Equal(t, res.CoordinatorId, got.CoordinatorID)
	})

	t.Run("found chain", func(t *testing.T) {
		k.AppendChain(sdkCtx, types.Chain{
			CoordinatorID:   res.CoordinatorId,
			LaunchTriggered: true,
		})
		got, found := launch.FindChain(sdkCtx, *k, true)
		require.True(t, found)
		require.Equal(t, res.CoordinatorId, got.CoordinatorID)
		got, found = launch.FindChain(sdkCtx, *k, false)
		require.True(t, found)
		require.Equal(t, res.CoordinatorId, got.CoordinatorID)
	})
}

func TestFindChainCoordinatorAccount(t *testing.T) {
	var (
		k, _, _, _, profileSrv, _, sdkCtx = setupMsgServer(t)

		ctx  = sdk.WrapSDKContext(sdkCtx)
		r    = rand.New(rand.NewSource(1))
		accs = simulation.RandomAccounts(r, 2)
	)

	// Create coordinator
	msgCreateCoord := sample.MsgCreateCoordinator(accs[0].Address.String())
	res, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoord)
	require.NoError(t, err)

	// Create chains
	chainID := k.AppendChain(sdkCtx, types.Chain{
		CoordinatorID: res.CoordinatorId,
	})
	chainWithoutCoordID := k.AppendChain(sdkCtx, types.Chain{
		CoordinatorID: 1000,
	})
	tests := []struct {
		name    string
		chainID uint64
		want    simulation.Account
		wantErr bool
	}{
		{
			name:    "valid chain coordinator",
			chainID: chainID,
			want:    accs[0],
		},
		{
			name:    "chain without coordinator",
			chainID: chainWithoutCoordID,
			wantErr: true,
		},
		{
			name:    "not found chain",
			chainID: 1000,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := launch.FindChainCoordinatorAccount(sdkCtx, *k, accs, tt.chainID)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestIsLaunchTriggeredChain(t *testing.T) {
	var (
		k, _, _, _, profileSrv, _, sdkCtx = setupMsgServer(t)

		ctx            = sdk.WrapSDKContext(sdkCtx)
		msgCreateCoord = sample.MsgCreateCoordinator(sample.Address())
	)

	// Create coordinator
	res, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoord)
	require.NoError(t, err)

	tests := []struct {
		name        string
		IsTriggered bool
	}{
		{
			name:        "is launch triggered chain",
			IsTriggered: true,
		},
		{
			name:        "is not launch triggered chain",
			IsTriggered: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chainID := k.AppendChain(sdkCtx, types.Chain{
				CoordinatorID:   res.CoordinatorId,
				LaunchTriggered: tt.IsTriggered,
			})
			got := launch.IsLaunchTriggeredChain(sdkCtx, *k, chainID)
			require.Equal(t, tt.IsTriggered, got)
		})
	}
}
