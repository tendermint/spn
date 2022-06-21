package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/x/claim/types"
)

func (k msgServer) ClaimInitial(goCtx context.Context, msg *types.MsgClaimInitial) (*types.MsgClaimInitialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// retrieve initial claim information
	initialClaim, found := k.GetInitialClaim(ctx)
	if !found {
		return nil, types.ErrInitialClaimNotFound
	}
	if !initialClaim.Enabled {
		return nil, types.ErrInitialClaimNotEnabled
	}

	if err := k.CompleteMission(ctx, initialClaim.MissionID, msg.Claimer); err != nil {
		return nil, sdkerrors.Wrap(types.ErrMissionCompleteFailure, err.Error())
	}

	return &types.MsgClaimInitialResponse{}, nil
}
