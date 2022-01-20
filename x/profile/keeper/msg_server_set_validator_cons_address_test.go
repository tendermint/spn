package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgSetValidatorConsAddress(t *testing.T) {
	var (
		validatorAddr = sample.Address()
		consensusAddr = sample.Address()
		ctx, k, srv   = setupMsgServer(t)
		wCtx          = sdk.WrapSDKContext(ctx)
	)
	k.SetValidator(ctx, types.Validator{
		Address:     validatorAddr,
		Description: types.ValidatorDescription{},
	})
	k.SetValidatorByConsAddress(ctx, types.ValidatorByConsAddress{
		ValidatorAddress: validatorAddr,
		ConsensusAddress: consensusAddr,
	})
	tests := []struct {
		name string
		msg  *types.MsgSetValidatorConsAddress
		err  error
	}{
		{
			name: "valida message",
			msg:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.SetValidatorConsAddress(wCtx, tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			_, found := k.GetValidator(ctx, tt.msg.Address)
			require.False(t, found, "validator was not removed")
		})
	}
}
