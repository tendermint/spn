package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/pkg/types"
	tc "github.com/tendermint/spn/testutil/constructor"
	"github.com/tendermint/spn/testutil/sample"
)

func TestMonitoringPacket_ValidateBasic(t *testing.T) {
	tests := []struct {
		name    string
		mp      types.MonitoringPacket
		wantErr bool
	}{
		{
			name: "empty is valid",
			mp:   types.MonitoringPacket{},
		},
		{
			name: "block height can be greater than block count",
			mp: types.MonitoringPacket{
				BlockHeight: 200,
				SignatureCounts: types.SignatureCounts{
					BlockCount: 100,
					Counts: []types.SignatureCount{
						tc.SignatureCount(t, sample.OperatorAddress(r), "1"),
						tc.SignatureCount(t, sample.OperatorAddress(r), "0.5"),
						tc.SignatureCount(t, sample.OperatorAddress(r), "5.5"),
					},
				},
			},
		},
		{
			name: "block height can be equal to block count",
			mp: types.MonitoringPacket{
				BlockHeight: 100,
				SignatureCounts: types.SignatureCounts{
					BlockCount: 100,
					Counts: []types.SignatureCount{
						tc.SignatureCount(t, sample.OperatorAddress(r), "1"),
						tc.SignatureCount(t, sample.OperatorAddress(r), "0.5"),
						tc.SignatureCount(t, sample.OperatorAddress(r), "5.5"),
					},
				},
			},
		},
		{
			name: "invalid signature counts should fail",
			mp: types.MonitoringPacket{
				BlockHeight: 1,
				SignatureCounts: types.SignatureCounts{
					BlockCount: 1,
					Counts: []types.SignatureCount{
						tc.SignatureCount(t, sample.OperatorAddress(r), "1"),
						tc.SignatureCount(t, sample.OperatorAddress(r), "0.5"),
						tc.SignatureCount(t, sample.OperatorAddress(r), "5.5"),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "block height lower than block count should fail",
			mp: types.MonitoringPacket{
				BlockHeight: 50,
				SignatureCounts: types.SignatureCounts{
					BlockCount: 100,
					Counts: []types.SignatureCount{
						tc.SignatureCount(t, sample.OperatorAddress(r), "1"),
						tc.SignatureCount(t, sample.OperatorAddress(r), "0.5"),
						tc.SignatureCount(t, sample.OperatorAddress(r), "5.5"),
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.mp.ValidateBasic()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
