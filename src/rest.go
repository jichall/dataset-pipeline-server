package main

//
// This files handles all the routes that have been requested by the client. It
// is a like a REST interface.
//

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

	rows, err := GetHandler().Select()

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Internal server error (500)")
		log.Printf("[!] Error trying to retrieve information %s", err.Error())
	}

	defer rows.Close()

	var row DataRow

	// Writing the table header
	fmt.Fprintf(w, "%s", "FILENAME\t\tPK\t\tSCORE\n")
	for rows.Next() {
		err = rows.Scan(&row.filename, &row.pk, &row.score)

		if err != nil {
			log.Printf("[!] Some error ocurred reading data from the DB. Cause %s",
				err.Error())
		} else {
			fmt.Fprintf(w, "%s\t\t\t%s\t\t%s\n", row.filename, row.pk, row.score)
		}
	}
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

			// TODO: Return the status of the upload through websockets
			fmt.Fprintf(w, "%s", "The data has been sent.")

			// This piece of code might below take longer on big files,
			// something else might be implemented to make this work faster and
			// non blocking.

			// Saves the file on the uploads folder
			go func() {
				filename := Save(file, handler.Filename)
				// Persist the file on the database
				Persist(filename)
			}()

		} else {
			// No file was sent, inform the client using websockets. Currently
			// I'm just sending him to the index page.
			handle("index", w, r)
		}
	}
}
