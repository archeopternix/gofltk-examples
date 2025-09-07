// Package main implements a Sudoku game interface using the go-fltk GUI library.
// This version features a refactored structure with separate Cell creation.
package main

import (
	"strconv" // For string to integer conversion

	"github.com/archeopternix/go-fltk" // GUI library for creating the Sudoku interface
)

// Constants defining the Sudoku grid properties and visual layout
const (
	gridSize   = 9  // Standard Sudoku grid size (9x9)
	subGrid    = 3  // Size of each 3x3 sub-grid within the main grid
	cellSize   = 40 // Pixel dimensions for each cell/button
	padding    = 10 // Padding around the entire grid within the window
	subPadding = 4  // Extra space between subgrids for visual separation
)

// midBlue defines the highlight color for selected cells using RGB values
var midBlue = fltk.ColorFromRgb(66, 135, 245)

// Cell represents an individual Sudoku cell with its value, UI widget, and state
type Cell struct {
	Value  string       // String representation of the cell's value (1-9 or empty)
	Widget *fltk.Button // FLTK button widget that represents the cell visually
	i      int          // Row index in the 9x9 grid (0-8)
	j      int          // Column index in the 9x9 grid (0-8)
	active bool         // Selection state indicating if this cell is currently selected
}

// NewCell creates and initializes a new Cell instance with proper positioning and styling
// g: Reference to the parent SudokuGrid for callback registration
// i, j: Grid coordinates (row and column indices)
func NewCell(g *SudokuGrid, i, j int) *Cell {
	// Calculate extra padding for subgrid separation
	extraX := (j / subGrid) * subPadding // Horizontal spacing between subgrids
	extraY := (i / subGrid) * subPadding // Vertical spacing between subgrids

	// Calculate absolute screen coordinates for cell placement
	cellX := j*cellSize + extraX + padding // X coordinate with subgrid and window padding
	cellY := i*cellSize + extraY + padding // Y coordinate with subgrid and window padding

	// Initialize Cell struct with default values
	cell := &Cell{
		Value:  "",                                                       // Start with empty value
		Widget: fltk.NewButton(cellX, cellY, cellSize-2, cellSize-2, ""), // Create button with slight size reduction for borders
		i:      i,                                                        // Store row index
		j:      j,                                                        // Store column index
		active: false,                                                    // Initially not selected
	}

	// Configure button appearance and behavior
	cell.Widget.SetAlign(fltk.ALIGN_CENTER) // Center text within the button
	cell.Widget.SetBox(fltk.BORDER_BOX)     // Use bordered box style

	// Apply alternating background colors for better subgrid visibility
	// Checkerboard pattern across subgrids for visual distinction
	if ((i/subGrid)+(j/subGrid))%2 == 1 {
		cell.Widget.SetColor(fltk.ColorFromRgb(230, 230, 230)) // Light gray background
	}

	// Set up click handler using copies of indices to avoid closure issues
	i_copy := i // Local copy for closure safety
	j_copy := j // Local copy for closure safety
	cell.Widget.SetCallback(func() {
		g.SelectCell(i_copy, j_copy, true) // Delegate selection logic to parent grid
		g.Redraw()                         // Request visual update
	})

	return cell // Return the fully configured cell
}

// SudokuGrid represents the complete 9x9 Sudoku board with cell management
type SudokuGrid struct {
	*fltk.Group                             // Embedded FLTK group for widget container functionality
	cells         [gridSize][gridSize]*Cell // 2D array of cell pointers
	activeCellRow int                       // Currently selected cell's row index (-1 if none)
	activeCellCol int                       // Currently selected cell's column index (-1 if none)
}

// NewSudokuGrid creates and initializes a new Sudoku grid widget
// x, y: Top-left coordinates for grid placement
// width, height: Dimensions of the grid area
func NewSudokuGrid(x, y, width, height int) *SudokuGrid {
	// Initialize grid with default values
	grid := &SudokuGrid{
		Group:         fltk.NewGroup(x, y, width, height), // Create FLTK group container
		activeCellRow: -1,                                 // Initialize with no active cell
		activeCellCol: -1,                                 // Initialize with no active cell
	}
	grid.Begin() // Start widget addition to group
	// Create all 81 cells in row-major order
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			grid.cells[i][j] = NewCell(grid, i, j) // Create and store each cell
		}
	}
	grid.End()  // Finish widget addition to group
	return grid // Return the fully populated grid
}

// SelectCell manages cell selection state with visual feedback
// i, j: Coordinates of cell to select/deselect
// active: Boolean indicating selection (true) or deselection (false)
func (g *SudokuGrid) SelectCell(i, j int, active bool) {
	// Deselect all other cells when selecting a new one (single selection mode)
	if active {
		for row := 0; row < gridSize; row++ {
			for col := 0; col < gridSize; col++ {
				if row != i || col != j { // Skip the target cell
					g.cells[row][col].active = false // Mark as inactive
					// Restore original background color based on subgrid position
					if ((g.cells[row][col].i/subGrid)+(g.cells[row][col].j/subGrid))%2 == 1 {
						g.cells[row][col].Widget.SetColor(fltk.ColorFromRgb(230, 230, 230)) // Light gray
					} else {
						g.cells[row][col].Widget.SetColor(fltk.BACKGROUND_COLOR) // Default background
					}
					g.cells[row][col].Widget.SetBox(fltk.BORDER_BOX) // Restore border style
				}
			}
		}
	}

	// Update selection state of target cell
	g.cells[i][j].active = active
	if active {
		// Visual indicators for selected cell
		g.cells[i][j].Widget.SetColor(midBlue)        // Apply highlight color
		g.cells[i][j].Widget.SetBox(fltk.THIN_UP_BOX) // Use raised border style
		g.activeCellRow = i                           // Store active row
		g.activeCellCol = j                           // Store active column
	} else {
		// Restore normal appearance when deselecting
		if ((g.cells[i][j].i/subGrid)+(g.cells[i][j].j/subGrid))%2 == 1 {
			g.cells[i][j].Widget.SetColor(fltk.ColorFromRgb(230, 230, 230)) // Light gray
		} else {
			g.cells[i][j].Widget.SetColor(fltk.BACKGROUND_COLOR) // Default background
		}
		g.cells[i][j].Widget.SetBox(fltk.BORDER_BOX) // Restore border style
		g.activeCellRow = -1                         // Clear active row
		g.activeCellCol = -1                         // Clear active column
	}
	g.Redraw() // Refresh display to show changes
}

// GetCell retrieves the string value of a specific cell
// i, j: Coordinates of the cell to query
// Returns: String representation of the cell's value (empty if unset)
func (g *SudokuGrid) GetCell(i, j int) string {
	return g.cells[i][j].Value // Direct access to stored value
}

// SetCell updates a cell's value with validation
// i, j: Coordinates of the cell to update
// text: Proposed new value (must be digit 1-9 to be accepted)
func (g *SudokuGrid) SetCell(i, j int, text string) {
	// Validate input: must be a number between 1 and 9
	val, err := strconv.Atoi(text) // Attempt conversion to integer
	if err != nil {
		text = "" // Clear invalid input
	}
	// Only accept values in the valid Sudoku range (1-9)
	if (val > 0) && (val <= 9) {
		g.cells[i][j].Value = text          // Store the string value
		g.cells[i][j].Widget.SetLabel(text) // Update button label
		g.Redraw()                          // Refresh display
	}
	// Note: Invalid values are silently ignored (no update occurs)
}

// main function: Application entry point that sets up and runs the GUI
func main() {
	// Calculate window dimensions based on grid properties and padding
	winWidth := gridSize*cellSize + 2*padding + (subGrid-1)*subPadding
	winHeight := gridSize*cellSize + 2*padding + (subGrid-1)*subPadding

	// Create main application window
	win := fltk.NewWindow(winWidth, winHeight)
	win.SetLabel("Sudoku Board - go-fltk (Read Only, Selectable)") // Window title

	// Create Sudoku grid centered within window padding
	sudokuGrid := NewSudokuGrid(padding, padding, winWidth-2*padding, winHeight-2*padding)

	// Set some initial values for demonstration purposes
	sudokuGrid.SetCell(0, 0, "5") // Top-left cell
	sudokuGrid.SetCell(1, 2, "3") // Row 1, Column 2
	sudokuGrid.SetCell(8, 8, "9") // Bottom-right cell

	// Complete window setup and start event loop
	win.End()  // Finalize window content
	win.Show() // Make window visible
	fltk.Run() // Start FLTK event processing (blocks until window closed)
}
