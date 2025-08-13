package main

import "testing"

func TestHello(t *testing.T) {
	text := Hello("Chris")
	want := "Hello, Chris"

	if text != want {
		t.Errorf("Got %q want %q", text, want)
	}
}
