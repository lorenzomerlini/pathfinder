package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getDirectories(homeDir string) []string {
	var directories []string
	files, err := os.ReadDir(homeDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, dir := range files {
		if dir.IsDir() {
			directories = append(directories, filepath.Join(homeDir, dir.Name()))
		}
	}
	return directories
}
