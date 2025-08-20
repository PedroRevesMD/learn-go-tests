package maps

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "Just a Test"}
	t.Run("Known Word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "Just a Test"
		assertStrings(t, got, want)
	})
	t.Run("Unknown Word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("Expected to get an error...")
		}
		assertError(t, got, ErrNotFound)

	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Got error %q want %q", got, want)
	}
}
