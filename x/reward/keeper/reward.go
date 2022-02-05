package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/x/reward/types"
)

// DistributeRewards distribute rewards based on the monitoring packet
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

	// only the monitored blocks relative to last reward height are rewarded
	blockRatio := (lastBlockHeight - rewardPool.CurrentRewardHeight) / (rewardPool.LastRewardHeight - rewardPool.CurrentRewardHeight)
	if blockRatio > 1 {
		blockRatio = 1
	}

	// distribute rewards to all block signers
	signatureCounts := rewardPool.SignatureCounts
	for _, signatureCount := range signatureCounts {
		//totalSignaturesRelative += signatureCount.RelativeSignatures

		// get the validator address from the cons address
		// if the validator is not registered, reward distribution is skipped
		// all funds are sent back to the coordinator
		validator, found := k.profileKeeper.GetValidator(ctx, signatureCount.ConsAddress)
		if found {
			// compute reward relative to the signature and block count
			// and update reward pool
			//signatureRatio = signatureCount.RelativeSignatures/signatureCounts.BlockCount
			//reward = floor(blockRatio*signatureRatio*rewardPool.Coins)
			//rewardPool.Coins = rewardPool.Coins.SubstractCoins(reward)
			//if negative(rewardPool.Coins) {
			//	panic
			//}

			// send rewards to the address
			account, err := sdk.AccAddressFromBech32(validator.Address)
			if err != nil {
				return spnerrors.Criticalf("can't parse address %s", err.Error())
			}
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, account, reward); err != nil {
				return spnerrors.Criticalf("can't send coins from module %s", err.Error())
			}
		}
	}
	// if the reward pool is closed or last reward height is reached
	// the remaining coins are refunded and reward pool is deleted
	//if closeRewardPool || lastBlockHeight >= rewardPool.LastRewardHeight
	//if bankKeeper.Transfer(ModuleAccount, rewardPool.Provider, rewardPool.Coins) != nil
	//	panic
	//delete(RewardPools, launchID)
	//
	//else
	// otherwise the refund is relative to the block ratio and the reward pool is updated
	// refundRation is blockCount
	//
	// this is sum of signaturesRelative values from validator to compute refund
	//var totalSignaturesRelative float
	//for signatureCount in signatureCounts
	//totalSignaturesRelative += signatureCount.RelativeSignatures
	//
	//refundRatio = (signatureCounts.BlockCount-totalSignaturesRelative)/signatureCounts.BlockCount
	//refund = floor(blockRatio*refundRatio*rewardPool.Coins)
	//rewardPool.Coins = rewardPool.Coins.SubstractCoins(reward)
	//if negative(rewardPool.Coins)
	//	panic
	//
	// send rewards to the address
	//if bankKeeper.Transfer(ModuleAccount, address, reward) != nil
	//	panic
	//
	// update the current reward height for next reward
	//rewardPool.CurrentRewardHeight = lastBlockHeight
	//save(RewardPools, launchID, rewardPool)

	return nil
}
