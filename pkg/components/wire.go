package components

import (
	"fmt"
	"image/color"
)

// WireConfig is configuration for a line on the canvas
type WireConfig struct {
	StartX, StartY int
	EndX, EndY     int
	Colour         color.Color
}

// NewWire returns a Wire config starting from specified x,y and ending at x,y
// Note: only straight lines currently supported
func NewWire(startX, startY, endX, endY int) WireConfig {
	return WireConfig{
		StartX: startX,
		StartY: startY,
		EndX:   endX,
		EndY:   endY,
		Colour: color.Black,
	}
}

// GetColour gets the colour to render the element in
func (w *WireConfig) GetColour() color.Color {
	return w.Colour
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
