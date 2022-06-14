package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"testing"
)

func TestKeeper_EmitCampaignAuctionCreated(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	type inputState struct {
		noCampaign    bool
		noCoordinator bool
		campaign      types.Campaign
		coordinator   profiletypes.Coordinator
	}

	// prepare addresses
	var addr []string
	for i := 0; i < 20; i++ {
		addr = append(addr, sample.Address(r))
	}

	tests := []struct {
		name        string
		inputState  inputState
		auctionId   uint64
		auctioneer  string
		sellingCoin sdk.Coin
		emitted     bool
		err         error
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// initialize input state
			if !tt.inputState.noCampaign {
				tk.CampaignKeeper.SetCampaign(ctx, tt.inputState.campaign)
			}
			if !tt.inputState.noCoordinator {
				tk.ProfileKeeper.SetCoordinator(ctx, tt.inputState.coordinator)
			}

			emitted, err := tk.CampaignKeeper.EmitCampaignAuctionCreated(ctx, tt.auctionId, tt.auctioneer, tt.sellingCoin)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
				require.EqualValues(t, tt.emitted, emitted)
			}

			// clean state
			if !tt.inputState.noCampaign {
				tk.CampaignKeeper.RemoveCampaign(ctx, tt.inputState.campaign.CampaignID)
			}
			if !tt.inputState.noCoordinator {
				tk.ProfileKeeper.RemoveCoordinator(ctx, tt.inputState.coordinator.CoordinatorID)
			}
		})
	}
}
