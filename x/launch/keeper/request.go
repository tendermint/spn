package keeper

import (
	"fmt"
	"strconv"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
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

// applyRequest approves the request and performs
// the launch information changes
func applyRequest(
	ctx sdk.Context,
	k Keeper,
	msg *types.MsgSettleRequest,
	request types.Request,
) error {
	cdc := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(cdc)

	var content types.RequestContent
	if err := cdc.UnpackAny(request.Content, &content); err != nil {
		return spnerrors.Critical(err.Error())
	}

	switch c := content.(type) {
	case *types.GenesisAccount:
		_, foundGenesis := k.GetGenesisAccount(ctx, msg.ChainID, c.Address)
		_, foundVested := k.GetVestedAccount(ctx, msg.ChainID, c.Address)
		if foundVested || foundGenesis {
			return spnerrors.Critical(fmt.Sprintf(
				"account %s for chain %s already exist",
				c.Address, msg.ChainID),
			)
		}
		k.SetGenesisAccount(ctx, *c)
	case *types.VestedAccount:
		_, foundGenesis := k.GetGenesisAccount(ctx, msg.ChainID, c.Address)
		_, foundVested := k.GetVestedAccount(ctx, msg.ChainID, c.Address)
		if foundVested || foundGenesis {
			return spnerrors.Critical(fmt.Sprintf(
				"account %s for chain %s already exist",
				c.Address, msg.ChainID),
			)
		}
		k.SetVestedAccount(ctx, *c)
	case *types.AccountRemoval:
		_, foundGenesis := k.GetGenesisAccount(ctx, msg.ChainID, c.Address)
		_, foundVested := k.GetVestedAccount(ctx, msg.ChainID, c.Address)
		if foundGenesis {
			k.RemoveGenesisAccount(ctx, msg.ChainID, c.Address)
		}
		if foundVested {
			k.RemoveVestedAccount(ctx, msg.ChainID, c.Address)
		}
		if !foundGenesis && !foundVested {
			return spnerrors.Critical(fmt.Sprintf(
				"account %s for chain %s not found",
				c.Address, msg.ChainID),
			)
		}
	case *types.GenesisValidator:
		if _, found := k.GetGenesisValidator(ctx, msg.ChainID, c.Address); found {
			return spnerrors.Critical(fmt.Sprintf(
				"genesis validator %s for chain %s already exist",
				c.Address, msg.ChainID),
			)
		}
		k.SetGenesisValidator(ctx, *c)
	case *types.ValidatorRemoval:
		if _, found := k.GetGenesisValidator(ctx, msg.ChainID, c.ValAddress); !found {
			return spnerrors.Critical(fmt.Sprintf(
				"genesis validator %s for chain %s not found",
				c.ValAddress, msg.ChainID),
			)
		}
		k.RemoveGenesisValidator(ctx, msg.ChainID, c.ValAddress)
	default:
		return spnerrors.Critical("unknown request content type")
	}
	return nil
}
