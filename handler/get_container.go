package handler

import (
	"errors"
	"fmt"

	"net/http"
	"strconv"

	"github.com/arcology-network/3rd-party/eth/common"
	"github.com/arcology-network/frontend-svc/backend"
	"github.com/julienschmidt/httprouter"
)

func GetContainer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !ValidateAccessToken(w, r) {
		return
	}

	address := params.ByName("address")
	if !common.IsHexAddress(address) {
		BadRequest(w, fmt.Errorf("invalid address: %s", address))
		return
	}

	id := params.ByName("id")
	key := params.ByName("key")
	if r.Form["type"] == nil {
		BadRequest(w, errors.New("type cannot be empty"))
		return
	}
	typ := r.Form["type"][0]
	if typ != "map" && typ != "array" && typ != "queue" {
		BadRequest(w, fmt.Errorf("unknown type %s got, expected map, array, or queue", typ))
		return
	}

	if typ == "array" || typ == "queue" {
		index, err := strconv.Atoi(key)
		if err != nil {
			BadRequest(w, fmt.Errorf("invalid key for array or queue: %s", key))
			return
		}

		if index < 0 {
			BadRequest(w, errors.New("index for array or queue should be greater than or equal to 0"))
			return
		}
	}

	height := -1
	var err error
	r.ParseForm()
	if r.Form["height"] != nil {
		heightStr := r.Form["height"][0]
		heightStr = checkLatest(heightStr)
		height, err = strconv.Atoi(heightStr)
		if err != nil {
			BadRequest(w, fmt.Errorf("invalid height: %s", r.Form["height"][0]))
			return
		}
	}
	fmt.Printf("address: %s, id: %s, key: %s, type: %s, height: %d\n", address, id, key, typ, height)

	data, err := backend.GetContainer(address, id, key, typ, height)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	ResponseOK(w, "data", data)
}
