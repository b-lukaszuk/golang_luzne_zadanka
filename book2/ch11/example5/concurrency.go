package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// each iteration of the loop will start a new goroutine
	// concurrent with the curent process
	for _, url := range urls {
		// goroutine (runs in parallel) is denoted in code: go doSomething()
		go func(u string) {
			// send statement operator: <-
			// result struct on the right, channel on the left
			resultChannel <- result{u, wc(u)}
		}(url) // all this happens concurrently
	}

	// all this happens one at a time
	for i := 0; i < len(urls); i++ {
		// recive expression: <-, channel on the right, variable on the left
		// writing to results map happens one at a time
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
