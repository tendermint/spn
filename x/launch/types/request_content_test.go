package types_test

import (
	"errors"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
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

	t.Run("should validate request with valid genesis account", func(t *testing.T) {
		requestContent := types.NewGenesisAccount(launchID, address, coins)
		require.NoError(t, requestContent.Validate())
	})

	t.Run("should validate request with valid vesting account", func(t *testing.T) {
		requestContent := types.NewVestingAccount(launchID, address, vestingOptions)
		require.NoError(t, requestContent.Validate())
	})

	t.Run("should validate request with valid genesis validator", func(t *testing.T) {
		requestContent := types.NewGenesisValidator(
			launchID,
			address,
			gentTx,
			consPubKey,
			selfDelegation,
			peer,
		)
		require.NoError(t, requestContent.Validate())
	})

	t.Run("should validate request with valid account removal", func(t *testing.T) {
		requestContent := types.NewAccountRemoval(address)
		require.NoError(t, requestContent.Validate())
	})

	t.Run("should validate request with valid validator removal", func(t *testing.T) {
		requestContent := types.NewValidatorRemoval(address)
		require.NoError(t, requestContent.Validate())
	})

	t.Run("should prevent validate request with unrecognized content", func(t *testing.T) {
		// request with no content
		requestContent := types.RequestContent{}
		require.Equal(t, requestContent.Validate(), errors.New("unrecognized request content"))
	})
}

func TestNewGenesisAccount(t *testing.T) {
	launchID := uint64(0)
	address := sample.Address(r)
	coins := sample.Coins(r)

	t.Run("should create a new genesis account", func(t *testing.T) {
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
	})
}

func TestNewVestingAccount(t *testing.T) {
	launchID := uint64(0)
	address := sample.Address(r)
	vestingOptions := sample.VestingOptions(r)

	t.Run("should create a new vesting account", func(t *testing.T) {
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
	})
}

func TestNewGenesisValidator(t *testing.T) {
	launchID := uint64(0)
	address := sample.Address(r)
	gentTx := sample.Bytes(r, 300)
	consPubKey := sample.Bytes(r, 30)
	selfDelegation := sample.Coin(r)
	peer := sample.GenesisValidatorPeer(r)

	t.Run("should create a new genesis validator", func(t *testing.T) {
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
	})
}

func TestNewAccountRemoval(t *testing.T) {
	address := sample.Address(r)
	requestContent := types.NewAccountRemoval(address)

	t.Run("should create a new account removal", func(t *testing.T) {
		accountRemoval := requestContent.GetAccountRemoval()
		require.NotNil(t, accountRemoval)
		require.EqualValues(t, address, accountRemoval.Address)

		require.Nil(t, requestContent.GetGenesisAccount())
		require.Nil(t, requestContent.GetVestingAccount())
		require.Nil(t, requestContent.GetGenesisValidator())
		require.Nil(t, requestContent.GetValidatorRemoval())
	})
}

func TestNewValidatorRemoval(t *testing.T) {
	address := sample.Address(r)
	requestContent := types.NewValidatorRemoval(address)

	t.Run("should create a new validator removal", func(t *testing.T) {
		validatorRemoval := requestContent.GetValidatorRemoval()
		require.NotNil(t, validatorRemoval)
		require.EqualValues(t, address, validatorRemoval.ValAddress)

		require.Nil(t, requestContent.GetGenesisAccount())
		require.Nil(t, requestContent.GetVestingAccount())
		require.Nil(t, requestContent.GetGenesisValidator())
		require.Nil(t, requestContent.GetAccountRemoval())
	})
}

func TestAccountRemoval_Validate(t *testing.T) {
	tests := []struct {
		name    string
		content types.AccountRemoval
		wantErr bool
	}{
		{
			name: "should prevent validate account removal with invalid address",
			content: types.AccountRemoval{
				Address: "invalid_address",
			},
			wantErr: true,
		},
		{
			name: "should validate valid account removal",
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
			name: "should prevent validate genesis account with invalid address",
			content: types.GenesisAccount{
				Address:  "invalid_address",
				LaunchID: launchID,
				Coins:    sample.Coins(r),
			},
			wantErr: true,
		},
		{
			name: "should prevent validate genesis account without coins",
			content: types.GenesisAccount{
				Address:  addr,
				LaunchID: launchID,
				Coins:    sdk.NewCoins(),
			},
			wantErr: true,
		},
		{
			name: "should prevent validate genesis account with invalid coins",
			content: types.GenesisAccount{
				Address:  addr,
				LaunchID: launchID,
				Coins:    sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
			},
			wantErr: true,
		},
		{
			name: "should validate valid genesis acocunt",
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
			name: "should validate valid genesis validator",
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
			name: "should prevent validate genesis validator with invalid address",
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
			name: "should prevent validate genesis validator with empty consensus public key",
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
			name: "should prevent validate genesis validator with empty gentx",
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
			name: "should prevent validate genesis validator with empty peer",
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
			name: "should prevent validate genesis validator with invalid self delegation",
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
			name: "should prevent validate genesis validator with zero self delegation",
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
			name: "should prevent validate validator removal with invalid validator address",
			content: types.ValidatorRemoval{
				ValAddress: "invalid_address",
			},
			wantErr: true,
		},
		{
			name: "should validate valid validator removal",
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

	option := *types.NewDelayedVesting(
		tc.Coins(t, "1000foo500bar"),
		tc.Coins(t, "500foo500bar"),
		time.Now().Unix(),
	)

	tests := []struct {
		name    string
		content types.VestingAccount
		wantErr bool
	}{
		{
			name: "should prevent validate vesting account with invalid address",
			content: types.VestingAccount{
				LaunchID:       launchID,
				Address:        "invalid_address",
				VestingOptions: option,
			},
			wantErr: true,
		},
		{
			name: "should prevent validate vesting account with invalid vesting option",
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
			name: "should validate valid vesting account",
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
