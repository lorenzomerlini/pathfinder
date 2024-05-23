package main

import (
	"sync"
)

func execute(dirs []string, query string, results chan<- string, maxGoroutines int, searchType string) {
	wg := new(sync.WaitGroup)
	semaphore := make(chan struct{}, maxGoroutines)

	for _, dir := range dirs {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(dir string) {
			defer wg.Done()
			SearchFile(dir, query, results, searchType)
			<-semaphore
		}(dir)
	}

	wg.Wait()
	close(results)
}
