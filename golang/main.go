package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "ping_request_count", Help: "No. of requests handled a Ping handler"})

func main() {
	prometheus.MustRegister(pingCounter)

	http.HandleFunc("/ping", handlePing)
	http.Handle("/metrics", promhttp.Handler())

	port := 8080
	log.Printf("golang backend running on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	pingCounter.Inc()
	w.WriteHeader(200)
	w.Write([]byte("pong"))
}
