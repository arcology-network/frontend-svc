package handler

import (
	"net/http"

	"github.com/arcology-network/frontend-svc/backend"
	"github.com/julienschmidt/httprouter"
)

func GetLatestHeight(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !ValidateAccessToken(w, r) {
		return
	}
	latestHeight, err := backend.GetLatestHeight()
	if err != nil {
		InternalServerError(w, err)
		return
	}

	ResponseOK(w, "latestheight", latestHeight)
}
