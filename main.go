package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var query string

		myFigure := figure.NewColorFigure("PathFinder Tool", "linux", "red", true)
		myFigure.Print()
		fmt.Println("Insert query: ")
		fmt.Scan(&query)

		homeDir, _ := os.UserHomeDir()
		dirs := getDirectories(homeDir)

		results := make(chan string)

		go func() {
			for result := range results {
				fmt.Println(result)
			}
		}()

		startTime := time.Now()
		execute(dirs, query, results)
		elapsedTime := time.Since(startTime)

		fmt.Printf("Time occourred: %s\n", &elapsedTime)

		fmt.Println("Would you like to do another search? (y/n): ")
		response, _ := reader.ReadString('\n')
		if response[0] != 'y' && response[0] != 'Y' {
			break
		}

		clearTerminal()
	}
}

func execute(dirs []string, query string, results chan<- string) {
	wg := new(sync.WaitGroup)

	for _, dir := range dirs {
		wg.Add(1)
		go SearchFile(dir, query, results, wg)
	}

	wg.Wait()
	close(results)
}

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

func SearchFile(directory string, query string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() && d.Name() == query {
			results <- path
		}
		return nil
	})
}

func clearTerminal() {

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
