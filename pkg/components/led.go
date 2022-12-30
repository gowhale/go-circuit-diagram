package components

import (
	"fmt"
	"go-circuit-diagram/pkg/common"
	"image"
	"image/color"
)

const (
	pixelEmpt   = 0
	pixelFill   = 1
	ledFileName = "images/led.png"

	North = "N"
	East  = "E"
	South = "S"
	West  = "W"
)

var ledPixels = [][]int{
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelFill, pixelFill},
	{pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt, pixelFill, pixelFill},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelFill},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
	{pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelFill, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt, pixelEmpt},
}

var directions = map[string]bool{
	North: true,
	East:  true,
	South: true,
	West:  true,
}

// LEDConfig is configuration for an LED component
type LEDConfig struct {
	startCoord common.Coordinate
	LedPixels  [][]int
	Colour     color.Color
	Direction  string
}

// NewLED returns a LED config starting from specified x,y
func NewLED(startCoord common.Coordinate, direction string) (LEDConfig, error) {
	if _, ok := directions[direction]; !ok {
		return LEDConfig{}, fmt.Errorf("direction not valid")
	}
	return LEDConfig{
		startCoord: startCoord,
		LedPixels:  ledPixels,
		Colour:     color.Black,
		Direction:  direction,
	}, nil
}

// GetColour gets the colour to render the element in
func (l *LEDConfig) GetColour() color.Color {
	return l.Colour
}

// GetCoordinates calculates cords to draw onto a canvas
func (l *LEDConfig) GetCoordinates() ([][]int, error) {
	cordsToDraw := [][]int{}
	for x := range l.LedPixels[0] {
		for y := range l.LedPixels {
			if l.LedPixels[y][x] == pixelFill {
				cordsToDraw = append(cordsToDraw, []int{x + l.startCoord.GetX(), y + l.startCoord.GetY()})
			}
		}
	}
	return cordsToDraw, nil
}

// GetCathode gets the LED's cathode coord
func (l *LEDConfig) GetCathode() common.Coordinate {
	return common.NewCord(l.startCoord.GetX()+5, l.startCoord.GetY())
}

// GetAnode gets the LED's anode coord
func (l *LEDConfig) GetAnode() common.Coordinate {
	return common.NewCord(l.startCoord.GetX()+5, l.startCoord.GetY()+len(ledPixels))
}

func validatePixelArray(pixelArray [][]int) error {
	width := len(pixelArray[0])
	for i, row := range pixelArray {
		if len(row) != width {
			return fmt.Errorf("row #%d is not same width as first row #1", i+1)
		}
	}
	return nil
}

// DrawLED will create pixel art of an LED
func DrawLED(o common.OS) error {
	return drawLEDImpl(o, ledPixels)
}

func drawLEDImpl(o common.OS, pixelArray [][]int) error {
	err := validatePixelArray(pixelArray)
	if err != nil {
		return err
	}

	width := len(pixelArray[0])
	height := len(pixelArray)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case pixelArray[y][x] == pixelEmpt: // upper left quadrant
				img.Set(x, y, color.White)
			case pixelArray[y][x] == pixelFill: // lower right quadrant
				img.Set(x, y, color.Black)
			default:
				return fmt.Errorf("pixel value not handled")
			}
		}
	}

	f, err := o.Create(ledFileName)
	if err != nil {
		return err
	}

	return o.Encode(f, img)
}
