package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/tendermint/spn/x/identity/types"
)

const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

// RegisterRoutes registers identity-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/identity/username/{identifier}", usernameHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/identity/username_from_address/{address}", usernameFromAddressHandler(clientCtx)).Methods("GET")
}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/identity/set_username"), setUsernameHandler(clientCtx)).Methods("POST")
}

func usernameHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get and encode the identifier param
		var params types.QueryUsernameRequest
		vars := mux.Vars(r)
		identifier := vars["identifier"]
		params.Identifier = identifier
		data, err := clientCtx.LegacyAmino.MarshalJSON(params)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, height, err := clientCtx.QueryWithData(fmt.Sprintf("/custom/%s/%s", types.QuerierRoute, types.QueryUsername), data)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func usernameFromAddressHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get and encode the address param
		var params types.QueryUsernameFromAddressRequest
		vars := mux.Vars(r)
		address := vars["address"]
		params.Address = address
		data, err := clientCtx.LegacyAmino.MarshalJSON(params)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, height, err := clientCtx.QueryWithData(fmt.Sprintf("/custom/%s/%s", types.QuerierRoute, types.QueryUsernameFromAddress), data)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}

type setUsernameRequest struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Creator  string       `json:"creator"`
	Username string       `json:"name"`
}

func setUsernameHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setUsernameRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg, err := types.NewMsgSetUsername(
			creator,
			req.Username,
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
