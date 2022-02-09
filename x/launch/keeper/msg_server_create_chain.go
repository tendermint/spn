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
	coordByAddress, found := k.profileKeeper.GetCoordinatorByAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}

	coord, _ := k.profileKeeper.GetCoordinator(ctx, coordByAddress.CoordinatorID)
	if !coord.Active {
		return nil, sdkerrors.Wrapf(profiletypes.ErrCoordInactive,
			"coordinator %d inactive", coord.CoordinatorID)
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
