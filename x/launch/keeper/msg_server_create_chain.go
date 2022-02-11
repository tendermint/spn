package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	spnerrors "github.com/tendermint/spn/pkg/errors"
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

	id, err := k.CreateNewChain(
		ctx,
		coordID,
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

	chain, found := k.GetChain(ctx, id)
	if !found {
		return nil, spnerrors.Criticalf("created chain not found: launch ID %d", id)
	}

	err = ctx.EventManager().EmitTypedEvent(&types.EventChainCreated{
		Chain: chain,
	})

	return &types.MsgCreateChainResponse{
		LaunchID: id,
	}, err
}
