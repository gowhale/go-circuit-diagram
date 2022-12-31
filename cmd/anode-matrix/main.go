// Package main creates a diagram from anode LED Matrix
package main

import (
	"fmt"
	"log"

	"github.com/gowhale/go-circuit-diagram/pkg/canvas"
	"github.com/gowhale/go-circuit-diagram/pkg/common"
	"github.com/gowhale/go-circuit-diagram/pkg/components"
)

func main() {
	createAnodeMatrix()
}

func addLEDS(board *canvas.Board, ledMatrix [][]components.LEDConfig, cols, rows int) {
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			led, err := components.NewLED(common.NewCord(40+(x*30), 40+(y*30)), components.North)
			if err != nil {
				log.Fatal(err)
			}
			board.AddElement(&led)

			connection := components.NewConnection(led.GetAnode())
			board.AddElement(&connection)

			ledMatrix[y][x] = led
		}
	}
}

func addVerticalWires(board *canvas.Board, ledMatrix [][]components.LEDConfig, cols, rows int) {
	for i := 0; i < cols; i++ {
		x := ledMatrix[0][i].GetAnode().X + 15
		startY := ledMatrix[0][i].GetAnode().Y - 30
		endY := ledMatrix[rows-1][i].GetAnode().Y
		wire := components.NewWire(common.NewCord(x, startY), common.NewCord(x, endY))
		board.AddElement(&wire)

		label := components.NewLabel(common.NewCord(x, ledMatrix[0][i].GetAnode().Y-40), fmt.Sprintf("%d%d", 9-i, i))
		board.AddElement(&label)
	}
}

func addHorizontalWires(board *canvas.Board, ledMatrix [][]components.LEDConfig, cols, rows int) {
	for i := 0; i < rows; i++ {
		y := ledMatrix[i][0].GetAnode().Y
		startX := ledMatrix[i][0].GetAnode().X - 20
		endX := ledMatrix[i][cols-1].GetAnode().X
		wire := components.NewWire(common.NewCord(startX, y), common.NewCord(endX, y))
		board.AddElement(&wire)

		label := components.NewLabel(common.NewCord(startX-10, y), fmt.Sprintf("%d%d", 9-i, i))
		board.AddElement(&label)
	}
}

func createAnodeMatrix() {
	realOS := &common.OSReal{}

	board := canvas.NewBoard("anode-matrix-16", 300, 300, 10)

	cols := 8
	rows := 8

	ledMatrix := make([][]components.LEDConfig, rows)
	for i := range ledMatrix {
		ledMatrix[i] = make([]components.LEDConfig, cols)
	}

	addLEDS(&board, ledMatrix, cols, rows)

	for y := range ledMatrix {
		for x := range ledMatrix[y] {
			top := ledMatrix[y][x].GetCathode()
			topLED := common.NewCord(top.X+15, top.Y)
			wire := components.NewWire(top, topLED)
			board.AddElement(&wire)

			connection := components.NewConnection(topLED)
			board.AddElement(&connection)
		}
	}

	addHorizontalWires(&board, ledMatrix, cols, rows)

	addVerticalWires(&board, ledMatrix, cols, rows)

	// horzonal gpio
	for _, led := range ledMatrix[0] {
		y := led.GetAnode().Y - 30
		x := led.GetAnode().X + 15
		newGPIO := components.NewGPIO(common.NewCord(x, y))
		board.AddElement(&newGPIO)
	}

	// vertical gpio
	for _, led := range ledMatrix {
		leftGPIO := common.NewCord(led[0].GetAnode().X-20, led[0].GetAnode().Y)
		newGPIO := components.NewGPIO(leftGPIO)
		board.AddElement(&newGPIO)
	}

	// Render board
	err := board.Draw(realOS)
	if err != nil {
		log.Fatalln(err)
	}
}
