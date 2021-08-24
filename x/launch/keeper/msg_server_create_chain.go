package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) CreateChain(goCtx context.Context, msg *types.MsgCreateChain) (*types.MsgCreateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the coordinator ID associated to the sender address
	coordID, found := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}

	// Initialize the chain
	chain := types.Chain{
		CoordinatorID:   coordID,
		CreatedAt:       ctx.BlockTime().Unix(),
		SourceURL:       msg.SourceURL,
		SourceHash:      msg.SourceHash,
		LaunchTriggered: false,
		LaunchTimestamp: 0,
	}

	// Initialize initial genesis
	if msg.GenesisURL == "" {
		chain.InitialGenesis = types.NewDefaultInitialGenesis()
	} else {
		chain.InitialGenesis = types.NewGenesisURL(msg.GenesisURL, msg.GenesisHash)
	}

	id := k.AppendChain(ctx, chain)

	return &types.MsgCreateChainResponse{
		Id: id,
	}, nil
}
