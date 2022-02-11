package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgCreateCoordinator(t *testing.T) {
	var (
		msg1        = sample.MsgCreateCoordinator(sample.Address())
		msg2        = sample.MsgCreateCoordinator(sample.Address())
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
			err:  types.ErrCoordAlreadyExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.CreateCoordinator(wCtx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			coordByAddr, found := k.GetCoordinatorByAddress(ctx, tt.msg.Address)
			require.True(t, found, "coordinator by address not found")
			require.EqualValues(t, tt.wantId, coordByAddr.CoordinatorID)
			require.EqualValues(t, tt.wantId, got.CoordinatorID)

			coord, found := k.GetCoordinator(ctx, coordByAddr.CoordinatorID)
			require.True(t, found, "coordinator id not found")
			require.EqualValues(t, tt.msg.Address, coord.Address)
			require.EqualValues(t, tt.msg.Description, coord.Description)
			require.EqualValues(t, coordByAddr.CoordinatorID, coord.CoordinatorID)
		})
	}
}
