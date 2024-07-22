package main

import (
	"sync"
)
// execute is used to implement goroutine calls - SearchFile will do the proper file / extension search
func execute(dirs []string, query string, results chan<- string, maxGoroutines int, searchType string) {
	wg := new(sync.WaitGroup)			// define a waitgroup for goroutines 
	semaphore := make(chan struct{}, maxGoroutines)	// use a semaphore to coordinate goroutines 

	for _, dir := range dirs {
		wg.Add(1)			// waitgroup + 1
		semaphore <- struct{}{}
		go func(dir string) {
			defer wg.Done()
			SearchFile(dir, query, results, searchType)	// function call SearchFile (call a goroutine)
			<-semaphore
		}(dir)
	}

	wg.Wait()			// wait for all the goroutines 
	close(results)			// close results channel
}
