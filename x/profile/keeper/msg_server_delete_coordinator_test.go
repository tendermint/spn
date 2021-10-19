package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgDeleteCoordinator(t *testing.T) {
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
		msg  types.MsgDeleteCoordinator
		err  error
	}{
		{
			name: "not found coordinator address",
			msg:  types.MsgDeleteCoordinator{Address: addr},
			err:  types.ErrCoordAddressNotFound,
		},
		{
			name: "delete coordinator",
			msg:  types.MsgDeleteCoordinator{Address: msgCoord.Address},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.DeleteCoordinator(wCtx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			_, found := k.GetCoordinatorByAddress(ctx, tt.msg.Address)
			require.False(t, found, "coordinator by address was not removed")

			_, found = k.GetCoordinator(ctx, got.CoordinatorId)
			require.False(t, found, "coordinator id not removed")
		})
	}
}
