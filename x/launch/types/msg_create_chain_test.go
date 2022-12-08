package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgCreateChain_ValidateBasic(t *testing.T) {
	// TODO check error types in test

	invalidGenesisHash := sample.MsgCreateChain(r, sample.Address(r), "foo.com", false, 0)
	invalidInitialGenesis := types.NewGenesisURL("foo.com", "NoHash")
	invalidGenesisHash.InitialGenesis = invalidInitialGenesis

	invalidConfigFile := sample.MsgCreateChain(r, sample.Address(r), "foo.com", false, 0)
	invalidInitialGenesis = types.NewGenesisConfig("")
	invalidConfigFile.InitialGenesis = invalidInitialGenesis

	invalidGenesisChainID := sample.MsgCreateChain(r, sample.Address(r), "", false, 0)
	invalidGenesisChainID.GenesisChainID = "invalid"

	msgInvalidCoins := sample.MsgCreateChain(r, sample.Address(r), "foo.com", false, 0)
	msgInvalidCoins.AccountBalance = sdk.Coins{sdk.Coin{Denom: "invalid", Amount: sdk.NewInt(-1)}}

	for _, tc := range []struct {
		desc  string
		msg   types.MsgCreateChain
		valid bool
	}{
		{
			desc:  "should validate valid message",
			msg:   sample.MsgCreateChain(r, sample.Address(r), "", false, 0),
			valid: true,
		},
		{
			desc:  "should validate valid message with genesis URL",
			msg:   sample.MsgCreateChain(r, sample.Address(r), "foo.com", false, 0),
			valid: true,
		},
		{
			desc:  "should prevent validate message with invalid address",
			msg:   sample.MsgCreateChain(r, "invalid", "", false, 0),
			valid: false,
		},
		{
			desc:  "should prevent validate message with invalid genesis hash for custom genesis",
			msg:   invalidGenesisHash,
			valid: false,
		},
		{
			desc:  "should prevent validate message with invalid file for GenesisConfig custom genesis",
			msg:   invalidConfigFile,
			valid: false,
		},
		{
			desc:  "should prevent validate message with invalid genesis chain ID",
			msg:   invalidGenesisChainID,
			valid: false,
		},
		{
			desc:  "should prevent chain with invalid coins structure",
			msg:   msgInvalidCoins,
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
