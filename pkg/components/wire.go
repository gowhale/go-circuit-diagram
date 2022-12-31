package components

import (
	"fmt"
	"image/color"

	"github.com/gowhale/go-circuit-diagram/pkg/common"
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

func verticalCoords(w *WireConfig) [][]int {
	lineCords := [][]int{}

	for y := w.startCoord.GetY(); y < w.endCoord.GetY()+1; y++ {
		lineCords = append(lineCords, []int{w.startCoord.GetX(), y})
	}
	if len(lineCords) == 0 {
		for y := w.endCoord.GetY(); y < w.startCoord.GetY()+1; y++ {
			lineCords = append(lineCords, []int{w.startCoord.GetX(), y})
		}
	}

	return lineCords
}

func horizontal(w *WireConfig) [][]int {
	lineCords := [][]int{}

	for x := w.startCoord.GetX(); x < w.endCoord.GetX()+1; x++ {
		lineCords = append(lineCords, []int{x, w.startCoord.GetY()})
	}
	if len(lineCords) == 0 {
		for x := w.endCoord.GetX(); x < w.startCoord.GetX()+1; x++ {
			lineCords = append(lineCords, []int{x, w.startCoord.GetY()})
		}
	}

	return lineCords
}

// GetCoordinates calculates cords to draw onto a canvas
func (w *WireConfig) GetCoordinates() ([][]int, error) {
	if w.startCoord.GetX() != w.endCoord.GetX() && w.startCoord.GetY() != w.endCoord.GetY() {
		return [][]int{}, fmt.Errorf("only straight lines, horizontal or vertical")
	}

	// vertical line
	if w.startCoord.GetX() == w.endCoord.GetX() {
		return verticalCoords(w), nil
	}

	// horizontal line

	return horizontal(w), nil
}
