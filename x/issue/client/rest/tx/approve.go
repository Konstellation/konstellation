package tx

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/konstellation/konstellation/x/issue/types"
)

type (
	// approveRequest defines the properties of issue create's body.
	approveRequest struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
		//Owner   string       `json:"owner" yaml:"owner"`
		Spender string    `json:"spender" yaml:"spender"`
		Amount  sdk.Coins `json:"amount" yaml:"amount"`
	}
)

func approveHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req approveRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		ownerAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		spenderAddr, err := sdk.AccAddressFromBech32(req.Spender)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgApprove(ownerAddr, spenderAddr, req.Amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, &msg)
	}
}