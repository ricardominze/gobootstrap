package telemetry

import "github.com/prometheus/client_golang/prometheus"

func MakeMetricRequest(httpRequests *prometheus.CounterVec, url string) {

	httpRequests.WithLabelValues(url)
}
