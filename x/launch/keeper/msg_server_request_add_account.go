package keeper

import (
	"context"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) RequestAddAccount(
	goCtx context.Context,
	msg *types.MsgRequestAddAccount,
) (*types.MsgRequestAddAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrChainNotFound, msg.ChainID)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrap(types.ErrTriggeredLaunch, msg.ChainID)
	}

	content, err := codec.NewAnyWithValue(&types.GenesisAccount{
		ChainID: msg.ChainID,
		Address: msg.Address,
		Coins:   msg.Coins,
	})
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCodecNotPacked, msg.String())
	}

	var (
		requestID uint64

		approved = false
		request  = types.Request{
			ChainID:   msg.ChainID,
			Creator:   msg.Address,
			CreatedAt: ctx.BlockTime().Unix(),
			Content:   content,
		}
	)
	coordAddress := k.profileKeeper.GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if msg.Address == coordAddress {
		err := applyRequest(ctx, k.Keeper, msg.ChainID, request)
		if err != nil {
			return nil, err
		}
		approved = true
	} else {
		requestID = k.AppendRequest(ctx, request)
	}

	return &types.MsgRequestAddAccountResponse{
		RequestID:    requestID,
		AutoApproved: approved,
	}, nil
}
