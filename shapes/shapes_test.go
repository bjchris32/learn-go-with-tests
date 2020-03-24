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
		name string
		shape Shape
		expected float64
	} {
		{ name: "Rectangle", shape: Rectangle { Width: 12, Height: 6 }, expected: 72.0 },
		{ name: "Circle", shape: Circle { Radius: 10 }, expected: 314.1592653589793 },
		{ name: "Triangle", shape: Triangle { Base: 12, Height: 6 }, expected: 36.0 },
	}

	for _, testItem := range areaTests {
		t.Run(testItem.name, func(t *testing.T) {
			got := testItem.shape.Area()
			if got != testItem.expected {
				t.Errorf("%#v got %g expected %g", testItem.shape, got, testItem.expected)
			}
		})

		// got := testItem.shape.Area()
		// if got != testItem.expected {
		// 	t.Errorf("got %g expected %g", got, testItem.expected)
		// }
	}
}
