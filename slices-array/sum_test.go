package slicesarray

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("It should be able to sum the numbers inside the array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		sut := Sum(numbers)
		got := 15

		if sut != got {
			t.Errorf("Test Error: The sum of the values are not equal, want %d, got %d", got, sut)
		}
	})

	t.Run("It should be able to sum the numbers inside the slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		sut := Sum(numbers)
		got := 6

		if sut != got {
			t.Errorf("Test Error: The sum of the values are not equal, want %d, got %d", got, sut)
		}
	})

	t.Run("It should be able to sum all the slices inside the function SumAll", func(t *testing.T) {

		sut := SumAll([]int{1, 2}, []int{3, 4, 5})
		got := []int{3, 12}

		if !reflect.DeepEqual(sut, got) {
			t.Errorf("Test Error: The sum of the values are not equal, want %d, got %d", got, sut)
		}
	})
}
