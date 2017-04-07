package selectPrompt

import (
	"github.com/nsf/termbox-go"
	"my-git/errors"
)

// Option - a struct representing select option.
// You need to specify text of the label and it's value.
type Option struct {
	Text     string
	Value    string
	Selected bool
	Custom   bool
}

func (o Option) render() string {
	if o.Selected {
		return "* " + o.Text
	}
	return "  " + o.Text
}

// Init - Show select prompt.
func Init(options []Option) Option {
	err := termbox.Init()
	errors.CheckError(err)

	termbox.SetOutputMode(termbox.OutputNormal)
	defer termbox.Close()

	sls := newRenderer(options)
	sls.print()

	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			switch ev.Key {
			case termbox.KeyEsc:
				return Option{}
			case termbox.KeyArrowUp:
				sls.selectPrevious()
				sls.print()
			case termbox.KeyArrowDown:
				sls.selectNext()
				sls.print()
			case termbox.KeyEnter:
				return sls.getSelected()
			}
		}
	}
}

type renderer struct {
	Options       []Option
	SelectedIndex int
}

func newRenderer(options []Option) renderer {
	sl := renderer{Options: options, SelectedIndex: -1}
	for i, option := range options {
		if option.Selected {
			sl.SelectedIndex = i
		}
	}
	return sl
}

func (r *renderer) selectPrevious() {
	if r.SelectedIndex > 0 {
		r.SelectedIndex--
	}
	for i := range r.Options {
		r.Options[i].Selected = (i == r.SelectedIndex)
	}
}

func (r *renderer) selectNext() {
	if r.SelectedIndex < len(r.Options)-1 {
		r.SelectedIndex++
	}

	for i := range r.Options {
		r.Options[i].Selected = (i == r.SelectedIndex)
	}
}

func (r *renderer) getSelected() Option {
	return r.Options[r.SelectedIndex]
}

func (r *renderer) print() {
	y := 0
	for _, option := range r.Options {
		printOption(y, option)
		y++
	}
}

func printOption(y int, o Option) {
	var color termbox.Attribute
	if o.Selected {
		color = termbox.ColorGreen
	} else if o.Custom {
		color = termbox.ColorMagenta
	} else {
		color = termbox.ColorWhite
	}
	print(y, o.render(), color)
}

func print(y int, text string, color termbox.Attribute) {
	x := 0
	for _, char := range text {
		termbox.SetCell(x, y, char, color, termbox.ColorDefault)
		x++
	}
	termbox.Flush()
}
