//!/bin/sh
package main

import (
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
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
	req, err := http.NewRequest("GET", "http://service2/hello", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}
	lowerCaseHeader := make(http.Header)
	for key, value := range r.Header {
		lowerCaseHeader[strings.ToLower(key)] = value
	}
	glog.Info("headers:", lowerCaseHeader)
	req.Header = lowerCaseHeader
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glog.Info("HTTP get failed with error: ", "error", err)
	} else {
		glog.Info("HTTP get succeeded")
	}
	if resp != nil {
		resp.Write(w)
	}
	glog.V(4).Infof("Respond in %d ms", delay)
}
