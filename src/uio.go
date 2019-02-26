package main

// The util io file contains operations that facilitate tasks on the project,
// such as saving a file or verifying its existence. It also reads JSON from
// files, persisting them on the database.

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

const (
	uploads string = "./uploads/"
)

type JSONData struct {
	Pk    string
	Score string
}

// Returns whether or not a file or directory exists
func Exists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}

// Saves a file named filename on the file system
func Save(file multipart.File, filename string) string {
	// Files are stored in the format <time - filename>, this way it's
	// less probable that an error of override occurs
	t := time.Now().Format(time.Stamp)
	dest := uploads + strings.Split(filename, ".")[0] + " - " + t + ".json"

	f, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("[!] Error trying to create a file: %v", err.Error())
		return ""
	}

	defer f.Close()
	io.Copy(f, file)

	return dest
}

func Persist(filename string) {
	// After copying all of the content inside f we have to get all the objects
	// inside it and persist on the database.
	file, err := os.Open(filename)

	if err != nil {
		log.Printf("[!] Couldn't open the file on persist. Cause %s",
			err.Error())
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var row JSONData
		var data []byte = scanner.Bytes()

		if len(data) != 0 {
			err = json.Unmarshal(data, &row)

			if err != nil {
				log.Printf("[!] Couldn't unmarshal data of file. Cause %s",
					err.Error())
			}

			_, err = GetHandler().Insert(strings.Split(filename, "/")[2],
				row.Pk, row.Score)

			if err != nil {
				log.Printf("[!] Couldn't persist file %s on the database. "+
					"Cause %s", filename, err.Error())
			}
		}
	}
}
