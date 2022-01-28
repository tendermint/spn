package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		return false, sdkerrors.Wrapf(types.ErrChainNotFound, "%d", launchID)
	}
	if !chain.LaunchTriggered {
		return false, sdkerrors.Wrapf(types.ErrNotTriggeredLaunch, "%d", launchID)
	}
	if chain.GenesisChainID != chainID {
		return false, sdkerrors.Wrap(types.ErrInvalidGenesisChainID, chainID)
	}

	// all validators must be present in the launch module and
	// the total amount of self-delegation from the provided validators
	// must reach at least 2/3 of the total self delegation for the chain
	valSetSelfDelegation := 0.0
	for _, validator := range validatorSet.Validators {
		valAddr := sdk.AccAddress(validator.Address.Bytes())
		launchValidator, found := k.GetGenesisValidator(ctx, launchID, valAddr.String())
		if !found {
			return false, sdkerrors.Wrap(types.ErrValidatorNotFound, valAddr.String())
		}
		valConsPubKey := validator.PubKey.Bytes()
		if bytes.Compare(launchValidator.ConsPubKey, valConsPubKey) != 0 { // nolint
			return false, sdkerrors.Wrap(types.ErrInvalidConsPubKey, validator.PubKey.Address().String())
		}
		valSetSelfDelegation += float64(launchValidator.SelfDelegation.Amount.Int64())
	}

	// check if 2/3 of total self-delegation is reached from the provided validator set
	// GetTotalSelfDelegation is the sum of all self delegation
	totalSelfDelegation := float64(k.GetTotalSelfDelegation(ctx, launchID))
	reached := valSetSelfDelegation >= (2.0/3.0)*totalSelfDelegation
	return reached, nil
}

// GetTotalSelfDelegation returns the sum of all self delegation
func (k Keeper) GetTotalSelfDelegation(ctx sdk.Context, launchID uint64) int64 {
	validators := k.GetAllGenesisValidatorByLaunchID(ctx, launchID)
	var totalSelfDelegation int64 = 0
	for _, validator := range validators {
		totalSelfDelegation += validator.SelfDelegation.Amount.Int64()
	}
	return totalSelfDelegation
}
