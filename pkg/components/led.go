// Package components contains code for drawing electronical components
package components

import (
	"fmt"
	"go-circuit-diagram/pkg/common"
	"image"
	"image/color"
)

const (
	pixelEmpty  = 0
	pixelFilled = 1
	ledFileName = "images/led.png"
)

var ledPixels = [][]int{
	{pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty},
	{pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty},
	{pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty},
	{pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty},
	{pixelFilled, pixelFilled, pixelFilled, pixelFilled, pixelFilled, pixelFilled, pixelFilled, pixelFilled, pixelFilled},
	{pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty},
	{pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty},
	{pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty},
	{pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty},
	{pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled},
	{pixelFilled, pixelFilled, pixelFilled, pixelFilled, pixelFilled, pixelFilled, pixelFilled, pixelFilled, pixelFilled},
	{pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty},
	{pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty},
	{pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty},
	{pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty},
	{pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty, pixelFilled, pixelEmpty, pixelEmpty, pixelEmpty, pixelEmpty},
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
			case pixelArray[y][x] == pixelEmpty: // upper left quadrant
				img.Set(x, y, color.White)
			case pixelArray[y][x] == pixelFilled: // lower right quadrant
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
