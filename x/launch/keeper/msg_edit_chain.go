package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	campaigntypes "github.com/tendermint/spn/x/project/types"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) EditChain(goCtx context.Context, msg *types.MsgEditChain) (*types.MsgEditChainResponse, error) {
	var (
		err error
		ctx = sdk.UnwrapSDKContext(goCtx)
	)

	// check if the metadata length is valid
	maxMetadataLength := k.MaxMetadataLength(ctx)
	if uint64(len(msg.Metadata)) > maxMetadataLength {
		return nil, sdkerrors.Wrapf(types.ErrInvalidMetadataLength,
			"metadata length %d is greater than maximum %d",
			len(msg.Metadata),
			maxMetadataLength,
		)
	}

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

	if len(msg.Metadata) > 0 {
		chain.Metadata = msg.Metadata
	}

	if msg.SetCampaignID {
		// check if chain already has id associated
		if chain.HasCampaign {
			return nil, sdkerrors.Wrapf(types.ErrChainHasCampaign,
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

		err = k.campaignKeeper.AddChainToCampaign(ctx, chain.CampaignID, chain.LaunchID)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrAddChainToCampaign, err.Error())
		}
	}

	k.SetChain(ctx, chain)
	return &types.MsgEditChainResponse{}, err
}
