package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profile "github.com/tendermint/spn/x/profile/types"
)

func TestMsgInitializeMainnet_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgInitializeMainnet
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgInitializeMainnet{
				Coordinator:    sample.Address(r),
				CampaignID:     sample.Uint64(r),
				SourceURL:      sample.String(r, 30),
				SourceHash:     sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
		},
		{
			name: "should prevent validation of msg with invalid address",
			msg: types.MsgInitializeMainnet{
				Coordinator:    "invalid_address",
				CampaignID:     sample.Uint64(r),
				SourceURL:      sample.String(r, 30),
				SourceHash:     sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: profile.ErrInvalidCoordAddress,
		},
		{
			name: "should prevent validation of msg with empty source URL",
			msg: types.MsgInitializeMainnet{
				Coordinator:    sample.Address(r),
				CampaignID:     sample.Uint64(r),
				SourceURL:      "",
				SourceHash:     sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: types.ErrInvalidMainnetInfo,
		},
		{
			name: "should prevent validation of msg with empty source hash",
			msg: types.MsgInitializeMainnet{
				Coordinator:    sample.Address(r),
				CampaignID:     sample.Uint64(r),
				SourceURL:      sample.String(r, 30),
				SourceHash:     "",
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: types.ErrInvalidMainnetInfo,
		},
		{
			name: "should prevent validation of msg with invalid chain id",
			msg: types.MsgInitializeMainnet{
				Coordinator:    sample.Address(r),
				CampaignID:     sample.Uint64(r),
				SourceURL:      sample.String(r, 30),
				SourceHash:     sample.String(r, 20),
				MainnetChainID: "invalid_chain_id",
			},
			err: types.ErrInvalidMainnetInfo,
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
