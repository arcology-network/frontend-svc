package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/arcology-network/frontend-svc/backend"
	"github.com/julienschmidt/httprouter"
)

func UpdateConfig(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !ValidateAccessToken(w, r) {
		return
	}

	for k, v := range r.Form {
		switch k {
		case "parallelism":
			p, err := strconv.Atoi(v[0])
			if err != nil {
				BadRequest(w, fmt.Errorf("invalid parallelism value: %s", v[0]))
				return
			}

			if p <= 0 {
				BadRequest(w, errors.New("parallelism must be greater than 0"))
				return
			}

			p, err = backend.SetParallelism(p)
			if err != nil {
				InternalServerError(w, err)
				return
			}

			fmt.Printf("returned p: %d\n", p)
		default:
			BadRequest(w, fmt.Errorf("unknown config parameter: %s", k))
			return
		}
	}

	ResponseOK(w)
}
