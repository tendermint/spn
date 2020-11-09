package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/genesis/types"
)

// GetProposalCount
func (k Keeper) GetProposalCount(ctx sdk.Context, chainID string) int32 {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetProposalCountKey(chainID))

	// If not set we return 0
	if value == nil {
		return 0
	}
	count := types.UnmarshalProposalCount(k.cdc, value)

	return count
}

// SetProposalCount
func (k Keeper) SetProposalCount(ctx sdk.Context, chainID string, count int32) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MarshalProposalCount(k.cdc, count)
	store.Set(types.GetProposalCountKey(chainID), bz)
}

// GetProposal
func (k Keeper) GetProposal(ctx sdk.Context, chainID string, proposalID int32) (proposal types.Proposal, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetProposalKey(chainID, proposalID))
	if value == nil {
		return proposal, false
	}
	proposal = types.UnmarshalProposal(k.cdc, value)

	return proposal, true
}

// SetProposal
func (k Keeper) SetProposal(ctx sdk.Context, proposal types.Proposal) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MarshalProposal(k.cdc, proposal)
	chainID := proposal.GetProposalInformation().GetChainID()
	proposalID := proposal.GetProposalInformation().GetProposalID()
	store.Set(types.GetProposalKey(chainID, proposalID), bz)
}

// GetApprovedProposals
func (k Keeper) GetApprovedProposals(ctx sdk.Context, chainID string) types.ProposalList {
	return k.getProposalList(ctx, types.GetApprovedProposalKey(chainID))
}

// SetApprovedProposals
func (k Keeper) SetApprovedProposals(ctx sdk.Context, chainID string, list types.ProposalList) {
	k.setProposalList(ctx, list, types.GetApprovedProposalKey(chainID))
}

// GetPendingProposals
func (k Keeper) GetPendingProposals(ctx sdk.Context, chainID string) types.ProposalList {
	return k.getProposalList(ctx, types.GetPendingProposalKey(chainID))
}

// SetPendingProposals
func (k Keeper) SetPendingProposals(ctx sdk.Context, chainID string, list types.ProposalList) {
	k.setProposalList(ctx, list, types.GetPendingProposalKey(chainID))
}

// GetRejectedProposals
func (k Keeper) GetRejectedProposals(ctx sdk.Context, chainID string) types.ProposalList {
	return k.getProposalList(ctx, types.GetRejectedProposalKey(chainID))
}

// SetRejectedProposals
func (k Keeper) SetRejectedProposals(ctx sdk.Context, chainID string, list types.ProposalList) {
	k.setProposalList(ctx, list, types.GetRejectedProposalKey(chainID))
}

func (k Keeper) getProposalList(ctx sdk.Context, key []byte) types.ProposalList {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(key)

	// If not set we return an empty proposal list
	if value == nil {
		return types.ProposalList{
			ProposalIDs: []int32{},
		}
	}
	list := types.UnmarshalProposalList(k.cdc, value)

	return list
}

func (k Keeper) setProposalList(ctx sdk.Context, list types.ProposalList, key []byte) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MarshalProposalList(k.cdc, list)
	store.Set(key, bz)
}