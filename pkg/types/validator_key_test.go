package types_test

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/tendermint/crypto"
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
	type args struct {
		keyJSONBytes []byte
	}
	tests := []struct {
		name      string
		args      args
		wantPvKey types.ValidatorKey
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPvKey, err := types.LoadValidatorKey(tt.args.keyJSONBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadValidatorKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPvKey, tt.wantPvKey) {
				t.Errorf("LoadValidatorKey() gotPvKey = %v, want %v", gotPvKey, tt.wantPvKey)
			}
		})
	}
}

func TestValidatorKey_GetConsAddress(t *testing.T) {
	type fields struct {
		Address crypto.Address
		PubKey  crypto.PubKey
		PrivKey crypto.PrivKey
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.ConsAddress
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := types.ValidatorKey{
				Address: tt.fields.Address,
				PubKey:  tt.fields.PubKey,
				PrivKey: tt.fields.PrivKey,
			}
			if got := v.GetConsAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConsAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatorKey_Sign(t *testing.T) {
	type fields struct {
		Address crypto.Address
		PubKey  crypto.PubKey
		PrivKey crypto.PrivKey
	}
	type args struct {
		nonce uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := types.ValidatorKey{
				Address: tt.fields.Address,
				PubKey:  tt.fields.PubKey,
				PrivKey: tt.fields.PrivKey,
			}
			got, err := v.Sign(tt.args.nonce)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sign() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatorKey_VerifySignature(t *testing.T) {
	type fields struct {
		Address crypto.Address
		PubKey  crypto.PubKey
		PrivKey crypto.PrivKey
	}
	type args struct {
		nonce uint64
		sig   []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := types.ValidatorKey{
				Address: tt.fields.Address,
				PubKey:  tt.fields.PubKey,
				PrivKey: tt.fields.PrivKey,
			}
			if got := v.VerifySignature(tt.args.nonce, tt.args.sig); got != tt.want {
				t.Errorf("VerifySignature() = %v, want %v", got, tt.want)
			}
		})
	}
}
