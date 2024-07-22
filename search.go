package main

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func SearchFile(directory string, query string, results chan<- string, searchType string) {

	filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {		// walk the selected directory
		if err != nil {
			return nil
		}
// two different approaches based on the search type
		if searchType == "file" {		// bool - true!
			if !d.IsDir() && d.Name() == query {
				results <- path
			}
		} else if searchType == "extension" {
			if !d.IsDir() && strings.HasSuffix(d.Name(), query) {
				results <- path
			}
		}
		return nil
	})
}
