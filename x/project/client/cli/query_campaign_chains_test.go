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

func (suite *QueryTestSuite) TestShowCampaignChains() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.CampaignState.CampaignChains

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		name         string
		idCampaignID uint64

		args []string
		err  error
		obj  types.CampaignChains
	}{
		{
			name:         "should allow valid query",
			idCampaignID: objs[0].CampaignID,

			args: common,
			obj:  objs[0],
		},
		{
			name:         "should fail if not found",
			idCampaignID: 100000,

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		suite.T().Run(tc.name, func(t *testing.T) {
			args := []string{
				strconv.Itoa(int(tc.idCampaignID)),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowCampaignChains(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetCampaignChainsResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.CampaignChains)
				require.Equal(t, tc.obj, resp.CampaignChains)
			}
		})
	}
}
