package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	valtypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmjson "github.com/tendermint/tendermint/libs/json"
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
		randPrivKey = ed25519.GenPrivKey()
		randPubKey  = randPrivKey.PubKey()
		randValKey  = valtypes.ValidatorKey{
			Address: randPubKey.Address(),
			PubKey:  randPubKey,
			PrivKey: randPrivKey,
		}
		consensusAddrWithoutAcc = sample.Address()
		ctx, k, srv             = setupMsgServer(t)
		wCtx                    = sdk.WrapSDKContext(ctx)
	)
	invalidValKey, err := tmjson.Marshal(randValKey)
	require.NoError(t, err)
	valKey, err := valtypes.LoadValidatorKey([]byte(validatorKey))
	require.NoError(t, err)
	signature, err := valKey.Sign(0)
	require.NoError(t, err)

	k.SetValidator(ctx, types.Validator{
		Address:     valKey.Address.String(),
		Description: types.ValidatorDescription{},
	})
	k.SetValidatorByConsAddress(ctx, types.ValidatorByConsAddress{
		ValidatorAddress: valKey.Address.String(),
		ConsensusAddress: valKey.GetConsAddress().String(),
	})

	k.SetValidatorByConsAddress(ctx, types.ValidatorByConsAddress{
		ValidatorAddress: sample.Address(),
		ConsensusAddress: consensusAddrWithoutAcc,
	})

	tests := []struct {
		name   string
		valKey valtypes.ValidatorKey
		msg    *types.MsgSetValidatorConsAddress
		err    error
	}{
		{
			name:   "valid message",
			valKey: valKey,
			msg: &types.MsgSetValidatorConsAddress{
				ValidatorKey: []byte(validatorKey),
				Signature:    signature,
			},
		},
		{
			name:   "invalid signature",
			valKey: valKey,
			msg: &types.MsgSetValidatorConsAddress{
				ValidatorKey: []byte(validatorKey),
				Signature:    "invalid_signature",
			},
			err: types.ErrInvalidValidatorSignature,
		},
		{
			name:   "invalid validator key",
			valKey: randValKey,
			msg: &types.MsgSetValidatorConsAddress{
				ValidatorKey: invalidValKey,
				Signature:    signature,
			},
			err: types.ErrValidatorConsAddressNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentNonce := uint64(1)
			oldConsNonce, hasConsNonce := k.GetConsensusKeyNonce(ctx, tt.valKey.GetConsAddress().String())
			if hasConsNonce {
				currentNonce = oldConsNonce.Nonce + 1
			}

			_, err := srv.SetValidatorConsAddress(wCtx, tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			validator, found := k.GetValidator(ctx, tt.valKey.Address.String())
			require.True(t, found, "validator was not saved")
			require.Equal(t, tt.valKey.Address.String(), validator.Address)
			require.Equal(t, tt.valKey.GetConsAddress().String(), validator.ConsensusAddress)

			valByConsAddr, found := k.GetValidatorByConsAddress(ctx, validator.ConsensusAddress)
			require.True(t, found, "validator by consensus address was not saved")
			require.Equal(t, tt.valKey.Address.String(), valByConsAddr.ValidatorAddress)
			require.Equal(t, tt.valKey.GetConsAddress().String(), valByConsAddr.ConsensusAddress)

			consNonce, found := k.GetConsensusKeyNonce(ctx, validator.ConsensusAddress)
			require.True(t, found, "validator consensus nonce was not saved")
			require.Equal(t, currentNonce, consNonce.Nonce)
			require.Equal(t, tt.valKey.GetConsAddress().String(), consNonce.ConsensusAddress)
		})
	}
}
