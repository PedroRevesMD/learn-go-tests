package mocking

import (
	"fmt"
	"io"
	"time"
)

const COUNTER = 3
const FINALMESSAGE = "Go!"

func Countdown(out io.Writer) {
	for i := COUNTER; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}

	fmt.Fprintf(out, FINALMESSAGE)
}
