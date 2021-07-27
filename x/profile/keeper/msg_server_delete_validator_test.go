package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgDeleteValidator(t *testing.T) {
	var addr = sample.AccAddress()
	srv, ctx := setupMsgServer(t)
	// TODO create a default validator to be deleted
	tests := []struct {
		name string
		msg  types.MsgDeleteValidator
		err  error
	}{
		{
			name: "delete a non-existing validator",
			msg: types.MsgDeleteValidator{
				Address: sample.AccAddress(),
			},
		}, {
			name: "delete an existing validator",
			msg: types.MsgDeleteValidator{
				Address: addr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.DeleteValidator(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
