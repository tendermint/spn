package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorDescription(t *testing.T) {
	var (
		addr           = sample.Address(r)
		msgCoord       = sample.MsgCreateCoordinator(sample.Address(r))
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)
	)
	_, err := ts.ProfileSrv.CreateCoordinator(ctx, &msgCoord)
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  types.MsgUpdateCoordinatorDescription
		err  error
	}{
		{
			name: "should prevent updating description of non existing coordinator",
			msg:  sample.MsgUpdateCoordinatorDescription(addr),
			err:  types.ErrCoordAddressNotFound,
		},
		{
			name: "should allow updating one value of coordinator description",
			msg: types.MsgUpdateCoordinatorDescription{
				Address: msgCoord.Address,
				Description: types.CoordinatorDescription{
					Identity: "update",
				},
			},
		},
		{
			name: "should allow updating all values of coordinator description",
			msg:  sample.MsgUpdateCoordinatorDescription(msgCoord.Address),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var oldCoord types.Coordinator
			var found bool
			if tt.err == nil {
				coordByAddr, err := tk.ProfileKeeper.GetCoordinatorByAddress(sdkCtx, tt.msg.Address)
				require.NoError(t, err, "coordinator by address not found")
				oldCoord, found = tk.ProfileKeeper.GetCoordinator(sdkCtx, coordByAddr.CoordinatorID)
				require.True(t, found, "coordinator not found")
			}

			_, err := ts.ProfileSrv.UpdateCoordinatorDescription(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			coordByAddr, err := tk.ProfileKeeper.GetCoordinatorByAddress(sdkCtx, tt.msg.Address)
			require.NoError(t, err, "coordinator by address not found")
			coord, found := tk.ProfileKeeper.GetCoordinator(sdkCtx, coordByAddr.CoordinatorID)
			require.True(t, found, "coordinator not found")
			require.EqualValues(t, tt.msg.Address, coord.Address)
			require.EqualValues(t, coordByAddr.CoordinatorID, coord.CoordinatorID)

			if len(tt.msg.Description.Identity) > 0 {
				require.EqualValues(t, tt.msg.Description.Identity, coord.Description.Identity)
			} else {
				require.EqualValues(t, oldCoord.Description.Identity, coord.Description.Identity)
			}

			if len(tt.msg.Description.Website) > 0 {
				require.EqualValues(t, tt.msg.Description.Website, coord.Description.Website)
			} else {
				require.EqualValues(t, oldCoord.Description.Website, coord.Description.Website)
			}

			if len(tt.msg.Description.Details) > 0 {
				require.EqualValues(t, tt.msg.Description.Details, coord.Description.Details)
			} else {
				require.EqualValues(t, oldCoord.Description.Details, coord.Description.Details)
			}
		})
	}
}
