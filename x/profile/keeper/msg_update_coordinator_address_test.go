package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorAddress(t *testing.T) {
	var (
		addr           = sample.Address()
		coord1         = sample.MsgCreateCoordinator(sample.Address())
		coord2         = sample.MsgCreateCoordinator(sample.Address())
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)
	)
	_, err := ts.ProfileSrv.CreateCoordinator(ctx, &coord1)
	require.NoError(t, err)
	_, err = ts.ProfileSrv.CreateCoordinator(ctx, &coord2)
	require.NoError(t, err)

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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ts.ProfileSrv.UpdateCoordinatorAddress(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			_, err = tk.ProfileKeeper.GetCoordinatorByAddress(sdkCtx, tt.msg.Address)
			require.ErrorIs(t, err, types.ErrCoordAddressNotFound, "old coordinator address was not removed")

			coordByAddr, err := tk.ProfileKeeper.GetCoordinatorByAddress(sdkCtx, tt.msg.NewAddress)
			require.NoError(t, err, "coordinator by address not found")
			require.EqualValues(t, tt.msg.NewAddress, coordByAddr.Address)

			coord, found := tk.ProfileKeeper.GetCoordinator(sdkCtx, coordByAddr.CoordinatorID)
			require.True(t, found, "coordinator id not found")
			require.EqualValues(t, tt.msg.NewAddress, coord.Address)
			require.EqualValues(t, coordByAddr.CoordinatorID, coord.CoordinatorID)
		})
	}
}
