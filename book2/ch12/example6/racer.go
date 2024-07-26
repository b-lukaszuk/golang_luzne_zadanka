package main

import (
	"fmt"
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

var tenSecondTimeout = 10 * time.Second

func Racer(urlA, urlB string) (winner string, err error) {
	return ConfigurableRacer(urlA, urlB, tenSecondTimeout)
}

func ConfigurableRacer(urlA, urlB string, timeout time.Duration) (winner string, err error) {
	// select allows you to wait on multiple channels
	// the first one to send a value "wins" and the code underneath the case
	// is executed
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	// time.After(timeout) returns a chan
	// and will send a signal down after the specified amount of time
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}
