package cli_test

import (
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/project/client/cli"
	"github.com/tendermint/spn/x/project/types"
)

func (suite *QueryTestSuite) TestShowProject() {
	ctx := suite.Network.Validators[0].ClientCtx
	projects := suite.ProjectState.Projects

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		name string
		id   string
		args []string
		err  error
		obj  types.Project
	}{
		{
			name: "should allow valid query",
			id:   fmt.Sprintf("%d", projects[0].ProjectID),
			args: common,
			obj:  projects[0],
		},
		{
			name: "should fail if not found",
			id:   "not_found",
			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		suite.T().Run(tc.name, func(t *testing.T) {
			args := []string{tc.id}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowProject(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetProjectResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				tc.obj.Metadata = []uint8(nil)
				require.Equal(t, nullify.Fill(tc.obj), nullify.Fill(resp.Project))
			}
		})
	}
}

func (suite *QueryTestSuite) TestListProject() {
	ctx := suite.Network.Validators[0].ClientCtx
	projects := suite.ProjectState.Projects

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
	suite.T().Run("should paginate by offset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(projects); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListProject(), args)
			require.NoError(t, err)
			var resp types.QueryAllProjectResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Project), step)
			require.Subset(t, nullify.Fill(projects), nullify.Fill(resp.Project))
		}
	})
	suite.T().Run("should paginate by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(projects); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListProject(), args)
			require.NoError(t, err)
			var resp types.QueryAllProjectResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Project), step)
			require.Subset(t, nullify.Fill(projects), nullify.Fill(resp.Project))
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("should paginate all", func(t *testing.T) {
		args := request(nil, 0, uint64(len(projects)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListProject(), args)
		require.NoError(t, err)
		var resp types.QueryAllProjectResponse
		require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(projects), int(resp.Pagination.Total))
		require.ElementsMatch(t, nullify.Fill(projects), nullify.Fill(resp.Project))
	})
}
