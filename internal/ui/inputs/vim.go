package inputs

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jgr142/zeno/internal/infra/logger"
	"github.com/rivo/tview"
)

type Mode int

const (
	NormalMode Mode = iota
	InsertMode
)

type Navigatable interface {
	NavigateDown()
	NavigateUp()
	GetCurrent() tview.Primitive
}

// VimController wraps an application with Vim-like motion behavior.
type VimController struct {
	app        *tview.Application
	mode       Mode
	focusStack []tview.Primitive
}

// NewVimController initializes a global vim motion controller.
func NewVimController(app *tview.Application) *VimController {
	vc := &VimController{
		app:  app,
		mode: NormalMode,
	}

	app.SetInputCapture(vc.inputHandler)
	return vc
}

func (vc *VimController) inputHandler(event *tcell.EventKey) *tcell.EventKey {
	if len(vc.focusStack) == 0 {
		focused := vc.app.GetFocus()
		vc.focusStack = append(vc.focusStack, focused)
	}

	curFocus := vc.peekFocus()
	vc.app.SetFocus(curFocus)
	nav, ok := curFocus.(Navigatable)

	switch vc.mode {
	case NormalMode:
		if ok {
			switch event.Rune() {
			case 'j':
				logger.Info("Navigating Down", "curFocus", fmt.Sprintf("%T", vc.peekFocus()))
				nav.NavigateDown()
				return nil
			case 'k':
				logger.Info("Navigating Up", "curFocus", fmt.Sprintf("%T", vc.peekFocus()))
				nav.NavigateUp()
				return nil
			case 'i', 'a':
				logger.Info("Switching to Insert Mode", "curFocus", fmt.Sprintf("%T", vc.peekFocus()))
				vc.mode = InsertMode
				return nil
			}
		}
		switch event.Key() {
		case tcell.KeyEscape:
			logger.Info("Returning to parent", "curFocus", fmt.Sprintf("%T", vc.peekFocus()), "nextFocus", fmt.Sprintf("%T", vc.peekParent()))
			vc.popFocus()
			return nil
		case tcell.KeyEnter:
			if ok {
				logger.Info("Entering Child", "curFocus", fmt.Sprintf("%T", vc.peekFocus()), "nextFocus", fmt.Sprintf("%T", nav.GetCurrent()))
				vc.pushFocus(nav.GetCurrent())
				return nil
			}
		}
	case InsertMode:
		switch event.Key() {
		case tcell.KeyEnter, tcell.KeyEscape:
			logger.Info("Switching to Normal Mode", "curFocus", fmt.Sprintf("%T", vc.peekFocus()))
			vc.mode = NormalMode
			return nil
		default:
			return event
		}
	}
	return event
}

func (vc *VimController) pushFocus(p tview.Primitive) {
	vc.focusStack = append(vc.focusStack, p)
}

func (vc *VimController) popFocus() {
	if len(vc.focusStack) > 0 {
		last := vc.focusStack[len(vc.focusStack)-1]
		vc.focusStack = vc.focusStack[:len(vc.focusStack)-1]
		vc.app.SetFocus(last)
	}
}

func (vc *VimController) peekFocus() tview.Primitive {
	if len(vc.focusStack) > 0 {
		return vc.focusStack[len(vc.focusStack)-1]
	}
	return nil
}

func (vc *VimController) peekParent() tview.Primitive {
	if len(vc.focusStack) > 1 {
		return vc.focusStack[len(vc.focusStack)-2]
	}
	return nil
}
