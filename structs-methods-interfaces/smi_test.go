package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	got := CalculatePerimeter(10.0, 10.0)
	want := 40.00

	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.CalculateArea()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 5.0}
		want := 50.00
		checkArea(t, rectangle, want)
	})

	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})

	// Table Driven Tests: good way to test and easy to add another struct (in this case a new shape)
	arrayTests := []struct {
		name         string
		shape        Shape
		areaExpected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 5}, areaExpected: 50},
		{name: "Circle", shape: Circle{Radius: 10}, areaExpected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 3, Height: 4}, areaExpected: 6},
	}

	for _, tt := range arrayTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.CalculateArea()
			if got != tt.areaExpected {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.areaExpected)
			}
		})
	}
}
