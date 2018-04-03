package main

import (
	//"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleFalcoHTTP)
	//http.HandleFunc("/metrics", handleMetrics)
	//http.HandleFunc("/custom-metrics-api", handleCustomMetricsAPI)
	http.HandleFunc("/deleteK8SPod", handleDeleteK8SPod)
	http.ListenAndServe(":80", nil)
}
