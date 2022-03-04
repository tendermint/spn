package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/x/reward/types"
)

// DistributeRewards distributes rewards based on the monitoring packet
// this `closeRewardPool` is a boolean that specifies if the reward pool
// must be closed after the reward distribution.
// In a model where rewards are distributed in a single round, this
// boolean is always `true`. In this way, if the monitoring packet is
// distributed earlier and `lastBlockHeight < rewardPool.LastRewardHeight`
// only a portion of the rewards will be distributed and the remaining is
// refunded to the reward's provider.
// When rewards are distributed periodically, this value is set to `false`
// so the reward pool is not closed as long as `lastBlockHeight` does not
// reach `rewardPool.LastRewardHeight`
func (k Keeper) DistributeRewards(
	ctx sdk.Context,
	launchID uint64,
	signatureCounts spntypes.SignatureCounts,
	lastBlockHeight uint64,
	closeRewardPool bool,
) error {
	// get the reward pool related to the chain
	rewardPool, found := k.GetRewardPool(ctx, launchID)
	if !found {
		return sdkerrors.Wrapf(types.ErrRewardPoolNotFound, "%d", launchID)
	}

	if rewardPool.Closed {
		return sdkerrors.Wrapf(types.ErrRewardPoolClosed, "%d", launchID)
	}

	provider, err := sdk.AccAddressFromBech32(rewardPool.Provider)
	if err != nil {
		return spnerrors.Criticalf("can't parse the provider address %s", err.Error())
	}

	// only the monitored blocks relative to last reward height are rewarded
	blockRatioNumerator := sdk.NewDec(int64(lastBlockHeight)).Sub(sdk.NewDec(int64(rewardPool.CurrentRewardHeight)))
	blockRatioDenominator := sdk.NewDec(int64(rewardPool.LastRewardHeight)).Sub(sdk.NewDec(int64(rewardPool.CurrentRewardHeight)))
	blockRatio := blockRatioNumerator.Quo(blockRatioDenominator)
	if blockRatio.GT(sdk.OneDec()) {
		blockRatio = sdk.OneDec()
	}

	// store the total relative signature distributed to calculate the refund for the round
	totalRelativeSignaturesDistributed := sdk.ZeroDec()

	// store rewards to distributes per address
	rewardsToDistribute := make(map[string]sdk.Coins)

	// calculate the total reward for all validators
	for _, signatureCount := range signatureCounts.Counts {

		// get the validator address from the cons address
		// if the validator is not registered, reward distribution is skipped
		// all funds are sent back to the coordinator
		validatorByOpAddr, found := k.profileKeeper.GetValidatorByOperatorAddress(ctx, string(signatureCount.ConsAddress))
		if found {
			validator, found := k.profileKeeper.GetValidator(ctx, validatorByOpAddr.ValidatorAddress)
			if !found {
				return spnerrors.Criticalf(
					"validator by consensus address not associated with a validator %s",
					validatorByOpAddr.ValidatorAddress,
				)
			}
			// calculate the total relative signature distributed to calculate the refund for the round
			totalRelativeSignaturesDistributed = totalRelativeSignaturesDistributed.Add(signatureCount.RelativeSignatures)

			// compute reward relative to the signature and block count
			// and update reward pool
			signatureRatio := signatureCount.RelativeSignatures.Quo(
				sdk.NewDecFromInt(sdk.NewIntFromUint64(signatureCounts.BlockCount)),
			)
			rewards, err := CalculateRewards(blockRatio, signatureRatio, rewardPool.RemainingCoins)
			if err != nil {
				return spnerrors.Criticalf("invalid reward: %s", err.Error())
			}
			rewardsToDistribute[validator.Address] = rewards
		}
	}

	// distribute the rewards to validators
	for address, rewards := range rewardsToDistribute {
		coins, isNegative := rewardPool.RemainingCoins.SafeSub(rewards)
		if isNegative {
			return spnerrors.Criticalf("negative reward pool: %s", rewardPool.RemainingCoins.String())
		}
		rewardPool.RemainingCoins = coins

		// send rewards to the address
		account, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return spnerrors.Criticalf("can't parse address %s", err.Error())
		}
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, account, rewards); err != nil {
			return spnerrors.Criticalf("send rewards error: %s", err.Error())
		}
	}

	// if the reward pool is closed or last reward height is reached
	// the remaining coins are refunded and reward pool is deleted
	if closeRewardPool || lastBlockHeight >= rewardPool.LastRewardHeight {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			provider,
			rewardPool.RemainingCoins); err != nil {
			return spnerrors.Criticalf("send rewards error: %s", err.Error())
		}

		// close the pool
		rewardPool.Closed = true
		rewardPool.RemainingCoins = rewardPool.RemainingCoins.Sub(rewardPool.RemainingCoins) // sub coins transferred
		k.SetRewardPool(ctx, rewardPool)
		return nil
	}

	// Otherwise, the refund is relative to the block ratio and the reward pool is updated
	// refundRatio is blockCount.
	// This is sum of signaturesRelative values from validator to compute refund
	blockCount := sdk.NewDecFromInt(sdk.NewIntFromUint64(signatureCounts.BlockCount))
	refundRatioNumerator := blockCount.Sub(totalRelativeSignaturesDistributed)
	refundRatio := refundRatioNumerator.Quo(blockCount)
	refund, err := CalculateRewards(blockRatio, refundRatio, rewardPool.RemainingCoins)
	if err != nil {
		return spnerrors.Criticalf("invalid reward: %s", err.Error())
	}

	// if refund is non-null, refund is sent to the provider
	if !refund.IsZero() {
		coins, isNegative := rewardPool.RemainingCoins.SafeSub(refund)
		if isNegative {
			return spnerrors.Criticalf("negative reward pool: %s", rewardPool.RemainingCoins.String())
		}
		rewardPool.RemainingCoins = coins

		// send rewards to the address
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			provider,
			rewardPool.RemainingCoins); err != nil {
			return spnerrors.Criticalf("send rewards error: %s", err.Error())
		}
	}

	// update the current reward height for next reward
	rewardPool.CurrentRewardHeight = lastBlockHeight
	k.SetRewardPool(ctx, rewardPool)
	return nil
}

// CalculateRewards calculates the reward relative to the signature and block ratio
func CalculateRewards(blockRatio, signatureRatio sdk.Dec, coins sdk.Coins) (sdk.Coins, error) {
	// ratio can't be greater than one
	if blockRatio.GT(sdk.OneDec()) {
		return nil, fmt.Errorf("block ratio is greater than 1 %s", blockRatio.String())
	}
	if signatureRatio.GT(sdk.OneDec()) {
		return nil, fmt.Errorf("signature ratio is greater than 1 %s", signatureRatio.String())
	}

	// if one ratio is zero, rewards are null
	if blockRatio.IsZero() || signatureRatio.IsZero() {
		return sdk.NewCoins(), nil
	}

	// calculate rewards
	rewards := sdk.NewCoins()
	for _, coin := range coins {
		amount := blockRatio.Mul(signatureRatio).Mul(coin.Amount.ToDec())
		coin.Amount = amount.TruncateInt()
		rewards = rewards.Add(coin)
	}
	return rewards, nil
}

// ConvertOperatorAddress returns the operator address with the spn account prefix
func ConvertOperatorAddress(addr string) (string, error) {
	_, decoded, err := bech32.DecodeAndConvert(addr)
	if err != nil {
		return "", err
	}
	return bech32.ConvertAndEncode(spntypes.AccountAddressPrefix, decoded)
}