package keeper

import (
	"context"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) RequestAddVestedAccount(
	goCtx context.Context,
	msg *types.MsgRequestAddVestedAccount,
) (*types.MsgRequestAddVestedAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrChainNotFound, msg.ChainID)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrap(types.ErrTriggeredLaunch, msg.ChainID)
	}

	content, err := codec.NewAnyWithValue(&types.VestedAccount{
		ChainID:         msg.ChainID,
		Address:         msg.Address,
		StartingBalance: msg.Coins,
		VestingOptions:  msg.Options,
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
	return &types.MsgRequestAddVestedAccountResponse{
		RequestID: requestID,
	}, nil
}
