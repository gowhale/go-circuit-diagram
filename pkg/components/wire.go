package components

import (
	"fmt"
	"go-circuit-diagram/pkg/common"
	"image/color"
)

// WireConfig is configuration for a line on the canvas
type WireConfig struct {
	startCoord common.Coordinate
	endCoord   common.Coordinate
	Colour     color.Color
}

// NewWire returns a Wire config starting from specified x,y and ending at x,y
// Note: only straight lines currently supported
func NewWire(startCoord, endCoord common.Coordinate) WireConfig {
	return WireConfig{
		startCoord: startCoord,
		endCoord:   endCoord,
		Colour:     color.Black,
	}
}

// GetColour gets the colour to render the element in
func (w *WireConfig) GetColour() color.Color {
	return w.Colour
}

// GetCoordinates calculates cords to draw onto a canvas
func (w *WireConfig) GetCoordinates() ([][]int, error) {
	if w.startCoord.X() != w.endCoord.X() && w.startCoord.Y() != w.endCoord.Y() {
		return [][]int{}, fmt.Errorf("only straight lines, horizontal or vertical")
	}

	lineCords := [][]int{}

	// vertical line
	if w.startCoord.X() == w.endCoord.X() {
		for y := w.startCoord.Y(); y < w.endCoord.Y(); y++ {
			lineCords = append(lineCords, []int{w.startCoord.X(), y})
		}
		return lineCords, nil
	}

	// horizontal line
	for x := w.startCoord.X(); x < w.endCoord.X(); x++ {
		lineCords = append(lineCords, []int{x, w.startCoord.Y()})
	}
	return lineCords, nil
}
