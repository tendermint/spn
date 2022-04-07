package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestRequestContent_Validate(t *testing.T) {
	launchID := uint64(0)
	address := sample.Address(r)
	coins := sample.Coins(r)
	vestingOptions := sample.VestingOptions(r)
	gentTx := sample.Bytes(r, 300)
	consPubKey := sample.Bytes(r, 30)
	selfDelegation := sample.Coin(r)
	peer := sample.GenesisValidatorPeer(r)

	requestContent := types.NewGenesisAccount(launchID, address, coins)
	require.NoError(t, requestContent.Validate())

	requestContent = types.NewVestingAccount(launchID, address, vestingOptions)
	require.NoError(t, requestContent.Validate())

	requestContent = types.NewGenesisValidator(
		launchID,
		address,
		gentTx,
		consPubKey,
		selfDelegation,
		peer,
	)
	require.NoError(t, requestContent.Validate())

	requestContent = types.NewAccountRemoval(address)
	require.NoError(t, requestContent.Validate())

	requestContent = types.NewValidatorRemoval(address)
	require.NoError(t, requestContent.Validate())

}

func TestNewGenesisAccount(t *testing.T) {
	launchID := uint64(0)
	address := sample.Address(r)
	coins := sample.Coins(r)

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
	address := sample.Address(r)
	vestingOptions := sample.VestingOptions(r)
	requestContent := types.NewVestingAccount(launchID, address, vestingOptions)

	vestingAccount := requestContent.GetVestingAccount()
	require.NotNil(t, vestingAccount)
	require.EqualValues(t, launchID, vestingAccount.LaunchID)
	require.EqualValues(t, address, vestingAccount.Address)
	require.Equal(t, vestingOptions, vestingAccount.VestingOptions)

	require.Nil(t, requestContent.GetGenesisAccount())
	require.Nil(t, requestContent.GetValidatorRemoval())
	require.Nil(t, requestContent.GetAccountRemoval())
	require.Nil(t, requestContent.GetValidatorRemoval())
}

func TestNewGenesisValidator(t *testing.T) {
	launchID := uint64(0)
	address := sample.Address(r)
	gentTx := sample.Bytes(r, 300)
	consPubKey := sample.Bytes(r, 30)
	selfDelegation := sample.Coin(r)
	peer := sample.GenesisValidatorPeer(r)
	requestContent := types.NewGenesisValidator(
		launchID,
		address,
		gentTx,
		consPubKey,
		selfDelegation,
		peer,
	)

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
	address := sample.Address(r)
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
	address := sample.Address(r)
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
				Address: sample.Address(r),
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
		addr     = sample.Address(r)
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
				Coins:    sample.Coins(r),
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
				Address:  sample.Address(r),
				LaunchID: launchID,
				Coins:    sample.Coins(r),
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
		addr     = sample.Address(r)
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
				GenTx:          sample.Bytes(r, 500),
				ConsPubKey:     sample.Bytes(r, 30),
				SelfDelegation: sample.Coin(r),
				Peer:           sample.GenesisValidatorPeer(r),
			},
		},
		{
			name: "invalid address",
			content: types.GenesisValidator{
				LaunchID:       launchID,
				Address:        "invalid_address",
				GenTx:          sample.Bytes(r, 500),
				ConsPubKey:     sample.Bytes(r, 30),
				SelfDelegation: sample.Coin(r),
				Peer:           sample.GenesisValidatorPeer(r),
			},
			wantErr: true,
		},
		{
			name: "empty consensus public key",
			content: types.GenesisValidator{
				LaunchID:       launchID,
				Address:        addr,
				GenTx:          sample.Bytes(r, 500),
				ConsPubKey:     nil,
				SelfDelegation: sample.Coin(r),
				Peer:           sample.GenesisValidatorPeer(r),
			},
			wantErr: true,
		},
		{
			name: "empty gentx",
			content: types.GenesisValidator{
				LaunchID:       launchID,
				Address:        addr,
				GenTx:          nil,
				ConsPubKey:     sample.Bytes(r, 30),
				SelfDelegation: sample.Coin(r),
				Peer:           sample.GenesisValidatorPeer(r),
			},
			wantErr: true,
		},
		{
			name: "empty peer",
			content: types.GenesisValidator{
				LaunchID:       launchID,
				Address:        addr,
				GenTx:          sample.Bytes(r, 500),
				ConsPubKey:     sample.Bytes(r, 30),
				SelfDelegation: sample.Coin(r),
				Peer:           types.Peer{},
			},
			wantErr: true,
		},
		{
			name: "invalid self delegation",
			content: types.GenesisValidator{
				LaunchID:   launchID,
				Address:    addr,
				GenTx:      sample.Bytes(r, 500),
				ConsPubKey: sample.Bytes(r, 30),
				SelfDelegation: sdk.Coin{
					Denom:  "",
					Amount: sdk.NewInt(10),
				},
				Peer: sample.GenesisValidatorPeer(r),
			},
			wantErr: true,
		},
		{
			name: "zero self delegation",
			content: types.GenesisValidator{
				LaunchID:   launchID,
				Address:    addr,
				GenTx:      sample.Bytes(r, 500),
				ConsPubKey: sample.Bytes(r, 30),
				SelfDelegation: sdk.Coin{
					Denom:  "stake",
					Amount: sdk.NewInt(0),
				},
				Peer: sample.GenesisValidatorPeer(r),
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
				ValAddress: sample.Address(r),
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

	option := *types.NewDelayedVesting(coinsStr(t, "1000foo500bar"), coinsStr(t, "500foo500bar"), time.Now().Unix())

	tests := []struct {
		name    string
		content types.VestingAccount
		wantErr bool
	}{
		{
			name: "invalid address",
			content: types.VestingAccount{
				LaunchID:       launchID,
				Address:        "invalid_address",
				VestingOptions: option,
			},
			wantErr: true,
		},
		{
			name: "invalid vesting option",
			content: types.VestingAccount{
				Address:  sample.Address(r),
				LaunchID: launchID,
				VestingOptions: *types.NewDelayedVesting(
					sample.Coins(r),
					sample.Coins(r),
					0,
				),
			},
			wantErr: true,
		},
		{
			name: "valid request content",
			content: types.VestingAccount{
				Address:        sample.Address(r),
				LaunchID:       launchID,
				VestingOptions: option,
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
