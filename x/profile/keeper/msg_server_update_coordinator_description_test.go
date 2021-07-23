package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorDescription(t *testing.T) {
	var (
		addr1 = sample.AccAddress()
	)
	tests := []struct {
		name string
		msg  types.MsgUpdateCoordinatorDescription
		want uint64
		err  error
	}{
		{
			name: "not found address",
			msg: types.MsgUpdateCoordinatorDescription{
				Address: addr1,
			},
			err: sdkerrors.Wrap(types.ErrCoordAddressNotFound, addr1),
		},
		// TODO: valid tests cases
	}
	srv, ctx := setupMsgServer(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.UpdateCoordinatorDescription(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.want, got.CoordinatorId)
		})
	}
}
