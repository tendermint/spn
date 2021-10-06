package simapp

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/spm/cosmoscmd"
	"github.com/tendermint/spn/app"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"
)

// New creates application instance with in-memory database and disabled logging.
func New(db tmdb.DB, dir string, logger log.Logger) app.SPNApp {
	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)

	a := app.New(
		logger,
		db,
		nil,
		false,
		map[int64]bool{},
		dir,
		0,
		encoding,
		simapp.EmptyAppOptions{},
	)
	// InitChain updates deliverState which is required when app.NewContext is called
	//a.InitChain(abci.RequestInitChain{
	//	ConsensusParams: defaultConsensusParams,
	//	AppStateBytes:   []byte("{}"),
	//})

	spnApp, ok := a.(app.SPNApp)
	if !ok {
		panic("AAHHH")
	}

	return spnApp
}

//var defaultConsensusParams = &abci.ConsensusParams{
//	Block: &abci.BlockParams{
//		MaxBytes: 200000,
//		MaxGas:   2000000,
//	},
//	Evidence: &tmproto.EvidenceParams{
//		MaxAgeNumBlocks: 302400,
//		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
//		MaxBytes:        10000,
//	},
//	Validator: &tmproto.ValidatorParams{
//		PubKeyTypes: []string{
//			tmtypes.ABCIPubKeyTypeEd25519,
//		},
//	},
//}
