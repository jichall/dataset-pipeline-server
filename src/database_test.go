package main

import (
	_ "os"
	"reflect"
	"testing"
)

const (
	testTable string = "CREATE TABLE TB_DATA (DATA_FILENAME CHAR(1024), " +
		"DATA_PK VARCHAR(1024) PRIMARY KEY, " +
		"DATA_SCORE VARCHAR(1024))"
)

type Row struct {
	filename string
	pk       string
	score    string
}

func TestDatabase(t *testing.T) {
	var testCases []Row = []Row{
		{
			filename: "file_1 - fictional hour.json",
			pk:       "101",
			score:    "1232",
		},
		{
			filename: "file_1 - fictional hour.json",
			pk:       "21",
			score:    "152",
		},
		{
			filename: "file_4 - fictional hour.json",
			pk:       "15",
			score:    "2230",
		},
		{
			filename: "file_3 - fictional hour.json",
			pk:       "13",
			score:    "5032",
		},
		{
			filename: "file_2 - fictional hour.json",
			pk:       "97",
			score:    "1232",
		},
		{
			filename: "file_2 - fictional hour.json",
			pk:       "103",
			score:    "19283",
		},
		{
			filename: "file_3 - fictional hour.json",
			pk:       "104",
			score:    "132",
		},
	}

	dh := GetHandler("test")

	stmt, err := dh.database.Prepare(table)

	if err != nil {
		t.Fatalf("[!] Error while creating the test storage table. Cause: %s",
			err.Error())
	}

	_, err = stmt.Exec()

	for i := range testCases {
		filename := testCases[i].filename
		pk := testCases[i].pk
		score := testCases[i].score

		_, err := dh.Insert(filename, pk, score)

		if err != nil {
			t.Fatalf("[!] Failed to insert data into the table. Cause: %s",
				err.Error())
		}
	}

	rows, err := GetHandler("test").Select()

	if err != nil {
		t.Fatalf("[!] Error trying to retrieve information %s", err.Error())
	}

	defer rows.Close()

	var row Row
	var index int = 0

	for rows.Next() {
		err = rows.Scan(&row.filename, &row.pk, &row.score)

		if err != nil {
			t.Fatalf("[!] Error trying to scan information. Cause %s",
				err.Error())
		}

		ok := reflect.DeepEqual(testCases[index], row)

		if ok == false {
			t.Fatalf("[!] Values does not correspond, got %v wants %v", row,
				testCases[index])
		}

		index++
	}
}
