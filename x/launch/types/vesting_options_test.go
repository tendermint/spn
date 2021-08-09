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

func TestDelayedVesting_Validate(t *testing.T) {
	tests := []struct {
		name   string
		option types.DelayedVesting
		err    error
	}{
		{
			name: "vesting without coins",
			option: types.DelayedVesting{
				Vesting: nil,
				EndTime: time.Now().Unix(),
			},
			err: sdkerrors.Wrap(types.ErrInvalidCoins,
				"invalid vesting coins for DelayedVesting"),
		}, {
			name: "vesting with invalid coins",
			option: types.DelayedVesting{
				Vesting: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
				EndTime: 0,
			},
			err: sdkerrors.Wrap(types.ErrInvalidCoins,
				"invalid vesting coins for DelayedVesting"),
		}, {
			name: "vesting with invalid timestamp",
			option: types.DelayedVesting{
				Vesting: sample.Coins(),
				EndTime: 0,
			},
			err: sdkerrors.Wrap(types.ErrInvalidTimestamp,
				"invalid end time for DelayedVesting"),
		}, {
			name: "valid account vesting",
			option: types.DelayedVesting{
				Vesting: sample.Coins(),
				EndTime: time.Now().Unix(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.option.Validate()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
