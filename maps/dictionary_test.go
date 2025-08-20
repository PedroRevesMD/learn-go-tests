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
		_, err := dictionary.Search("unknown")
		want := "Could not find the word you're looking..."

		if err == nil {
			t.Errorf("Expected to get an error...")
		}

		assertStrings(t, err.Error(), want)

	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
