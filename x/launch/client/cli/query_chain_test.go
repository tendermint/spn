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

	"github.com/tendermint/spn/x/launch/client/cli"
	"github.com/tendermint/spn/x/launch/types"
)

func (suite *QueryTestSuite) TestShowChain() {
	ctx := suite.Network.Validators[0].ClientCtx
	chains := suite.LaunchState.ChainList

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		id   string
		args []string
		err  error
		obj  types.Chain
	}{
		{
			desc: "should show an existing chain",
			id:   fmt.Sprintf("%d", chains[0].LaunchID),
			args: common,
			obj:  chains[0],
		},
		{
			desc: "should send error for a non existing chain",
			id:   "10",
			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		suite.T().Run(tc.desc, func(t *testing.T) {
			args := []string{tc.id}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowChain(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetChainResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Chain)
				require.Equal(t, tc.obj, resp.Chain)
			}
		})
	}
}

func (suite *QueryTestSuite) TestListChain() {
	ctx := suite.Network.Validators[0].ClientCtx
	chains := suite.LaunchState.ChainList

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
	suite.T().Run("should allow listing chains by offset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(chains); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListChain(), args)
			require.NoError(t, err)
			var resp types.QueryAllChainResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Chain), step)
			require.Subset(t, chains, resp.Chain)
		}
	})
	suite.T().Run("should allow listing chains by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(chains); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListChain(), args)
			require.NoError(t, err)
			var resp types.QueryAllChainResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Chain), step)
			require.Subset(t, chains, resp.Chain)
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("should allow listing all chains", func(t *testing.T) {
		args := request(nil, 0, uint64(len(chains)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListChain(), args)
		require.NoError(t, err)
		var resp types.QueryAllChainResponse
		require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(chains), int(resp.Pagination.Total))
		require.ElementsMatch(t, chains, resp.Chain)
	})
}
