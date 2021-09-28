package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) InitializeMainnet(goCtx context.Context, msg *types.MsgInitializeMainnet) (*types.MsgInitializeMainnetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, found := k.GetCampaign(ctx, msg.CampaignID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCampaignNotFound, "%v", msg.CampaignID)
	}

	if campaign.MainnetInitialized {
		return nil, sdkerrors.Wrapf(types.ErrMainnetInitialized, "%v", msg.CampaignID)
	}

	// Get the coordinator ID
	coordinatorID, found := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}
	if campaign.CoordinatorID != coordinatorID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the campaign is %v",
			campaign.CoordinatorID,
		))
	}

	// Create the mainnet chain for launch
	mainnetID, err := k.launchKeeper.CreateNewChain(
		ctx,
		coordinatorID,
		msg.MainnetChainID,
		msg.SourceURL,
		msg.SourceHash,
		"",
		"",
		true,
		msg.CampaignID,
		true,
	)
	if err != nil {
		return nil, spnerrors.Criticalf("cannot create the mainnet: %v", err.Error())
	}

	// Set mainnet as initialized and save the change
	campaign.MainnetID = mainnetID
	campaign.MainnetInitialized = true
	k.SetCampaign(ctx, campaign)

	return &types.MsgInitializeMainnetResponse{
		MainnetID: mainnetID,
	}, nil
}
