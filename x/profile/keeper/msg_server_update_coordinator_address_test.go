package keeper

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorAddress(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgUpdateCoordinatorAddress
		want uint64
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    "invalid_address",
				NewAddress: "cosmos1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm",
			},
			err: status.Error(codes.InvalidArgument,
				"invalid coordinator address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "invalid new address",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    "cosmos1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm",
				NewAddress: "invalid_address",
			},
			err: status.Error(codes.InvalidArgument,
				"invalid new coordinator address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "equal addresses",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    "cosmos1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm",
				NewAddress: "cosmos1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm",
			},
			err: status.Error(codes.InvalidArgument,
				"address are equal of new address (cosmos1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm)"),
		}, {
			name: "not found address",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    "cosmos1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm",
				NewAddress: "cosmos12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c",
			},
			err: status.Error(codes.NotFound,
				"coordinator address not found: cosmos1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"),
		},
		// TODO: valid tests cases
	}
	srv, ctx := setupMsgServer(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.UpdateCoordinatorAddress(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.want, got.CoordinatorId)
		})
	}
}
