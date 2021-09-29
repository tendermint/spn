package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
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
		false,
		0,
		false,
	)
	if err != nil {
		return nil, spnerrors.Criticalf("cannot create the chain: %v", err.Error())
	}

	if msg.CampaignID > 0 {
		campaign, found := k.campaignKeeper.GetCampaignChains(ctx, msg.CampaignID)
		if !found {
			campaign = campaigntypes.CampaignChains{
				CampaignID: msg.CampaignID,
				Chains:     []uint64{},
			}
		}
		campaign.Chains = append(campaign.Chains, id)
		k.campaignKeeper.SetCampaignChains(ctx, campaign)
	}

	return &types.MsgCreateChainResponse{
		Id: id,
	}, nil
}
