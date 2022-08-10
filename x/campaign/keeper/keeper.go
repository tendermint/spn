package keeper

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/tendermint/spn/x/campaign/types"
	launchtypes "github.com/tendermint/spn/x/launch/types"
)

type LaunchKeeper interface {
	GetChain(ctx sdk.Context, launchID uint64) (val launchtypes.Chain, found bool)
	CreateNewChain(
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
	) (uint64, error)
}

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      storetypes.StoreKey
		memKey        storetypes.StoreKey
		launchKeeper  LaunchKeeper
		bankKeeper    types.BankKeeper
		distrKeeper   types.DistributionKeeper
		profileKeeper types.ProfileKeeper
		paramSpace    paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	paramSpace paramtypes.Subspace,
	launchKeeper LaunchKeeper,
	bankKeeper types.BankKeeper,
	distrKeeper types.DistributionKeeper,
	profileKeeper types.ProfileKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramSpace:    paramSpace,
		launchKeeper:  launchKeeper,
		bankKeeper:    bankKeeper,
		distrKeeper:   distrKeeper,
		profileKeeper: profileKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
