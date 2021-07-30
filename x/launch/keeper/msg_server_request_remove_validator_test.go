package keeper

import (
	"testing"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestContentRemoveValidatorCodec(t *testing.T) {
	var err error
	cdc := sample.Codec()
	request := sample.Request("foo")
	content := &types.ContentRemoveValidator{
		ValAddress: sample.AccAddress(),
	}
	request.Content, err = codec.NewAnyWithValue(content)
	require.NoError(t, err)
	result, err := request.UnpackRequestRemoveValidator(cdc)
	require.NoError(t, err)
	require.EqualValues(t, content, result)
}

func TestMsgRequestRemoveValidator(t *testing.T) {
	var (
		addr1                     = sample.AccAddress()
		addr2                     = sample.AccAddress()
		addr3                     = sample.AccAddress()
		k, _, srv, _, sdkCtx, cdc = setupMsgServer(t)
		ctx                       = sdk.WrapSDKContext(sdkCtx)
		chains                    = createNChain(k, sdkCtx, 4)
	)
	tests := []struct {
		name string
		msg  types.MsgRequestRemoveValidator
		want uint64
		err  error
	}{
		{
			name: "invalid chain",
			msg: types.MsgRequestRemoveValidator{
				ChainID: "invalid_chain",
			},
			err: sdkerrors.Wrap(types.ErrChainIdNotFound, "invalid_chain"),
		}, {
			name: "add chain 1 request 1",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[0].ChainID,
				ValidatorAddress: addr1,
			},
			want: 0,
		}, {
			name: "add chain 1 request 2",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[1].ChainID,
				ValidatorAddress: addr2,
			},
			want: 0,
		}, {
			name: "add chain 1 request 3",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[1].ChainID,
				ValidatorAddress: addr2,
			},
			want: 1,
		}, {
			name: "add chain 2 request 1",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[2].ChainID,
				ValidatorAddress: addr3,
			},
			want: 0,
		}, {
			name: "add chain 2 request 2",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[2].ChainID,
				ValidatorAddress: addr3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.RequestRemoveValidator(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			request, found := k.GetRequest(sdkCtx, tt.msg.ChainID, got.RequestID)
			require.True(t, found, "request not found")
			require.Equal(t, tt.want, request.RequestID)

			content, err := request.UnpackRequestRemoveValidator(cdc)
			require.NoError(t, err)
			require.Equal(t, tt.msg.ValidatorAddress, content.ValAddress)
		})
	}
}
