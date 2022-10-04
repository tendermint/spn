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

	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/participation/client/cli"
	"github.com/tendermint/spn/x/participation/types"
)

func (suite *QueryTestSuite) TestShowAuctionUsedAllocations() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.ParticipationState.AuctionUsedAllocationsList

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		name        string
		idAddress   string
		idAuctionID uint64
		args        []string
		err         error
		obj         types.AuctionUsedAllocations
	}{
		{
			name:        "should find",
			idAddress:   objs[0].Address,
			idAuctionID: objs[0].AuctionID,
			args:        common,
			obj:         objs[0],
		},
		{
			name:        "should return not found",
			idAddress:   strconv.Itoa(100000),
			idAuctionID: 100000,
			args:        common,
			err:         status.Error(codes.NotFound, "not found"),
		},
	} {
		suite.T().Run(tc.name, func(t *testing.T) {
			args := []string{
				tc.idAddress,
				strconv.FormatUint(tc.idAuctionID, 10),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowAuctionUsedAllocations(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetAuctionUsedAllocationsResponse
				require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.AuctionUsedAllocations)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.AuctionUsedAllocations),
				)
			}
		})
	}
}

func (suite *QueryTestSuite) TestListAuctionUsedAllocations() {
	ctx := suite.Network.Validators[0].ClientCtx
	objs := suite.ParticipationState.AuctionUsedAllocationsList

	address := objs[0].Address

	request := func(addr string, next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			addr,
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
			args := request(address, nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListAuctionUsedAllocations(), args)
			require.NoError(t, err)
			var resp types.QueryAllAuctionUsedAllocationsResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.AuctionUsedAllocations), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.AuctionUsedAllocations),
			)
		}
	})
	suite.T().Run("should paginate by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(address, next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListAuctionUsedAllocations(), args)
			require.NoError(t, err)
			var resp types.QueryAllAuctionUsedAllocationsResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.AuctionUsedAllocations), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.AuctionUsedAllocations),
			)
			next = resp.Pagination.NextKey
		}
	})
	suite.T().Run("should paginate all", func(t *testing.T) {
		args := request(address, nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListAuctionUsedAllocations(), args)
		require.NoError(t, err)
		var resp types.QueryAllAuctionUsedAllocationsResponse
		require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.AuctionUsedAllocations),
		)
	})
}
