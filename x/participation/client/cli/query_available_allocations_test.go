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
	"github.com/tendermint/spn/x/participation/client/cli"
	"github.com/tendermint/spn/x/participation/types"
)

func TestShowAvailableAllocations(t *testing.T) {
	net, totalAlloc := networkWithDelegations(t)

	ctx := net.Validators[0].ClientCtx

	delAddr := net.Validators[0].Address.String()

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc       string
		delAddress string
		args       []string
		err        error
		alloc      uint64
	}{
		{
			desc:       "found",
			delAddress: delAddr,
			args:       common,
			alloc:      totalAlloc,
		},
		{
			desc:       "not found",
			delAddress: strconv.Itoa(100000),
			args:       common,
			err:        status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.delAddress,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowAvailableAllocations(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetAvailableAllocationsResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t,
					nullify.Fill(&tc.alloc),
					nullify.Fill(&resp.AvailableAllocations),
				)
			}
		})
	}
}
