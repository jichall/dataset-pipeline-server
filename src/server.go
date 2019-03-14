package main

import (
	"github.com/gorilla/mux"

	"log"
	"math"
	"net/http"
	"time"
)

func InitServer(addr string, port string) {

	r := mux.NewRouter()

	r.HandleFunc("/", HandleRoot)
	r.HandleFunc("/upload", HandleUpload)
	r.HandleFunc("/sent", HandleSent)
	// Rest API
	// /v1/new/?filename=xxx.txt&pk=xxx&score=xxx
	// /v1/get/xxx
	r.HandleFunc("/new/", HandleNewRecord)
	r.HandleFunc("/get/{pk}", HandleSelectRecord)

	server := http.Server{
		Addr:         addr + ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// size of the data that could be sent through the server (It's a int
		// field so therefore 1 << 31-1 on 32 bits computers)
		MaxHeaderBytes: math.MaxInt32,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("[!] The server couldn't be started because of %v", err)
	}
}
