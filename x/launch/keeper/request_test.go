package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

type RequestSample struct {
	Content types.RequestContent
	Creator string
	Status  types.Request_Status
}

func createRequestsFromSamples(
	k *keeper.Keeper,
	ctx sdk.Context,
	launchID uint64,
	samples []RequestSample,
) []types.Request {
	items := make([]types.Request, len(samples))
	for i, s := range samples {
		items[i] = sample.RequestWithContentAndCreator(r, launchID, s.Content, s.Creator)
		items[i].Status = s.Status
		id := k.AppendRequest(ctx, items[i])
		items[i].RequestID = id
	}
	return items
}

func createNRequest(k *keeper.Keeper, ctx sdk.Context, n int) []types.Request {
	items := make([]types.Request, n)
	for i := range items {
		items[i] = sample.Request(r, 0, sample.Address(r))
		id := k.AppendRequest(ctx, items[i])
		items[i].RequestID = id
	}
	return items
}

func TestRequestGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNRequest(tk.LaunchKeeper, ctx, 10)
	for _, item := range items {
		rst, found := tk.LaunchKeeper.GetRequest(ctx,
			item.LaunchID,
			item.RequestID,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestRequestRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNRequest(tk.LaunchKeeper, ctx, 10)
	for _, item := range items {
		tk.LaunchKeeper.RemoveRequest(ctx,
			item.LaunchID,
			item.RequestID,
		)
		_, found := tk.LaunchKeeper.GetRequest(ctx,
			item.LaunchID,
			item.RequestID,
		)
		require.False(t, found)
	}
}

func TestRequestGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNRequest(tk.LaunchKeeper, ctx, 10)
	require.ElementsMatch(t, items, tk.LaunchKeeper.GetAllRequest(ctx))
}

func TestRequestCounter(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNRequest(tk.LaunchKeeper, ctx, 10)
	counter := uint64(len(items)) + 1
	require.Equal(t, counter, tk.LaunchKeeper.GetRequestCounter(ctx, 0))
	require.Equal(t, uint64(1), tk.LaunchKeeper.GetRequestCounter(ctx, 1))
}

func TestCheckAccount(t *testing.T) {
	var (
		genesisAcc = sample.Address(r)
		vestingAcc = sample.Address(r)
		dupAcc     = sample.Address(r)
		notFound   = sample.Address(r)
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		launchID   = uint64(10)
	)

	ga := sample.GenesisAccount(
		r,
		launchID,
		genesisAcc,
	)
	tk.LaunchKeeper.SetGenesisAccount(ctx, ga)

	va := sample.VestingAccount(
		r,
		launchID,
		vestingAcc,
	)
	tk.LaunchKeeper.SetVestingAccount(ctx, va)

	// set duplicated entries
	ga.Address = dupAcc
	va.Address = dupAcc
	tk.LaunchKeeper.SetGenesisAccount(ctx, ga)
	tk.LaunchKeeper.SetVestingAccount(ctx, va)

	tests := []struct {
		name  string
		addr  string
		found bool
		err   error
	}{
		{
			name:  "account not found",
			addr:  notFound,
			found: false,
		}, {
			name:  "genesis account found",
			addr:  genesisAcc,
			found: true,
		}, {
			name:  "vesting account found",
			addr:  vestingAcc,
			found: true,
		}, {
			name: "critical error if duplicated accounts",
			addr: dupAcc,
			err:  spnerrors.ErrCritical,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, err := keeper.CheckAccount(ctx, *tk.LaunchKeeper, launchID, tt.addr)
			if tt.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, tt.err, err)
				return
			}

			require.Equal(t, found, tt.found)
		})
	}
}

func TestApplyRequest(t *testing.T) {
	var (
		genesisAcc     = sample.Address(r)
		vestingAcc     = sample.Address(r)
		validatorAcc   = sample.Address(r)
		ctx, tk, _     = testkeeper.NewTestSetup(t)
		launchID       = uint64(10)
		contents       = sample.AllRequestContents(r, launchID, genesisAcc, vestingAcc, validatorAcc)
		invalidContent = types.NewGenesisAccount(launchID, "", sdk.NewCoins())
	)
	tests := []struct {
		name    string
		request types.Request
		wantErr bool
	}{
		{
			name:    "test GenesisAccount content",
			request: sample.RequestWithContent(r, launchID, contents[0]),
		}, {
			name:    "test duplicated GenesisAccount content",
			request: sample.RequestWithContent(r, launchID, contents[0]),
			wantErr: true,
		}, {
			name:    "test genesis AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[1]),
		}, {
			name:    "test not found genesis AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[1]),
			wantErr: true,
		}, {
			name:    "test VestingAccount content",
			request: sample.RequestWithContent(r, launchID, contents[2]),
		}, {
			name:    "test duplicated VestingAccount content",
			request: sample.RequestWithContent(r, launchID, contents[2]),
			wantErr: true,
		}, {
			name:    "test vesting AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[3]),
		}, {
			name:    "test not found vesting AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[3]),
			wantErr: true,
		}, {
			name:    "test GenesisValidator content",
			request: sample.RequestWithContent(r, launchID, contents[4]),
		}, {
			name:    "test duplicated GenesisValidator content",
			request: sample.RequestWithContent(r, launchID, contents[4]),
			wantErr: true,
		}, {
			name:    "test ValidatorRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[5]),
		}, {
			name:    "test not found ValidatorRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[5]),
			wantErr: true,
		}, {
			name:    "test request with invalid parameters",
			request: sample.RequestWithContent(r, launchID, invalidContent),
			wantErr: true,
		},
		{
			name: "test request with no content",
			request: types.Request{
				Content: types.RequestContent{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := keeper.ApplyRequest(ctx, *tk.LaunchKeeper, launchID, tt.request)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			switch requestContent := tt.request.Content.Content.(type) {
			case *types.RequestContent_GenesisAccount:
				ga := requestContent.GenesisAccount
				_, found := tk.LaunchKeeper.GetGenesisAccount(ctx, launchID, ga.Address)
				require.True(t, found, "genesis account not found")
			case *types.RequestContent_VestingAccount:
				va := requestContent.VestingAccount
				_, found := tk.LaunchKeeper.GetVestingAccount(ctx, launchID, va.Address)
				require.True(t, found, "vesting account not found")
			case *types.RequestContent_AccountRemoval:
				ar := requestContent.AccountRemoval
				_, foundGenesis := tk.LaunchKeeper.GetGenesisAccount(ctx, launchID, ar.Address)
				require.False(t, foundGenesis, "genesis account not removed")
				_, foundVesting := tk.LaunchKeeper.GetVestingAccount(ctx, launchID, ar.Address)
				require.False(t, foundVesting, "vesting account not removed")
			case *types.RequestContent_GenesisValidator:
				ga := requestContent.GenesisValidator
				_, found := tk.LaunchKeeper.GetGenesisValidator(ctx, launchID, ga.Address)
				require.True(t, found, "genesis validator not found")
			case *types.RequestContent_ValidatorRemoval:
				vr := requestContent.ValidatorRemoval
				_, found := tk.LaunchKeeper.GetGenesisValidator(ctx, launchID, vr.ValAddress)
				require.False(t, found, "genesis validator not removed")
			}
		})
	}
}

func TestCheckRequest(t *testing.T) {
	var (
		genesisAcc                      = sample.Address(r)
		vestingAcc                      = sample.Address(r)
		validatorAcc                    = sample.Address(r)
		duplicatedAcc                   = sample.Address(r)
		ctx, tk, _                      = testkeeper.NewTestSetup(t)
		launchID                        = uint64(10)
		contents                        = sample.AllRequestContents(r, launchID, genesisAcc, vestingAcc, validatorAcc)
		invalidContent                  = types.NewGenesisAccount(launchID, "", sdk.NewCoins())
		duplicatedRequestGenesisContent = types.NewGenesisAccount(launchID, duplicatedAcc, sample.Coins(r))
		duplicatedRequestVestingContent = types.NewVestingAccount(launchID, duplicatedAcc, sample.VestingOptions(r))
		duplicatedRequestRemovalContent = types.NewAccountRemoval(duplicatedAcc)
	)

	tk.LaunchKeeper.SetGenesisAccount(ctx, types.GenesisAccount{
		LaunchID: launchID,
		Address:  duplicatedAcc,
		Coins:    nil,
	})

	tk.LaunchKeeper.SetVestingAccount(ctx, types.VestingAccount{
		LaunchID:       launchID,
		Address:        duplicatedAcc,
		VestingOptions: types.VestingOptions{},
	})

	tests := []struct {
		name    string
		request types.Request
		err     error
	}{
		{
			name:    "test GenesisAccount content",
			request: sample.RequestWithContent(r, launchID, contents[0]),
		}, {
			name:    "test duplicated GenesisAccount content",
			request: sample.RequestWithContent(r, launchID, contents[0]),
			err:     types.ErrAccountAlreadyExist,
		}, {
			name:    "test genesis AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[1]),
		}, {
			name:    "test not found genesis AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[1]),
			err:     types.ErrAccountNotFound,
		}, {
			name:    "test VestingAccount content",
			request: sample.RequestWithContent(r, launchID, contents[2]),
		}, {
			name:    "test duplicated VestingAccount content",
			request: sample.RequestWithContent(r, launchID, contents[2]),
			err:     types.ErrAccountAlreadyExist,
		}, {
			name:    "test vesting AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[3]),
		}, {
			name:    "test not found vesting AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[3]),
			err:     types.ErrAccountNotFound,
		}, {
			name:    "test GenesisValidator content",
			request: sample.RequestWithContent(r, launchID, contents[4]),
		}, {
			name:    "test duplicated GenesisValidator content",
			request: sample.RequestWithContent(r, launchID, contents[4]),
			err:     types.ErrValidatorAlreadyExist,
		}, {
			name:    "test ValidatorRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[5]),
		}, {
			name:    "test not found ValidatorRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[5]),
			err:     types.ErrValidatorNotFound,
		}, {
			name:    "test request with invalid parameters",
			request: sample.RequestWithContent(r, launchID, invalidContent),
			err:     spnerrors.ErrCritical,
		},
		{
			name:    "duplicated genesis and vesting account cause critical error",
			request: sample.RequestWithContent(r, launchID, duplicatedRequestGenesisContent),
			err:     spnerrors.ErrCritical,
		},
		{
			name:    "duplicated genesis and vesting account cause critical error",
			request: sample.RequestWithContent(r, launchID, duplicatedRequestVestingContent),
			err:     spnerrors.ErrCritical,
		},
		{
			name:    "duplicated genesis and vesting account cause critical error",
			request: sample.RequestWithContent(r, launchID, duplicatedRequestRemovalContent),
			err:     spnerrors.ErrCritical,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := keeper.CheckRequest(ctx, *tk.LaunchKeeper, launchID, tt.request)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			} else {
				err := keeper.ApplyRequest(ctx, *tk.LaunchKeeper, launchID, tt.request)
				require.NoError(t, err)
			}
		})
	}
}
