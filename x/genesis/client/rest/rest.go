package rest

import (
	"fmt"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
    MethodGet = "GET"
	MethodPost = "POST"
)

// RegisterRoutes registers genesis-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/genesis/chains", listChannelsHandler(clientCtx)).Methods(MethodGet)
	r.HandleFunc("/genesis/chain/{chain_id}", showChainHandler(clientCtx)).Methods(MethodGet)
}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/genesis/chain_create"), createChannelHandler(clientCtx)).Methods(MethodPost)
}

