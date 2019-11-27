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

	// Rest API
	// /v1/new/?filename={filename}&pk={pk}&score={pk}
	// /v1/get/{pk} -> returns data about the pk {pk}
	// /v1/get      -> returns all data sent to the server
	r.HandleFunc("/v1/get/", HandleSent)
	r.HandleFunc("/v1/new/", HandleNewRecord)
	r.HandleFunc("/v1/get/{pk}", HandleSelectRecord)

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
