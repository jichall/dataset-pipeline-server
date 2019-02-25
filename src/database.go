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

// A struct  that denotes a row in the database
type DataRow struct {
	pk         int
	filename   string
	uploadTime string
	path       string
}

func InitDatabase() {
	// If the database folder doesn't exist, create one
	if !Exists(folder) {
		log.Println("[+] Creating the database folder")
		os.Mkdir(folder, os.FileMode(0775))
	}

	db := GetHandler()

	var table string = "CREATE TABLE IF NOT EXISTS TB_FILE (FILE_PK INTEGER" +
		"PRIMARY KEY AUTO_INCREMENT, FILE_NAME CHAR(1024), FILE_DATE " +
		"CHAR(1024), FILE_PATH CHAR(1024));"

	log.Printf("[+] Creating the storage table")

	stmt, err := db.database.Prepare(table)
	_, err = stmt.Exec()

	if err != nil {
		log.Fatalf("[!] Some error occurred trying to create the file table %v",
			err.Error())
	}
}

// Gets a database handler, if any argument is provided it expects
// to be the test database
func GetHandler(params ...string) *DatabaseHandler {

	var db *sql.DB = nil
	var err error = nil

	if len(params) != 0 && params[0] == "test" {
		db, err = sql.Open("sqlite3", folder+"/test.db")
	} else {
		db, err = sql.Open("sqlite3", folder+"/"+database)
	}

	if err != nil {
		log.Fatalf("[!] Something went wrong trying to open the database %s",
			err.Error())
	}

	return &DatabaseHandler{db}
}

// Retrieve information from the default database. Additional settings may
// be included using the params argument such as a where clause.
func (db DatabaseHandler) Retrieve(params ...string) (sql.Result, error) {

	stmt, err := db.database.Prepare("SELECT * FROM TB_FILE" + params[0])

	if err != nil {
		return nil, err
	}

	return stmt.Exec()
}

// Adds a database entry
func (db DatabaseHandler) Insert(filename, uploadTime, filepath string) (sql.Result, error) {

	// A NULL value must be included to make usage of the FILE_PK variable that
	// is an auto increment field automatically in sqlite
	stmt, err := db.database.Prepare("INSERT INTO TB_FILE (FILE_PK, FILE_NAME, " +
		"FILE_DATE, FILE_PATH) VALUES (NULL, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	return stmt.Exec(filename, uploadTime, filepath)
}

// Delete a database entry
func (db DatabaseHandler) Delete(pk int) (sql.Result, error) {
	stmt, err := db.database.Prepare("DELETE FROM TB_FILE WHERE FILE_PK=?")

	if err != nil {
		return nil, err
	}

	return stmt.Exec(pk)
}
