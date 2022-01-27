package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/x/monitoringp/types"
)

// LastBlockHeight returns the last block height state param
func (k Keeper) LastBlockHeight(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyLastBlockHeight, &res)
	return
}

// ConsumerConsensusState returns the consumer consensus state param
func (k Keeper) ConsumerConsensusState(ctx sdk.Context) (res spntypes.ConsensusState) {
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
		k.LastBlockHeight(ctx),
		k.ConsumerChainID(ctx),
		k.ConsumerConsensusState(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
