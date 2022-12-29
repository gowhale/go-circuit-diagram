package components

import (
	"go-circuit-diagram/pkg/common"
	"image/color"
)

// GPIOConfig is configuration for a line on the canvas
type GPIOConfig struct {
	cord   common.Coordinate
	Colour color.Color
}

// NewGPIO returns a GPIO config starting at specified x,y
func NewGPIO(cord common.Coordinate) GPIOConfig {
	return GPIOConfig{
		cord:   cord,
		Colour: color.RGBA{0, 255, 0, 0xff},
	}
}

// GetColour gets the colour to render the element in
func (g *GPIOConfig) GetColour() color.Color {
	return g.Colour
}

// GetCoordinates calculates cords to draw onto a canvas
func (g *GPIOConfig) GetCoordinates() ([][]int, error) {
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
