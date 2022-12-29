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

// LEDConfig is configuration for an LED component
type LEDConfig struct {
	StartX, StartY int
	LedPixels      [][]int
}

// NewLED returns a LED config starting from specified x,y
func NewLED(startX, startY int) LEDConfig {
	return LEDConfig{
		StartX:    startX,
		StartY:    startY,
		LedPixels: ledPixels,
	}
}

// GetCoordinates calculates cords to draw onto a canvas
func (l *LEDConfig) GetCoordinates() ([][]int, error) {
	cordsToDraw := [][]int{}
	for x := range l.LedPixels[0] {
		for y := range l.LedPixels {
			if l.LedPixels[y][x] == pixelFill {
				cordsToDraw = append(cordsToDraw, []int{x + l.StartX, y + l.StartY})
			}
		}
	}
	return cordsToDraw, nil
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
