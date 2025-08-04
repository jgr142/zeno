package inputs

import "github.com/gdamore/tcell/v2"

type mode int

const (
	normalMode mode = iota
	insertMode
)

type Navigatable interface {
	NavigateDown()
	NavigateUp()
}

type vimMotions struct {
	Navigatable
	m mode
}

func NewVim(v Navigatable) *vimMotions {
	return &vimMotions{v, normalMode}
}

func (vm *vimMotions) VimInputHandler(event *tcell.EventKey) *tcell.EventKey {

	switch vm.m {
	case normalMode:
		switch event.Rune() {
		case 'j':
			vm.NavigateDown()
			return nil
		case 'k':
			vm.NavigateUp()
			return nil
		case 'i', 'a':
			vm.m = insertMode
			return nil
		default:
			return nil
		}
	case insertMode:
		switch event.Key() {
		case tcell.KeyEnter:
			vm.m = normalMode
		case tcell.KeyEscape:
			vm.m = normalMode
		default:
			return event
		}
	}
	return nil
}
