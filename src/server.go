package main

import (
	"log"
	"math"
	"net/http"
	"time"
)

func InitServer(addr string, port string) {

	server := http.Server{
		Addr:         addr + ":" + port,
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// size of the data that could be sent through the server (It's a int
		// field so therefore 1 << 31-1 on 32 bits computers)
		MaxHeaderBytes: math.MaxInt32,
	}

	// Maybe change it to a map (?)
	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/upload", HandleUpload)
	// TODO: Finish the sent request to retrieve every file ever sent to the
	// server.
	http.HandleFunc("/sent", HandleSent)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("[!] The server couldn't be started because of %v", err)
	}
}
