package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/campaign/types"
)

func (k Keeper) SpecialAllocationsBalance(
	goCtx context.Context,
	req *types.QuerySpecialAllocationsBalanceRequest,
) (*types.QuerySpecialAllocationsBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// get the campaign
	totalShareNumber := k.GetTotalShares(ctx)
	campaign, found := k.GetCampaign(ctx, req.CampaignID)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	// calculate special allocations balance from total supply
	genesisDistribution, err := campaign.SpecialAllocations.GenesisDistribution.CoinsFromTotalSupply(
		campaign.TotalSupply,
		totalShareNumber,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "genesis distribution can't be calculated: %s", err.Error())
	}
	claimableAirdrop, err := campaign.SpecialAllocations.ClaimableAirdrop.CoinsFromTotalSupply(
		campaign.TotalSupply,
		totalShareNumber,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "claimable airdrop can't be calculated: %s", err.Error())
	}

	return &types.QuerySpecialAllocationsBalanceResponse{
		GenesisDistribution: genesisDistribution,
		ClaimableAirdrop:    claimableAirdrop,
	}, nil
}
