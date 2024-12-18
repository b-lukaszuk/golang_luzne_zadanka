package main

import (
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func Racer(urlA, urlB string) (winner string) {
	aDuration := measureResponseTime(urlA)

	bDuration := measureResponseTime(urlB)

	if aDuration < bDuration {
		return urlA
	}

	return urlB
}
