package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) EditChain(goCtx context.Context, msg *types.MsgEditChain) (*types.MsgEditChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.LaunchID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainNotFound, "%d", msg.LaunchID)
	}

	// Check sender is the coordinator of the chain
	coordinatorID, found := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}
	if chain.CoordinatorID != coordinatorID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the chain is %d",
			chain.CoordinatorID,
		))
	}

	// Modify from provided values
	if msg.GenesisChainID != "" {
		chain.GenesisChainID = msg.GenesisChainID
	}
	if msg.SourceURL != "" {
		chain.SourceURL = msg.SourceURL
	}
	if msg.SourceHash != "" {
		chain.SourceHash = msg.SourceHash
	}
	if msg.InitialGenesis != nil {
		chain.InitialGenesis = *msg.InitialGenesis
	}

	k.SetChain(ctx, chain)

	err := ctx.EventManager().EmitTypedEvent(&types.EventChainUpdated{
		Chain: chain,
	})

	return &types.MsgEditChainResponse{}, err
}
