package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc/client/cli"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func (suite *QueryTestSuite) TestShowLaunchIDFromChannelID() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.MonitoringcState.LaunchIDsFromChannelID

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		name        string
		idChannelID string

		args []string
		err  error
		obj  types.LaunchIDFromChannelID
	}{
		{
			name:        "should allow valid query",
			idChannelID: objs[0].ChannelID,

			args: common,
			obj:  objs[0],
		},
		{
			name:        "should fail if not found",
			idChannelID: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		suite.T().Run(tc.name, func(t *testing.T) {
			args := []string{
				tc.idChannelID,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowLaunchIDFromChannelID(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetLaunchIDFromChannelIDResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.LaunchIDFromChannelID)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.LaunchIDFromChannelID),
				)
			}
		})
	}
}

func (suite *QueryTestSuite) TestListLaunchIDFromChannelID() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.MonitoringcState.LaunchIDsFromChannelID

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
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListLaunchIDFromChannelID(), args)
			require.NoError(t, err)
			var resp types.QueryAllLaunchIDFromChannelIDResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.LaunchIDFromChannelID), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.LaunchIDFromChannelID),
			)
		}
	})
	suite.T().Run("should paginate by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListLaunchIDFromChannelID(), args)
			require.NoError(t, err)
			var resp types.QueryAllLaunchIDFromChannelIDResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.LaunchIDFromChannelID), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.LaunchIDFromChannelID),
			)
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("should paginate all", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListLaunchIDFromChannelID(), args)
		require.NoError(t, err)
		var resp types.QueryAllLaunchIDFromChannelIDResponse
		require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.LaunchIDFromChannelID),
		)
	})
}
