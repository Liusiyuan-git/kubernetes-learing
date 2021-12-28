package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"week2/server"
)

func main() {
	srv := server.NewServer(":8080")

	errChan, err := srv.ListenAndServe()
	if err != nil {
		log.Println("web server start failed:", err)
		return
	}
	log.Println("web server start ok")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-errChan:
		log.Println("web server run failed:", err)
		return
	case <-c:
		log.Println("server program is exiting...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx)
	}

	if err != nil {
		log.Println("server program exit error:", err)
		return
	}
	log.Println("server program exit ok")
}
