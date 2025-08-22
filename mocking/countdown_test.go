package mocking

import (
	"bytes"
	"testing"
)

func TestCountdow(t *testing.T) {
	buffer := &bytes.Buffer{}
	spyer := &SpySleeper{}

	Countdown(buffer, spyer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

	if spyer.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spyer.Calls)
	}
}
