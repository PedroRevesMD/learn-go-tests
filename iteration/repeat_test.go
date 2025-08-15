package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("It should be able to Repeat a pattern of letters by using iteration", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("Test Error: Expected %q Received %q", expected, repeated)
		}
	})
}
