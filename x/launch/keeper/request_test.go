package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

func createRequests(
	keeper *keeper.Keeper,
	ctx sdk.Context,
	launchID uint64,
	contents []types.RequestContent,
) []types.Request {
	items := make([]types.Request, len(contents))
	for i, content := range contents {
		items[i] = sample.RequestWithContent(launchID, content)
		id := keeper.AppendRequest(ctx, items[i])
		items[i].RequestID = id
	}
	return items
}

func createNRequest(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Request {
	items := make([]types.Request, n)
	for i := range items {
		items[i] = sample.Request(0, sample.Address())
		id := keeper.AppendRequest(ctx, items[i])
		items[i].RequestID = id
	}
	return items
}

func TestRequestGet(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNRequest(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRequest(ctx,
			item.LaunchID,
			item.RequestID,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestRequestRemove(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNRequest(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRequest(ctx,
			item.LaunchID,
			item.RequestID,
		)
		_, found := keeper.GetRequest(ctx,
			item.LaunchID,
			item.RequestID,
		)
		require.False(t, found)
	}
}

func TestRequestGetAll(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNRequest(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllRequest(ctx))
}

func TestRequestCount(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNRequest(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetRequestCount(ctx, 0))
	require.Equal(t, uint64(0), keeper.GetRequestCount(ctx, 1))
}

func TestApplyRequest(t *testing.T) {
	var (
		genesisAcc     = sample.Address()
		vestingAcc     = sample.Address()
		validatorAcc   = sample.Address()
		k, ctx         = testkeeper.Launch(t)
		launchID       = uint64(10)
		contents       = sample.AllRequestContents(launchID, genesisAcc, vestingAcc, validatorAcc)
		invalidContent = types.NewGenesisAccount(launchID, "", sdk.NewCoins())
	)
	tests := []struct {
		name    string
		request types.Request
		wantErr bool
	}{
		{
			name:    "test GenesisAccount content",
			request: sample.RequestWithContent(launchID, contents[0]),
		}, {
			name:    "test duplicated GenesisAccount content",
			request: sample.RequestWithContent(launchID, contents[0]),
			wantErr: true,
		}, {
			name:    "test genesis AccountRemoval content",
			request: sample.RequestWithContent(launchID, contents[1]),
		}, {
			name:    "test not found genesis AccountRemoval content",
			request: sample.RequestWithContent(launchID, contents[1]),
			wantErr: true,
		}, {
			name:    "test VestingAccount content",
			request: sample.RequestWithContent(launchID, contents[2]),
		}, {
			name:    "test duplicated VestingAccount content",
			request: sample.RequestWithContent(launchID, contents[2]),
			wantErr: true,
		}, {
			name:    "test vesting AccountRemoval content",
			request: sample.RequestWithContent(launchID, contents[3]),
		}, {
			name:    "test not found vesting AccountRemoval content",
			request: sample.RequestWithContent(launchID, contents[3]),
			wantErr: true,
		}, {
			name:    "test GenesisValidator content",
			request: sample.RequestWithContent(launchID, contents[4]),
		}, {
			name:    "test duplicated GenesisValidator content",
			request: sample.RequestWithContent(launchID, contents[4]),
			wantErr: true,
		}, {
			name:    "test ValidatorRemoval content",
			request: sample.RequestWithContent(launchID, contents[5]),
		}, {
			name:    "test not found ValidatorRemoval content",
			request: sample.RequestWithContent(launchID, contents[5]),
			wantErr: true,
		}, {
			name:    "test request with invalid parameters",
			request: sample.RequestWithContent(launchID, invalidContent),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := keeper.ApplyRequest(ctx, *k, launchID, tt.request)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			switch requestContent := tt.request.Content.Content.(type) {
			case *types.RequestContent_GenesisAccount:
				ga := requestContent.GenesisAccount
				_, found := k.GetGenesisAccount(ctx, launchID, ga.Address)
				require.True(t, found, "genesis account not found")
			case *types.RequestContent_VestingAccount:
				va := requestContent.VestingAccount
				_, found := k.GetVestingAccount(ctx, launchID, va.Address)
				require.True(t, found, "vesting account not found")
			case *types.RequestContent_AccountRemoval:
				ar := requestContent.AccountRemoval
				_, foundGenesis := k.GetGenesisAccount(ctx, launchID, ar.Address)
				require.False(t, foundGenesis, "genesis account not removed")
				_, foundVesting := k.GetVestingAccount(ctx, launchID, ar.Address)
				require.False(t, foundVesting, "vesting account not removed")
			case *types.RequestContent_GenesisValidator:
				ga := requestContent.GenesisValidator
				_, found := k.GetGenesisValidator(ctx, launchID, ga.Address)
				require.True(t, found, "genesis validator not found")
			case *types.RequestContent_ValidatorRemoval:
				vr := requestContent.ValidatorRemoval
				_, found := k.GetGenesisValidator(ctx, launchID, vr.ValAddress)
				require.False(t, found, "genesis validator not removed")
			}
		})
	}
}
