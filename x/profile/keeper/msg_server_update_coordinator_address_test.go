package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorAddress(t *testing.T) {
	var (
		addr1 = sample.AccAddress()
		addr2 = sample.AccAddress()
		addr3 = sample.AccAddress()
	)
	tests := []struct {
		name string
		msg  types.MsgUpdateCoordinatorAddress
		want uint64
		err  error
	}{
		{
			name: "not found address",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    addr1,
				NewAddress: addr2,
			},
			err: sdkerrors.Wrap(types.ErrCoordAddressNotFound, addr1),
		}, {
			name: "new address already exist",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    addr3,
				NewAddress: addr2,
			},
			err: sdkerrors.Wrap(types.ErrCoordAlreadyExist, "new address already have a coordinator: 1"),
		},
		// TODO: valid tests cases
	}
	srv, ctx := setupMsgServer(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.UpdateCoordinatorAddress(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.want, got.CoordinatorId)
		})
	}
}
