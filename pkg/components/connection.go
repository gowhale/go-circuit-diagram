package components

import (
	"go-circuit-diagram/pkg/common"
	"image/color"
)

// ConnectionConfig is configuration for a line on the canvas
type ConnectionConfig struct {
	cord   common.Coordinate
	Colour color.Color
}

// NewGPIO returns a GPIO config starting at specified x,y
func NewConnection(cord common.Coordinate) ConnectionConfig {
	return ConnectionConfig{
		cord:   cord,
		Colour: color.Black,
	}
}

// GetColour gets the colour to render the element in
func (g *ConnectionConfig) GetColour() color.Color {
	return g.Colour
}

// GetCoordinates calculates cords to draw onto a canvas
func (g *ConnectionConfig) GetCoordinates() ([][]int, error) {
	cords := [][]int{}
	cords = append(cords, []int{g.cord.GetX() - 1, g.cord.GetY()})
	cords = append(cords, []int{g.cord.GetX() + 1, g.cord.GetY()})
	cords = append(cords, []int{g.cord.GetX() + 1, g.cord.GetY() + 1})
	cords = append(cords, []int{g.cord.GetX() - 1, g.cord.GetY() + 1})
	cords = append(cords, []int{g.cord.GetX(), g.cord.GetY() - 1})
	cords = append(cords, []int{g.cord.GetX(), g.cord.GetY() + 1})
	cords = append(cords, []int{g.cord.GetX() - 1, g.cord.GetY() - 1})
	cords = append(cords, []int{g.cord.GetX() + 1, g.cord.GetY() - 1})
	return cords, nil
}
