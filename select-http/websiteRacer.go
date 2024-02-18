package main

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	select { // the first channel to send a value "wins" and the code in it is executed
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
