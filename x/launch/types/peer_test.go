package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestPeer_Validate(t *testing.T) {
	tests := []struct {
		name    string
		peer    types.Peer
		wantErr bool
	}{
		{
			name:    "should validate valid new peer connection",
			peer:    types.NewPeerConn(sample.String(r, 3), sample.String(r, 10)),
			wantErr: false,
		},
		{
			name:    "should validate valid  new peer tunnel",
			peer:    types.NewPeerTunnel(sample.String(r, 3), sample.String(r, 5), sample.String(r, 10)),
			wantErr: false,
		},
		{
			name:    "should validate valid new peer empty",
			peer:    types.NewPeerEmpty(sample.String(r, 3)),
			wantErr: false,
		},
		{
			name:    "should prevent validate invalid peer",
			peer:    types.Peer{},
			wantErr: true,
		},
		{
			name:    "should prevent validate peer with empty peer id",
			peer:    types.NewPeerConn("", sample.String(r, 10)),
			wantErr: true,
		},
		{
			name:    "should prevent validate empty new peer connection address",
			peer:    types.NewPeerConn(sample.String(r, 3), ""),
			wantErr: true,
		},
		{
			name:    "should prevent validate empty new peer tunnel address",
			peer:    types.NewPeerTunnel(sample.String(r, 3), "", sample.String(r, 10)),
			wantErr: true,
		},
		{
			name:    "should prevent validate empty new peer tunnel name",
			peer:    types.NewPeerTunnel(sample.String(r, 3), sample.String(r, 10), ""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.peer.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
