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

	"github.com/tendermint/spn/testutil/network"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc/client/cli"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func networkWithMonitoringHistoryObjects(t *testing.T, n int) (*network.Network, []types.MonitoringHistory) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		monitoringHistory := types.MonitoringHistory{
			LaunchID: uint64(i),
		}
		nullify.Fill(&monitoringHistory)
		state.MonitoringHistoryList = append(state.MonitoringHistoryList, monitoringHistory)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.MonitoringHistoryList
}

func TestShowMonitoringHistory(t *testing.T) {
	net, objs := networkWithMonitoringHistoryObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc       string
		idLaunchID uint64

		args []string
		err  error
		obj  types.MonitoringHistory
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
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				strconv.Itoa(int(tc.idLaunchID)),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowMonitoringHistory(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetMonitoringHistoryResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.MonitoringHistory)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.MonitoringHistory),
				)
			}
		})
	}
}
