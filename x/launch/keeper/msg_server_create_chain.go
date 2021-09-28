package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	// Initialize the chain
	chain := types.Chain{
		CoordinatorID:   coordID,
		GenesisChainID:  msg.GenesisChainID,
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
