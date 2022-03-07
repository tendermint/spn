package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
)

// CreateNewChain creates a new chain in the store from the provided information
func (k Keeper) CreateNewChain(
	ctx sdk.Context,
	coordinatorID uint64,
	genesisChainID,
	sourceURL,
	sourceHash,
	genesisURL,
	genesisHash string,
	hasCampaign bool,
	campaignID uint64,
	isMainnet bool,
	metadata []byte,
) (uint64, error) {
	chain := types.Chain{
		CoordinatorID:   coordinatorID,
		GenesisChainID:  genesisChainID,
		CreatedAt:       ctx.BlockTime().Unix(),
		SourceURL:       sourceURL,
		SourceHash:      sourceHash,
		HasCampaign:     hasCampaign,
		CampaignID:      campaignID,
		IsMainnet:       isMainnet,
		LaunchTriggered: false,
		LaunchTimestamp: 0,
		Metadata:        metadata,
	}

	// Initialize initial genesis
	if genesisURL == "" {
		chain.InitialGenesis = types.NewDefaultInitialGenesis()
	} else {
		chain.InitialGenesis = types.NewGenesisURL(genesisURL, genesisHash)
	}

	if err := chain.Validate(); err != nil {
		return 0, err
	}

	// If the chain is associated to a campaign, campaign existence and coordinator is checked
	if hasCampaign {
		campaign, found := k.campaignKeeper.GetCampaign(ctx, campaignID)
		if !found {
			return 0, fmt.Errorf("campaign %d doesn't exist", campaignID)
		}
		if campaign.CoordinatorID != coordinatorID {
			return 0, fmt.Errorf(
				"chain coordinator %d and campaign coordinator %d don't match",
				coordinatorID,
				campaign.CoordinatorID,
			)
		}
	}

	// Append the chain to the store
	launchID := k.AppendChain(ctx, chain)

	// Register the chain to the campaign
	if hasCampaign {
		if err := k.campaignKeeper.AddChainToCampaign(ctx, campaignID, launchID); err != nil {
			return 0, err
		}
	}

	return launchID, nil
}

// GetChainCounter get the counter for chains
func (k Keeper) GetChainCounter(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ChainCounterKey)
	bz := store.Get(byteKey)

	// Counter doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetChainCounter set the counter for chains
func (k Keeper) SetChainCounter(ctx sdk.Context, counter uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ChainCounterKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, counter)
	store.Set(byteKey, bz)
}

// AppendChain appends a chain in the store with a new id and update the counter
func (k Keeper) AppendChain(ctx sdk.Context, chain types.Chain) uint64 {
	counter := k.GetChainCounter(ctx)

	// Set the ID of the appended value
	chain.LaunchID = counter

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKeyPrefix))
	appendedValue := k.cdc.MustMarshal(&chain)
	store.Set(types.ChainKey(chain.LaunchID), appendedValue)

	// Update chain counter
	k.SetChainCounter(ctx, counter+1)

	return counter
}

// SetChain set a specific chain in the store from its index
func (k Keeper) SetChain(ctx sdk.Context, chain types.Chain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKeyPrefix))
	b := k.cdc.MustMarshal(&chain)
	store.Set(types.ChainKey(chain.LaunchID), b)
}

// EnableMonitoringConnection sets a chain with MonitoringConnected set to true
func (k Keeper) EnableMonitoringConnection(ctx sdk.Context, launchID uint64) error {
	chain, found := k.GetChain(ctx, launchID)
	if !found {
		return types.ErrChainNotFound
	}

	chain.MonitoringConnected = true
	k.SetChain(ctx, chain)
	return nil
}

// GetChain returns a chain from its index
func (k Keeper) GetChain(ctx sdk.Context, launchID uint64) (val types.Chain, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKeyPrefix))

	b := store.Get(types.ChainKey(launchID))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllChain returns all chain
func (k Keeper) GetAllChain(ctx sdk.Context) (list []types.Chain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Chain
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
