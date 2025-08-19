package structs

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12.0, 6.0}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func checkArea(t testing.TB, shape Shape, want float64) float64 {
	t.Helper()
	got := shape.Area()

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}

	return got
}
