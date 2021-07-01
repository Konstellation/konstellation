package query

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/konstellation/kn-sdk/x/issue/query"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// HTTP request handler to query specified issues
func issueHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		denom := vars["denom"]

		res, height, err := cliCtx.QueryWithData(query.PathQueryIssue(denom), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx.WithHeight(height), res)
	}
}
