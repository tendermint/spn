package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/campaign/client/cli"
	"github.com/tendermint/spn/x/campaign/types"
)

func (suite *QueryTestSuite) TestShowMainnetAccount() {
	ctx := suite.Network.Validators[0].ClientCtx
	accs := suite.CampaignState.MainnetAccountList

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc         string
		idCampaignID uint64
		idAddress    string

		args []string
		err  error
		obj  types.MainnetAccount
	}{
		{
			desc:         "found",
			idCampaignID: accs[0].CampaignID,
			idAddress:    accs[0].Address,

			args: common,
			obj:  accs[0],
		},
		{
			desc:         "not found",
			idCampaignID: 100000,
			idAddress:    strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		suite.T().Run(tc.desc, func(t *testing.T) {
			args := []string{
				strconv.Itoa(int(tc.idCampaignID)),
				tc.idAddress,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowMainnetAccount(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetMainnetAccountResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.MainnetAccount)
				require.Equal(t, tc.obj, resp.MainnetAccount)
			}
		})
	}
}

func (suite *QueryTestSuite) TestListMainnetAccount() {
	ctx := suite.Network.Validators[0].ClientCtx
	accs := suite.CampaignState.MainnetAccountList

	campaignID := accs[0].CampaignID
	request := func(campaignID uint64, next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			strconv.FormatUint(campaignID, 10),
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
		for i := 0; i < len(accs); i += step {
			args := request(campaignID, nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMainnetAccount(), args)
			require.NoError(t, err)
			var resp types.QueryAllMainnetAccountResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.MainnetAccount), step)
			require.Subset(t, accs, resp.MainnetAccount)
		}
	})
	suite.T().Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(accs); i += step {
			args := request(campaignID, next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMainnetAccount(), args)
			require.NoError(t, err)
			var resp types.QueryAllMainnetAccountResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.MainnetAccount), step)
			require.Subset(t, accs, resp.MainnetAccount)
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("Total", func(t *testing.T) {
		args := request(campaignID, nil, 0, uint64(len(accs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMainnetAccount(), args)
		require.NoError(t, err)
		var resp types.QueryAllMainnetAccountResponse
		require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(accs), int(resp.Pagination.Total))
		require.ElementsMatch(t, accs, resp.MainnetAccount)
	})
}
