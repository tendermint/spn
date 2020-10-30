package rest

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
)

const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

// RegisterRoutes registers chat-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/chat/channel/{channel_id}", showChannelHandler(clientCtx)).Methods(MethodGet)
	r.HandleFunc("/chat/list_messages/{channel_id}", listMessagesHandler(clientCtx)).Methods(MethodGet)
}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/chat/create_channel"), createChannelHandler(clientCtx)).Methods(MethodPost)
	r.HandleFunc(fmt.Sprintf("/chat/send_message"), sendMessageHandler(clientCtx)).Methods(MethodPost)
}
