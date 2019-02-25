package main

// This file contains some helper functions around files, that being an html
// page.

import (
	"io/ioutil"
	"log"
)

const (
	root string = "./html/"
)

// Maybe rename this to a more specific thing later as it might have another
// kind of file as well.
type Page struct {
	Filename string
	Body     []byte
}

//
func Load(filename string) (*Page, error) {
	body, err := ioutil.ReadFile(Filepath(filename))

	if err != nil {
		log.Printf("Error %v", err.Error())
		return nil, err
	}

	return &Page{Filename: filename, Body: body}, nil
}

// Returns the filepath of a web page
func Filepath(filename string) string {
	return root + filename + ".html"
}
