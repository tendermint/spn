package cli_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/network"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/client/cli"
	"github.com/tendermint/spn/x/profile/types"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func networkWithConsensusKeyNonceObjects(t *testing.T, n int) (*network.Network, []types.ConsensusKeyNonce) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		consensusKeyNonce := types.ConsensusKeyNonce{
			ConsensusAddress: sample.ConsAddress(),
		}
		nullify.Fill(&consensusKeyNonce)
		state.ConsensusKeyNonceList = append(state.ConsensusKeyNonceList, consensusKeyNonce)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.ConsensusKeyNonceList
}

func TestShowConsensusKeyNonce(t *testing.T) {
	net, objs := networkWithConsensusKeyNonceObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc               string
		idConsensusAddress []byte

		args []string
		err  error
		obj  types.ConsensusKeyNonce
	}{
		{
			desc:               "found",
			idConsensusAddress: objs[0].ConsensusAddress,

			args: common,
			obj:  objs[0],
		},
		{
			desc:               "not found",
			idConsensusAddress: sample.ConsAddress(),

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				base64.StdEncoding.EncodeToString(tc.idConsensusAddress),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowConsensusKeyNonce(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetConsensusKeyNonceResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.ConsensusKeyNonce)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.ConsensusKeyNonce),
				)
			}
		})
	}
}
