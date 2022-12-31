package components

import (
	"image/color"
	"testing"

	"github.com/gowhale/go-circuit-diagram/pkg/common"

	"github.com/stretchr/testify/suite"
)

const (
	createFunc = "Create"
)

type wireTest struct {
	suite.Suite
}

func (*wireTest) SetupTest() {}

func Test_Wire(t *testing.T) {
	suite.Run(t, new(wireTest))
}

func (m *wireTest) Test_NewWire() {
	expectedW := WireConfig{
		startCoord: common.NewCord(0, 0),
		endCoord:   common.NewCord(0, 10),
		Colour:     color.Black,
	}
	w := NewWire(common.NewCord(0, 0), common.NewCord(0, 10))
	m.Equal(expectedW, w)
}

func (m *wireTest) Test_GetCoordinates_Vertical_Pass() {
	expectedCords := [][]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{0, 5},
		{0, 6},
		{0, 7},
		{0, 8},
		{0, 9},
		{0, 10},
	}

	w := NewWire(common.NewCord(0, 0), common.NewCord(0, 10))
	cords, err := w.GetCoordinates()
	m.Nil(err)
	m.Equal(expectedCords, cords)
}

func (m *wireTest) Test_GetCoordinates_Horizontal_Pass() {
	expectedCords := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
		{6, 0},
		{7, 0},
		{8, 0},
		{9, 0},
		{10, 0},
	}

	w := NewWire(common.NewCord(0, 0), common.NewCord(10, 0))
	cords, err := w.GetCoordinates()
	m.Nil(err)
	m.Equal(expectedCords, cords)
}

func (m *wireTest) Test_GetCoordinates_Error() {
	expectedCords := [][]int{}

	w := NewWire(common.NewCord(0, 0), common.NewCord(10, 1))
	cords, err := w.GetCoordinates()
	m.EqualError(err, "only straight lines, horizontal or vertical")
	m.Equal(expectedCords, cords)
}

func (m *wireTest) Test_GetColour() {
	w := NewWire(common.NewCord(0, 0), common.NewCord(0, 10))
	m.Equal(color.Black, w.GetColour())
}
