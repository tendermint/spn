package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/profile/client/cli"
	"github.com/tendermint/spn/x/profile/types"
)

func (suite *QueryTestSuite) TestShowCoordinatorByAddress() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.ProfileState.CoordinatorByAddressList

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		id   string
		args []string
		err  error
		obj  types.CoordinatorByAddress
	}{
		{
			desc: "should show an existing coordinator from address",
			id:   objs[0].Address,
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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowCoordinatorByAddress(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetCoordinatorByAddressResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.CoordinatorByAddress)
				require.Equal(t, tc.obj, resp.CoordinatorByAddress)
			}
		})
	}
}
