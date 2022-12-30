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

	board := canvas.NewBoard("test", 100, 100, 10)

	// Adding some LEDS
	led1, err := components.NewLED(common.NewCord(5, 5), components.North)
	if err != nil {
		log.Fatal(err)
	}
	board.AddElement(&led1)
	led2, err := components.NewLED(common.NewCord(25, 5), components.North)
	if err != nil {
		log.Fatal(err)
	}
	board.AddElement(&led2)
	led3, err := components.NewLED(common.NewCord(65, 50), components.North)
	if err != nil {
		log.Fatal(err)
	}
	board.AddElement(&led3)
	led4, err := components.NewLED(common.NewCord(5, 75), components.North)
	if err != nil {
		log.Fatal(err)
	}
	board.AddElement(&led4)
	led5, err := components.NewLED(common.NewCord(83, 5), components.South)
	if err != nil {
		log.Fatal(err)
	}
	board.AddElement(&led5)

	led6, err := components.NewLED(common.NewCord(45, 25), components.East)
	if err != nil {
		log.Fatal(err)
	}
	board.AddElement(&led6)

	led7, err := components.NewLED(common.NewCord(45, 75), components.West)
	if err != nil {
		log.Fatal(err)
	}
	board.AddElement(&led7)

	// Add some wires
	wire := components.NewWire(common.NewCord(30, 20), common.NewCord(30, 50))
	board.AddElement(&wire)
	wire2 := components.NewWire(common.NewCord(10, 50), common.NewCord(70, 50))
	board.AddElement(&wire2)
	wire3 := components.NewWire(common.NewCord(10, 20), common.NewCord(10, 50))
	board.AddElement(&wire3)
	wire4 := components.NewWire(common.NewCord(10, 65), common.NewCord(70, 65))
	board.AddElement(&wire4)
	wire5 := components.NewWire(common.NewCord(10, 65), common.NewCord(10, 75))
	board.AddElement(&wire5)
	wire6 := components.NewWire(common.NewCord(led3.GetCathode().X, led5.GetAnode().Y), led3.GetCathode())
	board.AddElement(&wire6)
	wire7 := components.NewWire(led5.GetAnode(), common.NewCord(led3.GetCathode().X, led5.GetAnode().Y))
	board.AddElement(&wire7)

	// Add some GPIOS
	gpio1 := components.NewGPIO(common.NewCord(30, 3))
	board.AddElement(&gpio1)

	gpio2 := components.NewGPIO(common.NewCord(10, 3))
	board.AddElement(&gpio2)

	gpio3 := components.NewGPIO(common.NewCord(10, 92))
	board.AddElement(&gpio3)

	gpio4 := components.NewGPIO(led5.GetCathode())
	board.AddElement(&gpio4)

	err = board.Draw(realOS)
	if err != nil {
		log.Fatalln(err)
	}
}
