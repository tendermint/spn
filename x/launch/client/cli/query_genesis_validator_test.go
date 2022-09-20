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

func (suite *QueryTestSuite) TestShowGenesisValidator() {
	ctx := suite.Network.Validators[0].ClientCtx
	accs := suite.LaunchState.GenesisValidators

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc      string
		idChainID string
		idAddress string

		args []string
		err  error
		obj  types.GenesisValidator
	}{
		{
			desc:      "should show an existing genesis validator",
			idChainID: strconv.Itoa(int(accs[0].LaunchID)),
			idAddress: accs[0].Address,

			args: common,
			obj:  accs[0],
		},
		{
			desc:      "should send error for a non existing genesis validator",
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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowGenesisValidator(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetGenesisValidatorResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.GenesisValidator)
				require.Equal(t, tc.obj, resp.GenesisValidator)
			}
		})
	}
}

func (suite *QueryTestSuite) TestListGenesisValidator() {
	ctx := suite.Network.Validators[0].ClientCtx
	accs := suite.LaunchState.GenesisValidators

	chainID := accs[0].LaunchID
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
	suite.T().Run("should allow listing genesis validators by offset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(accs); i += step {
			args := request(chainID, nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListGenesisValidator(), args)
			require.NoError(t, err)
			var resp types.QueryAllGenesisValidatorResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.GenesisValidator), step)
			require.Subset(t, accs, resp.GenesisValidator)
		}
	})
	suite.T().Run("should allow listing genesis validators by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(accs); i += step {
			args := request(chainID, next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListGenesisValidator(), args)
			require.NoError(t, err)
			var resp types.QueryAllGenesisValidatorResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.GenesisValidator), step)
			require.Subset(t, accs, resp.GenesisValidator)
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("should allow listing all genesis validators", func(t *testing.T) {
		args := request(chainID, nil, 0, uint64(len(accs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListGenesisValidator(), args)
		require.NoError(t, err)
		var resp types.QueryAllGenesisValidatorResponse
		require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(accs), int(resp.Pagination.Total))
		require.ElementsMatch(t, accs, resp.GenesisValidator)
	})
}
