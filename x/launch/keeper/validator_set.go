package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/launch/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// CheckValidatorSet checks the information about a validator
// set (used to create an IBC client) is valid
func (k Keeper) CheckValidatorSet(
	ctx sdk.Context,
	launchID uint64,
	chainID string,
	validatorSet tmtypes.ValidatorSet,
) error {
	// check chain ID
	chain, found := k.GetChain(ctx, launchID)
	if !found {
		return sdkerrors.Wrapf(types.ErrChainNotFound, "%d", launchID)
	}
	if !chain.LaunchTriggered {
		return sdkerrors.Wrapf(types.ErrNotTriggeredLaunch, "%d", launchID)
	}
	if chain.GenesisChainID != chainID {
		return sdkerrors.Wrap(types.ErrInvalidGenesisChainID, chainID)
	}

	// all validators must be present in the launch module and
	// the total amount of self-delegation from the provided validators
	// must reach at least 2/3 of the total self delegation for the chain
	valSetSelfDelegation := sdk.NewDec(0)
	for _, validator := range validatorSet.Validators {
		valConsPubKey := validator.PubKey.Bytes()
		launchValidator, found := k.GetGenesisValidatorByConsPubKey(ctx, launchID, valConsPubKey)
		if !found {
			return sdkerrors.Wrap(types.ErrInvalidConsPubKey, validator.PubKey.Address().String())
		}
		valSetSelfDelegation = valSetSelfDelegation.Add(launchValidator.SelfDelegation.Amount.ToDec())
	}

	// check if 2/3 of total self-delegation is reached from the provided validator set
	// GetTotalSelfDelegation is the sum of all self delegation
	totalSelfDelegation := k.GetTotalSelfDelegation(ctx, launchID)
	minSelfDelegation := totalSelfDelegation.Mul(sdk.NewDecWithPrec(6666, 4))
	if valSetSelfDelegation.LT(minSelfDelegation) {
		return sdkerrors.Wrap(types.ErrMinSelfDelegationNotReached, validatorSet.String())
	}
	return nil
}

// GetTotalSelfDelegation returns the sum of all self delegation
func (k Keeper) GetTotalSelfDelegation(ctx sdk.Context, launchID uint64) sdk.Dec {
	validators := k.GetAllGenesisValidatorByLaunchID(ctx, launchID)
	totalSelfDelegation := sdk.NewDec(0)
	for _, validator := range validators {
		totalSelfDelegation = totalSelfDelegation.Add(validator.SelfDelegation.Amount.ToDec())
	}
	return totalSelfDelegation
}
