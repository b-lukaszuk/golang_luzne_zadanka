package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Duration(delay))
			w.WriteHeader(http.StatusOK)
		}))
	return server
}

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one",
		func(t *testing.T) {

			slowServer := makeDelayedServer(20 * time.Millisecond)
			fastServer := makeDelayedServer(0 * time.Millisecond)

			// defer
			// will call serv.Close() at the end of the containing function
			defer slowServer.Close()
			defer fastServer.Close()

			slowURL := slowServer.URL
			fastURL := fastServer.URL

			fmt.Println("slowURL {}", slowURL)
			fmt.Println("fastURL {}", fastURL)

			want := fastURL
			got, err := Racer(slowURL, fastURL)

			if err != nil {
				t.Fatalf("did not expect an error but got one %v", err)
			}

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})

	t.Run("returns an error if a server doesn't respond within the specified time",
		func(t *testing.T) {
			server := makeDelayedServer(25 * time.Millisecond)

			defer server.Close()

			_, err := ConfigurableRacer(
				server.URL, server.URL, 20*time.Millisecond)

			if err == nil {
				t.Error("expected an error but didn't get one")
			}
		})
}
