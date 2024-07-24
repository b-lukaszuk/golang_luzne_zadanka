package concurrency

import (
	"time"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// each iteration of the loop will start a new goroutine
	// concurrent with the curent process
	for _, url := range urls {
		// goroutine (runs in parallel) is denoted in code: go doSomething()
		go func() {
			results[url] = wc(url)
		}()
	}
	time.Sleep(2 * time.Second) // wait for goroutines to fill results
	// still, all goroutines write their result to the same shared url
	return results
}
