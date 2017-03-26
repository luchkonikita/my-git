package main

import (
	"github.com/nsf/termbox-go"
)

// NewSelectorLines - create new SelectorLines struct from a list of strings.
func NewSelectorLines(lines []string) SelectorLines {
	sl := SelectorLines{lines: make([]SelectorLine, 0), selectedIndex: 0}
	for i, line := range lines {
		sl.lines = append(sl.lines, SelectorLine{line, i == 0})
	}
	return sl
}

// SelectorLines - a struct containing a list of lines.
type SelectorLines struct {
	lines         []SelectorLine
	selectedIndex int
}

// SelectPrevious - select previous line of SelectorLines.
func (sl *SelectorLines) SelectPrevious() {
	if sl.selectedIndex > 0 {
		sl.selectedIndex--
	}
	for i, line := range sl.lines {
		sl.lines[i] = SelectorLine{line.text, i == sl.selectedIndex}
	}
}

// SelectNext - select next line of SelectorLines.
func (sl *SelectorLines) SelectNext() {
	if sl.selectedIndex < len(sl.lines)-1 {
		sl.selectedIndex++
	}

	for i, line := range sl.lines {
		sl.lines[i] = SelectorLine{line.text, i == sl.selectedIndex}
	}
}

// SelectedText - return text of currently selected option.
func (sl *SelectorLines) SelectedText() string {
	return sl.lines[sl.selectedIndex].text
}

// Print - print lines of given SelectorLines struct.
func (sl *SelectorLines) Print() {
	y := 0
	for _, line := range sl.lines {
		printLine(y, line)
		y++
	}
}

// SelectorLine - a struct containing line text and `selected` flag.
type SelectorLine struct {
	text     string
	selected bool
}

func (sl SelectorLine) render() string {
	if sl.selected {
		return "* " + sl.text
	}
	return "  " + sl.text
}

func printLine(y int, sl SelectorLine) {
	color := termbox.ColorWhite
	if sl.selected {
		color = termbox.ColorGreen
	}
	Print(y, sl.render(), color)
}
