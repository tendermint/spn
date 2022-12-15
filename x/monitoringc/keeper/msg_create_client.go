package keeper

import (
	"context"
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	committypes "github.com/cosmos/ibc-go/v6/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"
	ignterrors "github.com/ignite/modules/pkg/errors"
	"github.com/tendermint/tendermint/light"

	"github.com/tendermint/spn/pkg/chainid"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func (k msgServer) CreateClient(goCtx context.Context, msg *types.MsgCreateClient) (*types.MsgCreateClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.launchKeeper.GetChain(ctx, msg.LaunchID)
	if !found {
		return nil, sdkerrors.Wrapf(launchtypes.ErrChainNotFound, "invalid launch ID %d", msg.LaunchID)
	}

	// initialize the client state
	clientState, err := k.initializeClientState(
		chain.GenesisChainID,
		msg.UnbondingPeriod,
		msg.RevisionHeight,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidClientState, err.Error())
	}
	if err := clientState.Validate(); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidClientState, err.Error())
	}

	// convert validator set
	tmValidatorSet, err := msg.ValidatorSet.ToTendermintValidatorSet()
	if err != nil {
		return nil, ignterrors.Criticalf("validated validator can't be converted %s", err.Error())
	}

	// verify the validator set
	err = k.launchKeeper.CheckValidatorSet(
		ctx,
		msg.LaunchID,
		chain.GenesisChainID,
		tmValidatorSet,
	)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidValidatorSet, "validator set can't be verified %s", err.Error())
	}

	// create the client from IBC keeper
	tmConsensusState, err := msg.ConsensusState.ToTendermintConsensusState()
	if err != nil {
		return nil, ignterrors.Criticalf("validated consensus state can't be converted %s", err.Error())
	}
	clientID, err := k.clientKeeper.CreateClient(ctx, clientState, &tmConsensusState)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrClientCreationFailure, err.Error())
	}

	// add the client ID as verified client ID
	k.AddVerifiedClientID(ctx, msg.LaunchID, clientID)
	k.SetLaunchIDFromVerifiedClientID(ctx, types.LaunchIDFromVerifiedClientID{
		LaunchID: msg.LaunchID,
		ClientID: clientID,
	})

	return &types.MsgCreateClientResponse{
		ClientID: clientID,
	}, nil
}

// initializeClientState initializes the client state provided for the IBC client
func (k msgServer) initializeClientState(
	chainID string,
	unbondingPeriod int64,
	revisionHeight uint64,
) (*ibctmtypes.ClientState, error) {
	_, revisionNumber, err := chainid.ParseGenesisChainID(chainID)
	if err != nil {
		return nil, err
	}

	return ibctmtypes.NewClientState(
		chainID,
		ibctmtypes.NewFractionFromTm(light.DefaultTrustLevel),
		time.Second*time.Duration(unbondingPeriod)-1,
		time.Second*time.Duration(unbondingPeriod),
		time.Minute*10,
		clienttypes.NewHeight(revisionNumber, revisionHeight),
		committypes.GetSDKSpecs(),
		[]string{"upgrade", "upgradedIBCState"},
		true,
		true,
	), nil
}
