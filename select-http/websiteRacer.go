package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	durationA := time.Since(startA)

	fmt.Println(durationA)

	startB := time.Now()
	http.Get(b)
	durationB := time.Since(startB)
	fmt.Println(durationB)

	if durationA < durationB {
		return a
	}
	return b
}
