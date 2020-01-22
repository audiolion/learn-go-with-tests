package structs_methods_interfaces

// Perimeter calculates the perimeter of a rectangle
// given the width and height.
func Perimeter(x float64, y float64) float64 {
	return 2 * (x + y)
}

// Area calculates the area of a rectangle
// given the width and height.
func Area(x float64, y float64) float64 {
	return x * y
}
