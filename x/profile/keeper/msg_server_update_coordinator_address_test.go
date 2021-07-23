package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorAddress(t *testing.T) {
	var (
		addr   = sample.AccAddress()
		coord1 = msgCreateCoordinator()
		coord2 = msgCreateCoordinator()
	)
	srv, ctx := setupMsgServer(t)
	if _, err := srv.CreateCoordinator(ctx, &coord1); err != nil {
		t.Fatal(err)
	}
	if _, err := srv.CreateCoordinator(ctx, &coord2); err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name string
		msg  types.MsgUpdateCoordinatorAddress
		want uint64
		err  error
	}{
		{
			name: "not found address",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    addr,
				NewAddress: addr,
			},
			err: sdkerrors.Wrap(types.ErrCoordAddressNotFound, addr),
		}, {
			name: "new address already exist",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    coord1.Address,
				NewAddress: coord2.Address,
			},
			err: sdkerrors.Wrap(types.ErrCoordAlreadyExist, "new address already have a coordinator: 1"),
		}, {
			name: "update first coordinator address update",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    coord1.Address,
				NewAddress: addr,
			},
			want: 0,
		}, {
			name: "update second coordinator address update",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    coord2.Address,
				NewAddress: coord1.Address,
			},
			want: 1,
		}, {
			name: "new address already updated",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    addr,
				NewAddress: coord1.Address,
			},
			err: sdkerrors.Wrap(types.ErrCoordAlreadyExist, "new address already have a coordinator: 0"),
		},
	}
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
