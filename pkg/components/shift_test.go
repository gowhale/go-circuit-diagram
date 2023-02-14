package components

import (
	"image/color"
	"testing"

	"github.com/gowhale/go-circuit-diagram/pkg/common"

	"github.com/stretchr/testify/suite"
)

type shiftTest struct {
	suite.Suite

	mockOS *common.MockOS
}

func (m *shiftTest) SetupTest() {
	m.mockOS = new(common.MockOS)
}

func TestShift(t *testing.T) {
	suite.Run(t, new(shiftTest))
}

func (m *shiftTest) Test_NewLED() {
	expectedL := ShiftConfig{
		startCoord: common.NewCord(0, 0),
		Colour:     color.Black,
		cords:      shiftPixels,
	}
	l := NewShift(common.NewCord(0, 0))
	m.Equal(expectedL, l)
}

func (m *shiftTest) Test_GetColour() {
	l := NewShift(common.NewCord(0, 0))
	m.Equal(color.Black, l.GetColour())
}

func (m *shiftTest) Test_GetCoordinates() {
	expectedL := ShiftConfig{
		startCoord: common.NewCord(0, 0),
		Colour:     color.Black,
		cords:      [][]int{{0, 0}, {1, 1}},
	}
	expectedCords := [][]int{{0, 1}, {1, 1}}
	cords, err := expectedL.GetCoordinates()
	m.Nil(err)
	m.Equal(expectedCords, cords)
}

func (m *shiftTest) Test_GetPins() {
	expectedL := ShiftConfig{
		startCoord: common.NewCord(0, 0),
		Colour:     color.Black,
		cords:      [][]int{{0, 0}, {1, 1}},
	}

	expectedCord := common.Coordinate{-1, 2}
	pinCord := expectedL.GetPin1()
	m.Equal(expectedCord, pinCord)

	expectedCord = common.Coordinate{-1, 6}
	pinCord = expectedL.GetPin2()
	m.Equal(expectedCord, pinCord)

	expectedCord = common.Coordinate{-1, 10}
	pinCord = expectedL.GetPin3()
	m.Equal(expectedCord, pinCord)

	expectedCord = common.Coordinate{-1, 14}
	pinCord = expectedL.GetPin4()
	m.Equal(expectedCord, pinCord)

	expectedCord = common.Coordinate{-1, 18}
	pinCord = expectedL.GetPin5()
	m.Equal(expectedCord, pinCord)

	expectedCord = common.Coordinate{-1, 18}
	pinCord = expectedL.GetPin5()
	m.Equal(expectedCord, pinCord)

	expectedCord = common.Coordinate{-1, 22}
	pinCord = expectedL.GetPin6()
	m.Equal(expectedCord, pinCord)

	expectedCord = common.Coordinate{-1, 26}
	pinCord = expectedL.GetPin7()
	m.Equal(expectedCord, pinCord)

	expectedCord = common.Coordinate{19, 2}
	pinCord = expectedL.GetPin11()
	m.Equal(expectedCord, pinCord)

	expectedCord = common.Coordinate{19, 6}
	pinCord = expectedL.GetPin12()
	m.Equal(expectedCord, pinCord)
}
