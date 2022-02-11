package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	valtypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

const validatorKey = `{
  "address": "B4AAC35ED4E14C09E530B10AF4DD604FAAC597C0",
  "pub_key": {
    "type": "tendermint/PubKeyEd25519",
    "value": "sYTsd7W1+SBtjD3BN/aTEDFvfRbZ9zdfpQH2Lk3MRK4="
  },
  "priv_key": {
    "type": "tendermint/PrivKeyEd25519",
    "value": "j45JhnCflEk3T6FC8LLuJqg9tPfCzJH+UYZY88xn+0exhOx3tbX5IG2MPcE39pMQMW99Ftn3N1+lAfYuTcxErg=="
  }
}`

func TestMsgSetValidatorConsAddress(t *testing.T) {
	var (
		ctx, k, srv = setupMsgServer(t)
		wCtx        = sdk.WrapSDKContext(ctx)
	)
	valKey, err := valtypes.LoadValidatorKey([]byte(validatorKey))
	require.NoError(t, err)
	signature, err := valKey.Sign(0, ctx.ChainID())
	require.NoError(t, err)

	k.SetValidator(ctx, types.Validator{
		Address:     valKey.Address.String(),
		Description: types.ValidatorDescription{},
	})

	tests := []struct {
		name   string
		pubKey valtypes.ValidatorConsPubKey
		msg    *types.MsgSetValidatorConsAddress
		err    error
	}{
		{
			name:   "invalid validator key",
			pubKey: valtypes.ValidatorConsPubKey{PubKey: valKey.PubKey},
			msg: &types.MsgSetValidatorConsAddress{
				ValidatorAddress:    valKey.Address.String(),
				ValidatorConsPubKey: []byte("invalid_key"),
				ValidatorKeyType:    "invalid_type",
				Signature:           signature,
				Nonce:               0,
				ChainID:             ctx.ChainID(),
			},
			err: spnerrors.ErrCritical,
		},
		{
			name:   "invalid validator nonce",
			pubKey: valtypes.ValidatorConsPubKey{PubKey: valKey.PubKey},
			msg: &types.MsgSetValidatorConsAddress{
				ValidatorAddress:    valKey.Address.String(),
				ValidatorConsPubKey: valKey.PubKey.Bytes(),
				ValidatorKeyType:    valKey.PubKey.Type(),
				Signature:           signature,
				Nonce:               1,
				ChainID:             ctx.ChainID(),
			},
			err: types.ErrInvalidValidatorNonce,
		},
		{
			name:   "invalid validator chain id",
			pubKey: valtypes.ValidatorConsPubKey{PubKey: valKey.PubKey},
			msg: &types.MsgSetValidatorConsAddress{
				ValidatorAddress:    valKey.Address.String(),
				ValidatorConsPubKey: valKey.PubKey.Bytes(),
				ValidatorKeyType:    valKey.PubKey.Type(),
				Signature:           signature,
				Nonce:               0,
				ChainID:             ctx.ChainID() + "_invalid",
			},
			err: types.ErrInvalidValidatorChainID,
		},
		{
			name:   "valid message",
			pubKey: valtypes.ValidatorConsPubKey{PubKey: valKey.PubKey},
			msg: &types.MsgSetValidatorConsAddress{
				ValidatorAddress:    valKey.Address.String(),
				ValidatorConsPubKey: valKey.PubKey.Bytes(),
				ValidatorKeyType:    valKey.PubKey.Type(),
				Signature:           signature,
				Nonce:               0,
				ChainID:             ctx.ChainID(),
			},
		},
		{
			name:   "duplicated cons pub key message",
			pubKey: valtypes.ValidatorConsPubKey{PubKey: valKey.PubKey},
			msg: &types.MsgSetValidatorConsAddress{
				ValidatorAddress:    sample.Address(),
				ValidatorConsPubKey: valKey.PubKey.Bytes(),
				ValidatorKeyType:    valKey.PubKey.Type(),
				Signature:           signature,
				Nonce:               1,
				ChainID:             ctx.ChainID(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentNonce := uint64(1)
			oldConsNonce, hasConsNonce := k.GetConsensusKeyNonce(ctx, tt.pubKey.GetConsAddress())
			if hasConsNonce {
				currentNonce = oldConsNonce.Nonce + 1
			}

			_, err := srv.SetValidatorConsAddress(wCtx, tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			validator, found := k.GetValidator(ctx, tt.msg.ValidatorAddress)
			require.True(t, found, "validator was not saved")
			require.Equal(t, tt.msg.ValidatorAddress, validator.Address)
			require.Equal(t, tt.pubKey.GetConsAddress().Bytes(), validator.ConsensusAddress)

			valByConsAddr, found := k.GetValidatorByConsAddress(ctx, validator.ConsensusAddress)
			require.True(t, found, "validator by consensus address was not saved")
			require.Equal(t, tt.msg.ValidatorAddress, valByConsAddr.ValidatorAddress)
			require.Equal(t, tt.pubKey.GetConsAddress().Bytes(), valByConsAddr.ConsensusAddress)

			consNonce, found := k.GetConsensusKeyNonce(ctx, validator.ConsensusAddress)
			require.True(t, found, "validator consensus nonce was not saved")
			require.Equal(t, currentNonce, consNonce.Nonce)
			require.Equal(t, tt.pubKey.GetConsAddress().Bytes(), consNonce.ConsensusAddress)
		})
	}
}
