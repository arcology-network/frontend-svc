package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Connect(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !ValidateAccessToken(w, r) {
		return
	}

	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)

	if r.Method == http.MethodGet {
		server := sess.Get("server")
		fmt.Printf("server in session: %s\n", server)
	} else if r.Method == http.MethodPost {
		server := r.Form.Get("server")
		sess.Set("server", server)
		fmt.Printf("server in form: %s\n", server)
	}
}
