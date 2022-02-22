package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	committypes "github.com/cosmos/ibc-go/v2/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/v2/modules/light-clients/07-tendermint/types"
	"github.com/tendermint/tendermint/light"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	"github.com/tendermint/spn/x/monitoringc/types"
)

const (
	// DefaultUnbondingPeriod is 21 days
	DefaultUnbondingPeriod = time.Hour * 24 * 21

	// DefaultTrustingPeriod must be lower than DefaultUnbondingPeriod
	DefaultTrustingPeriod = time.Hour*24*21 - 1
)

func (k msgServer) CreateClient(goCtx context.Context, msg *types.MsgCreateClient) (*types.MsgCreateClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.launchKeeper.GetChain(ctx, msg.LaunchID)
	if !found {
		return nil, sdkerrors.Wrapf(launchtypes.ErrChainNotFound, "invalid launch ID %d", msg.LaunchID)
	}

	// initialize the client state
	clientState := k.initializeClientState(chain.GenesisChainID)
	if err := clientState.Validate(); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidClientState, err.Error())
	}

	// convert validator set
	tmValidatorSet, err := msg.ValidatorSet.ToTendermintValidatorSet()
	if err != nil {
		return nil, spnerrors.Criticalf("validated validator can't be converted %s", err.Error())
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
		return nil, spnerrors.Criticalf("validated consensus state can't be converted %s", err.Error())
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
// TODO: Investigate configurable values
func (k msgServer) initializeClientState(
	chainID string,
	unbondingPeriod int64,
	revisionHeight uint64,
	) *ibctmtypes.ClientState {
	return ibctmtypes.NewClientState(
		chainID,
		ibctmtypes.NewFractionFromTm(light.DefaultTrustLevel),
		DefaultTrustingPeriod,
		unbondingPeriod,
		time.Minute*10,
		clienttypes.NewHeight(1, 1),
		committypes.GetSDKSpecs(),
		[]string{"upgrade", "upgradedIBCState"},
		true,
		true,
	)
}
