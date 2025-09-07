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

var midBlue = fltk.ColorFromRgb(66, 135, 245)
var lightGrey = fltk.ColorFromRgb(230, 230, 230)
var midGrey = fltk.BACKGROUND_COLOR

type SudokuGrid struct {
	*fltk.Group
	cells        [gridSize][gridSize]*fltk.Button
	selectedCell *fltk.Button
}

func NewSudokuGrid(x, y, width, height int) *SudokuGrid {
	grid := &SudokuGrid{
		Group: fltk.NewGroup(x, y, width, height),
	}
	grid.Begin()
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			extraX := (j / subGrid) * subPadding
			extraY := (i / subGrid) * subPadding
			cellX := j*cellSize + extraX
			cellY := i*cellSize + extraY

			cell := fltk.NewButton(cellX+x, cellY+y, cellSize-2, cellSize-2, "")
			cell.SetAlign(fltk.ALIGN_CENTER)
			cell.SetBox(fltk.BORDER_BOX)
			if ((i/subGrid)+(j/subGrid))%2 == 1 {
				cell.SetColor(lightGrey) // light gray
			} else {
				cell.SetColor(midGrey) // light gray
			}
			// Make selectable on click
			i_copy := i
			j_copy := j
			cell.SetCallback(func() {
				grid.selectCell(i_copy, j_copy)
				grid.Redraw()
			})
			grid.cells[i][j] = cell
		}
	}
	grid.End()
	return grid
}

func (g *SudokuGrid) selectCell(i, j int) {
	fmt.Println(i, ":", j)
	// Remove frame from previous selected cell
	if g.selectedCell != nil {
		g.selectedCell.SetBox(fltk.BORDER_BOX)
		if ((i/subGrid)+(j/subGrid))%2 == 1 {
			g.selectedCell.SetColor(midGrey)
		} else {
			g.selectedCell.SetColor(lightGrey) // light gray
		}
	}

	g.selectedCell = g.cells[i][j]
	// Set blue color for the new selected cell
	g.selectedCell.SetColor(midBlue)
	g.selectedCell.SetBox(fltk.THIN_UP_BOX)
	g.Redraw() // Redraw the SudokuGrid after selecting a cell
}

func (g *SudokuGrid) onCellSelected(cell *fltk.Button) {
	// Custom callback for when a cell is selected
	// You can replace this with your logic
	println("Cell selected!")
}

func (g *SudokuGrid) GetSelectedCell() *fltk.Button {
	return g.selectedCell
}

func (g *SudokuGrid) SetCell(i, j int, text string) {
	g.cells[i][j].SetLabel(text)
	g.Redraw()
}

func main() {
	winWidth := gridSize*cellSize + 2*padding + (subGrid-1)*subPadding
	winHeight := gridSize*cellSize + 2*padding + (subGrid-1)*subPadding

	win := fltk.NewWindow(winWidth, winHeight)
	win.SetLabel("Sudoku Board - go-fltk (Read Only, Selectable)")

	sudokuGrid := NewSudokuGrid(padding, padding, winWidth-2*padding, winHeight-2*padding)
	sudokuGrid.SetCell(0, 0, "5")
	sudokuGrid.SetCell(1, 2, "3")
	sudokuGrid.SetCell(8, 8, "9")

	win.End()
	win.Show()
	fltk.Run()
}
