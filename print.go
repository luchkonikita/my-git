package main

import (
	"github.com/nsf/termbox-go"
)

// Print - print line.
func Print(y int, text string, color termbox.Attribute) {
	x := 0
	for _, char := range text {
		termbox.SetCell(x, y, char, color, termbox.ColorDefault)
		x++
	}
	termbox.Flush()
}
