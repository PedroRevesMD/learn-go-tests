package mocking

import (
	"bytes"
	"testing"
)

func TestCountdow(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
