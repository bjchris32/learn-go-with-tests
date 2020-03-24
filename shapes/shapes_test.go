package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle { 10.0, 10.0 }
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		expected float64
	} {
		{ Rectangle {12, 6}, 72.0 },
		{ Circle { 10 }, 314.1592653589793 },
		{ Triangle { 12, 6 }, 36.0 },
	}

	for _, testItem := range areaTests {
		got := testItem.shape.Area()
		if got != testItem.expected {
			t.Errorf("got %g expected %g", got, testItem.expected)
		}
	}
}
