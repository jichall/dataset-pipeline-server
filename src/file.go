package main

// This file contains some helper functions around files, that being an html
// page.

import (
	"io/ioutil"
	"log"
	"strings"
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

// Looks for a web page ot asset on the fs and returns it embedded on the
// Page struct.
func Load(filename string) (*Page, error) {
	body, err := ioutil.ReadFile(Filepath(filename))

	if err != nil {
		return nil, err
	}

	return &Page{Filename: filename, Body: body}, nil
}

// Returns the filepath of a web page.
func Filepath(filename string) string {
	split := strings.Split(filename, ".")

	// The file isn't an HTML and probably is an asset because of the way that
	// HTML files are loaded, they are requested without its file extension.
	if len(split) > 1 {
		if debug {
			log.Printf("[+] Asset to load %s", filename)
		}
		return root + filename[1:]
	}

	return root + filename + ".html"
}
