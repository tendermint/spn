package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/x/mint/simulation"
	"github.com/tendermint/spn/x/mint/types"
)

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abonormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: sdkmath.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	var mintGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &mintGenesis)

	var (
		dec1 = sdk.MustNewDecFromStr("0.670000000000000000")
		dec2 = sdk.MustNewDecFromStr("0.200000000000000000")
		dec3 = sdk.MustNewDecFromStr("0.070000000000000000")
		dec4 = sdk.MustNewDecFromStr("0.170000000000000000")
		dec5 = sdk.MustNewDecFromStr("0.700000000000000000")
		dec6 = sdk.MustNewDecFromStr("0.060000000000000000")
		dec7 = sdk.MustNewDecFromStr("0.070000000000000000")
	)

	weightedAddresses := []types.WeightedAddress{
		{
			Address: "cosmos1jtzgjtywnea9egav6cxnjj9hp6m6xavccvhptt",
			Weight:  sdk.MustNewDecFromStr("0.007579683278078640"),
		},
		{
			Address: "cosmos1mkn7jj5ncvek4axtydnma37nsvweshkren4hsf",
			Weight:  sdk.MustNewDecFromStr("0.019381361604886090"),
		},
		{
			Address: "cosmos19a3q9ggrldtz4x562ku5dhs8ymw94sjjlapn37",
			Weight:  sdk.MustNewDecFromStr("0.008444926999667921"),
		},
		{
			Address: "cosmos1vzhn3wun33c3x82jvfkhjn0h9a337hjpqa4h00",
			Weight:  sdk.MustNewDecFromStr("0.014739408511639647"),
		},
		{
			Address: "cosmos19mut32skaqlfw30h0g9v6eafrvxwffu6xuccl4",
			Weight:  sdk.MustNewDecFromStr("0.019907312473387958"),
		},
		{
			Address: "cosmos1gq6efy5wtny5ga2j2pkdjuphn00eks7ullnkzd",
			Weight:  sdk.MustNewDecFromStr("0.014719612329779282"),
		},
		{
			Address: "cosmos1e52yky7qxek43w3wmcwnfev3qggwexg9d0894s",
			Weight:  sdk.MustNewDecFromStr("0.006553769588010266"),
		},
		{
			Address: "cosmos1fdz2l502a6s25ess2n0f799kc6qp48m20mc0hw",
			Weight:  sdk.MustNewDecFromStr("0.010760510739067874"),
		},
		{
			Address: "cosmos1q6vqfavjmy78u2ymp4lcv67ht0709ygtcqpt5h",
			Weight:  sdk.MustNewDecFromStr("0.013805657006441215"),
		},
		{
			Address: "cosmos1ca6dwkygsg97u7uckdu409pztfmtrfa6dx4tsu",
			Weight:  sdk.MustNewDecFromStr("0.020000000000000000"),
		},
		{
			Address: "cosmos138esrz5a32jhz9xnekhyz578zwg8dhjsv5djwh",
			Weight:  sdk.MustNewDecFromStr("0.006332758317578577"),
		},
		{
			Address: "cosmos1u8lz5wywvz0eg4405n4xczukklaz3upc9hxsna",
			Weight:  sdk.MustNewDecFromStr("0.016769438463264262"),
		},
		{
			Address: "cosmos1sguazn8q2ree4qhjwp3qgkfte36g36vtfztmq8",
			Weight:  sdk.MustNewDecFromStr("0.016486912293645013"),
		},
		{
			Address: "cosmos1jd8at3lh5sy4yvp2v5nmm36zqk9yh4r3wx4qvc",
			Weight:  sdk.MustNewDecFromStr("0.011221333052202303"),
		},
		{
			Address: "cosmos1syp379pdar2s3pfsjajnefjlq00czu309s5z2m",
			Weight:  sdk.MustNewDecFromStr("0.012547375725019854"),
		},
		{
			Address: "cosmos1e9vepggfm55sn8tu89uuxezfyg4ksp247vgcpw",
			Weight:  sdk.MustNewDecFromStr("0.005133490373095672"),
		},
		{
			Address: "cosmos1hd4v0atfwsyzjp7mss5chve0v0t4ukykh294cu",
			Weight:  sdk.MustNewDecFromStr("0.016310930079144450"),
		},
		{
			Address: "cosmos1mc8ngmxdr57hwaztw3fdwl807jcvr9s7wafl9h",
			Weight:  sdk.MustNewDecFromStr("0.005157396468835283"),
		},
		{
			Address: "cosmos147jfpxlrg3t3ju4fftceaaln5rqx2qqws7fujc",
			Weight:  sdk.MustNewDecFromStr("0.015186634308558717"),
		},
		{
			Address: "cosmos1l0y4wvqdr8smgr5y7hqt68ucgs6wtelvaxrsfa",
			Weight:  sdk.MustNewDecFromStr("0.015320081367030857"),
		},
		{
			Address: "cosmos1y5r6sh5pzdtkl7dh47gla3ahc29gz6um3pd409",
			Weight:  sdk.MustNewDecFromStr("0.008073123323036443"),
		},
		{
			Address: "cosmos1xx5lqm5e443rt8f5xqmkp6nkdjte3v3udgwhh6",
			Weight:  sdk.MustNewDecFromStr("0.020000000000000000"),
		},
		{
			Address: "cosmos18cmclgv68src3r37r85czymf8ukswxvg397ts9",
			Weight:  sdk.MustNewDecFromStr("0.013407917009802894"),
		},
		{
			Address: "cosmos1shpwvemq89zc5v8fxl3hpycucpys0f4xn26g6s",
			Weight:  sdk.MustNewDecFromStr("0.019716827416745468"),
		},
		{
			Address: "cosmos1943hl84hqstxmw6la5t24zdl5swyga5zufsk7u",
			Weight:  sdk.MustNewDecFromStr("0.019430003500044880"),
		},
		{
			Address: "cosmos14k4mpq5aqp3yrh2e7trpqppu6epzpudd0lvm8h",
			Weight:  sdk.MustNewDecFromStr("0.006513945114240174"),
		},
		{
			Address: "cosmos1l57qnrftjt3gryahhp6aaftngvdm3hxhc2d3zn",
			Weight:  sdk.MustNewDecFromStr("0.007391845925113882"),
		},
		{
			Address: "cosmos1m55gmlka3s5pcpalhh67lvxd544gnq9hcy9cmq",
			Weight:  sdk.MustNewDecFromStr("0.005000000000000000"),
		},
		{
			Address: "cosmos1xq3d3wt9q4td4r8smz55dhve8e5yfdaqvu5wp9",
			Weight:  sdk.MustNewDecFromStr("0.005905927546855497"),
		},
		{
			Address: "cosmos1hpvpldg2g0vcytjxfvs6pz95ztvctyue0rscf7",
			Weight:  sdk.MustNewDecFromStr("0.005000000000000000"),
		},
		{
			Address: "cosmos1vxxq87euk5jc4nd5wdj6ytf2txaqnuw5r5fwav",
			Weight:  sdk.MustNewDecFromStr("0.015577734590737776"),
		},
		{
			Address: "cosmos1803q2f46gr0wwzp0hu9sfeh6ranuhfrd7wtazm",
			Weight:  sdk.MustNewDecFromStr("0.016337747402005633"),
		},
		{
			Address: "cosmos1faggm2cfh5mzjjzcf9c5l2v47appddtp7ft80c",
			Weight:  sdk.MustNewDecFromStr("0.017649303914072022"),
		},
		{
			Address: "cosmos1y23cgd4x9mg7rwpsyaze05kcclzy288kp9lvjx",
			Weight:  sdk.MustNewDecFromStr("0.020000000000000000"),
		},
		{
			Address: "cosmos1389c6yvaw93zs3xfdgk8f7zcza25laq0h7hyym",
			Weight:  sdk.MustNewDecFromStr("0.017412338589465314"),
		},
		{
			Address: "cosmos1lpjthgnf5g0gq2k3pp3l9g6s40zva7t7q4y70r",
			Weight:  sdk.MustNewDecFromStr("0.018301886260426990"),
		},
		{
			Address: "cosmos139fkuuseq8j538m80p87nrxzvp3upc34pfhwu3",
			Weight:  sdk.MustNewDecFromStr("0.013784608707994140"),
		},
		{
			Address: "cosmos1wzxy3xcy24jsvxkyjtpy4zsvjf9pw25rcjecps",
			Weight:  sdk.MustNewDecFromStr("0.007006063945540094"),
		},
		{
			Address: "cosmos1egm90aq80hys9tvy7m8yr24c9tvxzrrkz0fzt2",
			Weight:  sdk.MustNewDecFromStr("0.005000000000000000"),
		},
		{
			Address: "cosmos1v32ux4smhrvmf6dpfs3u3wmuvwn5aqegrmvj9h",
			Weight:  sdk.MustNewDecFromStr("0.005000000000000000"),
		},
		{
			Address: "cosmos1kkp5tjyx3m3sk64pr95stnv08a53384yjmc3da",
			Weight:  sdk.MustNewDecFromStr("0.012540257524846348"),
		},
		{
			Address: "cosmos1x4yx5q09shge3p60rtyulgrv6cgl57pyss66d0",
			Weight:  sdk.MustNewDecFromStr("0.484591876249738564"),
		},
	}

	require.Equal(t, uint64(6311520), mintGenesis.Params.BlocksPerYear)
	require.Equal(t, dec1, mintGenesis.Params.GoalBonded)
	require.Equal(t, dec2, mintGenesis.Params.InflationMax)
	require.Equal(t, dec3, mintGenesis.Params.InflationMin)
	require.Equal(t, "stake", mintGenesis.Params.MintDenom)
	require.Equal(t, dec4, mintGenesis.Params.DistributionProportions.Staking)
	require.Equal(t, dec5, mintGenesis.Params.DistributionProportions.Incentives)
	require.Equal(t, dec6, mintGenesis.Params.DistributionProportions.DevelopmentFund)
	require.Equal(t, dec7, mintGenesis.Params.DistributionProportions.CommunityPool)
	require.Equal(t, "0stake", mintGenesis.Minter.BlockProvision(mintGenesis.Params).String())
	require.Equal(t, "0.170000000000000000", mintGenesis.Minter.NextAnnualProvisions(mintGenesis.Params, sdkmath.OneInt()).String())
	require.Equal(t, "0.169999926644441493", mintGenesis.Minter.NextInflationRate(mintGenesis.Params, sdk.OneDec()).String())
	require.Equal(t, "0.170000000000000000", mintGenesis.Minter.Inflation.String())
	require.Equal(t, "0.000000000000000000", mintGenesis.Minter.AnnualProvisions.String())
	require.Equal(t, weightedAddresses, mintGenesis.Params.DevelopmentFundRecipients)
}

// TestRandomizedGenState tests abnormal scenarios of applying RandomizedGenState.
func TestRandomizedGenState1(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)
	// all these tests will panic
	tests := []struct {
		simState module.SimulationState
		panicMsg string
	}{
		{ // panic => reason: incomplete initialization of the simState
			module.SimulationState{}, "invalid memory address or nil pointer dereference"},
		{ // panic => reason: incomplete initialization of the simState
			module.SimulationState{
				AppParams: make(simtypes.AppParams),
				Cdc:       cdc,
				Rand:      r,
			}, "assignment to entry in nil map"},
	}

	for _, tt := range tests {
		require.Panicsf(t, func() { simulation.RandomizedGenState(&tt.simState) }, tt.panicMsg)
	}
}
