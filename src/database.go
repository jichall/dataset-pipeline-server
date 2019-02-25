package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const (
	folder   string = "database"
	database string = "data.db"
)

type DatabaseHandler struct {
	database *sql.DB
}

type DataRow struct {
	pk         int
	filename   string
	uploadTime string
	//path       string
}

func InitDatabase() {
	// If the database folder doesn't exist, create one
	if !Exists(folder) {
		log.Println("[+] Creating the database folder")
		os.Mkdir(folder, os.FileMode(0775))
	}

	db, err := sql.Open("sqlite3", folder+"/"+database)

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

	if err != nil {
		log.Fatalf("[!] Some error occurred trying to create the file table %v",
			err.Error())
	}
}

// Gets a database handler
func GetHandler() *DatabaseHandler {
	db, err := sql.Open("sqlite3", folder+"/"+database)

	if err != nil {
		log.Fatalf("[!] Couldn't get a sqlite connection")
	}

	return &DatabaseHandler{db}
}

// Retrieve information from the default database
func (db DatabaseHandler) Retrieve() (sql.Result, error) {

	stmt, err := db.database.Prepare("SELECT * FROM TB_FILE")

	if err != nil {
		return nil, err
	}

	return stmt.Exec()
}

// Adds a database entry
func (db DatabaseHandler) Insert(filename, uploadTime, filepath string) (sql.Result, error) {

	stmt, err := db.database.Prepare("INSERT INTO TB_FILE (FILE_NAME, FILE_DATE," +
		"FILE_PATH) VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	return stmt.Exec(filename, uploadTime, filepath)
}

// Delete a database entry
func (db DatabaseHandler) Delete(pk int) (sql.Result, error) {
	stmt, err := db.database.Prepare("DELETE FROM TB_FILE WHERE TB_PK=?")

	if err != nil {
		return nil, err
	}

	return stmt.Exec(pk)
}
