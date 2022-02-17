package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgDeleteValidator(t *testing.T) {
	var (
		addr1       = sample.Address()
		addr2       = sample.Address()
		ctx, k, srv = setupMsgServer(t)
		wCtx        = sdk.WrapSDKContext(ctx)
	)
	k.SetValidator(ctx, types.Validator{
		Address:     addr2,
		Description: types.ValidatorDescription{},
	})
	tests := []struct {
		name string
		msg  types.MsgDeleteValidator
		err  error
	}{
		{
			name: "delete a non-existing validator",
			msg: types.MsgDeleteValidator{
				Address: addr1,
			},
			err: types.ErrValidatorNotFound,
		}, {
			name: "delete an existing validator",
			msg: types.MsgDeleteValidator{
				Address: addr2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.DeleteValidator(wCtx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			_, found := k.GetValidator(ctx, tt.msg.Address)
			require.False(t, found, "validator was not removed")
		})
	}
}
