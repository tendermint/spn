package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ignterrors "github.com/ignite/modules/pkg/errors"

	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) CreateChain(goCtx context.Context, msg *types.MsgCreateChain) (*types.MsgCreateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if the metadata length is valid
	maxMetadataLength := k.MaxMetadataLength(ctx)
	if uint64(len(msg.Metadata)) > maxMetadataLength {
		return nil, sdkerrors.Wrapf(types.ErrInvalidMetadataLength,
			"metadata length %d is greater than maximum %d",
			len(msg.Metadata),
			maxMetadataLength,
		)
	}

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
		msg.InitialGenesis,
		msg.HasProject,
		msg.ProjectID,
		false,
		msg.AccountBalance,
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
			return nil, ignterrors.Criticalf("invalid coordinator bech32 address %s", err.Error())
		}
		if err = k.distrKeeper.FundCommunityPool(ctx, creationFee, coordAddr); err != nil {
			return nil, sdkerrors.Wrap(types.ErrFundCommunityPool, err.Error())
		}
	}

	return &types.MsgCreateChainResponse{
		LaunchID: id,
	}, nil
}
