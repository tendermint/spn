package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/types"
)

func TestSpecialAllocationsBalance(t *testing.T) {
	var (
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)

		projectID                           = uint64(1)
		projectIDInvalidGenesisDistribution = uint64(2)
		projectIDInvalidClaimableAirdrop    = uint64(3)
	)

	tk.ProjectKeeper.SetTotalShares(ctx, 100)

	// initialize projects
	setProject := func(projectID uint64, genesisDistribution, claimableAirdrop types.Shares) {
		project := sample.Project(r, projectID)
		project.TotalSupply = tc.Coins(t, "1000foo,1000bar,1000baz")
		project.SpecialAllocations = types.NewSpecialAllocations(
			genesisDistribution,
			claimableAirdrop,
		)
		project.AllocatedShares = tc.Shares(t, "100foo,100bar,100baz,100bam")
		tk.ProjectKeeper.SetProject(ctx, project)
	}
	setProject(projectID,
		tc.Shares(t, "50foo,20bar,30bam"),
		tc.Shares(t, "50foo,100baz,40bam"),
	)
	setProject(projectIDInvalidGenesisDistribution,
		tc.Shares(t, "101foo"),
		tc.Shares(t, "50foo"),
	)
	setProject(projectIDInvalidClaimableAirdrop,
		tc.Shares(t, "50foo"),
		tc.Shares(t, "101foo"),
	)

	for _, tc := range []struct {
		desc          string
		request       *types.QuerySpecialAllocationsBalanceRequest
		response      *types.QuerySpecialAllocationsBalanceResponse
		errStatusCode codes.Code
	}{
		{
			desc:    "should fetch the balance of special allocations",
			request: &types.QuerySpecialAllocationsBalanceRequest{ProjectID: projectID},
			response: &types.QuerySpecialAllocationsBalanceResponse{
				GenesisDistribution: tc.Coins(t, "500foo,200bar"),
				ClaimableAirdrop:    tc.Coins(t, "500foo,1000baz"),
			},
		},
		{
			desc:          "should fail if project not found",
			request:       &types.QuerySpecialAllocationsBalanceRequest{ProjectID: 10000},
			errStatusCode: codes.NotFound,
		},
		{
			desc:          "should fail if genesis distribution is invalid",
			request:       &types.QuerySpecialAllocationsBalanceRequest{ProjectID: projectIDInvalidGenesisDistribution},
			errStatusCode: codes.Internal,
		},
		{
			desc:          "should fail if claimable airdrop is invalid",
			request:       &types.QuerySpecialAllocationsBalanceRequest{ProjectID: projectIDInvalidClaimableAirdrop},
			errStatusCode: codes.Internal,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.ProjectKeeper.SpecialAllocationsBalance(wctx, tc.request)
			if tc.errStatusCode != codes.OK {
				require.EqualValues(t, tc.errStatusCode, status.Code(err))
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}
