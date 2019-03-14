package main

//
// This files handles all the routes that have been requested by the client. It
// is a like a REST interface.
//

import (
	"github.com/gorilla/mux"

	"fmt"
	"log"
	"net/http"
	"strconv"
)

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
	fmt.Fprintf(w, "%s", "FILENAME\t\t\t\t\tPK\t\tSCORE\n")
	fmt.Fprintf(w, "%s", "--------\t\t\t\t\t--\t\t-----\n")

	for rows.Next() {
		err = rows.Scan(&row.filename, &row.pk, &row.score)

		if err != nil {
			log.Printf("[!] Some error ocurred reading data from the DB. Cause %s",
				err.Error())
		} else {
			fmt.Fprintf(w, "%-47s %-15s %s \n", row.filename, row.pk, row.score)
		}
	}
}

// REST API Below

func HandleNewRecord(w http.ResponseWriter, r *http.Request) {
	optional := r.URL.Query()

	filename := optional.Get("filename")
	pk := optional.Get("pk")
	score := optional.Get("score")

	_, err := strconv.Atoi(pk)
	_, err = strconv.Atoi(score)

	if err != nil {
		fmt.Fprintf(w, "%s", "Invalid request. Look at the documentation in "+
			"order to use the new record method.")
		return
	}

	res, err := GetHandler().Insert(filename, pk, score)

	if err != nil {
		fmt.Fprintf(w, "%s %s", "Error trying to persist the data sent. Cause ",
			err.Error())
	} else {
		rows, _ := res.RowsAffected()
		fmt.Fprintf(w, "Data was persisted successfully with %d rows affected",
			int(rows))
	}
}

func HandleSelectRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pk, ok := vars["pk"]

	if !ok {
		fmt.Fprintf(w, "%s", "Invalid request. Look at the documentation in "+
			"order to use the select record method.")
		return
	}

	// If the token can't be convertd to int it isn't a valid query
	// meaning it contains characters.
	_, err := strconv.Atoi(pk)
	if err != nil {
		fmt.Fprintf(w, "%s", "Invalid token sent, it must be numbers only.")
		return
	}

	// FIXME: There's a bug with pk that begins with zero.
	rows, err := GetHandler().Select("WHERE DATA_PK=" + pk)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Internal server error (500)")
		log.Printf("[!] Error trying to retrieve information %s", err.Error())
	}

	var row DataRow

	for rows.Next() {
		err = rows.Scan(&row.filename, &row.pk, &row.score)

		if err != nil {
			log.Printf("[!] Some error ocurred reading data from the DB. "+
				"Cause %s", err.Error())
		} else {
			fmt.Fprintf(w, "%s: [pk: %s score: %s]\n", row.filename, row.pk,
				row.score)
		}
	}
}
