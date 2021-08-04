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

	requestID := k.AppendRequest(ctx, types.Request{
		ChainID:   msg.ChainID,
		Creator:   msg.Address,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   content,
	})

	return &types.MsgRequestAddAccountResponse{
		RequestID: requestID,
	}, nil
}
