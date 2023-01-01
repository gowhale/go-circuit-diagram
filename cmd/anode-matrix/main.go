// Package main creates a diagram from anode LED Matrix
package main

import (
	"log"

	"github.com/gowhale/go-circuit-diagram/pkg/common"
	ledmatrix "github.com/gowhale/go-circuit-diagram/pkg/led-matrix"
)

func main() {
	rowPins := []int{1, 2, 3, 4, 5, 6, 7, 8}
	colPins := []int{9, 10, 11, 12, 14, 15, 16}
	err := ledmatrix.CreateAnodeMatrix(&common.OSReal{}, rowPins, colPins, "8-8-pins")
	if err != nil {
		log.Fatalln(err)
	}
}
