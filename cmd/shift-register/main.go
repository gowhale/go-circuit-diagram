// Package main will show a sample shift register components
package main

import (
	"log"

	"github.com/gowhale/go-circuit-diagram/pkg/canvas"
	"github.com/gowhale/go-circuit-diagram/pkg/common"
	"github.com/gowhale/go-circuit-diagram/pkg/components"
)

func main() {
	realOS := &common.OSReal{}

	board := canvas.NewBoard("shift-example", 30, 50, 10)

	shift := components.NewShift(common.NewCord(5, 5))
	board.AddElement(&shift)

	gpio1 := components.NewGPIO(shift.GetPin1())
	board.AddElement(&gpio1)

	gpio2 := components.NewGPIO(shift.GetPin2())
	board.AddElement(&gpio2)

	gpio3 := components.NewGPIO(shift.GetPin3())
	board.AddElement(&gpio3)

	gpio4 := components.NewGPIO(shift.GetPin4())
	board.AddElement(&gpio4)

	gpio5 := components.NewGPIO(shift.GetPin5())
	board.AddElement(&gpio5)

	gpio6 := components.NewGPIO(shift.GetPin6())
	board.AddElement(&gpio6)

	gpio7 := components.NewGPIO(shift.GetPin7())
	board.AddElement(&gpio7)

	gpio11 := components.NewGPIO(shift.GetPin11())
	board.AddElement(&gpio11)

	gpio12 := components.NewGPIO(shift.GetPin12())
	board.AddElement(&gpio12)

	gpio13 := components.NewGPIO(shift.GetPin13())
	board.AddElement(&gpio13)

	gpio14 := components.NewGPIO(shift.GetPin14())
	board.AddElement(&gpio14)

	gpio15 := components.NewGPIO(shift.GetPin15())
	board.AddElement(&gpio15)

	gpio16 := components.NewGPIO(shift.GetPin16())
	board.AddElement(&gpio16)

	gpio17 := components.NewGPIO(shift.GetPin17())
	board.AddElement(&gpio17)

	// Render board
	err := board.Draw(realOS)
	if err != nil {
		log.Fatalln(err)
	}
}
