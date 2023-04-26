package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/x/launch/client/cli"
	"github.com/tendermint/spn/x/launch/types"
)

func (suite *QueryTestSuite) TestListParamChange() {
	ctx := suite.Network.Validators[0].ClientCtx
	paramChanges := suite.LaunchState.ParamChanges

	chainID := paramChanges[0].LaunchID
	request := func(chainID uint64, next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			strconv.Itoa(int(chainID)),
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
	suite.T().Run("should allow listing param change by offset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(paramChanges); i += step {
			args := request(chainID, nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListParamChange(), args)
			require.NoError(t, err)
			var resp types.QueryAllParamChangeResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.ParamChanges), step)
			require.Subset(t, paramChanges, resp.ParamChanges)
		}
	})
	suite.T().Run("should allow listing param change by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(paramChanges); i += step {
			args := request(chainID, next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListParamChange(), args)
			require.NoError(t, err)
			var resp types.QueryAllParamChangeResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.ParamChanges), step)
			require.Subset(t, paramChanges, resp.ParamChanges)
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("should allow listing all param change", func(t *testing.T) {
		args := request(chainID, nil, 0, uint64(len(paramChanges)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListParamChange(), args)
		require.NoError(t, err)
		var resp types.QueryAllParamChangeResponse
		require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(paramChanges), int(resp.Pagination.Total))
		require.ElementsMatch(t, paramChanges, resp.ParamChanges)
	})
}
