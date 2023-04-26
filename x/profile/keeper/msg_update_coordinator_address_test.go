package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorAddress(t *testing.T) {
	var (
		addr           = sample.Address(r)
		coord1         = sample.MsgCreateCoordinator(sample.Address(r))
		coord2         = sample.MsgCreateCoordinator(sample.Address(r))
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
			name: "should prevent updating a non existing coordinator",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    addr,
				NewAddress: addr,
			},
			err: types.ErrCoordAddressNotFound,
		}, {
			name: "should prevent updating with an address already associated to a coordinator",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    coord1.Address,
				NewAddress: coord2.Address,
			},
			err: types.ErrCoordAlreadyExist,
		}, {
			name: "should allow updating coordinator address",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    coord1.Address,
				NewAddress: addr,
			},
		}, {
			name: "should allow updating coordinator address a second time",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    coord2.Address,
				NewAddress: coord1.Address,
			},
		}, {
			name: "should prevent updating from previous coordinator address",
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
