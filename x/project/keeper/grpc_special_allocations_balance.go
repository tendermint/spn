package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/project/types"
)

func (k Keeper) SpecialAllocationsBalance(
	goCtx context.Context,
	req *types.QuerySpecialAllocationsBalanceRequest,
) (*types.QuerySpecialAllocationsBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// get the project
	totalShareNumber := k.GetTotalShares(ctx)
	project, found := k.GetProject(ctx, req.ProjectID)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	// calculate special allocations balance from total supply
	genesisDistribution, err := project.SpecialAllocations.GenesisDistribution.CoinsFromTotalSupply(
		project.TotalSupply,
		totalShareNumber,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "genesis distribution can't be calculated: %s", err.Error())
	}
	claimableAirdrop, err := project.SpecialAllocations.ClaimableAirdrop.CoinsFromTotalSupply(
		project.TotalSupply,
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
