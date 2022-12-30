package components

import (
	"fmt"
	"go-circuit-diagram/pkg/common"
	"image"
	"image/color"
)

const (
	ledFileName = "images/led.png"

	// North Direction which the LED can point
	North = "N"
	// East Direction which the LED can point
	East = "E"
	// South Direction which the LED can point
	South = "S"
	// West Direction which the LED can point
	West = "W"
)

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

func reverse1DSlice(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func reverse2DSlice(arr [][]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func transformPixels(pixels [][]int, direction string) [][]int {
	newPixels := make([][]int, len(pixels))
	copy(newPixels, pixels)
	if direction == North {
		return pixels
	}
	if direction == South {
		reverse2DSlice(newPixels)
		return newPixels
	}

	newPixels = [][]int{}
	for range pixels[0] {
		newPixels = append(newPixels, []int{})
	}
	for _, row := range pixels {
		for k, pixel := range row {
			newPixels[k] = append(newPixels[k], pixel)
		}
	}
	if direction == West {
		return newPixels
	}
	for _, row := range newPixels {
		reverse1DSlice(row)
	}
	return newPixels
}

// NewLED returns a LED config starting from specified x,y
func NewLED(startCoord common.Coordinate, direction string) (LEDConfig, error) {
	if _, ok := directions[direction]; !ok {
		return LEDConfig{}, fmt.Errorf("direction not valid")
	}
	return LEDConfig{
		startCoord: startCoord,
		LedPixels:  transformPixels(ledPixels, direction),
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
	if l.Direction == North {
		return common.NewCord(l.startCoord.GetX()+5, l.startCoord.GetY())
	}
	if l.Direction == East {
		return common.NewCord(l.startCoord.GetX()-1, l.startCoord.GetY()+5)
	}
	if l.Direction == South {
		return common.NewCord(l.startCoord.GetX()+5, l.startCoord.GetY()+len(ledPixels))
	}
	return common.NewCord(l.startCoord.GetX()+len(ledPixels), l.startCoord.GetY()+5)
}

// GetAnode gets the LED's anode coord
func (l *LEDConfig) GetAnode() common.Coordinate {
	if l.Direction == North {
		return common.NewCord(l.startCoord.GetX()+5, l.startCoord.GetY()+len(ledPixels))
	}
	if l.Direction == East {
		return common.NewCord(l.startCoord.GetX()+len(ledPixels), l.startCoord.GetY()+5)
	}
	if l.Direction == South {
		return common.NewCord(l.startCoord.GetX()+5, l.startCoord.GetY())
	}
	return common.NewCord(l.startCoord.GetX()-1, l.startCoord.GetY()+5)
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
