package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
)

type CampaignKeeper interface {
	GetCampaignChains(ctx sdk.Context, campaignID uint64) (val types.CampaignChains, found bool)
	SetCampaignChains(ctx sdk.Context, campaignChains types.CampaignChains)
}

type ProfileKeeper interface {
	CoordinatorIDFromAddress(ctx sdk.Context, address string) (id uint64, found bool)
	GetCoordinatorAddressFromID(ctx sdk.Context, id uint64) (address string, found bool)
}
