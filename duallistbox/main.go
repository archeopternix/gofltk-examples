package main

import (
	"fmt"
	"strings"

	"github.com/archeopternix/go-fltk"
	"github.com/archeopternix/gofltk-duallistbox"
)

func main() {
	win := fltk.NewWindow(600, 400)
	win.SetLabel("DualListBox Example")

	dual := duallistbox.NewDualListBox(20, 20, 560, 320)

	// Set initial items
	dual.SetLeftItems([]string{"foo", "bar"})
	dual.SetRightItems([]string{"alpha", "beta", "gamma", "delta"})

	// Set custom titles
	dual.SetLeftTitle("Selected Items")
	dual.SetRightTitle("Available Items")

	// Register event handlers
	dual.RegisterMoveLeftHandler(func() {
		fmt.Printf("Moved to left: %s\n", strings.Join(dual.GetLeftItems(), ", "))
	})
	dual.RegisterMoveRightHandler(func() {
		fmt.Printf("Moved to right: %s\n", strings.Join(dual.GetRightItems(), ", "))
	})

	win.End()
	win.Show()
	fltk.Run()
}
