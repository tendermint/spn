package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				PortId: types.PortID,
				ConsumerClientID: &types.ConsumerClientID{
					ClientID: "29",
				},
				Params: types.DefaultParams(),
				ConnectionChannelID: &types.ConnectionChannelID{
					ChannelID: "67",
				},
				MonitoringInfo: &types.MonitoringInfo{},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "invalid params",
			genState: &types.GenesisState{
				PortId: types.PortID,
				ConsumerClientID: &types.ConsumerClientID{
					ClientID: "29",
				},
				Params: types.NewParams(
					1000,
					"foo", // chain id should be <chain-name>-<revision-number>
					sample.ConsensusState(0),
					false,
				),
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: false,
		},
		{
			desc: "invalid monitoring info",
			genState: &types.GenesisState{
				PortId: types.PortID,
				ConsumerClientID: &types.ConsumerClientID{
					ClientID: "29",
				},
				Params: types.DefaultParams(),
				ConnectionChannelID: &types.ConnectionChannelID{
					ChannelID: "67",
				},
				// Block count is lower than sum of relative signatures
				MonitoringInfo: &types.MonitoringInfo{
					SignatureCounts: spntypes.SignatureCounts{
						BlockCount: 1,
						Counts: []spntypes.SignatureCount{
							{
								ConsAddress:        []byte("foo"),
								RelativeSignatures: sdk.NewDec(10),
							},
						},
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
