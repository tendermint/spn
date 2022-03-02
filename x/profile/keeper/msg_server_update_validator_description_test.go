package keeper_test

import (
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateValidatorDescription(t *testing.T) {
	var (
		addr1          = sample.Address()
		addr2          = sample.Address()
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)
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
			oldValidator, oldFound := tk.ProfileKeeper.GetValidator(sdkCtx, tt.msg.Address)

			_, err := ts.ProfileSrv.UpdateValidatorDescription(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			validator, found := tk.ProfileKeeper.GetValidator(sdkCtx, tt.msg.Address)
			require.True(t, found, "validator not found")
			require.EqualValues(t, tt.msg.Address, validator.Address)

			if len(tt.msg.Description.Identity) > 0 {
				require.EqualValues(t, tt.msg.Description.Identity, validator.Description.Identity)
			} else if oldFound {
				require.EqualValues(t, oldValidator.Description.Identity, oldValidator.Description.Identity)
			}

			if len(tt.msg.Description.Website) > 0 {
				require.EqualValues(t, tt.msg.Description.Website, validator.Description.Website)
			} else if oldFound {
				require.EqualValues(t, oldValidator.Description.Website, oldValidator.Description.Website)
			}

			if len(tt.msg.Description.Details) > 0 {
				require.EqualValues(t, tt.msg.Description.Details, validator.Description.Details)
			} else if oldFound {
				require.EqualValues(t, oldValidator.Description.Details, oldValidator.Description.Details)
			}

			if len(tt.msg.Description.Moniker) > 0 {
				require.EqualValues(t, tt.msg.Description.Moniker, validator.Description.Moniker)
			} else if oldFound {
				require.EqualValues(t, oldValidator.Description.Moniker, oldValidator.Description.Moniker)
			}

			if len(tt.msg.Description.SecurityContact) > 0 {
				require.EqualValues(t, tt.msg.Description.SecurityContact, validator.Description.SecurityContact)
			} else if oldFound {
				require.EqualValues(t, oldValidator.Description.SecurityContact, oldValidator.Description.SecurityContact)
			}
		})
	}
}
