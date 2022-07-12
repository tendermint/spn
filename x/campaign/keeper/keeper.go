package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/tendermint/spn/x/campaign/types"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	rewardtypes "github.com/tendermint/spn/x/reward/types"
)

type LaunchKeeper interface {
	GetChain(ctx sdk.Context, launchID uint64) (val launchtypes.Chain, found bool)
	GetRequestCount(ctx sdk.Context, launchID uint64) (count uint64)
	GetGenesisValidatorCount(ctx sdk.Context, launchID uint64) (count uint64)
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

type RewardKeeper interface {
	GetRewardPool(ctx sdk.Context, launchID uint64) (val rewardtypes.RewardPool, found bool)
}

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		launchKeeper  LaunchKeeper
		bankKeeper    types.BankKeeper
		distrKeeper   types.DistributionKeeper
		profileKeeper types.ProfileKeeper
		rewardKeeper  RewardKeeper
		paramSpace    paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	launchKeeper LaunchKeeper,
	bankKeeper types.BankKeeper,
	distrKeeper types.DistributionKeeper,
	profileKeeper types.ProfileKeeper,
	rewardKeeper RewardKeeper,
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
		rewardKeeper:  rewardKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
