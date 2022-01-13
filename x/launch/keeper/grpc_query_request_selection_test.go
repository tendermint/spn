package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestQueryRequestSelection(t *testing.T) {
	var (
		coordinator1              = sample.Coordinator(sample.Address())
		k, pk, _, _, _, _, sdkCtx = setupMsgServer(t)
		ctx                       = sdk.WrapSDKContext(sdkCtx)
	)

	coordinator1.CoordinatorID = pk.AppendCoordinator(sdkCtx, coordinator1)
	chains := createNChainForCoordinator(k, sdkCtx, coordinator1.CoordinatorID, 1)

	requestSamples := make([]RequestSample, 6)
	for i := 0; i < 6; i++ {
		addr := sample.Address()
		requestSamples[i] = RequestSample{
			Content: sample.GenesisAccountContent(chains[0].LaunchID, addr),
			Creator: addr,
		}
	}
	_ = createRequestsFromSamples(k, sdkCtx, chains[0].LaunchID, requestSamples)

	tests := []struct {
		name           string
		req            types.QueryRequestSelectionRequest
		resultSequence []uint64
		err            error
	}{
		{
			name: "invalid input",
			req: types.QueryRequestSelectionRequest{
				LaunchID:   0,
				RequestIDs: "invalid,,",
			},
			err: status.Error(codes.InvalidArgument, "invalid argument"),
		},
		{
			name: "whole selection exists",
			req: types.QueryRequestSelectionRequest{
				LaunchID:   0,
				RequestIDs: "1,2",
			},
			resultSequence: []uint64{1, 2},
		},
		{
			name: "selection partially exists",
			req: types.QueryRequestSelectionRequest{
				LaunchID:   0,
				RequestIDs: "1-7",
			},
			resultSequence: []uint64{1, 2, 3, 4, 5, 6},
		},
		{
			name: "selection doesn't exist",
			req: types.QueryRequestSelectionRequest{
				LaunchID:   0,
				RequestIDs: "7-9",
			},
			resultSequence: []uint64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requests, err := k.RequestSelection(ctx, &tt.req)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, len(tt.resultSequence), len(requests.Request))

			for i, req := range requests.Request {
				if req.RequestID != tt.resultSequence[i] {
					t.Fail()
				}
			}
		})
	}
}
