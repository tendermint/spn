package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) CreateChain(goCtx context.Context, msg *types.MsgCreateChain) (*types.MsgCreateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the coordinator ID associated to the sender address
	coord, err := k.profileKeeper.GetActiveCoordinatorByAddress(ctx, msg.Coordinator)
	if err != nil {
		return nil, err
	}

	id, err := k.CreateNewChain(
		ctx,
		coord.CoordinatorID,
		msg.GenesisChainID,
		msg.SourceURL,
		msg.SourceHash,
		msg.GenesisURL,
		msg.GenesisHash,
		msg.HasCampaign,
		msg.CampaignID,
		false,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCreateChainFail, err.Error())
	}

	return &types.MsgCreateChainResponse{
		LaunchID: id,
	}, nil
}
