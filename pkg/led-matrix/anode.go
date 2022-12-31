package ledmatrix

import (
	"fmt"
	"log"

	"github.com/gowhale/go-circuit-diagram/pkg/canvas"
	"github.com/gowhale/go-circuit-diagram/pkg/common"
	"github.com/gowhale/go-circuit-diagram/pkg/components"
)

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

func addVerticalWires(board *canvas.Board, ledMatrix [][]components.LEDConfig, cols, rows int, pins []int) {
	for i := 0; i < cols; i++ {
		x := ledMatrix[0][i].GetAnode().X + 15
		startY := ledMatrix[0][i].GetAnode().Y - 30
		endY := ledMatrix[rows-1][i].GetAnode().Y
		wire := components.NewWire(common.NewCord(x, startY), common.NewCord(x, endY))
		board.AddElement(&wire)

		label := components.NewLabel(common.NewCord(x, ledMatrix[0][i].GetAnode().Y-40), fmt.Sprintf("%d", pins[i]))
		board.AddElement(&label)
	}
}

func addHorizontalWires(board *canvas.Board, ledMatrix [][]components.LEDConfig, cols, rows int, pins []int) {
	for i := 0; i < rows; i++ {
		y := ledMatrix[i][0].GetAnode().Y
		startX := ledMatrix[i][0].GetAnode().X - 20
		endX := ledMatrix[i][cols-1].GetAnode().X
		wire := components.NewWire(common.NewCord(startX, y), common.NewCord(endX, y))
		board.AddElement(&wire)

		pinAsString := fmt.Sprintf("%d", pins[i])
		log.Println(pinAsString)
		label := components.NewLabel(common.NewCord(startX-10, y), pinAsString)
		board.AddElement(&label)
	}
}

func connectLEDToVerticalWire(board *canvas.Board, ledMatrix [][]components.LEDConfig, cols, rows int, pins []int) {
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
}

func addVerticalGPIOs(board *canvas.Board, ledMatrix [][]components.LEDConfig) {
	for _, led := range ledMatrix {
		leftGPIO := common.NewCord(led[0].GetAnode().X-20, led[0].GetAnode().Y)
		newGPIO := components.NewGPIO(leftGPIO)
		board.AddElement(&newGPIO)
	}
}

func addHorizontalGPIOs(board *canvas.Board, ledMatrix [][]components.LEDConfig) {
	for _, led := range ledMatrix {
		leftGPIO := common.NewCord(led[0].GetAnode().X-20, led[0].GetAnode().Y)
		newGPIO := components.NewGPIO(leftGPIO)
		board.AddElement(&newGPIO)
	}
}

// CreateAnodeMatrix generates an image for an led matrix
func CreateAnodeMatrix(o common.OS, rowPins, colPins []int, imageName string) error {
	board := canvas.NewBoard(imageName, 40+(len(colPins)*30), 40+(len(rowPins)*30), 10)

	cols := len(colPins)
	rows := len(rowPins)

	ledMatrix := make([][]components.LEDConfig, rows)
	for i := range ledMatrix {
		ledMatrix[i] = make([]components.LEDConfig, cols)
	}

	addLEDS(&board, ledMatrix, cols, rows)

	connectLEDToVerticalWire(&board, ledMatrix, cols, rows, rowPins)

	addHorizontalWires(&board, ledMatrix, cols, rows, rowPins)

	addVerticalWires(&board, ledMatrix, cols, rows, colPins)

	addHorizontalGPIOs(&board, ledMatrix)

	addVerticalGPIOs(&board, ledMatrix)

	// Render board
	return board.Draw(o)
}
