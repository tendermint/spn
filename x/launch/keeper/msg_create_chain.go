package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) CreateChain(goCtx context.Context, msg *types.MsgCreateChain) (*types.MsgCreateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the coordinator ID associated to the sender address
	coordID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if err != nil {
		return nil, err
	}

	// TODO check metadata len

	id, err := k.CreateNewChain(
		ctx,
		coordID,
		msg.GenesisChainID,
		msg.SourceURL,
		msg.SourceHash,
		msg.GenesisURL,
		msg.GenesisHash,
		msg.HasCampaign,
		msg.CampaignID,
		false,
		msg.Metadata,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCreateChainFail, err.Error())
	}

	// Deduct chain creation fee if set
	creationFee := k.ChainCreationFee(ctx)
	if !creationFee.Empty() {
		coordAddr, err := sdk.AccAddressFromBech32(msg.Coordinator)
		if err != nil {
			return nil, spnerrors.Criticalf("invalid coordinator bech32 address %s", err.Error())
		}
		if err = k.distrKeeper.FundCommunityPool(ctx, creationFee, coordAddr); err != nil {
			return nil, err
		}
	}

	return &types.MsgCreateChainResponse{
		LaunchID: id,
	}, nil
}
