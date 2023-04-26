package keeper

import (
	"encoding/base64"

	sdkerrors "cosmossdk.io/errors"
	tmtypes "github.com/cometbft/cometbft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
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
	if chain.MonitoringConnected {
		return sdkerrors.Wrapf(types.ErrChainMonitoringConnected, "%d", launchID)
	}
	if chain.GenesisChainID != chainID {
		return sdkerrors.Wrap(types.ErrInvalidGenesisChainID, chainID)
	}

	validators, totalSelfDelegation := k.GetValidatorsAndTotalDelegation(ctx, launchID)

	// all validators must be present in the launch module and
	// the total amount of self-delegation from the provided validators
	// must reach at least 2/3 of the total self delegation for the chain
	valSetSelfDelegation := sdk.ZeroDec()
	for _, validator := range validatorSet.Validators {
		consPubKey := base64.StdEncoding.EncodeToString(validator.PubKey.Bytes())
		launchValidator, found := validators[consPubKey]
		if !found {
			return sdkerrors.Wrapf(
				types.ErrValidatorNotFound,
				"validator consensus pub key not found: %s",
				validator.PubKey.Address().String(),
			)
		}
		valSetSelfDelegation = valSetSelfDelegation.Add(sdk.NewDecFromInt(launchValidator.SelfDelegation.Amount))
	}

	// check if 2/3 of total self-delegation is reached from the provided validator set
	// GetTotalSelfDelegation is the sum of all self delegation
	minSelfDelegation := totalSelfDelegation.Mul(sdk.NewDecWithPrec(6666, 4))
	if valSetSelfDelegation.LT(minSelfDelegation) {
		return sdkerrors.Wrap(types.ErrMinSelfDelegationNotReached, validatorSet.String())
	}
	return nil
}
