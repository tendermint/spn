package keeper

import (
	"context"
	sdkerrors "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/pkg/errors"
	launchtypes "github.com/tendermint/spn/x/launch/types"

	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) InitializeMainnet(goCtx context.Context, msg *types.MsgInitializeMainnet) (*types.MsgInitializeMainnetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, found := k.GetCampaign(ctx, msg.CampaignID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCampaignNotFound, "%d", msg.CampaignID)
	}

	if campaign.MainnetInitialized {
		return nil, sdkerrors.Wrapf(types.ErrMainnetInitialized, "%d", msg.CampaignID)
	}

	if campaign.TotalSupply.Empty() {
		return nil, sdkerrors.Wrap(types.ErrInvalidTotalSupply, "total supply is empty")
	}

	// Get the coordinator ID associated to the sender address
	coordID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if err != nil {
		return nil, err
	}

	if campaign.CoordinatorID != coordID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the campaign is %d",
			campaign.CoordinatorID,
		))
	}

	initialGenesis := launchtypes.NewDefaultInitialGenesis()

	// Create the mainnet chain for launch
	mainnetID, err := k.launchKeeper.CreateNewChain(
		ctx,
		coordID,
		msg.MainnetChainID,
		msg.SourceURL,
		msg.SourceHash,
		&initialGenesis,
		true,
		msg.CampaignID,
		true,
		sdk.NewCoins(), // no enforced default for mainnet
		[]byte{},
	)
	if err != nil {
		return nil, ignterrors.Criticalf("cannot create the mainnet: %s", err.Error())
	}

	// Set mainnet as initialized and save the change
	campaign.MainnetID = mainnetID
	campaign.MainnetInitialized = true
	k.SetCampaign(ctx, campaign)

	err = ctx.EventManager().EmitTypedEvent(&types.EventCampaignMainnetInitialized{
		CampaignID:         campaign.CampaignID,
		CoordinatorAddress: msg.Coordinator,
		MainnetID:          campaign.MainnetID,
	})

	return &types.MsgInitializeMainnetResponse{
		MainnetID: mainnetID,
	}, err
}
