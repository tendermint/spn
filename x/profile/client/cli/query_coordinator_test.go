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

	"github.com/tendermint/spn/x/profile/client/cli"
	"github.com/tendermint/spn/x/profile/types"
)

func (suite *QueryTestSuite) TestShowCoordinator() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.ProfileState.Coordinators

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		id   string
		args []string
		err  error
		obj  types.Coordinator
	}{
		{
			desc: "should show an existing coordinator",
			id:   fmt.Sprintf("%d", objs[0].CoordinatorID),
			args: common,
			obj:  objs[0],
		},
		{
			desc: "should send error for a non existing coordinator",
			id:   "not_found",
			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		suite.T().Run(tc.desc, func(t *testing.T) {
			args := []string{tc.id}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowCoordinator(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetCoordinatorResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Coordinator)
				require.Equal(t, tc.obj, resp.Coordinator)
			}
		})
	}
}

func (suite *QueryTestSuite) TestListCoordinator() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.ProfileState.Coordinators

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
	suite.T().Run("should allow listing coordinators by offset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCoordinator(), args)
			require.NoError(t, err)
			var resp types.QueryAllCoordinatorResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Coordinator), step)
			require.Subset(t, objs, resp.Coordinator)
		}
	})
	suite.T().Run("should allow listing coordinators by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCoordinator(), args)
			require.NoError(t, err)
			var resp types.QueryAllCoordinatorResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Coordinator), step)
			require.Subset(t, objs, resp.Coordinator)
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("should allow listing all coordinators", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCoordinator(), args)
		require.NoError(t, err)
		var resp types.QueryAllCoordinatorResponse
		require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t, objs, resp.Coordinator)
	})
}
