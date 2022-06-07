package keeper

import (
	"encoding/binary"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/claim/types"
)

// GetMissionIDBytes returns the byte representation of the ID
func GetMissionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// SetMission set a specific mission in the store
func (k Keeper) SetMission(ctx sdk.Context, mission types.Mission) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	b := k.cdc.MustMarshal(&mission)
	store.Set(GetMissionIDBytes(mission.MissionID), b)
}

// GetMission returns a mission from its id
func (k Keeper) GetMission(ctx sdk.Context, id uint64) (val types.Mission, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	b := store.Get(GetMissionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMission removes a mission from the store
func (k Keeper) RemoveMission(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	store.Delete(GetMissionIDBytes(id))
}

// GetAllMission returns all mission
func (k Keeper) GetAllMission(ctx sdk.Context) (list []types.Mission) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mission
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CompleteMission triggers the completion of the mission and distribute the claimable portion of airdrop to the user
// the method fails if the mission has already been completed
func (k Keeper) CompleteMission(ctx sdk.Context, missionID uint64, address string) error {
	airdropSupply, found := k.GetAirdropSupply(ctx)
	if !found {
		return sdkerrors.Wrapf(types.ErrAirdropSupplyNotFound, "airdrop supply is not defined")
	}

	// retrieve mission
	mission, found := k.GetMission(ctx, missionID)
	if !found {
		return sdkerrors.Wrapf(types.ErrMissionNotFound, "mission %d not found", missionID)
	}

	// retrieve claim record of the user
	claimRecord, found := k.GetClaimRecord(ctx, address)
	if !found {
		return sdkerrors.Wrapf(types.ErrClaimRecordNotFound, "claim record not found for address %s", address)
	}

	// check if the mission is already complted for the claim record
	if claimRecord.IsMissionCompleted(missionID) {
		return sdkerrors.Wrapf(
			types.ErrMissionCompleted,
			"mission %d completed for address %s",
			missionID,
			address,
		)
	}
	claimRecord.CompletedMissions = append(claimRecord.CompletedMissions, missionID)

	// calculate claimable from mission weight and claim
	claimableAmount := mission.Weight.Mul(claimRecord.Claimable.ToDec()).TruncateInt()
	claimable := sdk.NewCoins(sdk.NewCoin(airdropSupply.Denom, claimableAmount))

	// decrease airdrop supply
	airdropSupply.Amount = airdropSupply.Amount.Sub(claimableAmount)
	if airdropSupply.Amount.IsNegative() {
		return spnerrors.Critical("airdrop supply is lower than total claimable")
	}

	// send claimable to the user
	claimer, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return spnerrors.Criticalf("invalid claimer address %s", err.Error())
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, claimer, claimable); err != nil {
		return spnerrors.Criticalf("can't send claimable coins %s", err.Error())
	}

	// update store
	k.SetAirdropSupply(ctx, airdropSupply)
	k.SetClaimRecord(ctx, claimRecord)

	return nil
}
