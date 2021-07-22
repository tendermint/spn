package keeper

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/account/types"
)

func TestMsgCreateCoordinator(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCreateCoordinator
		want uint64
		err  error
	}{
		{
			name: "valid coordinator 1",
			msg: types.MsgCreateCoordinator{
				Address: "cosmos1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm",
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
				Address: "cosmos12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c",
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
				Address: "cosmos12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c",
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
