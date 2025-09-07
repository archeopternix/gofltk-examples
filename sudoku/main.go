package main

import (
	"strconv"

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

type Cell struct {
	Value  string
	Widget *fltk.Button
	i      int
	j      int
	active bool
}

func NewCell(g *SudokuGrid, i, j int) *Cell {
	extraX := (j / subGrid) * subPadding
	extraY := (i / subGrid) * subPadding
	cellX := j*cellSize + extraX + padding
	cellY := i*cellSize + extraY + padding

	cell := &Cell{
		Value:  "",
		Widget: fltk.NewButton(cellX, cellY, cellSize-2, cellSize-2, ""),
		i:      i,
		j:      j,
		active: false,
	}

	cell.Widget.SetAlign(fltk.ALIGN_CENTER)
	cell.Widget.SetBox(fltk.BORDER_BOX)
	if ((i/subGrid)+(j/subGrid))%2 == 1 {
		cell.Widget.SetColor(fltk.ColorFromRgb(230, 230, 230)) // light gray
	}
	// Make selectable on click
	i_copy := i
	j_copy := j
	cell.Widget.SetCallback(func() {
		g.SelectCell(i_copy, j_copy, true)
		g.Redraw()
	})

	return cell
}

type SudokuGrid struct {
	*fltk.Group
	cells         [gridSize][gridSize]*Cell
	activeCellRow int
	activeCellCol int
}

func NewSudokuGrid(x, y, width, height int) *SudokuGrid {
	grid := &SudokuGrid{
		Group:         fltk.NewGroup(x, y, width, height),
		activeCellRow: -1, // Initialize with invalid values
		activeCellCol: -1,
	}
	grid.Begin()
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			grid.cells[i][j] = NewCell(grid, i, j)
		}
	}
	grid.End()
	return grid
}

func (g *SudokuGrid) SelectCell(i, j int, active bool) {
	// Deselect all cells
	if active {
		for row := 0; row < gridSize; row++ {
			for col := 0; col < gridSize; col++ {
				if row != i || col != j {
					g.cells[row][col].active = false
					if ((g.cells[row][col].i/subGrid)+(g.cells[row][col].j/subGrid))%2 == 1 {
						g.cells[row][col].Widget.SetColor(fltk.ColorFromRgb(230, 230, 230)) // light gray
					} else {
						g.cells[row][col].Widget.SetColor(fltk.BACKGROUND_COLOR)
					}
					g.cells[row][col].Widget.SetBox(fltk.BORDER_BOX)
				}
			}
		}
	}

	// Select the current cell
	g.cells[i][j].active = active
	if active {
		g.cells[i][j].Widget.SetColor(midBlue)
		g.cells[i][j].Widget.SetBox(fltk.THIN_UP_BOX)
		g.activeCellRow = i
		g.activeCellCol = j
	} else {
		if ((g.cells[i][j].i/subGrid)+(g.cells[i][j].j/subGrid))%2 == 1 {
			g.cells[i][j].Widget.SetColor(fltk.ColorFromRgb(230, 230, 230)) // light gray
		} else {
			g.cells[i][j].Widget.SetColor(fltk.BACKGROUND_COLOR)
		}
		g.cells[i][j].Widget.SetBox(fltk.BORDER_BOX)
		g.activeCellRow = -1
		g.activeCellCol = -1
	}
	g.Redraw()
}

func (g *SudokuGrid) GetCell(i, j int) string {
	return g.cells[i][j].Value
}

func (g *SudokuGrid) SetCell(i, j int, text string) {
	// test if text is a number
	val, err := strconv.Atoi(text)
	if err != nil {
		text = ""
	}
	if (val > 0) && (val <= 9) {
		g.cells[i][j].Value = text
		g.cells[i][j].Widget.SetLabel(text)
		g.Redraw()
	}
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
