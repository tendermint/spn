package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgDeleteCoordinator(t *testing.T) {
	var (
		addr1 = sample.AccAddress()
	)
	tests := []struct {
		name string
		msg  types.MsgDeleteCoordinator
		err  error
	}{
		{
			name: "not found coordinator address",
			msg:  types.MsgDeleteCoordinator{Address: addr1},
			err:  sdkerrors.Wrap(types.ErrCoordAddressNotFound, addr1),
		},
		// TODO: Add more test cases.
	}
	srv, ctx := setupMsgServer(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.DeleteCoordinator(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
