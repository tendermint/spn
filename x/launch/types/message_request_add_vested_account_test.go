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

func TestMsgRequestAddVestedAccount_ValidateBasic(t *testing.T) {
	var (
		chainID, _ = sample.ChainID(10)
	)

	option := *types.NewDelayedVesting(sample.Coins(), time.Now().Unix())

	tests := []struct {
		name string
		msg  types.MsgRequestAddVestedAccount
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRequestAddVestedAccount{
				Address:         "invalid_address",
				ChainID:         chainID,
				StartingBalance: sample.Coins(),
				Options:         option,
			},
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress,
				"invalid address (invalid_address): decoding bech32 failed: invalid index of 1"),
		},
		{
			name: "invalid chain id",
			msg: types.MsgRequestAddVestedAccount{
				Address:         sample.AccAddress(),
				ChainID:         "invalid_chain",
				StartingBalance: sample.Coins(),
				Options:         option,
			},
			err: sdkerrors.Wrap(types.ErrInvalidChainID, "invalid_chain"),
		},
		{
			name: "invalid coins",
			msg: types.MsgRequestAddVestedAccount{
				Address:         sample.AccAddress(),
				ChainID:         chainID,
				StartingBalance: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
				Options:         option,
			},
			err: sdkerrors.Wrap(types.ErrInvalidCoins,
				"invalid starting balance: 10: the coin list is invalid"),
		},
		{
			name: "invalid vesting options",
			msg: types.MsgRequestAddVestedAccount{
				Address:         sample.AccAddress(),
				ChainID:         chainID,
				StartingBalance: sample.Coins(),
				Options:         *types.NewDelayedVesting(sample.Coins(), 0),
			},
			err: sdkerrors.Wrapf(types.ErrInvalidVestingOption, ""),
		},
		{
			name: "valid message",
			msg: types.MsgRequestAddVestedAccount{
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
