package keeper

import (
	"bytes"
	"context"
	tmtypes "github.com/tendermint/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctmtypes "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func (k msgServer) CreateClient(goCtx context.Context, msg *types.MsgCreateClient) (*types.MsgCreateClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var clientState ibctmtypes.ClientState
	var consensusState ibctmtypes.ConsensusState

	// create the client from IBC keeper
	clientID, err := k.clientKeeper.CreateClient(ctx, &clientState, &consensusState)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateClientResponse{
		ClientID: clientID,
	}, nil
}

func checkValidatorSet(valSet tmtypes.ValidatorSet, consensusState ibctmtypes.ConsensusState) bool {
	valHash := consensusState.NextValidatorsHash
	return !bytes.Equal(valHash.Bytes(), valSet.Hash())
}