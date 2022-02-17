package keeper_test

import (
	"encoding/base64"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

func createNGenesisValidator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GenesisValidator {
	items := make([]types.GenesisValidator, n)
	for i := range items {
		items[i] = sample.GenesisValidator(uint64(i), strconv.Itoa(i))
		keeper.SetGenesisValidator(ctx, items[i])
	}
	return items
}

func createNGenesisValidatorByLaunchID(keeper *keeper.Keeper, ctx sdk.Context, launchID int) []types.GenesisValidator {
	items := make([]types.GenesisValidator, launchID)
	for i := range items {
		items[i] = sample.GenesisValidator(uint64(launchID), strconv.Itoa(i))
		keeper.SetGenesisValidator(ctx, items[i])
	}
	return items
}

func TestGenesisValidatorGet(t *testing.T) {
	k, ctx := testkeeper.Launch(t)
	items := createNGenesisValidator(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetGenesisValidator(ctx,
			item.LaunchID,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestGenesisValidatorRemove(t *testing.T) {
	k, ctx := testkeeper.Launch(t)
	items := createNGenesisValidator(k, ctx, 10)
	for _, item := range items {
		k.RemoveGenesisValidator(ctx,
			item.LaunchID,
			item.Address,
		)
		_, found := k.GetGenesisValidator(ctx,
			item.LaunchID,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestGenesisValidatorGetAll(t *testing.T) {
	k, ctx := testkeeper.Launch(t)
	items := createNGenesisValidator(k, ctx, 10)
	require.ElementsMatch(t, items, k.GetAllGenesisValidator(ctx))
}

func TestKeeper_GetValidatorsAndTotalDelegation(t *testing.T) {
	k, ctx := testkeeper.Launch(t)
	launchID := 10
	validators := createNGenesisValidatorByLaunchID(k, ctx, launchID)
	totalSelfDelegation := sdk.ZeroDec()
	validatorMap := make(map[string]types.GenesisValidator)
	for _, validator := range validators {
		consPubKey := base64.StdEncoding.EncodeToString(validator.ConsPubKey)
		validatorMap[consPubKey] = validator
		totalSelfDelegation = totalSelfDelegation.Add(validator.SelfDelegation.Amount.ToDec())
	}
	val, got := k.GetValidatorsAndTotalDelegation(ctx, uint64(launchID))
	require.Equal(t, totalSelfDelegation, got)
	require.EqualValues(t, validatorMap, val)
}
