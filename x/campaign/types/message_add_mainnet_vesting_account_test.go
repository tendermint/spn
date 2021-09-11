package types_test

import (
	"github.com/tendermint/spn/x/campaign/types"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
)

func TestMsgAddMainnetVestingAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgAddMainnetVestingAccount
		err  error
	}{
		{
			name: "valid address",
			msg: types.MsgAddMainnetVestingAccount{
				Coordinator:    sample.AccAddress(),
				CampaignID:     0,
				Shares:         sample.Shares(),
				VestingOptions: sample.ShareVestingOptions(),
			},
		},
		{
			name: "invalid address",
			msg: types.MsgAddMainnetVestingAccount{
				Coordinator:    "invalid_address",
				CampaignID:     0,
				Shares:         sample.Shares(),
				VestingOptions: sample.ShareVestingOptions(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid total supply",
			msg: types.MsgAddMainnetVestingAccount{
				Coordinator:    sample.AccAddress(),
				CampaignID:     0,
				Shares:         invalidShares,
				VestingOptions: sample.ShareVestingOptions(),
			},
			err: types.ErrInvalidAccountShare,
		},
		{
			name: "empty total supply",
			msg: types.MsgAddMainnetVestingAccount{
				Coordinator:    sample.AccAddress(),
				CampaignID:     0,
				Shares:         types.Shares{},
				VestingOptions: sample.ShareVestingOptions(),
			},
			err: types.ErrInvalidAccountShare,
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
