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

	// Get the coordinator ID associated to the sender address
	coordByAddress, found := k.profileKeeper.GetCoordinatorByAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}

	coord, found := k.profileKeeper.GetCoordinator(ctx, coordByAddress.CoordinatorID)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}

	if !coord.Active {
		return nil, sdkerrors.Wrapf(profiletypes.ErrCoordInactive,
			"coordinator %d inactive", coord.CoordinatorID)
	}

	if chain.CoordinatorID != coord.CoordinatorID {
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

	return &types.MsgEditChainResponse{}, nil
}
