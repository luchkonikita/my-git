package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	termbox.SetOutputMode(termbox.OutputNormal)

	defer termbox.Close()

	branches := GetBranches()
	options := NewSelectorLines(branches)

	options.Print()

loop:
	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyArrowUp:
				options.SelectPrevious()
				options.Print()
			case termbox.KeyArrowDown:
				options.SelectNext()
				options.Print()
			case termbox.KeyEnter:
				branch := options.SelectedText()
				SetBranch(branch)
			}
		}
	}
}
