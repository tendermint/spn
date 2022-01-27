package types_test

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func TestCreateSignMessage(t *testing.T) {
	tests := []struct {
		name  string
		nonce uint64
		want  []byte
	}{
		{
			name:  "with nonce and chain id",
			nonce: 10,
			want:  []byte{0x53, 0x74, 0x61, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x2f},
		},
		{
			name: "zero nonce",
			want: []byte{0x53, 0x74, 0x61, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2f},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := types.CreateSignMessage(tt.nonce)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestLoadValidatorKey(t *testing.T) {
	var (
		valKey = `{
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
		expectPrivKey = ed25519.PrivKey{0x8f, 0x8e, 0x49, 0x86, 0x70, 0x9f, 0x94, 0x49, 0x37, 0x4f, 0xa1, 0x42, 0xf0, 0xb2, 0xee, 0x26, 0xa8, 0x3d, 0xb4, 0xf7, 0xc2, 0xcc, 0x91, 0xfe, 0x51, 0x86, 0x58, 0xf3, 0xcc, 0x67, 0xfb, 0x47, 0xb1, 0x84, 0xec, 0x77, 0xb5, 0xb5, 0xf9, 0x20, 0x6d, 0x8c, 0x3d, 0xc1, 0x37, 0xf6, 0x93, 0x10, 0x31, 0x6f, 0x7d, 0x16, 0xd9, 0xf7, 0x37, 0x5f, 0xa5, 0x1, 0xf6, 0x2e, 0x4d, 0xcc, 0x44, 0xae}
	)

	tests := []struct {
		name         string
		keyJSONBytes []byte
		want         types.ValidatorKey
		err          error
	}{
		{
			name:         "valid key",
			keyJSONBytes: []byte(valKey),
			want: types.ValidatorKey{
				Address: expectPrivKey.PubKey().Address(),
				PubKey:  expectPrivKey.PubKey(),
				PrivKey: expectPrivKey,
			},
		},
		{
			name:         "invalid key",
			keyJSONBytes: sample.Bytes(100),
			want:         types.ValidatorKey{},
			err:          errors.New("error reading PrivValidator key: invalid character 'B' looking for beginning of value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPvKey, err := types.LoadValidatorKey(tt.keyJSONBytes)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, tt.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, gotPvKey)
		})
	}
}

func TestValidatorKey_GetConsAddress(t *testing.T) {
	var (
		mnemonic    = "cruel better below expect save rebuild judge drift basket stool paddle final gate behind dismiss dad stove uniform gift mean hat kit idea paddle"
		privKey     = ed25519.GenPrivKeyFromSecret([]byte(mnemonic))
		pubKey      = privKey.PubKey()
		randPrivKey = ed25519.GenPrivKey()
		randPubKey  = randPrivKey.PubKey()
	)
	tests := []struct {
		name   string
		valKey types.ValidatorKey
		want   string
	}{
		{
			name: "validator key",
			valKey: types.ValidatorKey{
				Address: pubKey.Address(),
				PubKey:  pubKey,
				PrivKey: privKey,
			},
			want: "cosmosvalcons1s80pwt3df76q68pr2srnc2qvc3gulr83caxuqe",
		},
		{
			name: "random priv key",
			valKey: types.ValidatorKey{
				Address: randPubKey.Address(),
				PubKey:  randPubKey,
				PrivKey: randPrivKey,
			},
			want: sdk.ConsAddress(randPubKey.Address()).String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.valKey.GetConsAddress()
			require.Equal(t, tt.want, got.String())
		})
	}
}

func TestValidatorKey_Sign(t *testing.T) {
	var (
		mnemonic = "cruel better below expect save rebuild judge drift basket stool paddle final gate behind dismiss dad stove uniform gift mean hat kit idea paddle"
		privKey  = ed25519.GenPrivKeyFromSecret([]byte(mnemonic))
		pubKey   = privKey.PubKey()
		valKey   = types.ValidatorKey{
			Address: pubKey.Address(),
			PubKey:  pubKey,
			PrivKey: privKey,
		}
		randPrivKey = ed25519.GenPrivKey()
		randPubKey  = randPrivKey.PubKey()
		randValKey  = types.ValidatorKey{
			Address: randPubKey.Address(),
			PubKey:  randPubKey,
			PrivKey: randPrivKey,
		}
	)
	validSig, err := valKey.Sign(10)
	require.NoError(t, err)
	validZeroNonceSig, err := valKey.Sign(0)
	require.NoError(t, err)
	randSig, err := randValKey.Sign(99)
	require.NoError(t, err)

	tests := []struct {
		name   string
		valKey types.ValidatorKey
		nonce  uint64
		want   string
		err    error
	}{
		{
			name:   "valid sign",
			valKey: valKey,
			nonce:  10,
			want:   validSig,
		},
		{
			name:   "zero nonce sign",
			valKey: valKey,
			nonce:  0,
			want:   validZeroNonceSig,
		},
		{
			name:   "random sign",
			valKey: randValKey,
			nonce:  99,
			want:   randSig,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.valKey.Sign(tt.nonce)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, tt.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestValidatorKey_VerifySignature(t *testing.T) {
	var (
		mnemonic = "cruel better below expect save rebuild judge drift basket stool paddle final gate behind dismiss dad stove uniform gift mean hat kit idea paddle"
		privKey  = ed25519.GenPrivKeyFromSecret([]byte(mnemonic))
		pubKey   = privKey.PubKey()
		valKey   = types.ValidatorKey{
			Address: pubKey.Address(),
			PubKey:  pubKey,
			PrivKey: privKey,
		}
		invalidPrivKey = ed25519.GenPrivKey()
		invalidPubKey  = privKey.PubKey()
		invalidValKey  = types.ValidatorKey{
			Address: invalidPubKey.Address(),
			PubKey:  invalidPubKey,
			PrivKey: invalidPrivKey,
		}
	)
	validSig, err := valKey.Sign(10)
	require.NoError(t, err)
	validZeroNonceSig, err := valKey.Sign(0)
	require.NoError(t, err)
	invalidSig, err := invalidValKey.Sign(0)
	require.NoError(t, err)

	tests := []struct {
		name   string
		valKey types.ValidatorKey
		nonce  uint64
		sig    string
		want   bool
	}{
		{
			name:   "valid check",
			valKey: valKey,
			sig:    validSig,
			nonce:  10,
			want:   true,
		},
		{
			name:   "zero nonce",
			valKey: valKey,
			sig:    validZeroNonceSig,
			nonce:  0,
			want:   true,
		},
		{
			name:   "random signature",
			valKey: valKey,
			sig:    sample.String(10),
			nonce:  0,
			want:   false,
		},
		{
			name:   "invalid validator key",
			valKey: invalidValKey,
			sig:    invalidSig,
			nonce:  0,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.valKey.VerifySignature(tt.nonce, tt.sig)
			require.Equal(t, tt.want, got)
		})
	}
}
