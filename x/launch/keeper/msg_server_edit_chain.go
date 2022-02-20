package keeper

import (
	"context"
	"fmt"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"

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
	coordID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if err != nil {
		return nil, err
	}

	if chain.CoordinatorID != coordID {
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

	if len(msg.Metadata) > 0 {
		chain.Metadata = msg.Metadata
	}

	if msg.HasCampaign {
		// check if chain already has id associated
		if chain.HasCampaign {
			return nil, sdkerrors.Wrapf(types.ErrChainCampaignAlreadyExist,
				"campaign with id %d already associated with chain %d",
				chain.CampaignID,
				chain.LaunchID,
			)
		}

		// check if chain coordinator is campaign coordinator
		campaign, found := k.campaignKeeper.GetCampaign(ctx, msg.CampaignID)
		if !found {
			return nil, sdkerrors.Wrapf(campaigntypes.ErrCampaignNotFound, "campaign with id %d not found", msg.CampaignID)
		}

		if campaign.CoordinatorID != chain.CoordinatorID {
			return nil, sdkerrors.Wrapf(profiletypes.ErrCoordInvalid,
				"coordinator of the campaign is %d, chain coordinator is %d",
				campaign.CoordinatorID,
				chain.CoordinatorID,
			)
		}

		chain.CampaignID = msg.CampaignID
		chain.HasCampaign = true
	}

	k.SetChain(ctx, chain)

	return &types.MsgEditChainResponse{}, nil
}
