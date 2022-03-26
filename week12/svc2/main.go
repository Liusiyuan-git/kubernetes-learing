//!/bin/sh
package main

import (
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", rootHandler)
	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatalf("start http server failed, error: %s\n", err.Error())
	}
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	glog.V(4).Info("entering v2 root handler")

	delay := randInt(10, 20)
	time.Sleep(time.Millisecond * time.Duration(delay))
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
	glog.V(4).Infof("Respond in %d ms", delay)
}
