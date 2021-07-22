package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/account/types"
)

func TestMsgCreateCoordinator(t *testing.T) {
	var (
		addr1 = sample.AccAddress()
		addr2 = sample.AccAddress()
	)
	tests := []struct {
		name string
		msg  types.MsgCreateCoordinator
		want uint64
		err  error
	}{
		{
			name: "valid coordinator 1",
			msg: types.MsgCreateCoordinator{
				Address: addr1,
				Description: &types.CoordinatorDescription{
					Identity: "identity_1",
					Website:  "website_1",
					Details:  "details_1",
				},
			},
			want: 0,
		}, {
			name: "valid coordinator 2",
			msg: types.MsgCreateCoordinator{
				Address: addr2,
				Description: &types.CoordinatorDescription{
					Identity: "identity_2",
					Website:  "website_2",
					Details:  "details_2",
				},
			},
			want: 1,
		}, {
			name: "already exist address",
			msg: types.MsgCreateCoordinator{
				Address: addr2,
				Description: &types.CoordinatorDescription{
					Identity: "identity_2",
					Website:  "website_2",
					Details:  "details_2",
				},
			},
			err: sdkerrors.Wrap(types.ErrCoordAlreadyExist, "coordinatorId: 1"),
		},
	}
	srv, ctx := setupMsgServer(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.CreateCoordinator(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.want, got.CoordinatorId)
		})
	}
}
