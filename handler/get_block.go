package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/arcology-network/frontend-svc/backend"
	"github.com/julienschmidt/httprouter"
)

func checkLatest(height string) string {
	if height == "latest" {
		height = "-1"
	}
	return height
}

func GetBlock(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !ValidateAccessToken(w, r) {
		return
	}
	heightStr := params.ByName("height")
	heightStr = checkLatest(heightStr)
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		BadRequest(w, fmt.Errorf("invalid block height: %s", params.ByName("height")))
		return
	}

	transactions := false
	if r.Form["transactions"] != nil {
		queryTransactions := r.Form["transactions"][0]
		if "true" == queryTransactions {
			transactions = true
		}
	}

	block, err := backend.GetBlock(height, transactions)
	if err != nil {
		InternalServerError(w, err)
		return
	}
	ResponseOK(w, "block", block)
}
