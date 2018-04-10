package main

import (
	//"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleFalcoHTTP)

	// Setup prometheus
	registerPrometheus()

	//http.HandleFunc("/custom-metrics-api", handleCustomMetricsAPI)

	http.HandleFunc("/deleteK8SPod", handleDeleteK8SPod)
	http.HandleFunc("/count", count)
	http.ListenAndServe(":80", nil)
}
