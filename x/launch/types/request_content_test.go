package types_test

import (
	"testing"
	"time"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestAccountRemovalCodec(t *testing.T) {
	var err error
	cdc := sample.Codec()
	chainID, _ := sample.ChainID(1)
	request := sample.Request(chainID)
	content := &types.AccountRemoval{
		Address: sample.AccAddress(),
	}
	request.Content, err = codec.NewAnyWithValue(content)
	require.NoError(t, err)
	result, err := request.UnpackAccountRemoval(cdc)
	require.NoError(t, err)
	require.EqualValues(t, content, result)

	invalidContent := &types.Request{}
	request.Content, err = codec.NewAnyWithValue(invalidContent)
	require.NoError(t, err)
	_, err = request.UnpackAccountRemoval(cdc)
	require.Error(t, err)
}

func TestGenesisValidatorCodec(t *testing.T) {
	var err error
	cdc := sample.Codec()
	chainID, _ := sample.ChainID(1)
	request := sample.Request(chainID)
	content := &types.GenesisValidator{
		Address: sample.AccAddress(),
		ChainID: chainID,
	}
	request.Content, err = codec.NewAnyWithValue(content)
	require.NoError(t, err)
	result, err := request.UnpackGenesisValidator(cdc)

	require.NoError(t, err)
	require.EqualValues(t, content, result)

	invalidContent := &types.Request{}
	request.Content, err = codec.NewAnyWithValue(invalidContent)
	require.NoError(t, err)

	_, err = request.UnpackGenesisValidator(cdc)
	require.Error(t, err)
}

func TestValidatorRemovalCodec(t *testing.T) {
	var err error
	cdc := sample.Codec()
	chainID, _ := sample.ChainID(1)
	request := sample.Request(chainID)
	content := &types.ValidatorRemoval{
		ValAddress: sample.AccAddress(),
	}
	request.Content, err = codec.NewAnyWithValue(content)
	require.NoError(t, err)
	result, err := request.UnpackValidatorRemoval(cdc)

	require.NoError(t, err)
	require.EqualValues(t, content, result)

	invalidContent := &types.Request{}
	request.Content, err = codec.NewAnyWithValue(invalidContent)
	require.NoError(t, err)

	_, err = request.UnpackValidatorRemoval(cdc)
	require.Error(t, err)

	_, err = request.UnpackGenesisValidator(cdc)
	require.Error(t, err)

	_, err = request.UnpackValidatorRemoval(cdc)
	require.Error(t, err)
}

func TestGenesisAccountCodec(t *testing.T) {
	var err error
	cdc := sample.Codec()
	chainID, _ := sample.ChainID(1)
	request := sample.Request(chainID)
	content := &types.GenesisAccount{
		Address: sample.AccAddress(),
		ChainID: chainID,
	}
	request.Content, err = codec.NewAnyWithValue(content)
	require.NoError(t, err)
	result, err := request.UnpackGenesisAccount(cdc)
	require.NoError(t, err)
	require.EqualValues(t, content, result)

	invalidContent := &types.Request{}
	request.Content, err = codec.NewAnyWithValue(invalidContent)
	require.NoError(t, err)
	_, err = request.UnpackGenesisAccount(cdc)
	require.Error(t, err)
}

func TestVestedAccountCodec(t *testing.T) {
	var err error
	cdc := sample.Codec()
	chainID, _ := sample.ChainID(1)
	request := sample.Request(chainID)
	content := &types.VestedAccount{
		Address: sample.AccAddress(),
	}
	request.Content, err = codec.NewAnyWithValue(content)
	require.NoError(t, err)
	result, err := request.UnpackVestedAccount(cdc)
	require.NoError(t, err)
	require.EqualValues(t, content, result)
	invalidContent := &types.Request{}
	request.Content, err = codec.NewAnyWithValue(invalidContent)
	require.NoError(t, err)
	_, err = request.UnpackVestedAccount(cdc)
	require.Error(t, err)
}

func TestAccountRemoval_Validate(t *testing.T) {
	tests := []struct {
		name    string
		content types.AccountRemoval
		wantErr bool
	}{
		{
			name: "invalid address",
			content: types.AccountRemoval{
				Address: "invalid_address",
			},
			wantErr: true,
		}, {
			name: "valid content",
			content: types.AccountRemoval{
				Address: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.content.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestGenesisAccount_Validate(t *testing.T) {
	var (
		addr       = sample.AccAddress()
		chainID, _ = sample.ChainID(10)
	)
	tests := []struct {
		name    string
		content types.GenesisAccount
		wantErr bool
	}{
		{
			name: "invalid address",
			content: types.GenesisAccount{
				Address: "invalid_address",
				ChainID: chainID,
				Coins:   sample.Coins(),
			},
			wantErr: true,
		}, {
			name: "invalid chain id",
			content: types.GenesisAccount{
				Address: sample.AccAddress(),
				ChainID: "invalid_chain",
				Coins:   sample.Coins(),
			},
			wantErr: true,
		}, {
			name: "message without coins",
			content: types.GenesisAccount{
				Address: addr,
				ChainID: chainID,
				Coins:   sdk.NewCoins(),
			},
			wantErr: true,
		}, {
			name: "message with invalid coins",
			content: types.GenesisAccount{
				Address: addr,
				ChainID: chainID,
				Coins:   sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
			},
			wantErr: true,
		}, {
			name: "valid message",
			content: types.GenesisAccount{
				Address: sample.AccAddress(),
				ChainID: chainID,
				Coins:   sample.Coins(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.content.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestGenesisValidator_Validate(t *testing.T) {
	var (
		addr       = sample.AccAddress()
		chainID, _ = sample.ChainID(10)
	)
	tests := []struct {
		name    string
		content types.GenesisValidator
		wantErr bool
	}{
		{
			name: "valid message",
			content: types.GenesisValidator{
				ChainID:        chainID,
				Address:        addr,
				GenTx:          sample.Bytes(500),
				ConsPubKey:     sample.Bytes(30),
				SelfDelegation: sample.Coin(),
				Peer:           sample.String(30),
			},
		}, {
			name: "invalid address",
			content: types.GenesisValidator{
				ChainID:        chainID,
				Address:        "invalid_address",
				GenTx:          sample.Bytes(500),
				ConsPubKey:     sample.Bytes(30),
				SelfDelegation: sample.Coin(),
				Peer:           sample.String(30),
			},
			wantErr: true,
		}, {
			name: "invalid chain ID",
			content: types.GenesisValidator{
				ChainID:        "invalid_chain_id",
				Address:        addr,
				GenTx:          sample.Bytes(500),
				ConsPubKey:     sample.Bytes(30),
				SelfDelegation: sample.Coin(),
				Peer:           sample.String(30),
			},
			wantErr: true,
		}, {
			name: "empty consensus public key",
			content: types.GenesisValidator{
				ChainID:        chainID,
				Address:        addr,
				GenTx:          sample.Bytes(500),
				ConsPubKey:     nil,
				SelfDelegation: sample.Coin(),
				Peer:           sample.String(30),
			},
			wantErr: true,
		}, {
			name: "empty gentx",
			content: types.GenesisValidator{
				ChainID:        chainID,
				Address:        addr,
				GenTx:          nil,
				ConsPubKey:     sample.Bytes(30),
				SelfDelegation: sample.Coin(),
				Peer:           sample.String(30),
			},
			wantErr: true,
		}, {
			name: "empty peer",
			content: types.GenesisValidator{
				ChainID:        chainID,
				Address:        addr,
				GenTx:          sample.Bytes(500),
				ConsPubKey:     sample.Bytes(30),
				SelfDelegation: sample.Coin(),
				Peer:           "",
			},
			wantErr: true,
		}, {
			name: "invalid self delegation",
			content: types.GenesisValidator{
				ChainID:    chainID,
				Address:    addr,
				GenTx:      sample.Bytes(500),
				ConsPubKey: sample.Bytes(30),
				SelfDelegation: sdk.Coin{
					Denom:  "",
					Amount: sdk.NewInt(10),
				},
				Peer: sample.String(30),
			},
			wantErr: true,
		}, {
			name: "zero self delegation",
			content: types.GenesisValidator{
				ChainID:    chainID,
				Address:    addr,
				GenTx:      sample.Bytes(500),
				ConsPubKey: sample.Bytes(30),
				SelfDelegation: sdk.Coin{
					Denom:  "stake",
					Amount: sdk.NewInt(0),
				},
				Peer: sample.String(30),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.content.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidatorRemoval_Validate(t *testing.T) {
	tests := []struct {
		name    string
		content types.ValidatorRemoval
		wantErr bool
	}{
		{
			name: "invalid validator address",
			content: types.ValidatorRemoval{
				ValAddress: "invalid_address",
			},
			wantErr: true,
		}, {
			name: "valid message",
			content: types.ValidatorRemoval{
				ValAddress: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.content.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestVestedAccount_Validate(t *testing.T) {
	var (
		addr       = sample.AccAddress()
		chainID, _ = sample.ChainID(10)
	)

	option, err := codec.NewAnyWithValue(&types.DelayedVesting{
		Vesting: sample.Coins(),
		EndTime: time.Now().Unix(),
	})
	require.NoError(t, err)

	invalidOption, err := codec.NewAnyWithValue(&types.Request{})
	require.NoError(t, err)

	tests := []struct {
		name    string
		content types.VestedAccount
		wantErr bool
	}{
		{
			name: "invalid address",
			content: types.VestedAccount{
				ChainID:         chainID,
				Address:         "invalid_address",
				StartingBalance: sample.Coins(),
				VestingOptions:  nil,
			},
			wantErr: true,
		}, {
			name: "invalid chain id",
			content: types.VestedAccount{
				Address:         sample.AccAddress(),
				ChainID:         "invalid_chain",
				StartingBalance: sample.Coins(),
				VestingOptions:  option,
			},
			wantErr: true,
		}, {
			name: "nil message option",
			content: types.VestedAccount{
				Address:         addr,
				ChainID:         chainID,
				StartingBalance: sample.Coins(),
				VestingOptions:  nil,
			},
			wantErr: true,
		}, {
			name: "invalid coins",
			content: types.VestedAccount{
				Address:         sample.AccAddress(),
				ChainID:         chainID,
				StartingBalance: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
				VestingOptions:  option,
			},
			wantErr: true,
		}, {
			name: "invalid message option",
			content: types.VestedAccount{
				Address:         addr,
				ChainID:         chainID,
				StartingBalance: sample.Coins(),
				VestingOptions:  invalidOption,
			},
			wantErr: true,
		}, {
			name: "valid message",
			content: types.VestedAccount{
				Address:         sample.AccAddress(),
				ChainID:         chainID,
				StartingBalance: sample.Coins(),
				VestingOptions:  option,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.content.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
