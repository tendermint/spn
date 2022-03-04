package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	committypes "github.com/cosmos/ibc-go/v2/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/v2/modules/light-clients/07-tendermint/types"
	"github.com/tendermint/tendermint/light"

	"github.com/tendermint/spn/pkg/chainid"
	"github.com/tendermint/spn/x/monitoringp/types"
)

// InitializeConsumerClient initializes the consumer IBC client and and set it in the store
func (k Keeper) InitializeConsumerClient(ctx sdk.Context) (string, error) {
	// initialize the client state
	clientState, err := k.initializeClientState(ctx, k.ConsumerChainID(ctx))
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrInvalidClientState, err.Error())
	}
	if err := clientState.Validate(); err != nil {
		return "", sdkerrors.Wrap(types.ErrInvalidClientState, err.Error())
	}

	// get consensus state from param
	tmConsensusState, err := k.ConsumerConsensusState(ctx).ToTendermintConsensusState()
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrInvalidConsensusState, err.Error())
	}

	// create IBC client for consumer
	clientID, err := k.clientKeeper.CreateClient(ctx, clientState, &tmConsensusState)
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrClientCreationFailure, err.Error())
	}

	// register the IBC client
	k.SetConsumerClientID(ctx, types.ConsumerClientID{
		ClientID: clientID,
	})

	return clientID, nil
}

// initializeClientState initializes the client state provided for the IBC client
func (k Keeper) initializeClientState(ctx sdk.Context, chainID string) (*ibctmtypes.ClientState, error) {
	_, revisionNumber, err := chainid.ParseGenesisChainID(chainID)
	if err != nil {
		return nil, err
	}

	unbondingPeriod := k.ConsumerUnbondingPeriod(ctx)
	revisionHeight := k.ConsumerRevisionHeight(ctx)

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
