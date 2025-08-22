package mocking

import (
	"fmt"
	"io"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const COUNTER = 3
const FINALMESSAGE = "Go!"

func Countdown(out io.Writer, sleeper *SpySleeper) {
	for i := COUNTER; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprintf(out, FINALMESSAGE)
}
