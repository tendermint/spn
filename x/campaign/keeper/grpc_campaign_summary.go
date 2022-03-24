package keeper

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CampaignSummaries(goCtx context.Context, req *types.QueryCampaignSummariesRequest) (*types.QueryCampaignSummariesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var campaignSummaries []types.CampaignSummary
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	campaignStore := prefix.NewStore(store, types.KeyPrefix(types.CampaignKey))

	pageRes, err := query.Paginate(campaignStore, req.Pagination, func(key []byte, value []byte) error {
		var campaign types.Campaign
		if err := k.cdc.Unmarshal(value, &campaign); err != nil {
			return err
		}

		campaignSummary, err := k.GetCampaignSummary(ctx, campaign)
		if err != nil {
			return err
		}

		campaignSummaries = append(campaignSummaries, campaignSummary)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryCampaignSummariesResponse{
		CampaignSummaries: campaignSummaries,
		Pagination: pageRes,
	}, nil
}

// GetCampaignSummary returns the campaign with summary attached to it like most recent chain and rewards attached to it
func (k Keeper) GetCampaignSummary(ctx sdk.Context, campaign types.Campaign) (cs types.CampaignSummary, err error) {
	cs.Campaign = campaign

	campaignChains, found := k.GetCampaignChains(ctx, campaign.CampaignID)
	if !found {
		return cs, fmt.Errorf("chain list not found for existing campaign %d", campaign.CampaignID)
	}

	// retrieve information about most recent chain
	chainCount := len(campaignChains.Chains)
	if chainCount > 0 {
		mostRecentChainID := campaignChains.Chains[chainCount-1]

		cs.HasMostRecentChain = true

		chain, found := k.launchKeeper.GetChain(ctx, mostRecentChainID)
		if !found {
			return cs, fmt.Errorf("chain not found for campaing chain %d", mostRecentChainID)
		}

		cs.MostRecentChain = types.MostRecentChain{
			LaunchID: mostRecentChainID,
			LaunchTriggered: chain.LaunchTriggered,
		}

		// retrieve information about rewards
		rewardPool, found := k.rewardKeeper.GetRewardPool(ctx, mostRecentChainID)
		if found {
			cs.Incentivized = true
			cs.Rewards = rewardPool.InitialCoins
			cs.RewardsDistributed = rewardPool.Closed
		}
	}

	return cs, nil
}