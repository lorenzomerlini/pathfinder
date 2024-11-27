package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var query string

		// print title
		myFigure := figure.NewColorFigure("PathFinder Tool", "linux", "red", true)
		myFigure.Print()

		// select type of search
		fmt.Println("Select search type:")
		fmt.Println("1. Search for a specific file")
		fmt.Println("2. Search for all files with a specific extension")
		searchType, _ := reader.ReadString('\n')
		searchType = strings.TrimSpace(searchType)

		// switch type of search
		switch searchType {
		case "1":
			searchType = "file"
			fmt.Println("Insert file name: ")
			query, _ = reader.ReadString('\n')
			query = strings.TrimSpace(query)

		case "2":
			searchType = "extension"
			fmt.Println("Insert file extension (e.g., .txt): ")
			query, _ = reader.ReadString('\n')
			query = strings.TrimSpace(query)

		default:
			fmt.Println("Invalid selection. Defaulting to file search.")
			searchType = "file"
			fmt.Println("Insert file name: ")
			query, _ = reader.ReadString('\n')
			query = strings.TrimSpace(query)
		}

		var startDir string
		for {

			fmt.Println("Select search directory:")
			fmt.Println("1. Home Directory")
			fmt.Println("2. Root Directory")
			selection, _ := reader.ReadString('\n')
			selection = strings.TrimSpace(selection) // remove newline char

			if selection == "1" || selection == "2" {
				switch selection {
				case "1":
					homeDir, _ := os.UserHomeDir()
					startDir = homeDir
				case "2":
					startDir = "/"

				}
				break
			} else {

				fmt.Println("Invalid selection. Please enter 1 or 2.")
			}
		}

		fmt.Printf("Search Type: %s\n", searchType)
		fmt.Printf("Query: %s\n", query)
		fmt.Printf("Starting Directory: %s\n", startDir)

		fmt.Println("Enter the maximum number of goroutines to use: ")
		maxGoroutinesinput, _ := reader.ReadString('\n')
		maxGoroutinesinput = strings.TrimSpace(maxGoroutinesinput) // remove newline char
		maxGoroutines, _ := strconv.Atoi(maxGoroutinesinput)
		if maxGoroutines <= 0 {

			fmt.Println("Invalid input. Using default value of 10 goroutines.")
			maxGoroutines = 10
		}

		dirs := getDirectories(startDir)

		results := make(chan string)
		var found []string

		go func() {
			for result := range results {
				found = append(found, result)
			}
		}()

		startTime := time.Now()
		execute(dirs, query, results, maxGoroutines, searchType)
		elapsedTime := time.Since(startTime)

		fmt.Println("Search executed successfully!")
		fmt.Printf("\nFound %d paths\n", len(found))
		fmt.Printf("Time occourred: %s\n", &elapsedTime)

		fmt.Println("Would you like to (p)rint the results, (s)ave them to a file, or (n)either? (p/s/n): ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		if response == "p" {
			for _, result := range found {
				fmt.Println(result)
			}
		} else if response == "s" {

			saveResults(found)
		}

		fmt.Println("Would you like to do another search? (y/n): ")
		newsearch, _ := reader.ReadString('\n')
		if newsearch[0] != 'y' && newsearch[0] != 'Y' {
			fmt.Println("Goodbye!")
			break
		}

		clearTerminal()
	}
}

func saveResults(results []string) {

	file, err := os.Create("paths")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, result := range results {
		_, err := writer.WriteString(result + "\n")
		if err != nil {
			fmt.Println("Error writing to file: ", err)
			return
		}
	}
	writer.Flush()
	fmt.Println("Results saved to paths.txt")
}
