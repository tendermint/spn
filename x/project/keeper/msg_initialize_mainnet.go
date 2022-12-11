package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/pkg/errors"

	launchtypes "github.com/tendermint/spn/x/launch/types"

	"github.com/tendermint/spn/x/project/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) InitializeMainnet(goCtx context.Context, msg *types.MsgInitializeMainnet) (*types.MsgInitializeMainnetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	project, found := k.GetProject(ctx, msg.ProjectID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrProjectNotFound, "%d", msg.ProjectID)
	}

	if project.MainnetInitialized {
		return nil, sdkerrors.Wrapf(types.ErrMainnetInitialized, "%d", msg.ProjectID)
	}

	if project.TotalSupply.Empty() {
		return nil, sdkerrors.Wrap(types.ErrInvalidTotalSupply, "total supply is empty")
	}

	// Get the coordinator ID associated to the sender address
	coordID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if err != nil {
		return nil, err
	}

	if project.CoordinatorID != coordID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the project is %d",
			project.CoordinatorID,
		))
	}

	initialGenesis := launchtypes.NewDefaultInitialGenesis()

	// Create the mainnet chain for launch
	mainnetID, err := k.launchKeeper.CreateNewChain(
		ctx,
		coordID,
		msg.MainnetChainID,
		msg.SourceURL,
		msg.SourceHash,
		initialGenesis,
		true,
		msg.ProjectID,
		true,
		sdk.NewCoins(), // no enforced default for mainnet
		[]byte{},
	)
	if err != nil {
		return nil, ignterrors.Criticalf("cannot create the mainnet: %s", err.Error())
	}

	// Set mainnet as initialized and save the change
	project.MainnetID = mainnetID
	project.MainnetInitialized = true
	k.SetProject(ctx, project)

	err = ctx.EventManager().EmitTypedEvent(&types.EventProjectMainnetInitialized{
		ProjectID:         project.ProjectID,
		CoordinatorAddress: msg.Coordinator,
		MainnetID:          project.MainnetID,
	})

	return &types.MsgInitializeMainnetResponse{
		MainnetID: mainnetID,
	}, err
}
