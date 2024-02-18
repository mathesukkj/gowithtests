package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test speed of two different websites", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("error if a website take > 10s to respond", func(t *testing.T) {
		serverA := makeDelayedServer(12 * time.Second)
		defer serverA.Close()

		serverB := makeDelayedServer(11 * time.Second)
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didnt get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
