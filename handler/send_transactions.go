package handler

import (
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/arcology-network/frontend-svc/backend"
	"github.com/julienschmidt/httprouter"
)

func SendTransactions(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !ValidateAccessToken(w, r) {
		return
	}

	var txs [][]byte
	for _, v := range r.Form {
		for _, txStr := range v {

			tx, err := hex.DecodeString(txStr)
			if err != nil {
				BadRequest(w, fmt.Errorf("%s is not a valid transaction", tx))
				return
			}

			txs = append(txs, tx)
		}
	}

	err := backend.SendTransactions(txs)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	ResponseOK(w)
}
