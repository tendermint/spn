package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgInitializeMainnet_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgInitializeMainnet
		err  error
	}{
		{
			name: "valid message",
			msg: types.MsgInitializeMainnet{
				Coordinator:    sample.Address(),
				CampaignID:     sample.Uint64(),
				SourceURL:      sample.String(30),
				SourceHash:     sample.String(20),
				MainnetChainID: sample.GenesisChainID(),
			},
		},
		{
			name: "invalid address",
			msg: types.MsgInitializeMainnet{
				Coordinator:    "invalid_address",
				CampaignID:     sample.Uint64(),
				SourceURL:      sample.String(30),
				SourceHash:     sample.String(20),
				MainnetChainID: sample.GenesisChainID(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty source URL",
			msg: types.MsgInitializeMainnet{
				Coordinator:    sample.Address(),
				CampaignID:     sample.Uint64(),
				SourceURL:      "",
				SourceHash:     sample.String(20),
				MainnetChainID: sample.GenesisChainID(),
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "empty source hash",
			msg: types.MsgInitializeMainnet{
				Coordinator:    sample.Address(),
				CampaignID:     sample.Uint64(),
				SourceURL:      sample.String(30),
				SourceHash:     "",
				MainnetChainID: sample.GenesisChainID(),
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid chain id",
			msg: types.MsgInitializeMainnet{
				Coordinator:    sample.Address(),
				CampaignID:     sample.Uint64(),
				SourceURL:      sample.String(30),
				SourceHash:     sample.String(20),
				MainnetChainID: "invalid_chain_id",
			},
			err: sdkerrors.ErrInvalidRequest,
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
