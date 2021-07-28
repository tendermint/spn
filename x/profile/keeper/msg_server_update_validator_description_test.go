package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
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

func TestMsgUpdateValidatorDescription(t *testing.T) {
	var (
		addr1       = sample.AccAddress()
		addr2       = sample.AccAddress()
		ctx, k, srv = setupMsgServerAndKeeper(t)
		wCtx        = sdk.WrapSDKContext(ctx)
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.UpdateValidatorDescription(wCtx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			validator, found := k.GetValidatorByAddress(ctx, tt.msg.Address)
			assert.True(t, found, "validator not found")
			assert.EqualValues(t, tt.msg.Address, validator.Address)

			if len(tt.msg.Description.Identity) > 0 {
				assert.EqualValues(t, tt.msg.Description.Identity, validator.Description.Identity)
			}
			if len(tt.msg.Description.Website) > 0 {
				assert.EqualValues(t, tt.msg.Description.Website, validator.Description.Website)
			}
			if len(tt.msg.Description.Details) > 0 {
				assert.EqualValues(t, tt.msg.Description.Details, validator.Description.Details)
			}
			if len(tt.msg.Description.Moniker) > 0 {
				assert.EqualValues(t, tt.msg.Description.Moniker, validator.Description.Moniker)
			}
			if len(tt.msg.Description.SecurityContact) > 0 {
				assert.EqualValues(t, tt.msg.Description.SecurityContact, validator.Description.SecurityContact)
			}
		})
	}
}
