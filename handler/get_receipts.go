package handler

import (
	"net/http"
	"strings"

	"github.com/arcology-network/frontend-svc/backend"
	"github.com/julienschmidt/httprouter"
)

func GetReceipts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !ValidateAccessToken(w, r) {
		return
	}

	hashes := params.ByName("hashes")

	executingDebugLogs := false
	if r.Form["executingDebugLogs"] != nil {
		executingDebugLogstr := r.Form["executingDebugLogs"][0]
		if "true" == executingDebugLogstr {
			executingDebugLogs = true
		}
	}

	receipts, err := backend.GetReceipts(strings.Split(hashes, ","), executingDebugLogs)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	ResponseOK(w, "receipts", receipts)
}
