package components

import (
	"image/color"
	"testing"

	"github.com/gowhale/go-circuit-diagram/pkg/common"

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
		cord:   common.NewCord(0, 0),
		Colour: color.RGBA{0, 255, 0, 0xff},
	}
	w := NewGPIO(common.NewCord(0, 0))
	g.Equal(expectedW, w)
}

func (g *gpioTest) Test_GetCoordinates() {
	expectedCords := [][]int{
		{4, 5},
		{6, 5},
		{6, 6},
		{4, 6},
		{5, 5},
		{5, 4},
		{5, 6},
		{4, 4},
		{6, 4}}
	w := NewGPIO(common.NewCord(5, 5))
	cords, err := w.GetCoordinates()
	g.Nil(err)
	g.Equal(expectedCords, cords)
}

func (g *gpioTest) Test_GetColour() {
	w := NewGPIO(common.NewCord(5, 5))
	g.Equal(color.RGBA{0, 255, 0, 0xff}, w.GetColour())
}
