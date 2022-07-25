package simulation_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/x/mint/simulation"
)

func TestParamChangest(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	expected := []struct {
		composedKey string
		key         string
		simValue    string
		subspace    string
	}{
		{"mint/InflationRateChange", "InflationRateChange", "\"0.230000000000000000\"", "mint"},
		{"mint/InflationMax", "InflationMax", "\"0.200000000000000000\"", "mint"},
		{"mint/InflationMin", "InflationMin", "\"0.070000000000000000\"", "mint"},
		{"mint/GoalBonded", "GoalBonded", "\"0.670000000000000000\"", "mint"},
		{"mint/DistributionProportions", "DistributionProportions", "{\"staking\":\"0.250000000000000000\",\"incentives\":\"0.210000000000000000\",\"development_fund\":\"0.350000000000000000\",\"community_pool\":\"0.190000000000000000\"}", "mint"},
		{"mint/DevelopmentFundRecipients", "DevelopmentFundRecipients", "[{\"address\":\"cosmos1qqkd0nt0jnkc76qvc9fderhr8lptmx56hs6t9s\",\"weight\":\"0.009966872261695090\"},{\"address\":\"cosmos1ygg6mlfr6g8dypnxgjlpdhk5z45659xue0udth\",\"weight\":\"0.005000000000000000\"},{\"address\":\"cosmos1vgtwfm6xge308aep829gnmk5jm4vj5t5rrq0ln\",\"weight\":\"0.020000000000000000\"},{\"address\":\"cosmos1jtzgjtywnea9egav6cxnjj9hp6m6xavccvhptt\",\"weight\":\"0.007579683278078640\"},{\"address\":\"cosmos1mkn7jj5ncvek4axtydnma37nsvweshkren4hsf\",\"weight\":\"0.019381361604886090\"},{\"address\":\"cosmos19a3q9ggrldtz4x562ku5dhs8ymw94sjjlapn37\",\"weight\":\"0.008444926999667921\"},{\"address\":\"cosmos1vzhn3wun33c3x82jvfkhjn0h9a337hjpqa4h00\",\"weight\":\"0.014739408511639647\"},{\"address\":\"cosmos19mut32skaqlfw30h0g9v6eafrvxwffu6xuccl4\",\"weight\":\"0.019907312473387958\"},{\"address\":\"cosmos1gq6efy5wtny5ga2j2pkdjuphn00eks7ullnkzd\",\"weight\":\"0.014719612329779282\"},{\"address\":\"cosmos1e52yky7qxek43w3wmcwnfev3qggwexg9d0894s\",\"weight\":\"0.006553769588010266\"},{\"address\":\"cosmos1fdz2l502a6s25ess2n0f799kc6qp48m20mc0hw\",\"weight\":\"0.010760510739067874\"},{\"address\":\"cosmos1q6vqfavjmy78u2ymp4lcv67ht0709ygtcqpt5h\",\"weight\":\"0.013805657006441215\"},{\"address\":\"cosmos1ca6dwkygsg97u7uckdu409pztfmtrfa6dx4tsu\",\"weight\":\"0.020000000000000000\"},{\"address\":\"cosmos138esrz5a32jhz9xnekhyz578zwg8dhjsv5djwh\",\"weight\":\"0.006332758317578577\"},{\"address\":\"cosmos1u8lz5wywvz0eg4405n4xczukklaz3upc9hxsna\",\"weight\":\"0.016769438463264262\"},{\"address\":\"cosmos1sguazn8q2ree4qhjwp3qgkfte36g36vtfztmq8\",\"weight\":\"0.016486912293645013\"},{\"address\":\"cosmos1jd8at3lh5sy4yvp2v5nmm36zqk9yh4r3wx4qvc\",\"weight\":\"0.011221333052202303\"},{\"address\":\"cosmos1syp379pdar2s3pfsjajnefjlq00czu309s5z2m\",\"weight\":\"0.012547375725019854\"},{\"address\":\"cosmos1e9vepggfm55sn8tu89uuxezfyg4ksp247vgcpw\",\"weight\":\"0.005133490373095672\"},{\"address\":\"cosmos1hd4v0atfwsyzjp7mss5chve0v0t4ukykh294cu\",\"weight\":\"0.016310930079144450\"},{\"address\":\"cosmos1mc8ngmxdr57hwaztw3fdwl807jcvr9s7wafl9h\",\"weight\":\"0.005157396468835283\"},{\"address\":\"cosmos147jfpxlrg3t3ju4fftceaaln5rqx2qqws7fujc\",\"weight\":\"0.015186634308558717\"},{\"address\":\"cosmos1l0y4wvqdr8smgr5y7hqt68ucgs6wtelvaxrsfa\",\"weight\":\"0.015320081367030857\"},{\"address\":\"cosmos1y5r6sh5pzdtkl7dh47gla3ahc29gz6um3pd409\",\"weight\":\"0.008073123323036443\"},{\"address\":\"cosmos1xx5lqm5e443rt8f5xqmkp6nkdjte3v3udgwhh6\",\"weight\":\"0.020000000000000000\"},{\"address\":\"cosmos18cmclgv68src3r37r85czymf8ukswxvg397ts9\",\"weight\":\"0.013407917009802894\"},{\"address\":\"cosmos1shpwvemq89zc5v8fxl3hpycucpys0f4xn26g6s\",\"weight\":\"0.019716827416745468\"},{\"address\":\"cosmos1943hl84hqstxmw6la5t24zdl5swyga5zufsk7u\",\"weight\":\"0.019430003500044880\"},{\"address\":\"cosmos14k4mpq5aqp3yrh2e7trpqppu6epzpudd0lvm8h\",\"weight\":\"0.006513945114240174\"},{\"address\":\"cosmos1l57qnrftjt3gryahhp6aaftngvdm3hxhc2d3zn\",\"weight\":\"0.007391845925113882\"},{\"address\":\"cosmos1m55gmlka3s5pcpalhh67lvxd544gnq9hcy9cmq\",\"weight\":\"0.005000000000000000\"},{\"address\":\"cosmos1xq3d3wt9q4td4r8smz55dhve8e5yfdaqvu5wp9\",\"weight\":\"0.005905927546855497\"},{\"address\":\"cosmos1hpvpldg2g0vcytjxfvs6pz95ztvctyue0rscf7\",\"weight\":\"0.005000000000000000\"},{\"address\":\"cosmos1vxxq87euk5jc4nd5wdj6ytf2txaqnuw5r5fwav\",\"weight\":\"0.015577734590737776\"},{\"address\":\"cosmos1803q2f46gr0wwzp0hu9sfeh6ranuhfrd7wtazm\",\"weight\":\"0.016337747402005633\"},{\"address\":\"cosmos1faggm2cfh5mzjjzcf9c5l2v47appddtp7ft80c\",\"weight\":\"0.017649303914072022\"},{\"address\":\"cosmos1y23cgd4x9mg7rwpsyaze05kcclzy288kp9lvjx\",\"weight\":\"0.020000000000000000\"},{\"address\":\"cosmos1389c6yvaw93zs3xfdgk8f7zcza25laq0h7hyym\",\"weight\":\"0.017412338589465314\"},{\"address\":\"cosmos1lpjthgnf5g0gq2k3pp3l9g6s40zva7t7q4y70r\",\"weight\":\"0.018301886260426990\"},{\"address\":\"cosmos139fkuuseq8j538m80p87nrxzvp3upc34pfhwu3\",\"weight\":\"0.013784608707994140\"},{\"address\":\"cosmos1wzxy3xcy24jsvxkyjtpy4zsvjf9pw25rcjecps\",\"weight\":\"0.007006063945540094\"},{\"address\":\"cosmos1egm90aq80hys9tvy7m8yr24c9tvxzrrkz0fzt2\",\"weight\":\"0.005000000000000000\"},{\"address\":\"cosmos1v32ux4smhrvmf6dpfs3u3wmuvwn5aqegrmvj9h\",\"weight\":\"0.005000000000000000\"},{\"address\":\"cosmos1kkp5tjyx3m3sk64pr95stnv08a53384yjmc3da\",\"weight\":\"0.012540257524846348\"},{\"address\":\"cosmos1x4yx5q09shge3p60rtyulgrv6cgl57pyss66d0\",\"weight\":\"0.005469628158639969\"},{\"address\":\"cosmos1s8ms2aza37fnet260ehy7nvw6dw4hsegpk9d45\",\"weight\":\"0.010804000690218799\"},{\"address\":\"cosmos19xfhazc6l7jfzfzdz7ed8mxuqa046lgl2d7c8p\",\"weight\":\"0.010784393959249722\"},{\"address\":\"cosmos1zhne8apjk0axgfzx9udtd7fej3c63dc8nddssv\",\"weight\":\"0.007509725733503326\"},{\"address\":\"cosmos155g77umtus4fpgas0rtvq5xvcg839qwpgnv8kd\",\"weight\":\"0.415057255446431658\"}]", "mint"},
	}

	paramChanges := simulation.ParamChanges(r)
	require.Len(t, paramChanges, 6)

	for i, p := range paramChanges {
		require.Equal(t, expected[i].composedKey, p.ComposedKey())
		require.Equal(t, expected[i].key, p.Key())
		require.Equal(t, expected[i].simValue, p.SimValue()(r))
		require.Equal(t, expected[i].subspace, p.Subspace())
	}
}
