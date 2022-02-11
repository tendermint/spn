package keeper

import (
	"fmt"

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
// so the reward pool is not closed as long as `lastBlockHeight` doesn’t
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
	provider, err := sdk.AccAddressFromBech32(rewardPool.Provider)
	if err != nil {
		return spnerrors.Criticalf("can't parse the provider address %s", err.Error())
	}

	// only the monitored blocks relative to last reward height are rewarded
	blockRatioA := sdk.NewDec(int64(lastBlockHeight)).Sub(sdk.NewDec(int64(rewardPool.CurrentRewardHeight)))
	blockRatioB := sdk.NewDec(int64(rewardPool.LastRewardHeight)).Sub(sdk.NewDec(int64(rewardPool.CurrentRewardHeight)))
	blockRatio := blockRatioA.Quo(blockRatioB)
	if blockRatio.GT(sdk.NewDec(1)) {
		blockRatio = sdk.NewDec(1)
	}

	totalRelativeSignaturesDistributed := sdk.NewDec(0)
	// distribute rewards to all block signers
	for _, signatureCount := range signatureCounts.Counts {

		// get the validator address from the cons address
		// if the validator is not registered, reward distribution is skipped
		// all funds are sent back to the coordinator
		validatorByConsAddr, found := k.profileKeeper.GetValidatorByConsAddress(ctx, signatureCount.ConsAddress)
		if found {
			validator, found := k.profileKeeper.GetValidator(ctx, validatorByConsAddr.ValidatorAddress)
			if !found {
				return spnerrors.Criticalf(
					"validator by consensus address not associated with a validator %s",
					validatorByConsAddr.ValidatorAddress,
				)
			}
			// calculate the total relative signature distributed for we calculate the refund
			totalRelativeSignaturesDistributed = totalRelativeSignaturesDistributed.Add(signatureCount.RelativeSignatures)

			// compute reward relative to the signature and block count
			// and update reward pool
			signatureRatio := signatureCount.RelativeSignatures.Quo(sdk.NewDec(int64(signatureCounts.BlockCount)))
			reward, err := CalculateReward(blockRatio, signatureRatio, rewardPool.Coins)
			if err != nil {
				return spnerrors.Criticalf("invalid reward: %s", err.Error())
			}
			coins, isNegative := rewardPool.Coins.SafeSub(reward)
			if isNegative {
				return spnerrors.Criticalf("negative reward pool: %s", rewardPool.Coins.String())
			}
			rewardPool.Coins = coins

			// send rewards to the address
			account, err := sdk.AccAddressFromBech32(validator.Address)
			if err != nil {
				return spnerrors.Criticalf("can't parse address %s", err.Error())
			}
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, account, reward); err != nil {
				return spnerrors.Criticalf("send rewards error: %s", err.Error())
			}
		}
	}
	// if the reward pool is closed or last reward height is reached
	// the remaining coins are refunded and reward pool is deleted
	if closeRewardPool || lastBlockHeight >= rewardPool.LastRewardHeight {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			provider,
			rewardPool.Coins); err != nil {
			return spnerrors.Criticalf("send rewards error: %s", err.Error())
		}
		k.RemoveRewardPool(ctx, launchID)
		return nil
	}

	// if the reward pool is closed or last reward height is reached
	// the remaining coins are refunded and reward pool is deleted
	if closeRewardPool || lastBlockHeight >= rewardPool.LastRewardHeight {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			provider,
			rewardPool.Coins); err != nil {
			return spnerrors.Criticalf("send rewards error: %s", err.Error())
		}
		k.RemoveRewardPool(ctx, launchID)
		return nil
	}

	// Otherwise, the refund is relative to the block ratio and the reward pool is updated
	// refundRation is blockCount.
	// This is sum of signaturesRelative values from validator to compute refund
	blockCount := sdk.NewDec(int64(signatureCounts.BlockCount))
	refundRatio := blockCount.Sub(totalRelativeSignaturesDistributed).Quo(blockCount)
	reward, err := CalculateReward(blockRatio, refundRatio, rewardPool.Coins)
	if err != nil {
		return spnerrors.Criticalf("invalid reward: %s", err.Error())
	}
	coins, isNegative := rewardPool.Coins.SafeSub(reward)
	if isNegative {
		return spnerrors.Criticalf("negative reward pool: %s", rewardPool.Coins.String())
	}
	rewardPool.Coins = coins

	// send rewards to the address
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		provider,
		rewardPool.Coins); err != nil {
		return spnerrors.Criticalf("send rewards error: %s", err.Error())
	}

	// update the current reward height for next reward
	rewardPool.CurrentRewardHeight = lastBlockHeight
	k.SetRewardPool(ctx, rewardPool)
	return nil
}

// CalculateReward calculates the reward relative to the signature and block count
func CalculateReward(blockRatio, ratio sdk.Dec, coins sdk.Coins) (sdk.Coins, error) {
	reward := sdk.NewCoins()
	for _, coin := range coins {
		refund := blockRatio.Mul(ratio).Mul(coin.Amount.ToDec())
		amount := coin.Amount.ToDec().Sub(refund)
		if amount.IsNegative() {
			return reward, fmt.Errorf("negative coin reward amount %s", amount.String())
		}
		reward = reward.Add(coin.SubAmount(refund.TruncateInt()))
	}
	return reward, nil
}
