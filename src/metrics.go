package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

var responseMetric = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name:    "request_duration_milliseconds",
		Help:    "Request latency distribution",
		Buckets: prometheus.ExponentialBuckets(10.0, 1.13, 40),
	},
)

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
	hdFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hd_errors_total",
		Help: "Number of hard-disk errors.",
	})
	stupidCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "stupid_counter",
		Help: "Just a stupid incrementing counter.",
	})
	infections = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "infections_total",
		Help: "Number of Infections Received.",
	})
)

func setup() {
	cpuTemp.Set(123.4)
	hdFailures.Inc()
	stupidCount.Inc()
}

func registerPrometheus() {
	setup()
	http.Handle("/metrics", prometheus.Handler())
	prometheus.MustRegister(responseMetric)
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
	prometheus.MustRegister(stupidCount)
	prometheus.MustRegister(infections)
}

func count(w http.ResponseWriter, r *http.Request) {
	stupidCount.Inc()
	msg := fmt.Println("Incremented count")
	w.Header().Set("content-type", "text/plain")
	w.Write([]byte(msg))
}
