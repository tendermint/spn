package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc/client/cli"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func (suite *QueryTestSuite) TestShowVerifiedClientIds() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.MonitoringcState.VerifiedClientIDs

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc       string
		idLaunchID uint64

		args []string
		err  error
		obj  types.VerifiedClientID
	}{
		{
			desc:       "found",
			idLaunchID: objs[0].LaunchID,

			args: common,
			obj:  objs[0],
		},
		{
			desc:       "not found",
			idLaunchID: 100000,

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		suite.T().Run(tc.desc, func(t *testing.T) {
			args := []string{
				strconv.Itoa(int(tc.idLaunchID)),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowVerifiedClientIds(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetVerifiedClientIdsResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.ClientIds)
				require.Equal(t,
					nullify.Fill(&tc.obj.ClientIDs),
					nullify.Fill(&resp.ClientIds),
				)
			}
		})
	}
}
