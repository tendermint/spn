package cli_test

import (
	"fmt"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/network"
	"github.com/tendermint/spn/x/campaign/client/cli"
	"github.com/tendermint/spn/x/campaign/types"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"testing"
)

func TestMaximumShares(t *testing.T) {
	// state does not need to be set because
	// maximum shares is set during InitGenesis
	net := network.New(t, network.DefaultConfig())

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		obj  uint64
	}{
		// there is no invalid request for this query
		{
			desc: "found",
			args: common,
			obj:  spntypes.TotalShareNumber,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdMaximumShares(), tc.args)
			require.NoError(t, err)
			var resp types.QueryMaximumSharesResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NotNil(t, resp.MaximumShares)
			require.Equal(t, tc.obj, resp.MaximumShares)
		})
	}
}
