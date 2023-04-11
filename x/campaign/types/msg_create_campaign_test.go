package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profile "github.com/tendermint/spn/x/profile/types"
)

func TestMsgCreateCampaign_ValidateBasic(t *testing.T) {
	invalidCoins := sdk.Coins{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}

	tests := []struct {
		name string
		msg  types.MsgCreateCampaign
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgCreateCampaign{
				Coordinator:  sample.Address(r),
				CampaignName: sample.CampaignName(r),
				TotalSupply:  sample.TotalSupply(r),
				Metadata:     sample.Metadata(r, 20),
			},
		},
		{
			name: "should prevent validation of msg with invalid address",
			msg: types.MsgCreateCampaign{
				Coordinator:  "invalid_address",
				CampaignName: sample.CampaignName(r),
				TotalSupply:  sample.TotalSupply(r),
				Metadata:     sample.Metadata(r, 20),
			},
			err: profile.ErrInvalidCoordAddress,
		},
		{
			name: "should prevent validation of msg with invalid campaign name",
			msg: types.MsgCreateCampaign{
				Coordinator:  sample.Address(r),
				CampaignName: invalidCampaignName,
				TotalSupply:  sample.TotalSupply(r),
				Metadata:     sample.Metadata(r, 20),
			},
			err: types.ErrInvalidCampaignName,
		},
		{
			name: "should prevent validation of msg with invalid total supply",
			msg: types.MsgCreateCampaign{
				Coordinator:  sample.Address(r),
				CampaignName: sample.CampaignName(r),
				TotalSupply:  invalidCoins,
				Metadata:     sample.Metadata(r, 20),
			},
			err: types.ErrInvalidTotalSupply,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
