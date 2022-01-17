package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/pkg/ibctypes"
	"github.com/tendermint/spn/x/monitoringp/types"
)

// ConsumerConsensusState returns the consumer consensus state param
func (k Keeper) ConsumerConsensusState(ctx sdk.Context) (res ibctypes.ConsensusState) {
	k.paramstore.Get(ctx, types.KeyConsumerConsensusState, &res)
	return
}

// ConsumerChainID returns the consumer chain ID param
func (k Keeper) ConsumerChainID(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyConsumerChainID, &res)
	return
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.ConsumerChainID(ctx),
		k.ConsumerConsensusState(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
