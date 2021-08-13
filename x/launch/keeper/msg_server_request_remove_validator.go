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
			fmt.Sprintf("coordinator %d not found for chain %s", chain.CoordinatorID, chain.ChainID))
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

	var requestID uint64
	approved := false
	request := types.Request{
		ChainID:   msg.ChainID,
		Creator:   msg.ValidatorAddress,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   content,
	}

	if msg.Creator == coordAddress {
		err := applyRequest(ctx, k.Keeper, msg.ChainID, request)
		if err != nil {
			return nil, err
		}
		approved = true
	} else {
		requestID = k.AppendRequest(ctx, request)
	}

	return &types.MsgRequestRemoveValidatorResponse{
		RequestID:    requestID,
		AutoApproved: approved,
	}, nil
}
