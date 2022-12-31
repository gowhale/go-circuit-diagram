package components

import (
	"image/color"

	"github.com/gowhale/go-circuit-diagram/pkg/common"

	"github.com/gowhale/led-matrix-golang/pkg/matrix"
)

// LabelConfig represents two components attached
type LabelConfig struct {
	cord   common.Coordinate
	Colour color.Color
	text   string
}

// NewLabel returns a LabelConfig at specified x,y
func NewLabel(cord common.Coordinate, text string) LabelConfig {
	return LabelConfig{
		cord:   cord,
		Colour: color.Black,
		text:   text,
	}
}

// GetColour gets the colour to render the element in
func (g *LabelConfig) GetColour() color.Color {
	return g.Colour
}

// GetCoordinates calculates cords to draw onto a canvas
func (g *LabelConfig) GetCoordinates() ([][]int, error) {
	cords := [][]int{}
	pixels, err := matrix.ConcatanateLetters(g.text)
	if err != nil {
		return [][]int{}, err
	}
	for y := 0; y < len(pixels); y++ {
		for x := 0; x < len(pixels[y]); x++ {
			if pixels[y][x] == pixelFill {
				cords = append(cords, []int{x + g.cord.X - len(pixels[0])/2, y + g.cord.Y - len(pixels)/2})
			}
		}
	}
	return cords, nil
}
