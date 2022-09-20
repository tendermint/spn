package cli_test

import (
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/campaign/client/cli"
	"github.com/tendermint/spn/x/campaign/types"
)

func (suite *QueryTestSuite) TestShowCampaign() {
	ctx := suite.Network.Validators[0].ClientCtx
	campaigns := suite.CampaignState.Campaigns

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		id   string
		args []string
		err  error
		obj  types.Campaign
	}{
		{
			desc: "found",
			id:   fmt.Sprintf("%d", campaigns[0].CampaignID),
			args: common,
			obj:  campaigns[0],
		},
		{
			desc: "not found",
			id:   "not_found",
			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		suite.T().Run(tc.desc, func(t *testing.T) {
			args := []string{tc.id}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowCampaign(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetCampaignResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				tc.obj.Metadata = []uint8(nil)
				require.Equal(t, nullify.Fill(tc.obj), nullify.Fill(resp.Campaign))
			}
		})
	}
}

func (suite *QueryTestSuite) TestListCampaign() {
	ctx := suite.Network.Validators[0].ClientCtx
	campaigns := suite.CampaignState.Campaigns

	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	suite.T().Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(campaigns); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCampaign(), args)
			require.NoError(t, err)
			var resp types.QueryAllCampaignResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Campaign), step)
			require.Subset(t, nullify.Fill(campaigns), nullify.Fill(resp.Campaign))
		}
	})
	suite.T().Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(campaigns); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCampaign(), args)
			require.NoError(t, err)
			var resp types.QueryAllCampaignResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Campaign), step)
			require.Subset(t, nullify.Fill(campaigns), nullify.Fill(resp.Campaign))
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(campaigns)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCampaign(), args)
		require.NoError(t, err)
		var resp types.QueryAllCampaignResponse
		require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(campaigns), int(resp.Pagination.Total))
		require.ElementsMatch(t, nullify.Fill(campaigns), nullify.Fill(resp.Campaign))
	})
}
