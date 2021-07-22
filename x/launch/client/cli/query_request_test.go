package cli_test

import (
	"fmt"
	"github.com/tendermint/spn/testutil/sample"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/testutil/network"
	"github.com/tendermint/spn/x/launch/client/cli"
	"github.com/tendermint/spn/x/launch/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithRequestObjects(t *testing.T, n int) (*network.Network, []*types.Request) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		request := sample.Request("foo")
		request.RequestID = uint64(i)
		state.RequestList = append(
			state.RequestList,
			request,
		)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.RequestList
}

func TestShowRequest(t *testing.T) {
	net, objs := networkWithRequestObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc        string
		idChainID   string
		idRequestID uint64

		args []string
		err  error
		obj  *types.Request
	}{
		{
			desc:        "found",
			idChainID:   objs[0].ChainID,
			idRequestID: objs[0].RequestID,

			args: common,
			obj:  objs[0],
		},
		{
			desc:        "not found",
			idChainID:   strconv.Itoa(100000),
			idRequestID: 100000,

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idChainID,
				strconv.Itoa(int(tc.idRequestID)),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowRequest(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetRequestResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Request)

				// Cached value is cleared when the any type is encoded into the store
				tc.obj.Content.ClearCachedValue()

				require.Equal(t, tc.obj, resp.Request)
			}
		})
	}
}

func TestListRequest(t *testing.T) {
	net, objs := networkWithRequestObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
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
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request("foo", nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRequest(), args)
			require.NoError(t, err)
			var resp types.QueryAllRequestResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			for j := i; j < len(objs) && j < i+step; j++ {
				// Cached value is cleared when the any type is encoded into the store
				objs[j].Content.ClearCachedValue()

				assert.Equal(t, objs[j], resp.Request[j-i])
			}
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request("foo", next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRequest(), args)
			require.NoError(t, err)
			var resp types.QueryAllRequestResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			for j := i; j < len(objs) && j < i+step; j++ {
				// Cached value is cleared when the any type is encoded into the store
				objs[j].Content.ClearCachedValue()

				assert.Equal(t, objs[j], resp.Request[j-i])
			}
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request("foo", nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRequest(), args)
		require.NoError(t, err)
		var resp types.QueryAllRequestResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))

		// Cached value is cleared when the any type is encoded into the store
		for _, obj := range objs {
			obj.Content.ClearCachedValue()
		}

		require.Equal(t, objs, resp.Request)
	})
}
