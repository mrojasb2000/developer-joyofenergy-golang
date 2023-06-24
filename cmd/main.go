package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	serverPort = ":8000"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Joyofenery"))
	})
	s := &http.Server{
		Addr:           serverPort,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Info(fmt.Sprintf("Server litering on port: %s", serverPort))
	s.ListenAndServe()
}
