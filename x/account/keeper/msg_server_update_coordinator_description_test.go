package keeper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/account/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_msgServer_UpdateCoordinatorDescription(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgUpdateCoordinatorDescription
		want uint64
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdateCoordinatorDescription{
				Address: "invalid_address",
				Description: &types.CoordinatorDescription{
					Identity: "identity_invalid_address",
					Website:  "website_invalid_address",
					Details:  "details_invalid_address",
				},
			},
			err: status.Error(codes.InvalidArgument,
				"invalid coordinator address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "not found address",
			msg: types.MsgUpdateCoordinatorDescription{
				Address: "cosmos1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm",
				Description: &types.CoordinatorDescription{
					Identity: "identity_1",
					Website:  "website_1",
					Details:  "details_1",
				},
			},
			err: status.Error(codes.NotFound,
				"coordinator address not found: cosmos1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"),
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
