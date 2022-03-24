package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/campaign/types"
)

func (k Keeper) CampaignSummary(goCtx context.Context, req *types.QueryCampaignSummaryRequest) (*types.QueryCampaignSummaryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, found := k.GetCampaign(ctx, req.CampaignID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}
	campaignSummary, err := k.GetCampaignSummary(ctx, campaign)

	return &types.QueryCampaignSummaryResponse{
		CampaignSummary: campaignSummary,
	}, err
}

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
		Pagination:        pageRes,
	}, nil
}

// GetCampaignSummary returns the campaign with summary attached to it like most recent chain and rewards attached to it
// TODO: add tests https://github.com/tendermint/spn/issues/650
func (k Keeper) GetCampaignSummary(ctx sdk.Context, campaign types.Campaign) (cs types.CampaignSummary, err error) {
	cs.Campaign = campaign

	campaignChains, found := k.GetCampaignChains(ctx, campaign.CampaignID)
	if !found {
		return cs, fmt.Errorf("chain list not found for existing campaign %d", campaign.CampaignID)
	}

	// retrieve information about most recent chain
	chainCount := len(campaignChains.Chains)
	if chainCount > 0 {
		mostRecentLaunchID := campaignChains.Chains[chainCount-1]

		cs.HasMostRecentChain = true

		chain, found := k.launchKeeper.GetChain(ctx, mostRecentLaunchID)
		if !found {
			return cs, fmt.Errorf("chain not found for campaign chain %d", mostRecentLaunchID)
		}

		cs.MostRecentChain = types.MostRecentChain{
			LaunchID:        mostRecentLaunchID,
			LaunchTriggered: chain.LaunchTriggered,
			SourceURL:       chain.SourceURL,
			SourceHash:      chain.SourceHash,
			RequestNb:       k.launchKeeper.GetRequestCount(ctx, mostRecentLaunchID),
			ValidatorNb:     k.launchKeeper.GetGenesisValidatorCount(ctx, mostRecentLaunchID),
		}

		// retrieve information about rewards
		rewardPool, found := k.rewardKeeper.GetRewardPool(ctx, mostRecentLaunchID)
		if found {
			cs.Incentivized = true
			cs.Rewards = rewardPool.InitialCoins
			cs.RewardsDistributed = rewardPool.Closed
		}
	}

	return cs, nil
}
