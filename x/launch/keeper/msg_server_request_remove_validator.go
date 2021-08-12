package keeper

import (
	"context"
	"fmt"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) RequestRemoveValidator(
	goCtx context.Context,
	msg *types.MsgRequestRemoveValidator,
) (*types.MsgRequestRemoveValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrChainNotFound, msg.ChainID)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrap(types.ErrTriggeredLaunch, msg.ChainID)
	}

	coordAddress, found := k.profileKeeper.GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if !found {
		return nil, spnerrors.Critical(
			fmt.Sprintf("Coordinator id not found: %d", chain.CoordinatorID))
	}
	if msg.Creator != msg.ValidatorAddress && msg.Creator != coordAddress {
		return nil, sdkerrors.Wrap(types.ErrNoAddressPermission, msg.Creator)
	}

	content, err := codec.NewAnyWithValue(&types.ValidatorRemoval{
		ValAddress: msg.ValidatorAddress,
	})
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCodecNotPacked, msg.String())
	}

	requestID := k.AppendRequest(ctx, types.Request{
		ChainID:   msg.ChainID,
		Creator:   msg.ValidatorAddress,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   content,
	})

	return &types.MsgRequestRemoveValidatorResponse{
		RequestID: requestID,
	}, nil
}
