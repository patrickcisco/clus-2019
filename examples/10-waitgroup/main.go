// https://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://developer.cisco.com",
		"https://google.com",
		"https://postman-echo.com/get?foo1=bar1&foo2=bar2",
	}
	jsonResponses := make(chan string)

	// A WaitGroup waits for a collection of goroutines to finish.
	// The main goroutine calls Add to set the number of
	// goroutines to wait for. Then each of the goroutines
	// runs and calls Done when finished. At the same time,
	// Wait can be used to block until all goroutines have finished.
	//
	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		// go routine
		go func(url string) {
			defer wg.Done()
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			} else {
				jsonResponses <- res.Status
			}
		}(url)
	}

	go func() {
		for response := range jsonResponses {
			fmt.Println(response)
		}
	}()

	wg.Wait()
}
