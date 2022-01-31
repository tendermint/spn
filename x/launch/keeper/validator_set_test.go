package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	"github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"
)

func TestKeeper_CheckValidatorSet(t *testing.T) {
	var (
		k, _, _, _, _, _, ctx = setupMsgServer(t)

		validators           = []crypto.PubKey{sample.PubKey(), sample.PubKey(), sample.PubKey()}
		validatorSet         = tmtypes.ValidatorSet{}
		validatorNotFoundSet = tmtypes.ValidatorSet{}
		invalidValidatorSet  = tmtypes.ValidatorSet{}
	)
	notTriggeredLaunchID := k.AppendChain(ctx, types.Chain{
		CoordinatorID:   0,
		LaunchTriggered: false,
		GenesisChainID:  "spn-1",
	})
	invalidChainIDLaunchID := k.AppendChain(ctx, types.Chain{
		CoordinatorID:   0,
		LaunchTriggered: true,
		GenesisChainID:  "spn-10",
	})
	launchID := k.AppendChain(ctx, types.Chain{
		CoordinatorID:   0,
		LaunchTriggered: true,
		GenesisChainID:  "spn-1",
	})
	for _, validator := range validators {
		addr := sdk.AccAddress(validator.Address().Bytes())
		k.SetGenesisValidator(ctx, types.GenesisValidator{
			LaunchID:       launchID,
			Address:        addr.String(),
			ConsPubKey:     validator.Bytes(),
			SelfDelegation: sdk.NewCoin("spn", sdk.NewInt(1000)),
		})
		validatorSet.Validators = append(validatorSet.Validators,
			tmtypes.NewValidator(validator, 0),
		)
	}
	validatorNotFoundSet.Validators = append(
		validatorSet.Validators,
		tmtypes.NewValidator(sample.PubKey(), 0),
	)
	invalidValidatorSet.Validators = validatorSet.Validators[:1]
	type args struct {
		launchID     uint64
		chainID      string
		validatorSet tmtypes.ValidatorSet
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "chain not found",
			args: args{
				launchID:     999,
				chainID:      "spn-1",
				validatorSet: validatorSet,
			},
			err: types.ErrChainNotFound,
		},
		{
			name: "chain not triggered launch",
			args: args{
				launchID:     notTriggeredLaunchID,
				chainID:      "spn-1",
				validatorSet: validatorSet,
			},
			err: types.ErrNotTriggeredLaunch,
		},
		{
			name: "invalid genesis chain id",
			args: args{
				launchID:     invalidChainIDLaunchID,
				chainID:      "spn-1",
				validatorSet: validatorSet,
			},
			err: types.ErrInvalidGenesisChainID,
		},
		{
			name: "validator not found",
			args: args{
				launchID:     launchID,
				chainID:      "spn-1",
				validatorSet: validatorNotFoundSet,
			},
			err: types.ErrValidatorNotFound,
		},
		{
			name: "invalid validator set",
			args: args{
				launchID:     launchID,
				chainID:      "spn-1",
				validatorSet: invalidValidatorSet,
			},
			err: types.ErrMinSelfDelegationNotReached,
		},
		{
			name: "valid validator set",
			args: args{
				launchID:     launchID,
				chainID:      "spn-1",
				validatorSet: validatorSet,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := k.CheckValidatorSet(ctx, tt.args.launchID, tt.args.chainID, tt.args.validatorSet)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestKeeper_GetTotalSelfDelegation(t *testing.T) {
	k, ctx := testkeeper.Launch(t)
	launchID := 10
	validators := createNGenesisValidatorByLaunchID(k, ctx, launchID)
	var totalSelfDelegation int64 = 0
	for _, validator := range validators {
		totalSelfDelegation += validator.SelfDelegation.Amount.Int64()
	}
	require.Equal(t, totalSelfDelegation, k.GetTotalSelfDelegation(ctx, uint64(launchID)))
}
