package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func msgCreateCoordinator() types.MsgCreateCoordinator {
	addr := sample.AccAddress()
	return types.MsgCreateCoordinator{
		Address: addr,
		Description: &types.CoordinatorDescription{
			Identity: addr,
			Website:  "https://cosmos.network/" + addr,
			Details:  addr + " details",
		},
	}
}

func TestMsgCreateCoordinator(t *testing.T) {
	var (
		msg1        = msgCreateCoordinator()
		msg2        = msgCreateCoordinator()
		ctx, k, srv = setupMsgServer(t)
		wCtx        = sdk.WrapSDKContext(ctx)
	)
	tests := []struct {
		name   string
		msg    types.MsgCreateCoordinator
		wantId uint64
		err    error
	}{
		{
			name:   "valid coordinator 1",
			msg:    msg1,
			wantId: 0,
		}, {
			name:   "valid coordinator 2",
			msg:    msg2,
			wantId: 1,
		}, {
			name: "already exist address",
			msg:  msg2,
			err:  sdkerrors.Wrap(types.ErrCoordAlreadyExist, "coordinatorId: 1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.CreateCoordinator(wCtx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			coordByAddr, found := k.GetCoordinatorByAddress(ctx, tt.msg.Address)
			assert.True(t, found, "coordinator by address not found")
			assert.EqualValues(t, tt.wantId, coordByAddr.CoordinatorId)

			coord := k.GetCoordinator(ctx, coordByAddr.CoordinatorId)
			assert.True(t, found, "coordinator id not found")
			assert.EqualValues(t, tt.msg.Address, coord.Address)
			assert.EqualValues(t, tt.msg.Description, coord.Description)
			assert.EqualValues(t, coordByAddr.CoordinatorId, coord.CoordinatorId)
		})
	}
}
