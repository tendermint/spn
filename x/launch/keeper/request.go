package keeper

import (
	"encoding/binary"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/pkg/errors"

	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

// GetRequestCounter get request counter for a specific chain ID
func (k Keeper) GetRequestCounter(ctx sdk.Context, launchID uint64) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestCounterKeyPrefix))
	bz := store.Get(types.RequestCounterKey(launchID))

	// Counter doesn't exist: no element
	if bz == nil {
		return 1
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRequestCounter set the total number of request for a chain
func (k Keeper) SetRequestCounter(ctx sdk.Context, launchID, counter uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestCounterKeyPrefix))
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, counter)
	store.Set(types.RequestCounterKey(launchID), bz)
}

// SetRequest set a specific request in the store from its index
func (k Keeper) SetRequest(ctx sdk.Context, request types.Request) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))
	b := k.cdc.MustMarshal(&request)
	store.Set(types.RequestKey(
		request.LaunchID,
		request.RequestID,
	), b)
}

// AppendRequest appends a request for a chain in the store with a new id and update the counter
func (k Keeper) AppendRequest(ctx sdk.Context, request types.Request) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))

	counter := k.GetRequestCounter(ctx, request.LaunchID)
	request.RequestID = counter

	b := k.cdc.MustMarshal(&request)
	store.Set(types.RequestKey(
		request.LaunchID,
		request.RequestID,
	), b)

	// increment the counter
	k.SetRequestCounter(ctx, request.LaunchID, counter+1)

	return counter
}

// GetRequest returns a request from its index
func (k Keeper) GetRequest(
	ctx sdk.Context,
	launchID,
	requestID uint64,
) (val types.Request, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))

	b := store.Get(types.RequestKey(
		launchID,
		requestID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRequest removes a request from the store
func (k Keeper) RemoveRequest(
	ctx sdk.Context,
	launchID,
	requestID uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))
	store.Delete(types.RequestKey(
		launchID,
		requestID,
	))
}

// GetAllRequest returns all request
func (k Keeper) GetAllRequest(ctx sdk.Context) (list []types.Request) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Request
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CheckAccount check account inconsistency and return
// if an account exists for genesis or vesting accounts
func CheckAccount(ctx sdk.Context, k Keeper, launchID uint64, address string) (bool, error) {
	_, foundGenesis := k.GetGenesisAccount(ctx, launchID, address)
	_, foundVesting := k.GetVestingAccount(ctx, launchID, address)
	if foundGenesis && foundVesting {
		return false, ignterrors.Critical(
			fmt.Sprintf("account %s for chain %d found in vesting and genesis accounts",
				address, launchID),
		)
	}
	return foundGenesis || foundVesting, nil
}

// ApplyRequest approves the request and performs
// the launch information changes
func ApplyRequest(
	ctx sdk.Context,
	k Keeper,
	chain types.Chain,
	request types.Request,
	coord profiletypes.Coordinator,
) error {
	err := CheckRequest(ctx, k, chain.LaunchID, request)
	if err != nil {
		return err
	}

	switch requestContent := request.Content.Content.(type) {
	case *types.RequestContent_GenesisAccount:
		ga := requestContent.GenesisAccount
		if !chain.AccountBalance.Empty() {
			ga.Coins = chain.AccountBalance
		}
		k.SetGenesisAccount(ctx, *ga)
		err = ctx.EventManager().EmitTypedEvent(&types.EventGenesisAccountAdded{
			Address:            ga.Address,
			Coins:              ga.Coins,
			LaunchID:           chain.LaunchID,
			CoordinatorAddress: coord.Address,
		})

	case *types.RequestContent_VestingAccount:
		va := requestContent.VestingAccount
		if !chain.AccountBalance.Empty() {
			switch opt := va.VestingOptions.Options.(type) { //nolint:gocritic
			case *types.VestingOptions_DelayedVesting:
				dv := opt.DelayedVesting
				va = &types.VestingAccount{
					Address:  va.Address,
					LaunchID: va.LaunchID,
					VestingOptions: *types.NewDelayedVesting(
						chain.AccountBalance,
						chain.AccountBalance,
						dv.EndTime,
					),
				}
			}
		}
		k.SetVestingAccount(ctx, *va)
		err = ctx.EventManager().EmitTypedEvent(&types.EventVestingAccountAdded{
			Address:            va.Address,
			VestingOptions:     va.VestingOptions,
			LaunchID:           chain.LaunchID,
			CoordinatorAddress: coord.Address,
		})

	case *types.RequestContent_AccountRemoval:
		ar := requestContent.AccountRemoval
		k.RemoveGenesisAccount(ctx, chain.LaunchID, ar.Address)
		k.RemoveVestingAccount(ctx, chain.LaunchID, ar.Address)
		err = ctx.EventManager().EmitTypedEvent(&types.EventAccountRemoved{
			Address:            ar.Address,
			LaunchID:           chain.LaunchID,
			CoordinatorAddress: coord.Address,
		})

	case *types.RequestContent_GenesisValidator:
		ga := requestContent.GenesisValidator
		k.SetGenesisValidator(ctx, *ga)
		err = ctx.EventManager().EmitTypedEvent(&types.EventValidatorAdded{
			Address:            ga.Address,
			GenTx:              ga.GenTx,
			ConsPubKey:         ga.ConsPubKey,
			SelfDelegation:     ga.SelfDelegation,
			Peer:               ga.Peer,
			LaunchID:           chain.LaunchID,
			HasCampaign:        chain.HasCampaign,
			CampaignID:         chain.CampaignID,
			CoordinatorAddress: coord.Address,
		})

	case *types.RequestContent_ValidatorRemoval:
		vr := requestContent.ValidatorRemoval
		k.RemoveGenesisValidator(ctx, chain.LaunchID, vr.ValAddress)
		err = ctx.EventManager().EmitTypedEvent(&types.EventValidatorRemoved{
			GenesisValidatorAccount: vr.ValAddress,
			LaunchID:                chain.LaunchID,
			HasCampaign:             chain.HasCampaign,
			CampaignID:              chain.CampaignID,
			CoordinatorAddress:      coord.Address,
		})

	}
	return err
}

// CheckRequest verifies that a request can be applied
func CheckRequest(
	ctx sdk.Context,
	k Keeper,
	launchID uint64,
	request types.Request,
) error {
	if err := request.Content.Validate(launchID); err != nil {
		return ignterrors.Critical(err.Error())
	}

	switch requestContent := request.Content.Content.(type) {
	case *types.RequestContent_GenesisAccount:
		ga := requestContent.GenesisAccount
		found, err := CheckAccount(ctx, k, launchID, ga.Address)
		if err != nil {
			return err
		}
		if found {
			return sdkerrors.Wrapf(types.ErrAccountAlreadyExist,
				"account %s for chain %d already exist",
				ga.Address, launchID,
			)
		}
	case *types.RequestContent_VestingAccount:
		va := requestContent.VestingAccount
		found, err := CheckAccount(ctx, k, launchID, va.Address)
		if err != nil {
			return err
		}
		if found {
			return sdkerrors.Wrapf(types.ErrAccountAlreadyExist,
				"account %s for chain %d already exist",
				va.Address, launchID,
			)
		}
	case *types.RequestContent_AccountRemoval:
		ar := requestContent.AccountRemoval
		found, err := CheckAccount(ctx, k, launchID, ar.Address)
		if err != nil {
			return err
		}
		if !found {
			return sdkerrors.Wrapf(types.ErrAccountNotFound,
				"account %s for chain %d not found",
				ar.Address, launchID,
			)
		}
	case *types.RequestContent_GenesisValidator:
		ga := requestContent.GenesisValidator
		if _, found := k.GetGenesisValidator(ctx, launchID, ga.Address); found {
			return sdkerrors.Wrapf(types.ErrValidatorAlreadyExist,
				"genesis validator %s for chain %d already exist",
				ga.Address, launchID,
			)
		}
	case *types.RequestContent_ValidatorRemoval:
		vr := requestContent.ValidatorRemoval
		if _, found := k.GetGenesisValidator(ctx, launchID, vr.ValAddress); !found {
			return sdkerrors.Wrapf(types.ErrValidatorNotFound,
				"genesis validator %s for chain %d not found",
				vr.ValAddress, launchID,
			)
		}
	case *types.RequestContent_ChangeParam:
		// currently no stateful checks can be performed on param
	}

	return nil
}
