package types_test

import (
	"testing"
	"time"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestAddVestedAccount_ValidateBasic(t *testing.T) {
	var (
		addr       = sample.AccAddress()
		chainID, _ = sample.ChainID(10)
	)

	option, err := codec.NewAnyWithValue(&types.DelayedVesting{
		Vesting: sample.Coins(),
		EndTime: time.Now().Unix(),
	})
	require.NoError(t, err)

	invalidOption, err := codec.NewAnyWithValue(&types.Request{})
	require.NoError(t, err)

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
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid chain id",
			msg: types.MsgRequestAddVestedAccount{
				Address:         sample.AccAddress(),
				ChainID:         "invalid_chain",
				StartingBalance: sample.Coins(),
				Options:         option,
			},
			err: types.ErrInvalidChainID,
		}, {
			name: "nil message option",
			msg: types.MsgRequestAddVestedAccount{
				Address:         addr,
				ChainID:         chainID,
				StartingBalance: sample.Coins(),
				Options:         nil,
			},
			err: types.ErrInvalidAccountOption,
		}, {
			name: "invalid coins",
			msg: types.MsgRequestAddVestedAccount{
				Address:         sample.AccAddress(),
				ChainID:         chainID,
				StartingBalance: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
				Options:         option,
			},
			err: types.ErrInvalidCoins,
		}, {
			name: "invalid message option",
			msg: types.MsgRequestAddVestedAccount{
				Address:         addr,
				ChainID:         chainID,
				StartingBalance: sample.Coins(),
				Options:         invalidOption,
			},
			err: types.ErrInvalidAccountOption,
		}, {
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
