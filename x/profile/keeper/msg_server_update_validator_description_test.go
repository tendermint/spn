package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateValidatorDescription(t *testing.T) {
	var (
		addr1       = sample.AccAddress()
		addr2       = sample.AccAddress()
		ctx, k, srv = setupMsgServer(t)
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
				Description: sample.ValidatorDescription(addr1),
			},
		}, {
			name: "update a existing validator",
			msg: types.MsgUpdateValidatorDescription{
				Address:     addr1,
				Description: sample.ValidatorDescription(addr2),
			},
		}, {
			name: "update and create another validator",
			msg: types.MsgUpdateValidatorDescription{
				Address:     addr2,
				Description: sample.ValidatorDescription(addr2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldValidator, oldFound := k.GetValidator(ctx, tt.msg.Address)

			_, err := srv.UpdateValidatorDescription(wCtx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			validator, found := k.GetValidator(ctx, tt.msg.Address)
			assert.True(t, found, "validator not found")
			assert.EqualValues(t, tt.msg.Address, validator.Address)

			if len(tt.msg.Description.Identity) > 0 {
				assert.EqualValues(t, tt.msg.Description.Identity, validator.Description.Identity)
			} else if oldFound {
				require.EqualValues(t, oldValidator.Description.Identity, oldValidator.Description.Identity)
			}

			if len(tt.msg.Description.Website) > 0 {
				assert.EqualValues(t, tt.msg.Description.Website, validator.Description.Website)
			} else if oldFound {
				require.EqualValues(t, oldValidator.Description.Website, oldValidator.Description.Website)
			}

			if len(tt.msg.Description.Details) > 0 {
				assert.EqualValues(t, tt.msg.Description.Details, validator.Description.Details)
			} else if oldFound {
				require.EqualValues(t, oldValidator.Description.Details, oldValidator.Description.Details)
			}

			if len(tt.msg.Description.Moniker) > 0 {
				assert.EqualValues(t, tt.msg.Description.Moniker, validator.Description.Moniker)
			} else if oldFound {
				require.EqualValues(t, oldValidator.Description.Moniker, oldValidator.Description.Moniker)
			}

			if len(tt.msg.Description.SecurityContact) > 0 {
				assert.EqualValues(t, tt.msg.Description.SecurityContact, validator.Description.SecurityContact)
			} else if oldFound {
				require.EqualValues(t, oldValidator.Description.SecurityContact, oldValidator.Description.SecurityContact)
			}
		})
	}
}
