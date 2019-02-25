package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	_ "os"
	"strconv"
	"testing"
	"time"
)

func TestInsertion(*testing.T) {

	db, err := sql.Open("sqlite3", "./database/test.db")

	if err != nil {
		log.Fatalf("[!] Something went wrong trying to open the database %s",
			err.Error())
	}

	var table string = "CREATE TABLE IF NOT EXISTS TB_FILE (FILE_PK INTEGER" +
		"PRIMARY KEY AUTO_INCREMENT, FILE_NAME CHAR(1024), FILE_DATE " +
		"CHAR(1024), FILE_PATH CHAR(1024));"

	log.Printf("[+] Creating the storage table")

	stmt, err := db.Prepare(table)
	_, err = stmt.Exec()

	dh := DatabaseHander{db}

	for i := 0; i < 1000; i++ {
		_, err := dh.Insert("json_file_"+strconv.Itoa(i),
			time.Now().Format("%D/%M/%Y %h:%m:s"), "/uploads/")

		if err != nil {
			log.Fatalf(err.Error())
		}
	}
}

func TestRemoval(*testing.T) {

}
