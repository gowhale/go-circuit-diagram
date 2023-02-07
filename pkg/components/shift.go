package components

import (
	"image/color"

	"github.com/gowhale/go-circuit-diagram/pkg/common"
)

// ShiftConfig is configuration for an LED component
type ShiftConfig struct {
	startCoord common.Coordinate
	cords      [][]int
	Colour     color.Color
}

// NewLED returns a LED config starting from specified x,y
func NewShift(startCoord common.Coordinate) ShiftConfig {
	return ShiftConfig{
		startCoord: startCoord,
		cords:      shiftPixels,
		Colour:     color.Black,
	}
}

// GetColour gets the colour to render the element in
func (s *ShiftConfig) GetColour() color.Color {
	return s.Colour
}

// GetCoordinates calculates cords to draw onto a canvas
func (s *ShiftConfig) GetCoordinates() ([][]int, error) {
	cordsToDraw := [][]int{}
	for x := range s.cords[0] {
		for y := range s.cords {
			if s.cords[y][x] == pixelFill {
				cordsToDraw = append(cordsToDraw, []int{x + s.startCoord.GetX(), y + s.startCoord.GetY()})
			}
		}
	}
	return cordsToDraw, nil
}

func (s *ShiftConfig) GetPin1() common.Coordinate {
	return common.NewCord(s.startCoord.GetX()-1, s.startCoord.GetY()+2)
}

func (s *ShiftConfig) GetPin2() common.Coordinate {
	return common.NewCord(s.startCoord.GetX()-1, s.startCoord.GetY()+6)
}
