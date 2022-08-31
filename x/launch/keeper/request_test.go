package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/errors"
	"github.com/stretchr/testify/require"

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

	t.Run("should get a request", func(t *testing.T) {
		for _, item := range items {
			rst, found := tk.LaunchKeeper.GetRequest(ctx,
				item.LaunchID,
				item.RequestID,
			)
			require.True(t, found)
			require.Equal(t, item, rst)
		}
	})
}

func TestRequestRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNRequest(tk.LaunchKeeper, ctx, 10)

	t.Run("should remove a request", func(t *testing.T) {
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
	})
}

func TestRequestGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNRequest(tk.LaunchKeeper, ctx, 10)

	t.Run("should get all requests", func(t *testing.T) {
		require.ElementsMatch(t, items, tk.LaunchKeeper.GetAllRequest(ctx))
	})
}

func TestRequestCounter(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNRequest(tk.LaunchKeeper, ctx, 10)

	t.Run("should get request counter", func(t *testing.T) {
		counter := uint64(len(items)) + 1
		require.Equal(t, counter, tk.LaunchKeeper.GetRequestCounter(ctx, 0))
		require.Equal(t, uint64(1), tk.LaunchKeeper.GetRequestCounter(ctx, 1))
	})
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
			name:  "should return false if genesis or vesting account is not found",
			addr:  notFound,
			found: false,
		}, {
			name:  "should return true if genesis account found",
			addr:  genesisAcc,
			found: true,
		}, {
			name:  "vesting return true if account found",
			addr:  vestingAcc,
			found: true,
		}, {
			name: "should return critical error if duplicated genesis and vesting accounts",
			addr: dupAcc,
			err:  ignterrors.ErrCritical,
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
		coord          = sample.Coordinator(r, sample.Address(r))
		coordID        = uint64(3)
		genesisAcc     = sample.Address(r)
		vestingAcc     = sample.Address(r)
		validatorAcc   = sample.Address(r)
		ctx, tk, _     = testkeeper.NewTestSetup(t)
		launchID       = uint64(10)
		contents       = sample.AllRequestContents(r, launchID, genesisAcc, vestingAcc, validatorAcc)
		invalidContent = types.NewGenesisAccount(launchID, "", sdk.NewCoins())
	)

	coord.CoordinatorID = coordID
	tk.ProfileKeeper.SetCoordinator(ctx, coord)
	chain := sample.Chain(r, launchID, coordID)
	tk.LaunchKeeper.SetChain(ctx, chain)

	tests := []struct {
		name    string
		request types.Request
		wantErr bool
	}{
		{
			name:    "should allow applying GenesisAccount content",
			request: sample.RequestWithContent(r, launchID, contents[0]),
		},
		{
			name:    "should prevent applying duplicated GenesisAccount content",
			request: sample.RequestWithContent(r, launchID, contents[0]),
			wantErr: true,
		},
		{
			name:    "should allow applying genesis AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[1]),
		},
		{
			name:    "should prevent applying AccountRemoval when account not found",
			request: sample.RequestWithContent(r, launchID, contents[1]),
			wantErr: true,
		},
		{
			name:    "should allow applying VestingAccount content",
			request: sample.RequestWithContent(r, launchID, contents[2]),
		},
		{
			name:    "should prevent applying duplicated VestingAccount content",
			request: sample.RequestWithContent(r, launchID, contents[2]),
			wantErr: true,
		},
		{
			name:    "should allow applying vesting AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[3]),
		},
		{
			name:    "should prevent applying vesting AccountRemoval content when account not found",
			request: sample.RequestWithContent(r, launchID, contents[3]),
			wantErr: true,
		},
		{
			name:    "should allow applying GenesisValidator content",
			request: sample.RequestWithContent(r, launchID, contents[4]),
		},
		{
			name:    "should prevent applying duplicated GenesisValidator content",
			request: sample.RequestWithContent(r, launchID, contents[4]),
			wantErr: true,
		},
		{
			name:    "should allow applying ValidatorRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[5]),
		},
		{
			name:    "should prevent applying ValidatorRemoval when validator not found",
			request: sample.RequestWithContent(r, launchID, contents[5]),
			wantErr: true,
		},
		{
			name:    "should prevent applying invalid request content",
			request: sample.RequestWithContent(r, launchID, invalidContent),
			wantErr: true,
		},
		{
			name: "should prevent applying empty request content",
			request: types.Request{
				Content: types.RequestContent{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := keeper.ApplyRequest(ctx, *tk.LaunchKeeper, chain, tt.request, coord)
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
		coord                           = sample.Coordinator(r, sample.Address(r))
		coordID                         = uint64(3)
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

	coord.CoordinatorID = coordID
	tk.ProfileKeeper.SetCoordinator(ctx, coord)
	chain := sample.Chain(r, launchID, coordID)
	tk.LaunchKeeper.SetChain(ctx, chain)

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
			name:    "should validate valid GenesisAccount content",
			request: sample.RequestWithContent(r, launchID, contents[0]),
		},
		{
			name:    "should prevent validate duplicated GenesisAccount content",
			request: sample.RequestWithContent(r, launchID, contents[0]),
			err:     types.ErrAccountAlreadyExist,
		},
		{
			name:    "should validate valid genesis AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[1]),
		},
		{
			name:    "should prevent validate AccountRemoval with no account",
			request: sample.RequestWithContent(r, launchID, contents[1]),
			err:     types.ErrAccountNotFound,
		},
		{
			name:    "should validate valid VestingAccount content",
			request: sample.RequestWithContent(r, launchID, contents[2]),
		},
		{
			name:    "should prevent validate duplicated VestingAccount content",
			request: sample.RequestWithContent(r, launchID, contents[2]),
			err:     types.ErrAccountAlreadyExist,
		},
		{
			name:    "should validate valid vesting AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[3]),
		},
		{
			name:    "should validate vesting AccountRemoval content with vesting account not found",
			request: sample.RequestWithContent(r, launchID, contents[3]),
			err:     types.ErrAccountNotFound,
		},
		{
			name:    "should validate valid GenesisValidator content",
			request: sample.RequestWithContent(r, launchID, contents[4]),
		},
		{
			name:    "should prevent validate duplicated GenesisValidator content",
			request: sample.RequestWithContent(r, launchID, contents[4]),
			err:     types.ErrValidatorAlreadyExist,
		},
		{
			name:    "should validate valid ValidatorRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[5]),
		},
		{
			name:    "should prevent validate ValidatorRemoval content with no validator to remove",
			request: sample.RequestWithContent(r, launchID, contents[5]),
			err:     types.ErrValidatorNotFound,
		},
		{
			name:    "should prevent validate request content with invalid parameters",
			request: sample.RequestWithContent(r, launchID, invalidContent),
			err:     ignterrors.ErrCritical,
		},
		{
			name:    "should prevent validate with critical error genesis account request content with genesis and vesting account",
			request: sample.RequestWithContent(r, launchID, duplicatedRequestGenesisContent),
			err:     ignterrors.ErrCritical,
		},
		{
			name:    "should prevent validate with critical error vesting account request content with genesis and vesting account",
			request: sample.RequestWithContent(r, launchID, duplicatedRequestVestingContent),
			err:     ignterrors.ErrCritical,
		},
		{
			name:    "should prevent validate with critical error account removal request content with genesis and vesting account",
			request: sample.RequestWithContent(r, launchID, duplicatedRequestRemovalContent),
			err:     ignterrors.ErrCritical,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := keeper.CheckRequest(ctx, *tk.LaunchKeeper, launchID, tt.request)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			} else {
				err := keeper.ApplyRequest(ctx, *tk.LaunchKeeper, chain, tt.request, coord)
				require.NoError(t, err)
			}
		})
	}
}
