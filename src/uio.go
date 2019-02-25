package main

// This file contains util io operations present across the project and also
// does the persistence job.

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"
)

const (
	uploads string = "./uploads/"
)

// Returns whether or not a file or directory exists
func Exists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}

// Saves a file on the fs and also on the database
func Save(file multipart.File, filename string) {
	// Files are stored in the format <time - filename>, this way it's
	// less probable that an error of override occurs
	t := time.Now().Format(time.Stamp)
	dest := uploads + t + " - " + filename

	if Exists(dest) {
		// TODO
	}

	f, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("[!] Error trying to create a file: %v", err.Error())
	}

	defer f.Close()
	io.Copy(f, file)

	// TODO: Persist the file on the database
	_, err = GetHandler().Insert(filename, t, dest)

	if err != nil {
		log.Printf("[!] Couldn't persist file %s on the database", filename)
	}
}
