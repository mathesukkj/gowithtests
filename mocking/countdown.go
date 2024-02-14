package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
