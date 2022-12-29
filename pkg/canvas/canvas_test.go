package canvas

import (
	"fmt"
	"testing"

	"go-circuit-diagram/pkg/common"
	"go-circuit-diagram/pkg/components"

	fruit "github.com/gowhale/go-test-data/pkg/fruits"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	createFunc = "Create"
)

type mainTest struct {
	suite.Suite

	mockOS  *common.MockOS
	testLed components.LEDConfig
}

func (m *mainTest) SetupTest() {
	m.mockOS = new(common.MockOS)
	m.testLed = components.NewLED(0, 0)
}

func TestMainTest(t *testing.T) {
	suite.Run(t, new(mainTest))
}

func (m *mainTest) Test_NewBoard() {
	expectedB := Board{
		name:   fruit.Apple,
		width:  100,
		height: 100,
	}
	b := NewBoard(fruit.Apple, 100, 100)
	m.Equal(expectedB, b)
}

func (m *mainTest) Test_AddElement() {
	b := NewBoard(fruit.Apple, 100, 100)
	m.Equal(0, len(b.elements))
	b.AddElement(&m.testLed)
	m.Equal(1, len(b.elements))
}

func (m *mainTest) TestDrawLED_Pass() {
	m.mockOS.On(createFunc, fmt.Sprintf("images/%s.png", fruit.Apple)).Return(nil, nil)
	m.mockOS.On("Encode", mock.Anything, mock.Anything).Return(nil)

	b := NewBoard(fruit.Apple, 100, 100)
	b.AddElement(&m.testLed)
	err := b.Draw(m.mockOS)
	m.Nil(err)
}

func (m *mainTest) TestDrawLED_Create_Error() {
	m.mockOS.On(createFunc, fmt.Sprintf("images/%s.png", fruit.Apple)).Return(nil, fmt.Errorf("create err"))

	b := NewBoard(fruit.Apple, 100, 100)
	b.AddElement(&m.testLed)
	err := b.Draw(m.mockOS)
	m.EqualError(err, "create err")
}
