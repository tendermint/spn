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

	"github.com/tendermint/spn/x/launch/client/cli"
	"github.com/tendermint/spn/x/launch/types"
)

func (suite *QueryTestSuite) TestShowVestingAccount() {
	ctx := suite.Network.Validators[0].ClientCtx
	accs := suite.LaunchState.VestingAccountList

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc      string
		idChainID string
		idAddress string

		args []string
		err  error
		obj  types.VestingAccount
	}{
		{
			desc:      "should show an existing vesting account",
			idChainID: strconv.Itoa(int(accs[0].LaunchID)),
			idAddress: accs[0].Address,

			args: common,
			obj:  accs[0],
		},
		{
			desc:      "should send error for a non existing vesting account",
			idChainID: strconv.Itoa(100000),
			idAddress: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		suite.T().Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idChainID,
				tc.idAddress,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowVestingAccount(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetVestingAccountResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.VestingAccount)
				require.Equal(t, tc.obj, resp.VestingAccount)
			}
		})
	}
}

func (suite *QueryTestSuite) TestListVestingAccount() {
	ctx := suite.Network.Validators[0].ClientCtx
	accs := suite.LaunchState.VestingAccountList

	chainID := strconv.Itoa(int(accs[0].LaunchID))
	request := func(chainID string, next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			chainID,
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
	suite.T().Run("should allow listing vesting accounts by offset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(accs); i += step {
			args := request(chainID, nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListVestingAccount(), args)
			require.NoError(t, err)
			var resp types.QueryAllVestingAccountResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.VestingAccount), step)
			require.Subset(t, accs, resp.VestingAccount)
		}
	})
	suite.T().Run("should allow listing vesting accounts by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(accs); i += step {
			args := request(chainID, next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListVestingAccount(), args)
			require.NoError(t, err)
			var resp types.QueryAllVestingAccountResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.VestingAccount), step)
			require.Subset(t, accs, resp.VestingAccount)
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("should allow listing all vesting accounts", func(t *testing.T) {
		args := request(chainID, nil, 0, uint64(len(accs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListVestingAccount(), args)
		require.NoError(t, err)
		var resp types.QueryAllVestingAccountResponse
		require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(accs), int(resp.Pagination.Total))
		require.ElementsMatch(t, accs, resp.VestingAccount)
	})
}
