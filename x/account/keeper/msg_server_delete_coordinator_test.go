package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/account/types"
)

func TestMsgDeleteCoordinator(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgDeleteCoordinator
		err  error
	}{
		// TODO: Add test cases.
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
