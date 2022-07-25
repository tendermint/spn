package cli_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/claim/client/cli"
	"github.com/tendermint/spn/x/claim/types"
)

func (suite *QueryTestSuite) TestShowClaimRecord() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.ClaimState.ClaimRecords

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc    string
		address string

		args []string
		err  error
		obj  types.ClaimRecord
	}{
		{
			desc:    "found",
			address: objs[0].Address,

			args: common,
			obj:  objs[0],
		},
		{
			desc:    "not found",
			address: sample.Address(sample.Rand()),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		suite.T().Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.address,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowClaimRecord(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetClaimRecordResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.ClaimRecord)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.ClaimRecord),
				)
			}
		})
	}
}

func (suite *QueryTestSuite) TestListClaimRecord() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.ClaimState.ClaimRecords

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
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListClaimRecord(), args)
			require.NoError(t, err)
			var resp types.QueryAllClaimRecordResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.ClaimRecord), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.ClaimRecord),
			)
		}
	})
	suite.T().Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListClaimRecord(), args)
			require.NoError(t, err)
			var resp types.QueryAllClaimRecordResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.ClaimRecord), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.ClaimRecord),
			)
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListClaimRecord(), args)
		require.NoError(t, err)
		var resp types.QueryAllClaimRecordResponse
		require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.ClaimRecord),
		)
	})
}
