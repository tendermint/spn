package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgDisableCoordinator(t *testing.T) {
	var (
		addr        = sample.Address()
		msgCoord    = sample.MsgCreateCoordinator(sample.Address())
		ctx, k, srv = setupMsgServer(t)
		wCtx        = sdk.WrapSDKContext(ctx)
	)
	if _, err := srv.CreateCoordinator(wCtx, &msgCoord); err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name string
		msg  types.MsgDisableCoordinator
		err  error
	}{
		{
			name: "not found coordinator address",
			msg:  types.MsgDisableCoordinator{Address: addr},
			err:  types.ErrCoordAddressNotFound,
		},
		{
			name: "successfully disable coordinator",
			msg:  types.MsgDisableCoordinator{Address: msgCoord.Address},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.DisableCoordinator(wCtx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			_, err = k.GetCoordinatorByAddress(ctx, tt.msg.Address)
			require.ErrorIs(t, err, types.ErrCoordAddressNotFound)

			coord, found := k.GetCoordinator(ctx, got.CoordinatorID)
			require.True(t, found)
			require.EqualValues(t, false, coord.Active)

		})
	}
}
