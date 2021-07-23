package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) CreateChain(goCtx context.Context, msg *types.MsgCreateChain) (*types.MsgCreateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the coordinator ID associated to the sender address
	coordId, found := k.profileKeeper.CoordinatorIdFromAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrCoordinatorNotExist, msg.Coordinator)
	}

	// Compute the chain id
	chainNameCount, found := k.GetChainNameCount(ctx, msg.ChainName)
	if !found {
		chainNameCount = types.ChainNameCount{
			msg.ChainName,
			0,
		}
	}
	chainID := types.ChainIdFromChainName(msg.ChainName, chainNameCount.Count)
	chainNameCount.Count++

	// chainID must always be unique by design
	// if it already exists then something is wrong in the protocol
	_, found = k.GetChain(ctx, chainID)
	if found {
		panic(fmt.Sprintf("chain id %s already exists", chainID))
	}

	// Initialize the chain
	chain := types.Chain{
		ChainID: chainID,
		CoordinatorID: coordId,
		CreatedAt: ctx.BlockTime().Unix(),
		SourceURL: msg.SourceURL,
		SourceHash: msg.SourceHash,
		LaunchTriggered: false,
		LaunchTimestamp: 0,
	}

	// Store values
	k.SetChain(ctx, chain)
	k.SetChainNameCount(ctx, chainNameCount)

	return &types.MsgCreateChainResponse{
		ChainID: chainID,
	}, nil
}
