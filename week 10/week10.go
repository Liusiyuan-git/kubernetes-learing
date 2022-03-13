//!/bin/sh
package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
	"week10/metrics"
	_ "week10/metrics"
)

func main() {
	mux := http.NewServeMux()
	metrics.Register()
	mux.HandleFunc("/hello", rootHandler)
	mux.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatalf("start http server failed, error: %s\n", err.Error())
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	user := r.URL.Query().Get("user")
	delay := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(delay))
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello %s", user))
	} else {
		io.WriteString(w, fmt.Sprintf("hello [stranger]"))
	}
}
