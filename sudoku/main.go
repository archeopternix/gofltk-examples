package main

import (
	"fmt"

	"github.com/archeopternix/go-fltk"
)

const (
	gridSize   = 9
	subGrid    = 3
	cellSize   = 40
	padding    = 10
	subPadding = 4 // extra space between subgrids
)

var selectedCell *fltk.Output // Global variable to store the selected cell

func selectCell(i, j int) {
	fmt.Println(i, ":", j)
	// Remove frame from previous selected cell
	if selectedCell != nil {
		selectedCell.SetColor(fltk.BACKGROUND_COLOR)
		selectedCell.SetBox(fltk.BORDER_BOX)
	}

	selectedCell = cells[i][j]
	// Set blue frame for the new selected cell
	selectedCell.SetColor(fltk.BACKGROUND_COLOR)
	selectedCell.SetBox(fltk.THIN_UP_BOX)
	selectedCell.SetSelectionColor(fltk.ColorFromRgb(0, 120, 255)) // blue frame

	// Trigger your callback here
	onCellSelected(selectedCell)
}

func onCellSelected(cell *fltk.Output) {
	// Custom callback for when a cell is selected
	// You can replace this with your logic
	println("Cell selected!")
}

var cells [gridSize][gridSize]*fltk.Output

func main() {
	winWidth := gridSize*cellSize + 2*padding + (subGrid-1)*subPadding
	winHeight := gridSize*cellSize + 2*padding + (subGrid-1)*subPadding

	win := fltk.NewWindow(winWidth, winHeight)
	win.SetLabel("Sudoku Board - go-fltk (Read Only, Selectable)")

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			extraX := (j / subGrid) * subPadding
			extraY := (i / subGrid) * subPadding
			x := padding + j*cellSize + extraX
			y := padding + i*cellSize + extraY

			cell := fltk.NewOutput(x, y, cellSize-2, cellSize-2, "")
			cell.SetAlign(fltk.ALIGN_CENTER)
			cell.SetBox(fltk.BORDER_BOX)
			if ((i/subGrid)+(j/subGrid))%2 == 1 {
				cell.SetColor(fltk.ColorFromRgb(230, 230, 230)) // light gray
			}
			// Make selectable on click
			cell.SetCallback(func() {
				selectCell(i, j)
				win.Redraw()
			})
			cells[i][j] = cell
		}
	}

	win.End()
	win.Show()
	fltk.Run()
}
