package inputs

import "github.com/gdamore/tcell/v2"

// Navigatable defines components that can be navigated up and down.
type Navigatable interface {
	NavigateDown()
	NavigateUp()
}

// Insertable defines components that can enter an "insert mode".
type Insertable interface {
	OnInsertMode()
}

// VimInputHandler creates a tview input capture function that handles basic
// Vim-like keybindings for components that implement Navigatable and optionally
// Insertable.
func VimInputHandler(v Navigatable) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'j':
			v.NavigateDown()
			return nil
		case 'k':
			v.NavigateUp()
			return nil
		case 'i', 'a':
			if insertable, ok := v.(Insertable); ok {
				insertable.OnInsertMode()
				return nil
			}
		}
		return event
	}
}
