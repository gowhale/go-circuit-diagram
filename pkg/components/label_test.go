package components

import (
	"go-circuit-diagram/pkg/common"
	"image/color"
	"testing"

	"github.com/stretchr/testify/suite"
)

type labelTest struct {
	suite.Suite
}

func (*labelTest) SetupTest() {}

func Test_Label(t *testing.T) {
	suite.Run(t, new(labelTest))
}

func (g *labelTest) Test_NewLabel() {
	expectedL := LabelConfig{
		cord:   common.NewCord(0, 0),
		Colour: color.Black,
		text:   " ",
	}
	w := NewLabel(common.NewCord(0, 0), " ")
	g.Equal(expectedL, w)
}

func (g *labelTest) Test_GetCoordinates() {
	expectedCords := [][]int{
		{3, 3},
		{3, 4},
		{3, 5},
		{3, 6},
		{3, 7},
		{4, 7},
		{5, 7},
		{6, 7},
		{7, 7}}
	w := NewLabel(common.NewCord(5, 5), "l")
	cords, err := w.GetCoordinates()
	g.Nil(err)
	g.Equal(expectedCords, cords)
}

func (g *labelTest) Test_GetColour() {
	w := NewLabel(common.NewCord(0, 0), "l")
	g.Equal(color.Black, w.GetColour())
}
