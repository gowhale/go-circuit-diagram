// Package canvas is contains code for drawing components on
package canvas

import (
	"fmt"
	"go-circuit-diagram/pkg/common"
	"go-circuit-diagram/pkg/components"
	"image"
	"image/color"
	"log"
)

// Board represents a board to add components to
type Board struct {
	name          string
	width, height int
	magnification int
	elements      []components.Element
}

// NewBoard returns a new Board struct to add components to
func NewBoard(name string, width, height, magnification int) Board {
	return Board{
		name:          name,
		width:         width,
		height:        height,
		magnification: magnification,
	}
}

// AddElement adds an elemnt which is rendered when Draw is called
func (b *Board) AddElement(elem components.Element) {
	b.elements = append(b.elements, elem)
}

func enlargeCoordintes(coords [][]int, scale int) [][]int {
	newCoords := [][]int{}
	for _, cord := range coords {
		x := cord[0]
		y := cord[1]
		for i := 0; i < scale; i++ {
			for j := 0; j < scale; j++ {
				newCoords = append(newCoords, []int{(x * (scale - 1)) + x + i, (y * (scale - 1)) + y + j})
			}
		}
	}

	return newCoords
}

func (b *Board) fillCoordinates(img *image.RGBA, elem components.Element) error {
	cords, err := elem.GetCoordinates()
	if err != nil {
		return err
	}
	cords = enlargeCoordintes(cords, b.magnification)
	warning := false
	for _, cord := range cords {
		if cord[0] < (b.width*b.magnification) && cord[1] < (b.height*b.magnification) {
			img.Set(cord[0], cord[1], elem.GetColour())
		} else {
			warning = true
		}
	}
	if warning {
		log.Println("WARNING: some of this elements contents will not be shown as out of bounds")
	}
	return nil
}

// Draw creates image with all components drawn on
func (b *Board) Draw(o common.OS) error {
	width := b.width * b.magnification
	height := b.height * b.magnification

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.White)
		}
	}

	for _, elem := range b.elements {
		if err := b.fillCoordinates(img, elem); err != nil {
			return err
		}
	}

	f, err := o.Create(fmt.Sprintf("images/%s.png", b.name))
	if err != nil {
		return err
	}

	return o.Encode(f, img)
}
