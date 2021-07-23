package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func validatorDescription(desc string) *types.ValidatorDescription {
	return &types.ValidatorDescription{
		Identity:        desc,
		Moniker:         "moniker " + desc,
		Website:         "https://cosmos.network/" + desc,
		SecurityContact: "foo",
		Details:         desc + " details",
	}
}

func Test_msgServer_UpdateValidatorDescription(t *testing.T) {
	var (
		addr1 = sample.AccAddress()
		addr2 = sample.AccAddress()
	)
	tests := []struct {
		name string
		msg  types.MsgUpdateValidatorDescription
		err  error
	}{
		{
			name: "update and create a new validator",
			msg: types.MsgUpdateValidatorDescription{
				Address:     addr1,
				Description: validatorDescription(addr1),
			},
		}, {
			name: "update a existing validator",
			msg: types.MsgUpdateValidatorDescription{
				Address:     addr1,
				Description: validatorDescription(addr2),
			},
		}, {
			name: "update and create anotherw validator",
			msg: types.MsgUpdateValidatorDescription{
				Address:     addr2,
				Description: validatorDescription(addr2),
			},
		},
	}
	srv, ctx := setupMsgServer(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.UpdateValidatorDescription(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
