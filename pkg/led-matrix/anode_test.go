// Package ledmatrix contains logic for creating LED matrix circuit diagrams
package ledmatrix

import (
	"testing"

	"github.com/gowhale/go-circuit-diagram/pkg/canvas"
	"github.com/gowhale/go-circuit-diagram/pkg/common"
	"github.com/gowhale/go-circuit-diagram/pkg/components"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type matrixTest struct {
	suite.Suite

	mockOS *common.MockOS
}

func (a *matrixTest) SetupTest() {
	a.mockOS = new(common.MockOS)
}

func Test_Matrix(t *testing.T) {
	suite.Run(t, new(matrixTest))
}

func (a *matrixTest) Test_addLEDS() {
	b := &canvas.Board{}
	elements := [][]components.LEDConfig{
		{components.LEDConfig{}, components.LEDConfig{}},
		{components.LEDConfig{}, components.LEDConfig{}},
	}
	addLEDS(b, elements, 2, 2)
	a.Equal(8, b.ElementCount())
}

func (a *matrixTest) Test_addVerticalWires() {
	b := &canvas.Board{}
	elements := [][]components.LEDConfig{
		{components.LEDConfig{}, components.LEDConfig{}},
		{components.LEDConfig{}, components.LEDConfig{}},
	}
	addVerticalWires(b, elements, 2, 2, []int{1, 2, 3, 4})
	a.Equal(4, b.ElementCount())
}

func (a *matrixTest) Test_addHorizontalWires() {
	b := &canvas.Board{}
	elements := [][]components.LEDConfig{
		{components.LEDConfig{}, components.LEDConfig{}},
		{components.LEDConfig{}, components.LEDConfig{}},
	}
	addHorizontalWires(b, elements, 2, 2, []int{1, 2, 3, 4})
	a.Equal(4, b.ElementCount())
}

func (a *matrixTest) Test_connectLEDToVerticalWire() {
	b := &canvas.Board{}
	elements := [][]components.LEDConfig{
		{components.LEDConfig{}, components.LEDConfig{}},
		{components.LEDConfig{}, components.LEDConfig{}},
	}
	connectLEDToVerticalWire(b, elements)
	a.Equal(8, b.ElementCount())
}

func (a *matrixTest) Test_addHorizontalGPIOs() {
	b := &canvas.Board{}
	elements := [][]components.LEDConfig{
		{components.LEDConfig{}, components.LEDConfig{}},
		{components.LEDConfig{}, components.LEDConfig{}},
	}
	addHorizontalGPIOs(b, elements)
	a.Equal(2, b.ElementCount())
}

func (a *matrixTest) Test_addVerticalGPIOs() {
	b := &canvas.Board{}
	elements := [][]components.LEDConfig{
		{components.LEDConfig{}, components.LEDConfig{}},
		{components.LEDConfig{}, components.LEDConfig{}},
	}
	addVerticalGPIOs(b, elements)
	a.Equal(2, b.ElementCount())
}

func (a *matrixTest) Test_CreateAnodeMatrix() {
	a.mockOS.On("Create", "images/test.png").Return(nil, nil)
	a.mockOS.On("Encode", mock.Anything, mock.Anything).Return(nil)
	err := CreateAnodeMatrix(a.mockOS, []int{1, 2}, []int{1, 2}, "test")
	a.Nil(err)
}
