package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/launch/types"
)

// GetRequestCount get the total number of request for a specific chain ID
func (k Keeper) GetRequestCount(ctx sdk.Context, chainID string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestCountKeyPrefix))
	bz := store.Get(types.RequestCountKey(chainID))

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to uint64
		panic("cannot decode count")
	}

	return count
}

// SetRequestCount set the total number of request for a chain
func (k Keeper) SetRequestCount(ctx sdk.Context, chainID string, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestCountKeyPrefix))
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(types.RequestCountKey(chainID), bz)
}

// SetRequest set a specific request in the store from its index
func (k Keeper) SetRequest(ctx sdk.Context, request types.Request) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))
	b := k.cdc.MustMarshalBinaryBare(&request)
	store.Set(types.RequestKey(
		request.ChainID,
		request.RequestID,
	), b)
}

// AppendRequest appends a request for a chain in the store with a new id and update the count
func (k Keeper) AppendRequest(ctx sdk.Context, request types.Request) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))

	count := k.GetRequestCount(ctx, request.ChainID)
	request.RequestID = count

	b := k.cdc.MustMarshalBinaryBare(&request)
	store.Set(types.RequestKey(
		request.ChainID,
		request.RequestID,
	), b)

	// increment the count
	k.SetRequestCount(ctx, request.ChainID, count+1)

	return count
}

// GetRequest returns a request from its index
func (k Keeper) GetRequest(
	ctx sdk.Context,
	chainID string,
	requestID uint64,
) (val types.Request, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))

	b := store.Get(types.RequestKey(
		chainID,
		requestID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveRequest removes a request from the store
func (k Keeper) RemoveRequest(
	ctx sdk.Context,
	chainID string,
	requestID uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))
	store.Delete(types.RequestKey(
		chainID,
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
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// checkAccount check account inconsistency and return
// if an account exists for genesis or vested accounts
func checkAccount(ctx sdk.Context, k Keeper, chainID, address string) (bool, error) {
	_, foundGenesis := k.GetGenesisAccount(ctx, chainID, address)
	_, foundVested := k.GetVestedAccount(ctx, chainID, address)
	if foundGenesis && foundVested {
		return false, spnerrors.Critical(
			fmt.Sprintf("account %s for chain %s found in vested and genesis accounts",
				address, chainID),
		)
	}
	return foundGenesis || foundVested, nil
}

// ApplyRequest approves the request and performs
// the launch information changes
func ApplyRequest(
	ctx sdk.Context,
	k Keeper,
	chainID string,
	request types.Request,
) error {
	if err := request.Content.Validate(); err != nil {
		return spnerrors.Critical(err.Error())
	}

	switch requestContent := request.Content.Content.(type) {
	case *types.RequestContent_GenesisAccount:
		ga := requestContent.GenesisAccount
		found, err := checkAccount(ctx, k, chainID, ga.Address)
		if err != nil {
			return err
		}
		if found {
			return sdkerrors.Wrapf(types.ErrAccountAlreadyExist,
				"account %s for chain %s already exist",
				ga.Address, chainID,
			)
		}
		k.SetGenesisAccount(ctx, *ga)
	case *types.RequestContent_VestedAccount:
		va := requestContent.VestedAccount
		found, err := checkAccount(ctx, k, chainID, va.Address)
		if err != nil {
			return err
		}
		if found {
			return sdkerrors.Wrapf(types.ErrAccountAlreadyExist,
				"account %s for chain %s already exist",
				va.Address, chainID,
			)
		}
		k.SetVestedAccount(ctx, *va)
	case *types.RequestContent_AccountRemoval:
		ar := requestContent.AccountRemoval
		found, err := checkAccount(ctx, k, chainID, ar.Address)
		if err != nil {
			return err
		}
		if !found {
			return sdkerrors.Wrapf(types.ErrAccountNotFound,
				"account %s for chain %s not found",
				ar.Address, chainID,
			)
		}
		k.RemoveGenesisAccount(ctx, chainID, ar.Address)
		k.RemoveVestedAccount(ctx, chainID, ar.Address)
	case *types.RequestContent_GenesisValidator:
		ga := requestContent.GenesisValidator
		if _, found := k.GetGenesisValidator(ctx, chainID, ga.Address); found {
			return sdkerrors.Wrapf(types.ErrValidatorAlreadyExist,
				"genesis validator %s for chain %s already exist",
				ga.Address, chainID,
			)
		}
		k.SetGenesisValidator(ctx, *ga)
	case *types.RequestContent_ValidatorRemoval:
		vr := requestContent.ValidatorRemoval
		if _, found := k.GetGenesisValidator(ctx, chainID, vr.ValAddress); !found {
			return sdkerrors.Wrapf(types.ErrValidatorNotFound,
				"genesis validator %s for chain %s not found",
				vr.ValAddress, chainID,
			)
		}
		k.RemoveGenesisValidator(ctx, chainID, vr.ValAddress)
	default:
		return spnerrors.Critical("unknown request content type")
	}
	return nil
}
