package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	committypes "github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	"github.com/tendermint/tendermint/light"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctmtypes "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
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

	// create the client from IBC keeper
	clientID, err := k.clientKeeper.CreateClient(ctx, clientState, &msg.ConsensusState)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrClientCreationFailure, err.Error())
	}

	// add the client ID as verified client ID
	k.SetVerifiedClientID(ctx, types.VerifiedClientID{
		ClientID: clientID,
		LaunchID: msg.LaunchID,
	})

	return &types.MsgCreateClientResponse{
		ClientID: clientID,
	}, nil
}

// initializeClientState initializes the client state provided for the IBC client
// TODO: Investigate configurable values
func (k msgServer) initializeClientState(chainID string) *ibctmtypes.ClientState {
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
