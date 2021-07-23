package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorDescription(t *testing.T) {
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
		msg  types.MsgUpdateCoordinatorDescription
		err  error
	}{
		{
			name: "not found address",
			msg: types.MsgUpdateCoordinatorDescription{
				Address: addr,
			},
			err: sdkerrors.Wrap(types.ErrCoordAddressNotFound, addr),
		}, {
			name: "update one value",
			msg: types.MsgUpdateCoordinatorDescription{
				Address: coord.Address,
				Description: &types.CoordinatorDescription{
					Identity: "update",
				},
			},
		}, {
			name: "update all values",
			msg: types.MsgUpdateCoordinatorDescription{
				Address: coord.Address,
				Description: &types.CoordinatorDescription{
					Identity: "update",
					Website:  "update",
					Details:  "update",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.UpdateCoordinatorDescription(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
