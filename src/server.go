package main

import (
	"log"
	"math"
	"mime"
	"net/http"
	"strings"
	"time"
)

var (
	pages map[string]func(http.ResponseWriter, *http.Request)
)

// Dataset Pipeline Server Handler
type DPServerHandler struct{}

func InitServer(addr string, port string) {

	server := http.Server{
		Addr:         addr + ":" + port,
		Handler:      &DPServerHandler{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// size of the data that could be sent through the server (It's a int
		// field so therefore 1 << 31-1 on 32 bits computers)
		MaxHeaderBytes: math.MaxInt32,
	}

	pages = make(map[string]func(http.ResponseWriter, *http.Request))

	pages["/"] = HandleRoot
	pages["/index"] = HandleRoot
	pages["/upload"] = HandleUpload
	pages["/sent"] = HandleSent

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("[!] The server couldn't be started because of %v", err)
	}
}

func (*DPServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Does the request is inside the pages map? If yes, handle it by the
	// function defined for the URL key.
	if h, ok := pages[r.URL.String()]; ok {
		h(w, r)
	} else {
		// Verify the MIME type of the requested file
		filepath := strings.Split(r.URL.String(), "/")
		filename := ""

		var mimeType string

		if len(filepath) > 2 {
			filename = filepath[len(filepath)-1]
			extension := strings.Split(filename, ".")

			mimeType = mime.TypeByExtension("." + extension[len(extension)-1])
		}

		if debug {
			if filename == "" {
				filename = filepath[1]
			}

			log.Printf("[+] Requested file %s. MIME type: %s\n", filename,
				mimeType)
		}

		// Loads the file and send it back with the correct MIME type
		file, err := Load(r.URL.String())

		if err == nil {
			w.Header().Add("Content-Type", mimeType)
			w.Write(file.Body)
		} else {
			log.Printf("[!] Error requesting the file %s. Cause %s",
				filepath[1], err.Error())
		}
	}
}
