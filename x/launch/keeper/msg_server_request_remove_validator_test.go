package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestRemoveValidator(t *testing.T) {
	var (
		//addr     = sample.AccAddress()
		srv, ctx = setupMsgServer(t)
	)
	tests := []struct {
		name string
		msg  types.MsgRequestRemoveValidator
		err  error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.RequestRemoveValidator(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			// TODO add more assertion
		})
	}
}
