package middleware

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func TelemetryMetricMiddleware(next http.Handler, metric *prometheus.CounterVec) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metric.WithLabelValues(r.URL.Path).Inc()
		next.ServeHTTP(w, r)
	})
}
