package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/tendermint/spn/x/chat/types"
)

type createChannelRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
	Name    string       `json:"name"`
	Subject string       `json:"subject"`
	Payload string       `json:"payload"`
}

func createChannelHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createChannelRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// Decode the address of the creator
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Create and send the message
		msg, err := types.NewMsgCreateChannel(
			creator,
			req.Name,
			req.Subject,
			[]byte(req.Payload),
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type sendMessageRequest struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	ChannelID   string       `json:"channel_id"`
	Creator     string       `json:"creator"`
	Content     string       `json:"content"`
	Tags        []string     `json:"tags"`
	PollOptions []string     `json:"poll_options"`
	Payload     string       `json:"payload"`
}

func sendMessageHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req sendMessageRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// Decode the channel ID
		channelID, err := strconv.Atoi(req.ChannelID)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Decode the address of the creator
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Create and send the message
		msg, err := types.NewMsgSendMessage(
			int32(channelID),
			creator,
			req.Content,
			req.Tags,
			req.PollOptions,
			[]byte(req.Payload),
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
