package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgCreateCoordinator(t *testing.T) {
	var (
		msg1           = sample.MsgCreateCoordinator(sample.Address(r))
		msg2           = sample.MsgCreateCoordinator(sample.Address(r))
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)
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
			got, err := ts.ProfileSrv.CreateCoordinator(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			coordByAddr, err := tk.ProfileKeeper.GetCoordinatorByAddress(sdkCtx, tt.msg.Address)
			require.NoError(t, err)
			require.EqualValues(t, tt.wantId, coordByAddr.CoordinatorID)
			require.EqualValues(t, tt.wantId, got.CoordinatorID)

			coord, found := tk.ProfileKeeper.GetCoordinator(sdkCtx, coordByAddr.CoordinatorID)
			require.True(t, found, "coordinator id not found")
			require.EqualValues(t, tt.msg.Address, coord.Address)
			require.EqualValues(t, tt.msg.Description, coord.Description)
			require.EqualValues(t, coordByAddr.CoordinatorID, coord.CoordinatorID)
			require.EqualValues(t, true, coord.Active)
		})
	}
}
