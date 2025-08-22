package mocking

import (
	"fmt"
	"io"
)

const COUNTER = 3
const FINALMESSAGE = "Go!"

func Countdown(out io.Writer) {
	for i := COUNTER; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprintf(out, FINALMESSAGE)
}
