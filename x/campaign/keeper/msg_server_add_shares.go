package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) AddShares(goCtx context.Context, msg *types.MsgAddShares) (*types.MsgAddSharesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, found := k.GetCampaign(ctx, msg.CampaignID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCampaignNotFound, "%d", msg.CampaignID)
	}

	// Get the coordinator ID associated to the sender address
	coordByAddress, found := k.profileKeeper.GetCoordinatorByAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}

	coord, _ := k.profileKeeper.GetCoordinator(ctx, coordByAddress.CoordinatorID)
	if !coord.Active {
		return nil, sdkerrors.Wrapf(profiletypes.ErrCoordInactive,
			"coordinator %d inactive", coord.CoordinatorID)
	}

	if campaign.CoordinatorID != coord.CoordinatorID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the campaign is %d",
			campaign.CoordinatorID,
		))
	}

	// check if the account already exists
	account, found := k.GetMainnetAccount(ctx, campaign.CampaignID, msg.Address)
	if !found {
		// if not, create the account
		account = types.MainnetAccount{
			CampaignID: campaign.CampaignID,
			Address:    msg.Address,
			Shares:     types.EmptyShares(),
		}
	}
	// increase the account shares
	account.Shares = types.IncreaseShares(account.Shares, msg.Shares)

	// increase the campaign shares
	campaign.AllocatedShares = types.IncreaseShares(campaign.AllocatedShares, msg.Shares)
	if types.IsTotalSharesReached(campaign.AllocatedShares, campaign.TotalShares) {
		return nil, sdkerrors.Wrapf(types.ErrTotalSharesLimit, "%d", msg.CampaignID)
	}

	k.SetCampaign(ctx, campaign)
	k.SetMainnetAccount(ctx, account)

	return &types.MsgAddSharesResponse{}, nil
}
