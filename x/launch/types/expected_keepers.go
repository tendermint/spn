package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
)

type CampaignKeeper interface {
	GetCampaign(ctx sdk.Context, id uint64) (campaigntypes.Campaign, bool)
	AddChainToCampaign(ctx sdk.Context, campaignID, chainID uint64) error
	GetCampaignChains(ctx sdk.Context, campaignID uint64) (val campaigntypes.CampaignChains, found bool)
	SetCampaignChains(ctx sdk.Context, campaignChains campaigntypes.CampaignChains)
}

type ProfileKeeper interface {
	CoordinatorIDFromAddress(ctx sdk.Context, address string) (id uint64, found bool)
	GetCoordinatorAddressFromID(ctx sdk.Context, id uint64) (address string, found bool)
}
