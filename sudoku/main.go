package main

import (
	"github.com/archeopternix/go-fltk"
)

const (
	gridSize = 9
	cellSize = 40
	padding  = 10
)

func main() {
	win := fltk.NewWindow(gridSize*cellSize+2*padding, gridSize*cellSize+2*padding)
	win.SetLabel("Sudoku Board - go-fltk")

	// 2D array of input boxes
	var cells [gridSize][gridSize]*fltk.Input

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			x := padding + j*cellSize
			y := padding + i*cellSize
			input := fltk.NewInput(x, y, cellSize-2, cellSize-2, "")
			//	input.SetMaximumSize(cellSize-2, cellSize-2)
			input.SetAlign(fltk.ALIGN_CENTER)
			//	input.SetTextSize(20)
			input.SetCallback(func() {
				val := input.Value()
				if len(val) > 1 || (len(val) == 1 && (val[0] < '1' || val[0] > '9')) {
					input.SetValue("")
				}
			})
			cells[i][j] = input
		}
	}

	win.End()
	win.Show()
	fltk.Run()
}
