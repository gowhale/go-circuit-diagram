package components

// Element ensures a struct can be painted onto a canvas
type Element interface {
	// GetCordinates returns cordinates an element is located at.
	GetCoordinates() [][]int
}
