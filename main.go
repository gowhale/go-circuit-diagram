// Package main runs the shopping list
package main

import (
	"go-circuit-diagram/pkg/common"
	"go-circuit-diagram/pkg/components"
	"log"
)

func main() {
	err := components.DrawLED(&common.OSReal{})
	if err != nil {
		log.Fatalln(err)
	}
}
