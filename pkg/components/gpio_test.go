package components

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/suite"
)

type gpioTest struct {
	suite.Suite
}

func (*gpioTest) SetupTest() {}

func Test_GPIO(t *testing.T) {
	suite.Run(t, new(gpioTest))
}

func (g *gpioTest) Test_NewGPIO() {
	expectedW := GPIOConfig{
		StartX: 0,
		StartY: 0,
		Colour: color.RGBA{0, 255, 0, 0xff},
	}
	w := NewGPIO(0, 0)
	g.Equal(expectedW, w)
}

func (g *gpioTest) Test_GetCoordinates() {
	expectedCords := [][]int{
		{4, 5},
		{6, 5},
		{6, 6},
		{4, 6},
		{5, 4},
		{5, 6},
		{4, 4},
		{6, 4}}
	w := NewGPIO(5, 5)
	cords, err := w.GetCoordinates()
	g.Nil(err)
	g.Equal(expectedCords, cords)
}

func (g *gpioTest) Test_GetColour() {
	w := NewGPIO(5, 5)
	g.Equal(color.RGBA{0, 255, 0, 0xff}, w.GetColour())
}
