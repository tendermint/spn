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
		addr  = sample.AccAddress()
		coord = msgCreateCoordinator()
	)
	srv, ctx := setupMsgServer(t)
	if _, err := srv.CreateCoordinator(ctx, &coord); err != nil {
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
			msg:  types.MsgDeleteCoordinator{Address: coord.Address},
		},
	}
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
