package main

import (
	"fmt"
	"html/template"

	"log"
	_ "math"
	"net/http"
)

// Handles all the requests sent by the client
func handle(title string, w http.ResponseWriter, r *http.Request) {

	page, err := Load(title)

	if err == nil {
		render(w, Filepath(title), page)
	} else {
		log.Printf("[!] The page \"%s\" couldn't be loaded. Reason: %v",
			Filepath(title), err)

	}
}

// Renders the webpage using the template available through the
// html/template package
func render(w http.ResponseWriter, filepath string, p *Page) {
	t, err := template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = t.Execute(w, p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	handle("index", w, r)
}

func HandleSent(w http.ResponseWriter, r *http.Request) {
	// Read from the database every file that has been uploaded and sent it
	// back to the client.
	handle("sent", w, r)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	// Only accepts POST method when uploading
	if r.Method == "POST" {
		// Max file size equals to the size in bytes of an int 64, maybe it is
		// to big (?).
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("upload")

		if handler.Size > 0 {
			// There's no way to receive multiple files because of the incorrect
			// request made by the page. There has to be implemented on the web
			// page a way to submit n-input files (but I'm not doing it now and
			// letting it as a TODO.
			log.Printf("[+] File(s) received: %s|%d\n", handler.Filename,
				handler.Size)

			if err != nil {
				log.Fatalf("[!] Some error has ocurred: %v", err.Error())
			}

			defer file.Close()
			// Saves the file on the uploads folder
			Save(file, handler.Filename)

			// TODO: Return the status of the upload through websockets
			fmt.Fprintf(w, "%s", "The data has been sent.")

		} else {
			// No file was sent, inform the client using websockets. Currently
			// I'm just sending him to the index page.
			handle("index", w, r)
		}
	}
}
