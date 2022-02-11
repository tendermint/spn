package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) RequestAddAccount(
	goCtx context.Context,
	msg *types.MsgRequestAddAccount,
) (*types.MsgRequestAddAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.LaunchID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainNotFound, "%d", msg.LaunchID)
	}

	if chain.IsMainnet {
		return nil, sdkerrors.Wrapf(
			types.ErrAddMainnetAccount,
			"the chain %d is a mainnet",
			msg.LaunchID,
		)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%d", msg.LaunchID)
	}

	coordAddress, found := k.profileKeeper.GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainInactive,
			"the chain %d coordinator has been deleted", chain.LaunchID)
	}

	content := types.NewGenesisAccount(msg.LaunchID, msg.Address, msg.Coins)
	request := types.Request{
		LaunchID:  msg.LaunchID,
		Creator:   msg.Creator,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   content,
	}

	var requestID uint64
	var err error
	approved := false
	if msg.Creator == coordAddress {
		err = ApplyRequest(ctx, k.Keeper, msg.LaunchID, request)
		if err != nil {
			return nil, err
		}
		approved = true
		err = ctx.EventManager().EmitTypedEvent(&types.EventGenesisAccountAdded{
			GenesisAccount: *content.GetGenesisAccount(),
		})
	} else {
		requestID = k.AppendRequest(ctx, request)
		err = ctx.EventManager().EmitTypedEvent(&types.EventRequestCreated{
			Request: request,
		})
	}

	return &types.MsgRequestAddAccountResponse{
		RequestID:    requestID,
		AutoApproved: approved,
	}, err
}
