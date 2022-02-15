package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func setupMsgServer(t testing.TB) (
	*keeper.Keeper,
	*launchkeeper.Keeper,
	*profilekeeper.Keeper,
	bankkeeper.Keeper,
	types.MsgServer,
	profiletypes.MsgServer,
	launchtypes.MsgServer,
	sdk.Context,
) {
	_, launchKeeper, profileKeeper, rewardKeeper, _, bankKeeper, _, ctx := testkeeper.AllKeepers(t)

	return rewardKeeper,
		launchKeeper,
		profileKeeper,
		bankKeeper,
		keeper.NewMsgServerImpl(*rewardKeeper),
		profilekeeper.NewMsgServerImpl(*profileKeeper),
		launchkeeper.NewMsgServerImpl(*launchKeeper),
		ctx
}

// coinsFromString returns a sdk.Coins from a string
func coinsFromString(t testing.TB, str string) sdk.Coins {
	coins, err := sdk.ParseCoinsNormalized(str)
	require.NoError(t, err)
	return coins
}

// decFromString returns a sdk.Dec from a string
func decFromString(t testing.TB, str string) sdk.Dec {
	dec, err := sdk.NewDecFromStr(str)
	require.NoError(t, err)
	return dec
}

// signatureCount returns a signature count object for test from a cons address and a decimal string for relative signatures
func signatureCount(t *testing.T, consAddr []byte, relSig string) spntypes.SignatureCount {
	return spntypes.SignatureCount{
		ConsAddress:        consAddr,
		RelativeSignatures: decFromString(t, relSig),
	}
}

// signatureCounts returns a signature counts object for tests from a a block count and list of signature counts
func signatureCounts(blockCount uint64, sc ...spntypes.SignatureCount) spntypes.SignatureCounts {
	return spntypes.SignatureCounts{
		BlockCount: blockCount,
		Counts:     sc,
	}
}
