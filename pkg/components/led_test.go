package components

import (
	"fmt"
	"testing"

	"go-circuit-diagram/pkg/common"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ledTest struct {
	suite.Suite

	mockOS *common.MockOS
}

func (m *ledTest) SetupTest() {
	m.mockOS = new(common.MockOS)
}

func TestMainTest(t *testing.T) {
	suite.Run(t, new(ledTest))
}

func (m *ledTest) TestDrawLED_Pass() {
	m.mockOS.On(createFunc, ledFileName).Return(nil, nil)
	m.mockOS.On("Encode", mock.Anything, mock.Anything).Return(nil)

	err := DrawLED(m.mockOS)
	m.Nil(err)
}

func (m *ledTest) Test_NewLED() {
	expectedL := LEDConfig{
		StartX:    0,
		StartY:    0,
		LedPixels: ledPixels,
	}
	l := NewLED(0, 0)
	m.Equal(expectedL, l)
}

func (m *ledTest) TestDrawLED_Encode_Error() {
	m.mockOS.On(createFunc, ledFileName).Return(nil, nil)
	m.mockOS.On("Encode", mock.Anything, mock.Anything).Return(fmt.Errorf("encode err"))

	err := DrawLED(m.mockOS)
	m.EqualError(err, "encode err")
}

func (m *ledTest) TestDrawLED_Create_Error() {
	m.mockOS.On(createFunc, ledFileName).Return(nil, fmt.Errorf("create err"))

	err := DrawLED(m.mockOS)
	m.EqualError(err, "create err")
}

func (m *ledTest) TestDrawLED_Invalid_Pixels_Error() {
	err := drawLEDImpl(m.mockOS, [][]int{
		{pixelEmpt, pixelFill},
		{pixelEmpt},
	})
	m.EqualError(err, "row #2 is not same width as first row #1")
}

func (m *ledTest) TestDrawLED_Invalid_Pixel_Config_Error() {
	err := drawLEDImpl(m.mockOS, [][]int{
		{pixelEmpt, pixelFill},
		{pixelEmpt, 21},
	})
	m.EqualError(err, "pixel value not handled")
}
