package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("It should be able to Repeat a pattern of letters by using iteration", func(t *testing.T) {
		repeated, _ := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("Test Error: Expected %q Received %q", expected, repeated)
		}
	})

	t.Run("It should be able to accept the number of iterations in the loop", func(t *testing.T) {
		repeated, _ := Repeat("a", 4)
		expected := "aaaa"

		if repeated != expected {
			t.Errorf("Test Error: Expected %q Received %q", expected, repeated)
		}

		if len(repeated) != 4 {
			t.Errorf("Test Error: The length of the string is not compatible")
		}
	})

	t.Run("It should be able to throw an error when the input is a negative number", func(t *testing.T) {
		_, err := Repeat("a", -1)

		if err == nil {
			t.Errorf("Test Error: The input is a negative number, therefore impossible to make a loop")
		}
	})
}
