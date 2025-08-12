package main

import "testing"

func TestHello(t *testing.T) {
	text := Hello()
	want := "Hello, World!"

	if text != want {
		t.Errorf("Got %q want %q", text, want)
	}
}
