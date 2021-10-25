package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/spn/x/launch/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc            codec.BinaryCodec
		storeKey       sdk.StoreKey
		memKey         sdk.StoreKey
		paramstore     paramtypes.Subspace
		profileKeeper  types.ProfileKeeper
		campaignKeeper types.CampaignKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	profileKeeper types.ProfileKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		profileKeeper: profileKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetCampaignKeeper set the campaign keeper interface of the module
func (k *Keeper) SetCampaignKeeper(campaignKeeper types.CampaignKeeper) {
	if k.campaignKeeper != nil {
		panic("campaign keeper already set for launch module")
	}
	k.campaignKeeper = campaignKeeper
}

// GetProfileKeeper get the profile keeper interface of the module
func (k *Keeper) GetProfileKeeper() types.ProfileKeeper {
	return k.profileKeeper
}
