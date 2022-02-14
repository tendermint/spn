package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/reward/types"
)

func (k msgServer) SetRewards(goCtx context.Context, msg *types.MsgSetRewards) (*types.MsgSetRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// determine if the chain exists
	chain, found := k.launchKeeper.GetChain(ctx, msg.LaunchID)
	if !found {
		return nil, sdkerrors.Wrapf(launchtypes.ErrChainNotFound, "%d", msg.LaunchID)
	}
	// check coordinator
	coordByAddr, found := k.profileKeeper.GetCoordinatorByAddress(ctx, msg.Provider)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Provider)
	}
	if coordByAddr.CoordinatorID != chain.CoordinatorID {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCoordinatorID, "%d", coordByAddr.CoordinatorID)
	}
	// reward can't be changed once launch is triggered
	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(launchtypes.ErrTriggeredLaunch, "%d", msg.LaunchID)
	}

	provider, err := sdk.AccAddressFromBech32(msg.Provider)
	if err != nil {
		return nil, spnerrors.Criticalf("can't parse provider address %s", err.Error())
	}

	rewardPool, found := k.GetRewardPool(ctx, msg.LaunchID)
	if !found {
		// create the reward pool and transfer tokens if not created yet
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, provider, types.ModuleName, msg.Coins); err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
		}
		rewardPool = types.NewRewardPool(msg.LaunchID, 0)
	} else if err := SetBalance(ctx, k.bankKeeper, provider, msg.Coins, rewardPool.Coins); err != nil {
		return nil, err
	}
	rewardPool.Coins = msg.Coins
	rewardPool.Provider = msg.Provider
	rewardPool.LastRewardHeight = msg.LastRewardHeight
	k.SetRewardPool(ctx, rewardPool)

	return &types.MsgSetRewardsResponse{}, nil
}

// SetBalance set balance to Coins on the module account
// calling the transfer depending on the balance difference
func SetBalance(
	ctx sdk.Context,
	bankKeeper types.BankKeeper,
	provider sdk.AccAddress,
	coins sdk.Coins,
	poolCoins sdk.Coins,
) error {
	if coins.DenomsSubsetOf(poolCoins) && coins.IsEqual(poolCoins) {
		return nil
	}
	if poolCoins != nil && !poolCoins.IsZero() {
		if err := bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			provider,
			poolCoins,
		); err != nil {
			return spnerrors.Critical(err.Error())
		}
	}
	if coins != nil && !coins.IsZero() {
		if err := bankKeeper.SendCoinsFromAccountToModule(
			ctx,
			provider,
			types.ModuleName,
			coins,
		); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
		}
	}
	return nil
}