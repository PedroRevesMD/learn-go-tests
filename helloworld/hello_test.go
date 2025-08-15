package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying Hello to People", func(t *testing.T) {
		text := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, text, want)
	})

	t.Run("Saying Hello World if the name isnt provided", func(t *testing.T) {
		text := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, text, want)
	})

	t.Run("Saying Hello in Spanish", func(t *testing.T) {
		text := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, text, want)
	})

	t.Run("Saying Hello in French", func(t *testing.T) {
		text := Hello("Elodie", "French")
		want := "Bonjour, Elodie"

		assertCorrectMessage(t, text, want)
	})
}

func assertCorrectMessage(t testing.TB, text, want string) {
	t.Helper()
	if text != want {
		t.Errorf("Got %q want %q", text, want)
	}
}
