// Package main runs the shopping list
package main

import (
	"go-circuit-diagram/pkg/canvas"
	"go-circuit-diagram/pkg/common"
	"go-circuit-diagram/pkg/components"
	"log"
)

func main() {
	realOS := &common.OSReal{}

	err := components.DrawLED(realOS)
	if err != nil {
		log.Fatalln(err)
	}

	board := canvas.NewBoard("test", 100, 100)

	// Adding some LEDS
	led1 := components.NewLED(5, 5)
	board.AddElement(&led1)
	led2 := components.NewLED(25, 5)
	board.AddElement(&led2)
	led3 := components.NewLED(65, 50)
	board.AddElement(&led3)
	led4 := components.NewLED(5, 75)
	board.AddElement(&led4)

	// Add some wires
	wire := components.NewWire(30, 20, 30, 50)
	board.AddElement(&wire)
	wire2 := components.NewWire(10, 50, 70, 50)
	board.AddElement(&wire2)
	wire3 := components.NewWire(10, 20, 10, 50)
	board.AddElement(&wire3)
	wire4 := components.NewWire(10, 65, 70, 65)
	board.AddElement(&wire4)
	wire5 := components.NewWire(10, 65, 10, 75)
	board.AddElement(&wire5)

	// Add some GPIOS
	gpio1 := components.NewGPIO(30, 3)
	board.AddElement(&gpio1)

	gpio2 := components.NewGPIO(10, 3)
	board.AddElement(&gpio2)

	gpio3 := components.NewGPIO(10, 92)
	board.AddElement(&gpio3)

	err = board.Draw(realOS)
	if err != nil {
		log.Fatalln(err)
	}
}
