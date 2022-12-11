package keeper

import (
	"context"
	"errors"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/pkg/errors"

	"github.com/tendermint/spn/x/project/types"
)

func (k msgServer) CreateProject(goCtx context.Context, msg *types.MsgCreateProject) (*types.MsgCreateProjectResponse, error) {
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

	// Validate provided totalSupply
	totalSupplyRange := k.TotalSupplyRange(ctx)
	if err := types.ValidateTotalSupply(msg.TotalSupply, totalSupplyRange); err != nil {
		if errors.Is(err, types.ErrInvalidSupplyRange) {
			return nil, ignterrors.Critical(err.Error())
		}

		return nil, sdkerrors.Wrap(types.ErrInvalidTotalSupply, err.Error())
	}

	// Deduct project creation fee if set
	creationFee := k.ProjectCreationFee(ctx)
	if !creationFee.Empty() {
		coordAddr, err := sdk.AccAddressFromBech32(msg.Coordinator)
		if err != nil {
			return nil, ignterrors.Criticalf("invalid coordinator bech32 address %s", err.Error())
		}
		if err = k.distrKeeper.FundCommunityPool(ctx, creationFee, coordAddr); err != nil {
			return nil, sdkerrors.Wrap(types.ErrFundCommunityPool, err.Error())
		}
	}

	// Append the new project
	project := types.NewProject(
		0,
		msg.ProjectName,
		coordID,
		msg.TotalSupply,
		msg.Metadata,
		ctx.BlockTime().Unix(),
	)
	projectID := k.AppendProject(ctx, project)

	// Initialize the list of project chains
	k.SetProjectChains(ctx, types.ProjectChains{
		ProjectID: projectID,
		Chains:     []uint64{},
	})

	err = ctx.EventManager().EmitTypedEvent(&types.EventProjectCreated{
		ProjectID:         projectID,
		CoordinatorAddress: msg.Coordinator,
		CoordinatorID:      coordID,
	})

	return &types.MsgCreateProjectResponse{ProjectID: projectID}, err
}
