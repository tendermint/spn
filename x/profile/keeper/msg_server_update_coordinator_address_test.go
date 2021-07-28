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

func TestMsgUpdateCoordinatorAddress(t *testing.T) {
	var (
		addr        = sample.AccAddress()
		coord1      = msgCreateCoordinator()
		coord2      = msgCreateCoordinator()
		ctx, k, srv = setupMsgServer(t)
		wCtx        = sdk.WrapSDKContext(ctx)
	)
	if _, err := srv.CreateCoordinator(wCtx, &coord1); err != nil {
		t.Fatal(err)
	}
	if _, err := srv.CreateCoordinator(wCtx, &coord2); err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name string
		msg  types.MsgUpdateCoordinatorAddress
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
		}, {
			name: "update second coordinator address update",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    coord2.Address,
				NewAddress: coord1.Address,
			},
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
			_, err := srv.UpdateCoordinatorAddress(wCtx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			_, found := k.GetCoordinatorByAddress(ctx, tt.msg.Address)
			assert.False(t, found, "old coordinator address was not removed")

			coordByAddr, found := k.GetCoordinatorByAddress(ctx, tt.msg.NewAddress)
			assert.True(t, found, "coordinator by address not found")
			assert.EqualValues(t, tt.msg.NewAddress, coordByAddr.Address)

			coord := k.GetCoordinator(ctx, coordByAddr.CoordinatorId)
			assert.True(t, found, "coordinator id not found")
			assert.EqualValues(t, tt.msg.NewAddress, coord.Address)
			assert.EqualValues(t, coordByAddr.CoordinatorId, coord.CoordinatorId)
		})
	}
}
