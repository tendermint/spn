package rest

import (
	"net/http"

	proto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"

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

		// Get and decode payload if defined
		var payload *proto.Message
		if req.Payload != "" {
			err = jsonpb.UnmarshalString(req.Payload, *payload)
			if err != nil {
				return
			}
		} else {
			payload = nil
		}

		// Create and send the message
		msg, err := types.NewMsgCreateChannel(
			creator,
			req.Name,
			req.Subject,
			payload,
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
