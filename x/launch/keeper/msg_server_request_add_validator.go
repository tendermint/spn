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

func (k msgServer) RequestAddValidator(
	goCtx context.Context,
	msg *types.MsgRequestAddValidator,
) (*types.MsgRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrChainNotFound, msg.ChainID)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrap(types.ErrTriggeredLaunch, msg.ChainID)
	}

	content, err := codec.NewAnyWithValue(&types.GenesisValidator{
		ChainID:        msg.ChainID,
		Address:        msg.ValAddress,
		GenTx:          msg.GenTx,
		ConsPubKey:     msg.ConsPubKey,
		SelfDelegation: msg.SelfDelegation,
		Peer:           msg.Peer,
	})
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCodecNotPacked, msg.String())
	}

	coordAddress, found := k.profileKeeper.GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if !found {
		return nil, spnerrors.Critical(
			fmt.Sprintf("coordinator %d not found for chain %s", chain.CoordinatorID, chain.ChainID))
	}

	request := types.Request{
		ChainID:   msg.ChainID,
		Creator:   msg.ValAddress,
		CreatedAt: ctx.BlockTime().Unix(),
		Content:   content,
	}

	var requestID uint64
	approved := false
	if msg.ValAddress == coordAddress {
		err := applyRequest(ctx, k.Keeper, msg.ChainID, request)
		if err != nil {
			return nil, err
		}
		approved = true
	} else {
		requestID = k.AppendRequest(ctx, request)
	}

	return &types.MsgRequestResponse{
		RequestID:    requestID,
		AutoApproved: approved,
	}, nil
}
