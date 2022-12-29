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
	err = board.Draw(realOS)
	if err != nil {
		log.Fatalln(err)
	}
}
