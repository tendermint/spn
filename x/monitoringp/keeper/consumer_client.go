package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	committypes "github.com/cosmos/ibc-go/v2/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/v2/modules/light-clients/07-tendermint/types"
	"github.com/tendermint/tendermint/light"

	"github.com/tendermint/spn/x/monitoringp/types"
)

const (
	// DefaultUnbondingPeriod is 21 days
	DefaultUnbondingPeriod = time.Hour * 24 * 21

	// DefaultTrustingPeriod must be lower than DefaultUnbondingPeriod
	DefaultTrustingPeriod = time.Hour*24*21 - 1
)

// InitializeConsumerClient initializes the consumer IBC client and and set it in the store
func (k Keeper) InitializeConsumerClient(ctx sdk.Context) (string, error) {
	// initialize the client state
	clientState := k.initializeClientState(k.ConsumerChainID(ctx))
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
// TODO: Investigate configurable values
func (k Keeper) initializeClientState(chainID string) *ibctmtypes.ClientState {
	return ibctmtypes.NewClientState(
		chainID,
		ibctmtypes.NewFractionFromTm(light.DefaultTrustLevel),
		DefaultTrustingPeriod,
		DefaultUnbondingPeriod,
		time.Minute*10,
		clienttypes.NewHeight(1, 1),
		committypes.GetSDKSpecs(),
		[]string{"upgrade", "upgradedIBCState"},
		true,
		true,
	)
}
