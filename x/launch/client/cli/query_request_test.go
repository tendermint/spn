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

	"github.com/tendermint/spn/x/launch/client/cli"
	"github.com/tendermint/spn/x/launch/types"
)

func (suite *QueryTestSuite) TestShowRequest() {
	ctx := suite.Network.Validators[0].ClientCtx
	requests := suite.LaunchState.Requests

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc        string
		idLaunchID  string
		idRequestID uint64

		args []string
		err  error
		obj  types.Request
	}{
		{
			desc:        "should show an existing request",
			idLaunchID:  strconv.Itoa(int(requests[0].LaunchID)),
			idRequestID: requests[0].RequestID,

			args: common,
			obj:  requests[0],
		},
		{
			desc:        "should send error for a non existing request",
			idLaunchID:  strconv.Itoa(100000),
			idRequestID: 100000,

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		suite.T().Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idLaunchID,
				strconv.Itoa(int(tc.idRequestID)),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowRequest(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetRequestResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Request)
				require.Equal(t, tc.obj, resp.Request)
			}
		})
	}
}

func (suite *QueryTestSuite) TestListRequest() {
	ctx := suite.Network.Validators[0].ClientCtx
	requests := suite.LaunchState.Requests

	request := func(launchID string, next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			launchID,
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
	suite.T().Run("should allow listing requests by offset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(requests); i += step {
			args := request("0", nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRequest(), args)
			require.NoError(t, err)
			var resp types.QueryAllRequestResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Request), step)
			require.Subset(t, requests, resp.Request)
		}
	})
	suite.T().Run("should allow listing requests by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(requests); i += step {
			args := request("0", next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRequest(), args)
			require.NoError(t, err)
			var resp types.QueryAllRequestResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Request), step)
			require.Subset(t, requests, resp.Request)
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("should allow listing all requests", func(t *testing.T) {
		args := request("0", nil, 0, uint64(len(requests)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRequest(), args)
		require.NoError(t, err)
		var resp types.QueryAllRequestResponse
		require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(requests), int(resp.Pagination.Total))
		require.ElementsMatch(t, requests, resp.Request)
	})
}
