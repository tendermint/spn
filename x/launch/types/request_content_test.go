package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestNewGenesisAccount(t *testing.T) {
	launchID := uint64(0)
	address := sample.Address()
	coins := sample.Coins()
	requestContent := types.NewGenesisAccount(launchID, address, coins)

	genesisAccount := requestContent.GetGenesisAccount()
	require.NotNil(t, genesisAccount)
	require.EqualValues(t, launchID, genesisAccount.LaunchID)
	require.EqualValues(t, address, genesisAccount.Address)
	require.True(t, coins.IsEqual(genesisAccount.Coins))

	require.Nil(t, requestContent.GetVestingAccount())
	require.Nil(t, requestContent.GetValidatorRemoval())
	require.Nil(t, requestContent.GetAccountRemoval())
	require.Nil(t, requestContent.GetValidatorRemoval())
}

func TestNewVestingAccount(t *testing.T) {
	launchID := uint64(0)
	address := sample.Address()
	startingBalance := sample.Coins()
	vestingOptions := sample.VestingOptions()
	requestContent := types.NewVestingAccount(launchID, address, startingBalance, vestingOptions)

	vestingAccount := requestContent.GetVestingAccount()
	require.NotNil(t, vestingAccount)
	require.EqualValues(t, launchID, vestingAccount.LaunchID)
	require.EqualValues(t, address, vestingAccount.Address)
	require.True(t, startingBalance.IsEqual(vestingAccount.StartingBalance))
	require.Equal(t, vestingOptions, vestingAccount.VestingOptions)

	require.Nil(t, requestContent.GetGenesisAccount())
	require.Nil(t, requestContent.GetValidatorRemoval())
	require.Nil(t, requestContent.GetAccountRemoval())
	require.Nil(t, requestContent.GetValidatorRemoval())
}

func TestNewGenesisValidator(t *testing.T) {
	launchID := uint64(0)
	address := sample.Address()
	gentTx := sample.Bytes(300)
	consPubKey := sample.Bytes(30)
	selfDelegation := sample.Coin()
	peer := sample.String(30)
	requestContent := types.NewGenesisValidator(launchID, address, gentTx, consPubKey, selfDelegation, peer)

	genesisValidator := requestContent.GetGenesisValidator()
	require.NotNil(t, genesisValidator)
	require.EqualValues(t, launchID, genesisValidator.LaunchID)
	require.EqualValues(t, address, genesisValidator.Address)
	require.EqualValues(t, gentTx, genesisValidator.GenTx)
	require.EqualValues(t, consPubKey, genesisValidator.ConsPubKey)
	require.True(t, selfDelegation.IsEqual(genesisValidator.SelfDelegation))
	require.EqualValues(t, peer, genesisValidator.Peer)

	require.Nil(t, requestContent.GetGenesisAccount())
	require.Nil(t, requestContent.GetVestingAccount())
	require.Nil(t, requestContent.GetAccountRemoval())
	require.Nil(t, requestContent.GetValidatorRemoval())
}

func TestNewAccountRemoval(t *testing.T) {
	address := sample.Address()
	requestContent := types.NewAccountRemoval(address)

	accountRemoval := requestContent.GetAccountRemoval()
	require.NotNil(t, accountRemoval)
	require.EqualValues(t, address, accountRemoval.Address)

	require.Nil(t, requestContent.GetGenesisAccount())
	require.Nil(t, requestContent.GetVestingAccount())
	require.Nil(t, requestContent.GetGenesisValidator())
	require.Nil(t, requestContent.GetValidatorRemoval())
}

func TestNewValidatorRemoval(t *testing.T) {
	address := sample.Address()
	requestContent := types.NewValidatorRemoval(address)

	validatorRemoval := requestContent.GetValidatorRemoval()
	require.NotNil(t, validatorRemoval)
	require.EqualValues(t, address, validatorRemoval.ValAddress)

	require.Nil(t, requestContent.GetGenesisAccount())
	require.Nil(t, requestContent.GetVestingAccount())
	require.Nil(t, requestContent.GetGenesisValidator())
	require.Nil(t, requestContent.GetAccountRemoval())
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
		},
		{
			name: "valid content",
			content: types.AccountRemoval{
				Address: sample.Address(),
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
		addr     = sample.Address()
		launchID = uint64(0)
	)
	tests := []struct {
		name    string
		content types.GenesisAccount
		wantErr bool
	}{
		{
			name: "invalid address",
			content: types.GenesisAccount{
				Address:  "invalid_address",
				LaunchID: launchID,
				Coins:    sample.Coins(),
			},
			wantErr: true,
		},
		{
			name: "request content without coins",
			content: types.GenesisAccount{
				Address:  addr,
				LaunchID: launchID,
				Coins:    sdk.NewCoins(),
			},
			wantErr: true,
		},
		{
			name: "request content with invalid coins",
			content: types.GenesisAccount{
				Address:  addr,
				LaunchID: launchID,
				Coins:    sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
			},
			wantErr: true,
		},
		{
			name: "valid request content",
			content: types.GenesisAccount{
				Address:  sample.Address(),
				LaunchID: launchID,
				Coins:    sample.Coins(),
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
		addr     = sample.Address()
		launchID = uint64(0)
	)
	tests := []struct {
		name    string
		content types.GenesisValidator
		wantErr bool
	}{
		{
			name: "valid request content",
			content: types.GenesisValidator{
				LaunchID:       launchID,
				Address:        addr,
				GenTx:          sample.Bytes(500),
				ConsPubKey:     sample.Bytes(30),
				SelfDelegation: sample.Coin(),
				Peer:           sample.String(30),
			},
		},
		{
			name: "invalid address",
			content: types.GenesisValidator{
				LaunchID:       launchID,
				Address:        "invalid_address",
				GenTx:          sample.Bytes(500),
				ConsPubKey:     sample.Bytes(30),
				SelfDelegation: sample.Coin(),
				Peer:           sample.String(30),
			},
			wantErr: true,
		},
		{
			name: "empty consensus public key",
			content: types.GenesisValidator{
				LaunchID:       launchID,
				Address:        addr,
				GenTx:          sample.Bytes(500),
				ConsPubKey:     nil,
				SelfDelegation: sample.Coin(),
				Peer:           sample.String(30),
			},
			wantErr: true,
		},
		{
			name: "empty gentx",
			content: types.GenesisValidator{
				LaunchID:       launchID,
				Address:        addr,
				GenTx:          nil,
				ConsPubKey:     sample.Bytes(30),
				SelfDelegation: sample.Coin(),
				Peer:           sample.String(30),
			},
			wantErr: true,
		},
		{
			name: "empty peer",
			content: types.GenesisValidator{
				LaunchID:       launchID,
				Address:        addr,
				GenTx:          sample.Bytes(500),
				ConsPubKey:     sample.Bytes(30),
				SelfDelegation: sample.Coin(),
				Peer:           "",
			},
			wantErr: true,
		},
		{
			name: "invalid self delegation",
			content: types.GenesisValidator{
				LaunchID:   launchID,
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
		},
		{
			name: "zero self delegation",
			content: types.GenesisValidator{
				LaunchID:   launchID,
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
		},
		{
			name: "valid request content",
			content: types.ValidatorRemoval{
				ValAddress: sample.Address(),
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

func TestVestingAccount_Validate(t *testing.T) {
	launchID := uint64(0)

	option := *types.NewDelayedVesting(sample.Coins(), time.Now().Unix())

	tests := []struct {
		name    string
		content types.VestingAccount
		wantErr bool
	}{
		{
			name: "invalid address",
			content: types.VestingAccount{
				LaunchID:        launchID,
				Address:         "invalid_address",
				StartingBalance: sample.Coins(),
				VestingOptions:  option,
			},
			wantErr: true,
		},
		{
			name: "invalid coins",
			content: types.VestingAccount{
				Address:         sample.Address(),
				LaunchID:        launchID,
				StartingBalance: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
				VestingOptions:  option,
			},
			wantErr: true,
		},
		{
			name: "valid request content",
			content: types.VestingAccount{
				Address:         sample.Address(),
				LaunchID:        launchID,
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
