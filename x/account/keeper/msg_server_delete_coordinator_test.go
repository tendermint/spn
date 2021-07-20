package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/account/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestMsgDeleteCoordinator(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgDeleteCoordinator
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgDeleteCoordinator{
				Address: "invalid_address",
			},
			err: status.Error(codes.InvalidArgument,
				"invalid coordinator address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "not found coordinator address",
			msg: types.MsgDeleteCoordinator{
				Address: "cosmos12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c",
			},
			err: status.Error(codes.NotFound,
				"coordinator address not found: cosmos12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c"),
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
