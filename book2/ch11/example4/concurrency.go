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
		go func(u string) {
			// 2 goroutines may write they result to results at the same time
			// ergo, race condition
			// test with race flag detects such conditions: go test -race
			results[u] = wc(u)
		}(url)
	}
	time.Sleep(2 * time.Second) // wait for goroutines to fill results
	// we make sure that url is fixed (to u) per iteration loop
	return results
}
