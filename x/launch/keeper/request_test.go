package keeper

import (
	"strconv"
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	spnerrors "github.com/tendermint/spn/pkg/errors"
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
		coordinatorAcc        = sample.AccAddress()
		genesisAcc            = sample.AccAddress()
		vestedAcc             = sample.AccAddress()
		validatorAcc          = sample.AccAddress()
		k, _, _, _, sdkCtx, _ = setupMsgServer(t)
		chainID, _            = sample.ChainID(10)
		contents              = sample.AllRequestContents(chainID, genesisAcc, vestedAcc, validatorAcc)
		requests              = createRequests(k, sdkCtx, chainID, contents)
		invalidContent, _     = codectypes.NewAnyWithValue(&types.Request{})
		invalidContentID      = k.AppendRequest(sdkCtx, *sample.RequestWithContent(chainID, invalidContent))
	)
	tests := []struct {
		name string
		msg  types.MsgSettleRequest
		err  error
	}{
		{
			name: "test GenesisAccount content",
			msg: types.MsgSettleRequest{
				ChainID:     chainID,
				RequestID:   requests[0].RequestID,
				Coordinator: coordinatorAcc,
				Approve:     true,
			},
		}, {
			name: "test duplicated GenesisAccount content",
			msg: types.MsgSettleRequest{
				ChainID:     chainID,
				RequestID:   requests[0].RequestID,
				Coordinator: coordinatorAcc,
				Approve:     true,
			},
			err: sdkerrors.Wrapf(types.ErrAccountAlreadyExist,
				"account %s for chain %s already exist", genesisAcc, chainID),
		}, {
			name: "test genesis AccountRemoval content",
			msg: types.MsgSettleRequest{
				ChainID:     chainID,
				RequestID:   requests[1].RequestID,
				Coordinator: coordinatorAcc,
				Approve:     true,
			},
		}, {
			name: "test not found genesis AccountRemoval content",
			msg: types.MsgSettleRequest{
				ChainID:     chainID,
				RequestID:   requests[1].RequestID,
				Coordinator: coordinatorAcc,
				Approve:     true,
			},
			err: sdkerrors.Wrapf(types.ErrAccountNotFound,
				"account %s for chain %s not found", genesisAcc, chainID),
		}, {
			name: "test VestedAccount content",
			msg: types.MsgSettleRequest{
				ChainID:     chainID,
				RequestID:   requests[2].RequestID,
				Coordinator: coordinatorAcc,
				Approve:     true,
			},
		}, {
			name: "test duplicated VestedAccount content",
			msg: types.MsgSettleRequest{
				ChainID:     chainID,
				RequestID:   requests[2].RequestID,
				Coordinator: coordinatorAcc,
				Approve:     true,
			},
			err: sdkerrors.Wrapf(types.ErrAccountAlreadyExist,
				"account %s for chain %s already exist", vestedAcc, chainID),
		}, {
			name: "test vested AccountRemoval content",
			msg: types.MsgSettleRequest{
				ChainID:     chainID,
				RequestID:   requests[3].RequestID,
				Coordinator: coordinatorAcc,
				Approve:     true,
			},
		}, {
			name: "test not found vested AccountRemoval content",
			msg: types.MsgSettleRequest{
				ChainID:     chainID,
				RequestID:   requests[3].RequestID,
				Coordinator: coordinatorAcc,
				Approve:     true,
			},
			err: sdkerrors.Wrapf(types.ErrAccountNotFound,
				"account %s for chain %s not found", vestedAcc, chainID),
		}, {
			name: "test GenesisValidator content",
			msg: types.MsgSettleRequest{
				ChainID:     chainID,
				RequestID:   requests[4].RequestID,
				Coordinator: coordinatorAcc,
				Approve:     true,
			},
		}, {
			name: "test duplicated GenesisValidator content",
			msg: types.MsgSettleRequest{
				ChainID:     chainID,
				RequestID:   requests[4].RequestID,
				Coordinator: coordinatorAcc,
				Approve:     true,
			},
			err: sdkerrors.Wrapf(types.ErrValidatorAlreadyExist,
				"genesis validator %s for chain %s already exist", validatorAcc, chainID),
		}, {
			name: "test ValidatorRemoval content",
			msg: types.MsgSettleRequest{
				ChainID:     chainID,
				RequestID:   requests[5].RequestID,
				Coordinator: coordinatorAcc,
				Approve:     true,
			},
		}, {
			name: "test not found ValidatorRemoval content",
			msg: types.MsgSettleRequest{
				ChainID:     chainID,
				RequestID:   requests[5].RequestID,
				Coordinator: coordinatorAcc,
				Approve:     true,
			},
			err: sdkerrors.Wrapf(types.ErrValidatorNotFound,
				"genesis validator %s for chain %s not found", validatorAcc, chainID),
		}, {
			name: "invalid request",
			msg: types.MsgSettleRequest{
				ChainID:     chainID,
				RequestID:   invalidContentID,
				Coordinator: coordinatorAcc,
				Approve:     true,
			},
			err: spnerrors.Critical(
				"no concrete type registered for type URL /tendermint.spn.launch.Request against interface *types.RequestContent"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, found := k.GetRequest(sdkCtx, tt.msg.ChainID, tt.msg.RequestID)
			require.True(t, found)

			err := applyRequest(sdkCtx, *k, &tt.msg, request)
			if tt.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, tt.err, err)
				require.Equal(t, tt.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}
