package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestAddVestingAccount_ValidateBasic(t *testing.T) {
	var (
		chainID = uint64(10)
	)

	option := *types.NewDelayedVesting(sample.Coins(), time.Now().Unix())

	tests := []struct {
		name string
		msg  types.MsgRequestAddVestingAccount
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRequestAddVestingAccount{
				Address:         "invalid_address",
				ChainID:         chainID,
				StartingBalance: sample.Coins(),
				Options:         option,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid coins",
			msg: types.MsgRequestAddVestingAccount{
				Address:         sample.AccAddress(),
				ChainID:         chainID,
				StartingBalance: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
				Options:         option,
			},
			err: types.ErrInvalidCoins,
		},
		{
			name: "invalid message option",
			msg: types.MsgRequestAddVestingAccount{
				Address:         sample.AccAddress(),
				ChainID:         chainID,
				StartingBalance: sample.Coins(),
				Options:         *types.NewDelayedVesting(sample.Coins(), 0),
			},
			err: types.ErrInvalidVestingOption,
		},
		{
			name: "valid message",
			msg: types.MsgRequestAddVestingAccount{
				Address:         sample.AccAddress(),
				ChainID:         chainID,
				StartingBalance: sample.Coins(),
				Options:         option,
			},
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
