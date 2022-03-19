package middleware

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus"
)

func PrometheusMiddleware(duration prometheus.Observer, next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		defer func(begin time.Time) {
			duration.Observe(time.Since(begin).Seconds())
		}(time.Now())
		next(w, r, params)
	}
}
