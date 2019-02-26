package main

import (
	_ "os"
	"strconv"
	"testing"
)

func TestInsertion(t *testing.T) {

	db := GetHandler("test")

	stmt, err := db.database.Prepare(table)

	if err != nil {
		t.Fatalf("[!] Error while creating the test storage table. Cause: %s",
			err.Error())
	}

	_, err = stmt.Exec()

	for i := 0; i < 4; i++ {
		filename := "json_file_" + strconv.Itoa(i)
		pk := strconv.Itoa(i + i*i)
		score := strconv.Itoa(54 + i*3 ^ 100)
		_, err := db.Insert(filename, pk, score)

		if err != nil {
			t.Fatalf("[!] Failed to insert data into the table. Cause: %s",
				err.Error())
		}
	}
}

func TestSelect(t *testing.T) {}
