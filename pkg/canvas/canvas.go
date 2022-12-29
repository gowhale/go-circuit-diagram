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

type Board struct {
	name          string
	width, height int
	elements      []components.Element
}

func NewBoard(name string, width, height int) Board {
	return Board{
		name:   name,
		width:  width,
		height: height,
	}
}

func (b *Board) AddElement(elem components.Element) {
	b.elements = append(b.elements, elem)
}

func (b *Board) fillCoordinates(img *image.RGBA, cords [][]int) {
	warning := false
	for _, cord := range cords {
		if cord[0] < b.width && cord[1] < b.height {
			img.Set(cord[0], cord[1], color.Black)
		} else {
			warning = true
		}
	}
	if warning {
		log.Println("WARNING: some of this elements contents will not be shown as out of bounds")
	}
}

func (b *Board) Draw(o common.OS) error {
	width := b.width
	height := b.height

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
		cordsToDraw := elem.GetCoordinates()
		b.fillCoordinates(img, cordsToDraw)
	}

	f, err := o.Create(fmt.Sprintf("images/%s.png", b.name))
	if err != nil {
		return err
	}

	return o.Encode(f, img)
}
