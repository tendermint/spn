package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/x/participation/types"
)

func TestGetTierFromID(t *testing.T) {
	list := types.DefaultParticipationTierList

	tests := []struct {
		name  string
		id    uint64
		found bool
	}{
		{
			name:  "find tier id 1",
			id:    1,
			found: true,
		},
		{
			name:  "find tier id 2",
			id:    2,
			found: true,
		},
		{
			name:  "tier not found",
			id:    111111,
			found: false,
		},
	}

	for _, tt := range tests {
		_, found := types.GetTierFromID(list, tt.id)
		require.Equal(t, tt.found, found)
	}
}
