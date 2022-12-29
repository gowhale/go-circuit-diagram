package components

import "fmt"

// WireConfig is configuration for a line on the canvas
type WireConfig struct {
	StartX, StartY int
	EndX, EndY     int
}

// NewLED returns a LED config starting from specified x,y
func NewWire(startX, startY, endX, endY int) WireConfig {
	return WireConfig{
		StartX: startX,
		StartY: startY,
		EndX:   endX,
		EndY:   endY,
	}
}

// GetCoordinates calculates cords to draw onto a canvas
func (w *WireConfig) GetCoordinates() ([][]int, error) {
	if w.StartX != w.EndX && w.StartY != w.EndY {
		return [][]int{}, fmt.Errorf("only straight lines, horizontal or vertical")
	}

	lineCords := [][]int{}

	// vertical line
	if w.StartX == w.EndX {
		for y := w.StartY; y < w.EndY; y++ {
			lineCords = append(lineCords, []int{w.StartX, y})
		}
		return lineCords, nil
	}

	// horizontal line
	for x := w.StartX; x < w.EndX; x++ {
		lineCords = append(lineCords, []int{x, w.StartY})
	}
	return lineCords, nil
}
