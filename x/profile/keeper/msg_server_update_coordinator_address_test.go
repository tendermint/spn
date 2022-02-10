package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorAddress(t *testing.T) {
	var (
		addr         = sample.Address()
		addr2        = sample.Address()
		coord1       = sample.MsgCreateCoordinator(sample.Address())
		coord2       = sample.MsgCreateCoordinator(sample.Address())
		disableCoord = sample.MsgCreateCoordinator(sample.Address())
		disableMsg   = sample.MsgDisableCoordinator(disableCoord.Address)
		ctx, k, srv  = setupMsgServer(t)
		wCtx         = sdk.WrapSDKContext(ctx)
	)
	if _, err := srv.CreateCoordinator(wCtx, &coord1); err != nil {
		t.Fatal(err)
	}
	if _, err := srv.CreateCoordinator(wCtx, &coord2); err != nil {
		t.Fatal(err)
	}
	if _, err := srv.CreateCoordinator(wCtx, &disableCoord); err != nil {
		t.Fatal(err)
	}
	if _, err := srv.DisableCoordinator(wCtx, &disableMsg); err != nil {
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
			err: types.ErrCoordAddressNotFound,
		}, {
			name: "new address already exist",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    coord1.Address,
				NewAddress: coord2.Address,
			},
			err: types.ErrCoordAlreadyExist,
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
			err: types.ErrCoordAlreadyExist,
		}, {
			name: "inactive coordinator",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    disableCoord.Address,
				NewAddress: addr2,
			},
			err: types.ErrCoordAddressNotFound,
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
			require.False(t, found, "old coordinator address was not removed")

			coordByAddr, found := k.GetCoordinatorByAddress(ctx, tt.msg.NewAddress)
			require.True(t, found, "coordinator by address not found")
			require.EqualValues(t, tt.msg.NewAddress, coordByAddr.Address)

			coord, found := k.GetCoordinator(ctx, coordByAddr.CoordinatorID)
			require.True(t, found, "coordinator id not found")
			require.EqualValues(t, tt.msg.NewAddress, coord.Address)
			require.EqualValues(t, coordByAddr.CoordinatorID, coord.CoordinatorID)
		})
	}
}
