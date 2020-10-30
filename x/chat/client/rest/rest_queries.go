package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/tendermint/spn/x/chat/types"
)

func showChannelHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get and encode the identifier param
		var params types.QueryShowChannelRequest
		vars := mux.Vars(r)
		channelID, err := strconv.Atoi(vars["channel_id"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		params.Id = int32(channelID)
		data, err := clientCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, height, err := clientCtx.QueryWithData(fmt.Sprintf("/custom/%s/%s", types.QuerierRoute, types.QueryShowChannel), data)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func listMessagesHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get and encode the identifier param
		var params types.QueryListMessagesRequest
		vars := mux.Vars(r)
		channelID, err := strconv.Atoi(vars["channel_id"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		params.ChannelId = int32(channelID)
		data, err := clientCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, height, err := clientCtx.QueryWithData(fmt.Sprintf("/custom/%s/%s", types.QuerierRoute, types.QueryListMessages), data)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}
