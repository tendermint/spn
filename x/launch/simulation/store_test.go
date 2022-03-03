package simulation_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	"github.com/tendermint/spn/x/launch/keeper"
	launchsimulation "github.com/tendermint/spn/x/launch/simulation"
	"github.com/tendermint/spn/x/launch/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func setupMsgServer(t testing.TB) (
	sdk.Context,
	testkeeper.TestKeepers,
	types.MsgServer,
	profiletypes.MsgServer,
	campaigntypes.MsgServer,
) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	return ctx,
		tk,
		keeper.NewMsgServerImpl(*tk.LaunchKeeper),
		profilekeeper.NewMsgServerImpl(*tk.ProfileKeeper),
		campaignkeeper.NewMsgServerImpl(*tk.CampaignKeeper)
}

func TestFindAccount(t *testing.T) {
	var (
		r    = sample.Rand()
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
			got, err := launchsimulation.FindAccount(accs, tt.address)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFindChainCoordinatorAccount(t *testing.T) {
	var (
		sdkCtx, tk, _, profileSrv, _ = setupMsgServer(t)

		ctx  = sdk.WrapSDKContext(sdkCtx)
		r    = sample.Rand()
		accs = simulation.RandomAccounts(r, 2)
	)

	// Create coordinator
	msgCreateCoord := sample.MsgCreateCoordinator(accs[0].Address.String())
	res, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoord)
	require.NoError(t, err)

	// Create coordinator and disable
	msgCreateCoord = sample.MsgCreateCoordinator(accs[1].Address.String())
	resDisable, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoord)
	require.NoError(t, err)
	msgDisableCoord := sample.MsgDisableCoordinator(accs[1].Address.String())
	_, err = profileSrv.DisableCoordinator(ctx, &msgDisableCoord)
	require.NoError(t, err)

	// Create chains
	chainID := tk.LaunchKeeper.AppendChain(sdkCtx, types.Chain{
		CoordinatorID: res.CoordinatorID,
	})
	chainWithoutCoordID := tk.LaunchKeeper.AppendChain(sdkCtx, types.Chain{
		CoordinatorID: 1000,
	})

	chainWithDisableCoord := tk.LaunchKeeper.AppendChain(sdkCtx, types.Chain{
		CoordinatorID: resDisable.CoordinatorID,
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
			name:    "chain with disabled coordinator",
			chainID: chainWithDisableCoord,
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
			got, err := launchsimulation.FindChainCoordinatorAccount(sdkCtx, *tk.LaunchKeeper, accs, tt.chainID)
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
		sdkCtx, tk, _, profileSrv, _ = setupMsgServer(t)

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
			chainID := tk.LaunchKeeper.AppendChain(sdkCtx, types.Chain{
				CoordinatorID:   res.CoordinatorID,
				LaunchTriggered: tt.IsTriggered,
			})
			got := launchsimulation.IsLaunchTriggeredChain(sdkCtx, *tk.LaunchKeeper, chainID)
			require.Equal(t, tt.IsTriggered, got)
		})
	}
}

func TestFindRandomChain(t *testing.T) {
	var (
		sdkCtx, tk, _, profileSrv, _ = setupMsgServer(t)

		r              = sample.Rand()
		ctx            = sdk.WrapSDKContext(sdkCtx)
		msgCreateCoord = sample.MsgCreateCoordinator(sample.Address())
	)

	t.Run("no chains", func(t *testing.T) {
		_, found := launchsimulation.FindRandomChain(r, sdkCtx, *tk.LaunchKeeper, true, false)
		require.False(t, found)
		_, found = launchsimulation.FindRandomChain(r, sdkCtx, *tk.LaunchKeeper, false, false)
		require.False(t, found)
	})

	// Create coordinator
	res, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoord)
	require.NoError(t, err)

	t.Run("chain without coordinator", func(t *testing.T) {
		tk.LaunchKeeper.AppendChain(sdkCtx, types.Chain{
			CoordinatorID:   1000,
			LaunchTriggered: true,
		})
		_, found := launchsimulation.FindRandomChain(r, sdkCtx, *tk.LaunchKeeper, true, false)
		require.False(t, found)
	})

	t.Run("chain with no mainnet", func(t *testing.T) {
		tk.LaunchKeeper.AppendChain(sdkCtx, types.Chain{
			CoordinatorID: res.CoordinatorID,
			IsMainnet:     true,
		})
		_, found := launchsimulation.FindRandomChain(r, sdkCtx, *tk.LaunchKeeper, false, true)
		require.False(t, found)

		tk.LaunchKeeper.AppendChain(sdkCtx, types.Chain{
			CoordinatorID: res.CoordinatorID,
		})
		c, found := launchsimulation.FindRandomChain(r, sdkCtx, *tk.LaunchKeeper, false, true)
		require.True(t, found)
		require.False(t, c.IsMainnet)
	})

	t.Run("not launch triggered chain", func(t *testing.T) {
		tk.LaunchKeeper.AppendChain(sdkCtx, types.Chain{
			CoordinatorID:   res.CoordinatorID,
			LaunchTriggered: false,
		})
		_, found := launchsimulation.FindRandomChain(r, sdkCtx, *tk.LaunchKeeper, true, false)
		require.False(t, found)
		got, found := launchsimulation.FindRandomChain(r, sdkCtx, *tk.LaunchKeeper, false, false)
		require.True(t, found)
		require.Equal(t, res.CoordinatorID, got.CoordinatorID)
	})

	t.Run("found chain", func(t *testing.T) {
		tk.LaunchKeeper.AppendChain(sdkCtx, types.Chain{
			CoordinatorID:   res.CoordinatorID,
			LaunchTriggered: true,
		})
		got, found := launchsimulation.FindRandomChain(r, sdkCtx, *tk.LaunchKeeper, true, false)
		require.True(t, found)
		require.Equal(t, res.CoordinatorID, got.CoordinatorID)
		got, found = launchsimulation.FindRandomChain(r, sdkCtx, *tk.LaunchKeeper, false, false)
		require.True(t, found)
		require.Equal(t, res.CoordinatorID, got.CoordinatorID)
	})
}

func TestFindRandomValidator(t *testing.T) {
	var (
		sdkCtx, tk, _, profileSrv, _ = setupMsgServer(t)

		ctx  = sdk.WrapSDKContext(sdkCtx)
		r    = sample.Rand()
		accs = simulation.RandomAccounts(r, 2)
	)

	t.Run("empty validators", func(t *testing.T) {
		gotSimAcc, gotVal, gotFound := launchsimulation.FindRandomValidator(r, sdkCtx, *tk.LaunchKeeper, accs)

		require.False(t, gotFound)
		require.Equal(t, simtypes.Account{}, gotSimAcc)
		require.Equal(t, types.GenesisValidator{}, gotVal)
	})

	t.Run("chain triggered launch", func(t *testing.T) {
		chainID := tk.LaunchKeeper.AppendChain(sdkCtx, types.Chain{
			LaunchTriggered: true,
		})
		tk.LaunchKeeper.SetGenesisValidator(sdkCtx, sample.GenesisValidator(chainID, sample.Address()))

		gotSimAcc, gotVal, gotFound := launchsimulation.FindRandomValidator(r, sdkCtx, *tk.LaunchKeeper, accs)
		require.False(t, gotFound)
		require.Equal(t, simtypes.Account{}, gotSimAcc)
		require.Equal(t, types.GenesisValidator{}, gotVal)
	})

	t.Run("chain without coordinator", func(t *testing.T) {
		chainID := tk.LaunchKeeper.AppendChain(sdkCtx, sample.Chain(0, 1000))
		tk.LaunchKeeper.SetGenesisValidator(sdkCtx, sample.GenesisValidator(chainID, sample.Address()))

		gotSimAcc, gotVal, gotFound := launchsimulation.FindRandomValidator(r, sdkCtx, *tk.LaunchKeeper, accs)
		require.False(t, gotFound)
		require.Equal(t, simtypes.Account{}, gotSimAcc)
		require.Equal(t, types.GenesisValidator{}, gotVal)
	})

	t.Run("chain without coordinator account", func(t *testing.T) {
		msgCreateCoord := sample.MsgCreateCoordinator(sample.Address())
		res, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoord)
		require.NoError(t, err)
		chainID := tk.LaunchKeeper.AppendChain(sdkCtx, sample.Chain(0, res.CoordinatorID))
		tk.LaunchKeeper.SetGenesisValidator(sdkCtx, sample.GenesisValidator(chainID, sample.Address()))

		gotSimAcc, gotVal, gotFound := launchsimulation.FindRandomValidator(r, sdkCtx, *tk.LaunchKeeper, accs)
		require.False(t, gotFound)
		require.Equal(t, simtypes.Account{}, gotSimAcc)
		require.Equal(t, types.GenesisValidator{}, gotVal)
	})

	t.Run("got a valid validator", func(t *testing.T) {
		msgCreateCoord := sample.MsgCreateCoordinator(accs[0].Address.String())
		res, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoord)
		require.NoError(t, err)
		chainID := tk.LaunchKeeper.AppendChain(sdkCtx, sample.Chain(0, res.CoordinatorID))
		validator := sample.GenesisValidator(chainID, sample.Address())
		tk.LaunchKeeper.SetGenesisValidator(sdkCtx, validator)

		gotSimAcc, gotVal, gotFound := launchsimulation.FindRandomValidator(r, sdkCtx, *tk.LaunchKeeper, accs)
		require.True(t, gotFound)
		require.Equal(t, accs[0], gotSimAcc)
		require.Equal(t, validator, gotVal)
	})
}

func TestFindRandomRequest(t *testing.T) {
	var (
		sdkCtx, tk, _, profileSrv, _ = setupMsgServer(t)

		r   = sample.Rand()
		ctx = sdk.WrapSDKContext(sdkCtx)
	)

	t.Run("empty requests", func(t *testing.T) {
		gotRequest, gotFound := launchsimulation.FindRandomRequest(r, sdkCtx, *tk.LaunchKeeper)
		require.Equal(t, types.Request{}, gotRequest)
		require.False(t, gotFound)
	})

	t.Run("no chain request", func(t *testing.T) {
		tk.LaunchKeeper.AppendRequest(sdkCtx, types.Request{
			LaunchID: 10000,
			Creator:  sample.Address(),
		})
		gotRequest, gotFound := launchsimulation.FindRandomRequest(r, sdkCtx, *tk.LaunchKeeper)
		require.Equal(t, types.Request{}, gotRequest)
		require.False(t, gotFound)
	})

	t.Run("launch triggered chain request", func(t *testing.T) {
		chainID := tk.LaunchKeeper.AppendChain(sdkCtx, types.Chain{
			LaunchTriggered: true,
		})
		tk.LaunchKeeper.AppendRequest(sdkCtx, sample.Request(chainID, sample.Address()))
		gotRequest, gotFound := launchsimulation.FindRandomRequest(r, sdkCtx, *tk.LaunchKeeper)
		require.Equal(t, types.Request{}, gotRequest)
		require.False(t, gotFound)
	})

	t.Run("chain without coordinator", func(t *testing.T) {
		chainID := tk.LaunchKeeper.AppendChain(sdkCtx, types.Chain{
			CoordinatorID:   10000,
			LaunchTriggered: true,
		})
		tk.LaunchKeeper.AppendRequest(sdkCtx, sample.Request(chainID, sample.Address()))
		gotRequest, gotFound := launchsimulation.FindRandomRequest(r, sdkCtx, *tk.LaunchKeeper)
		require.Equal(t, types.Request{}, gotRequest)
		require.False(t, gotFound)
	})

	t.Run("get a valid request", func(t *testing.T) {
		msgCreateCoord := sample.MsgCreateCoordinator(sample.Address())
		res, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoord)
		require.NoError(t, err)

		chainID := tk.LaunchKeeper.AppendChain(sdkCtx, sample.Chain(0, res.CoordinatorID))
		request := sample.Request(chainID, sample.Address())
		tk.LaunchKeeper.AppendRequest(sdkCtx, request)

		gotRequest, gotFound := launchsimulation.FindRandomRequest(r, sdkCtx, *tk.LaunchKeeper)
		require.Equal(t, request, gotRequest)
		require.True(t, gotFound)
	})
}
