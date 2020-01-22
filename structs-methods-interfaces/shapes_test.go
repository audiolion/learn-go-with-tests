package structs_methods_interfaces

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestRectangle_Area(t *testing.T) {

	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := rectangle.Area()
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

func TestCircle_Area(t *testing.T) {
	circle := Circle{10}
	got := circle.Area()
	want := 314.1592653589793

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestCircle_Circumference(t *testing.T) {
	circle := Circle{10}
	got := circle.Circumference()
	want := 62.83185307179586

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestCircle_Diameter(t *testing.T) {
	circle := Circle{10}
	got := circle.Diameter()
	want := 20.0

	if got != want {
		t.Errorf("got %.2f want %2.f", got, want)
	}
}
