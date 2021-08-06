package keeper

import (
	"context"
	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) RequestAddValidator(goCtx context.Context, msg *types.MsgRequestAddValidator) (*types.MsgRequestAddValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrChainNotFound, msg.ChainID)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrap(types.ErrTriggeredLaunch, msg.ChainID)
	}

	content, err := codec.NewAnyWithValue(&types.GenesisValidator{
		ChainID: msg.ChainID,
		Address: msg.ValAddress,
		GenTx:   msg.GenTx,
		ConsPubKey: msg.ConsPubKey,
		SelfDelegation: msg.SelfDelegation,
		Peer: msg.Peer,
	})
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCodecNotPacked, msg.String())
	}

	requestID := k.AppendRequest(ctx, types.Request{
		ChainID:   msg.ChainID,
		Creator:   msg.ValAddress,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   content,
	})

	return &types.MsgRequestAddValidatorResponse{
		RequestID: requestID,
	}, nil
}
