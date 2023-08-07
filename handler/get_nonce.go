package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/arcology-network/evm/common"
	"github.com/arcology-network/frontend-svc/backend"
	"github.com/julienschmidt/httprouter"
)

func GetNonce(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !ValidateAccessToken(w, r) {
		return
	}

	address := params.ByName("address")
	if !common.IsHexAddress(address) {
		BadRequest(w, fmt.Errorf("invalid address: %s", address))
		return
	}

	height := -1
	var err error
	if r.Form["height"] != nil {
		heightStr := r.Form["height"][0]
		heightStr = checkLatest(heightStr)

		height, err = strconv.Atoi(heightStr)
		if err != nil {
			BadRequest(w, fmt.Errorf("invalid height: %s", r.Form["height"][0]))
			return
		}
	}
	fmt.Printf("address: %s, height: %d\n", address, height)

	nonce, err := backend.GetNonce(address, height)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	ResponseOK(w, "nonce", nonce)
}
