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
	led1 := components.NewLED(5, 5)
	board.AddElement(&led1)
	led2 := components.NewLED(25, 5)
	board.AddElement(&led2)
	led3 := components.NewLED(65, 50)
	board.AddElement(&led3)
	led4 := components.NewLED(5, 75)
	board.AddElement(&led4)
	err = board.Draw(realOS)
	if err != nil {
		log.Fatalln(err)
	}
}
