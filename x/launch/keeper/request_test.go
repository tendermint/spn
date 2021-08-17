package keeper

import (
	"strconv"
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createRequests(
	keeper *Keeper,
	ctx sdk.Context,
	chainID string,
	contents []*codectypes.Any,
) []types.Request {
	items := make([]types.Request, len(contents))
	for i, content := range contents {
		items[i] = *sample.RequestWithContent(chainID, content)
		id := keeper.AppendRequest(ctx, items[i])
		items[i].RequestID = id
	}
	return items
}

func createNRequest(keeper *Keeper, ctx sdk.Context, n int) []types.Request {
	items := make([]types.Request, n)
	for i := range items {
		items[i] = *sample.Request("foo")
		id := keeper.AppendRequest(ctx, items[i])
		items[i].RequestID = id
	}
	return items
}

func TestRequestGet(t *testing.T) {
	keeper, _, ctx, _ := setupKeeper(t)
	items := createNRequest(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRequest(ctx,
			item.ChainID,
			item.RequestID,
		)
		assert.True(t, found)

		// Cached value is cleared when the any type is encoded into the store
		item.Content.ClearCachedValue()

		assert.Equal(t, item, rst)
	}
}
func TestRequestRemove(t *testing.T) {
	keeper, _, ctx, _ := setupKeeper(t)
	items := createNRequest(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRequest(ctx,
			item.ChainID,
			item.RequestID,
		)
		_, found := keeper.GetRequest(ctx,
			item.ChainID,
			item.RequestID,
		)
		assert.False(t, found)
	}
}

func TestRequestGetAll(t *testing.T) {
	keeper, _, ctx, _ := setupKeeper(t)
	items := createNRequest(keeper, ctx, 10)

	// Cached value is cleared when the any type is encoded into the store
	for _, item := range items {
		item.Content.ClearCachedValue()
	}

	assert.Equal(t, items, keeper.GetAllRequest(ctx))
}

func TestRequestCount(t *testing.T) {
	keeper, _, ctx, _ := setupKeeper(t)
	items := createNRequest(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetRequestCount(ctx, "foo"))
	assert.Equal(t, uint64(0), keeper.GetRequestCount(ctx, "bar"))
}

func TestApplyRequest(t *testing.T) {
	var (
		genesisAcc            = sample.AccAddress()
		vestedAcc             = sample.AccAddress()
		validatorAcc          = sample.AccAddress()
		k, _, _, _, sdkCtx, _ = setupMsgServer(t)
		chainID, _            = sample.ChainID(10)
		contents              = sample.AllRequestContents(chainID, genesisAcc, vestedAcc, validatorAcc)
		missingContent, _     = codectypes.NewAnyWithValue(&types.GenesisAccount{ChainID: chainID})
		invalidContent, _     = codectypes.NewAnyWithValue(&types.Request{})
	)
	tests := []struct {
		name    string
		request types.Request
		wantErr bool
	}{
		{
			name:    "test GenesisAccount content",
			request: *sample.RequestWithContent(chainID, contents[0]),
		}, {
			name:    "test duplicated GenesisAccount content",
			request: *sample.RequestWithContent(chainID, contents[0]),
			wantErr: true,
		}, {
			name:    "test genesis AccountRemoval content",
			request: *sample.RequestWithContent(chainID, contents[1]),
		}, {
			name:    "test not found genesis AccountRemoval content",
			request: *sample.RequestWithContent(chainID, contents[1]),
			wantErr: true,
		}, {
			name:    "test VestedAccount content",
			request: *sample.RequestWithContent(chainID, contents[2]),
		}, {
			name:    "test duplicated VestedAccount content",
			request: *sample.RequestWithContent(chainID, contents[2]),
			wantErr: true,
		}, {
			name:    "test vested AccountRemoval content",
			request: *sample.RequestWithContent(chainID, contents[3]),
		}, {
			name:    "test not found vested AccountRemoval content",
			request: *sample.RequestWithContent(chainID, contents[3]),
			wantErr: true,
		}, {
			name:    "test GenesisValidator content",
			request: *sample.RequestWithContent(chainID, contents[4]),
		}, {
			name:    "test duplicated GenesisValidator content",
			request: *sample.RequestWithContent(chainID, contents[4]),
			wantErr: true,
		}, {
			name:    "test ValidatorRemoval content",
			request: *sample.RequestWithContent(chainID, contents[5]),
		}, {
			name:    "test not found ValidatorRemoval content",
			request: *sample.RequestWithContent(chainID, contents[5]),
			wantErr: true,
		}, {
			name:    "test request with invalid parameters",
			request: *sample.RequestWithContent(chainID, missingContent),
			wantErr: true,
		}, {
			name:    "invalid request",
			request: *sample.RequestWithContent(chainID, invalidContent),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := applyRequest(sdkCtx, *k, chainID, tt.request)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			var content types.RequestContent
			cdc := codectypes.NewInterfaceRegistry()
			types.RegisterInterfaces(cdc)
			err = cdc.UnpackAny(tt.request.Content, &content)
			require.NoError(t, err)

			switch c := content.(type) {
			case *types.GenesisAccount:
				_, found := k.GetGenesisAccount(sdkCtx, chainID, c.Address)
				require.True(t, found, "genesis account not found")
			case *types.VestedAccount:
				_, found := k.GetVestedAccount(sdkCtx, chainID, c.Address)
				require.True(t, found, "vested account not found")
			case *types.AccountRemoval:
				_, foundGenesis := k.GetGenesisAccount(sdkCtx, chainID, c.Address)
				require.False(t, foundGenesis, "genesis account not removed")
				_, foundVested := k.GetVestedAccount(sdkCtx, chainID, c.Address)
				require.False(t, foundVested, "vested account not removed")
			case *types.GenesisValidator:
				_, found := k.GetGenesisValidator(sdkCtx, chainID, c.Address)
				require.True(t, found, "genesis validator not found")
			case *types.ValidatorRemoval:
				_, found := k.GetGenesisValidator(sdkCtx, chainID, c.ValAddress)
				require.False(t, found, "genesis validator not removed")
			}
		})
	}
}
