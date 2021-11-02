package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/launch/types"
)

// GetRequestCount get the total number of request for a specific chain ID
func (k Keeper) GetRequestCount(ctx sdk.Context, launchID uint64) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestCountKeyPrefix))
	bz := store.Get(types.RequestCountKey(launchID))

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRequestCount set the total number of request for a chain
func (k Keeper) SetRequestCount(ctx sdk.Context, launchID, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestCountKeyPrefix))
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(types.RequestCountKey(launchID), bz)
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

// AppendRequest appends a request for a chain in the store with a new id and update the count
func (k Keeper) AppendRequest(ctx sdk.Context, request types.Request) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))

	count := k.GetRequestCount(ctx, request.LaunchID)
	request.RequestID = count

	b := k.cdc.MustMarshal(&request)
	store.Set(types.RequestKey(
		request.LaunchID,
		request.RequestID,
	), b)

	// increment the count
	k.SetRequestCount(ctx, request.LaunchID, count+1)

	return count
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
		return false, spnerrors.Critical(
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
	launchID uint64,
	request types.Request,
) error {
	if err := request.Content.Validate(); err != nil {
		return spnerrors.Critical(err.Error())
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
		k.SetGenesisAccount(ctx, *ga)
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
		k.SetVestingAccount(ctx, *va)
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
		k.RemoveGenesisAccount(ctx, launchID, ar.Address)
		k.RemoveVestingAccount(ctx, launchID, ar.Address)
	case *types.RequestContent_GenesisValidator:
		ga := requestContent.GenesisValidator
		if _, found := k.GetGenesisValidator(ctx, launchID, ga.Address); found {
			return sdkerrors.Wrapf(types.ErrValidatorAlreadyExist,
				"genesis validator %s for chain %d already exist",
				ga.Address, launchID,
			)
		}
		k.SetGenesisValidator(ctx, *ga)
	case *types.RequestContent_ValidatorRemoval:
		vr := requestContent.ValidatorRemoval
		if _, found := k.GetGenesisValidator(ctx, launchID, vr.ValAddress); !found {
			return sdkerrors.Wrapf(types.ErrValidatorNotFound,
				"genesis validator %s for chain %d not found",
				vr.ValAddress, launchID,
			)
		}
		k.RemoveGenesisValidator(ctx, launchID, vr.ValAddress)
	default:
		return spnerrors.Critical("unknown request content type")
	}
	return nil
}
