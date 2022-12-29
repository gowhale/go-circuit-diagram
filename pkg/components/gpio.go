package components

import "image/color"

// GPIOConfig is configuration for a line on the canvas
type GPIOConfig struct {
	StartX, StartY int
	Colour         color.Color
}

// NewGPIO returns a GPIO config starting at specified x,y
func NewGPIO(startX, startY int) GPIOConfig {
	return GPIOConfig{
		StartX: startX,
		StartY: startY,
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
	cords = append(cords, []int{g.StartX - 1, g.StartY})
	cords = append(cords, []int{g.StartX + 1, g.StartY})
	cords = append(cords, []int{g.StartX + 1, g.StartY + 1})
	cords = append(cords, []int{g.StartX - 1, g.StartY + 1})
	cords = append(cords, []int{g.StartX, g.StartY - 1})
	cords = append(cords, []int{g.StartX, g.StartY + 1})
	cords = append(cords, []int{g.StartX - 1, g.StartY - 1})
	cords = append(cords, []int{g.StartX + 1, g.StartY - 1})
	return cords, nil
}
