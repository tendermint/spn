package keeper

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// CheckValidatorSet check the information about a validator
// set (used to create an IBC client) are valid
func (k Keeper) CheckValidatorSet(
	ctx sdk.Context,
	launchID uint64,
	chainID string,
	validatorSet tmtypes.ValidatorSet,
) (bool, error) {
	// check chain ID
	chain, found := k.GetChain(ctx, launchID)
	if !found {
		return false, types.ErrChainNotFound
	}
	if !chain.LaunchTriggered {
		return false, types.ErrNotTriggeredLaunch
	}
	if chain.GenesisChainID != chainID {
		return false, types.ErrInvalidGenesisChainID
	}

	// all validators must be present in the launch module and
	// the total amount of self-delegation from the provided validators
	// must reach at least 2/3 of the total self delegation for the chain
	var valSetSelfDelegation int64 = 0
	for _, validator := range validatorSet.Validators {
		launchValidator, found := k.GetGenesisValidator(ctx, launchID, validator.Address.String())
		if !found {
			return false, types.ErrValidatorNotFound
		}
		valConsPubKey := validator.PubKey.Bytes()
		if bytes.Compare(launchValidator.ConsPubKey, valConsPubKey) != 0 {
			return false, fmt.Errorf("invalid consensus pub key %s", validator.PubKey.Address().String())
		}
		valSetSelfDelegation += launchValidator.SelfDelegation.Amount.Int64()
	}

	// check if 2/3 of total self-delegation is reached from the provided validator set
	// GetTotalSelfDelegation is the sum of all self delegation
	totalSelfDelegation, err := k.GetTotalSelfDelegation(ctx, launchID)
	if err != nil {
		return false, err
	}
	reached := valSetSelfDelegation > (2/3)*totalSelfDelegation
	return reached, nil
}

// GetTotalSelfDelegation returns the sum of all self delegation
func (k Keeper) GetTotalSelfDelegation(ctx sdk.Context, launchID uint64) (int64, error) {
	return 0, nil
}
