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

	"github.com/tendermint/spn/x/project/client/cli"
	"github.com/tendermint/spn/x/project/types"
)

func (suite *QueryTestSuite) TestShowProjectChains() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.ProjectState.ProjectChains

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		name         string
		idProjectID uint64

		args []string
		err  error
		obj  types.ProjectChains
	}{
		{
			name:         "should allow valid query",
			idProjectID: objs[0].ProjectID,

			args: common,
			obj:  objs[0],
		},
		{
			name:         "should fail if not found",
			idProjectID: 100000,

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		suite.T().Run(tc.name, func(t *testing.T) {
			args := []string{
				strconv.Itoa(int(tc.idProjectID)),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowProjectChains(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetProjectChainsResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.ProjectChains)
				require.Equal(t, tc.obj, resp.ProjectChains)
			}
		})
	}
}
