package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/account/types"
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
		msg1 = msgCreateCoordinator()
		msg2 = msgCreateCoordinator()
	)
	tests := []struct {
		name string
		msg  types.MsgCreateCoordinator
		want uint64
		err  error
	}{
		{
			name: "valid coordinator 1",
			msg:  msg1,
			want: 0,
		}, {
			name: "valid coordinator 2",
			msg:  msg2,
			want: 1,
		}, {
			name: "already exist address",
			msg:  msg2,
			err:  sdkerrors.Wrap(types.ErrCoordAlreadyExist, "coordinatorId: 1"),
		},
	}
	srv, ctx := setupMsgServer(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.CreateCoordinator(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.want, got.CoordinatorId)
		})
	}
}
