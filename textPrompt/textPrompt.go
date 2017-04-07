package textPrompt

import (
	"github.com/nsf/termbox-go"
	"my-git/errors"
)

const (
	space   = 0x0020
	propmtY = 0
	editorY = 1
)

// Init - Show select prompt.
func Init(prompt string) string {
	err := termbox.Init()
	errors.CheckError(err)
	termbox.SetOutputMode(termbox.OutputNormal)
	defer termbox.Close()

	printLine(propmtY, stringToRunes(prompt), termbox.ColorGreen)

	var input []rune

	printLineWithCursor(editorY, input, termbox.ColorWhite)

	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			switch ev.Key {
			case termbox.KeyEsc:
				return ""
			case termbox.KeyEnter:
				return string(input)
			case termbox.KeyBackspace:
			case termbox.KeyBackspace2:
				if len(input) > 0 {
					input = input[:len(input)-1]
				}
			default:
				// Process only basic latin characters.
				if ev.Ch >= 32 && ev.Ch <= 126 {
					input = append(input, ev.Ch)
				}
			}

			printLineWithCursor(editorY, input, termbox.ColorWhite)
		}
	}
}

func stringToRunes(s string) []rune {
	return []rune(s)
}

// TODO: Move to a shared place...
func printLine(y int, text []rune, color termbox.Attribute) {
	width, _ := termbox.Size()

	for x := 0; x < width; x++ {
		if x < len(text) {
			termbox.SetCell(x, y, text[x], color, termbox.ColorDefault)
		} else {
			termbox.SetCell(x, y, space, color, termbox.ColorDefault)
		}
	}

	termbox.Flush()
}

func printLineWithCursor(y int, text []rune, color termbox.Attribute) {
	printLine(y, text, color)
	termbox.SetCursor(len(text), y)
	termbox.Flush()
}
