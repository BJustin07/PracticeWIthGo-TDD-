package structMethodsAndArrays

import "testing"

//func TestPerimeter(t *testing.T) {
//	rectangle := Rectangle{width: 20.00, height: 14.845}
//	got := Perimeter(rectangle)
//	want := 69.69
//	if got != want {
//		t.Errorf("Perimeter() = %.2f; want %.2f", got, want)
//	}
//}

//func TestArea(t *testing.T) {
//	checkArea := func(t testing.TB, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//		if got != want {
//			t.Errorf("got %.2f, want %.2f", got, want)
//		}
//	}
//	t.Run("Calculate area of a rectangle", func(t *testing.T) {
//		rectangle := Rectangle{12, 14}
//		checkArea(t, rectangle, 52)
//	})
//	t.Run("Calculate area of a circle", func(t *testing.T) {
//		circle := Circle{10}
//		checkArea(t, circle, 314.1592653589793)
//	})
//}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Rectangle", shape: Rectangle{width: 12, height: 6}, hasArea: 36.0},
		{name: "Triangle", shape: Triangle{base: 12, height: 6}, hasArea: 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %.2f, want %.2f", tt.shape, got, tt.hasArea)
		}
	}
}
