package main

import (
	_ "os"
	"strconv"
	"testing"
	"time"
)

func TestInsertion(t *testing.T) {

	db := GetHandler("test")

	var table string = "CREATE TABLE IF NOT EXISTS TB_FILE (FILE_PK INTEGER " +
		"PRIMARY KEY, FILE_NAME CHAR(1024), FILE_DATE " +
		"CHAR(1024), FILE_PATH CHAR(1024));"

	stmt, err := db.database.Prepare(table)

	if err != nil {
		t.Fatalf("[!] Error while creating the test storage table. Cause: %s",
			err.Error())
	}

	_, err = stmt.Exec()

	for i := 0; i < 4; i++ {
		filename := "json_file_" + strconv.Itoa(i)
		_, err := db.Insert(filename, time.Now().String(),
			"/uploads/"+filename)

		if err != nil {
			t.Fatalf("[!] Failed to insert data into the table. Cause: %s",
				err.Error())
		}
	}
}

func TestRemoval(t *testing.T) {

	db := GetHandler("test")

	res, err := db.Insert("test_removal", time.Now().String(), "")

	if err != nil {
		t.Fatalf("[!] Failed to insert data into the table. Cause: %s",
			err.Error())
	}

	// I gonna ignore the error for this expression for now (FIXME)
	pk, _ := res.LastInsertId()

	_, err = db.Delete(int(pk))

	if err != nil {
		t.Fatalf("[!] Failed to delete data from the table. Cause: %s",
			err.Error())
	}
}

func TestRetrieve(t *testing.T) {}
