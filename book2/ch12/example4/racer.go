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

func ping(url string) chan struct{} {
	// do not use: var ch
	// to initialize chanel, cause zero value for it is nil
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func Racer(urlA, urlB string) (winner string) {
	// select allows you to wait on multiple channels
	// the first one to send a value "wins" and the code underneath the case
	// is executed
	select {
	case <-ping(urlA):
		return urlA
	case <-ping(urlB):
		return urlB
	}
}
