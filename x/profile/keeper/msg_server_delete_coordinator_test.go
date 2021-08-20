package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgDeleteCoordinator(t *testing.T) {
	var (
		addr        = sample.AccAddress()
		msgCoord    = sample.MsgCreateCoordinator(sample.AccAddress())
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
			err:  sdkerrors.Wrap(types.ErrCoordAddressNotFound, addr),
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
			assert.False(t, found, "coordinator by address was not removed")

			found = k.HasCoordinator(ctx, got.CoordinatorId)
			assert.False(t, found, "coordinator id not removed")
		})
	}
}
