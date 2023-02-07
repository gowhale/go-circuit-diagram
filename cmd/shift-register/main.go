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

	// Render board
	err := board.Draw(realOS)
	if err != nil {
		log.Fatalln(err)
	}
}
