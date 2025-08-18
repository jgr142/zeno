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

// VimDecorator wraps an application with Vim-like motion behavior.
type VimDecorator struct {
	*tview.Pages
	mode       Mode
	statusBar  *tview.TextView
	app        *tview.Application
	focusStack []tview.Primitive
}

// NewVimDecorator initializes a global vim motion controller.
func NewVimDecorator(app *tview.Application) *VimDecorator {
	status := tview.NewTextView().
		SetDynamicColors(true)

	status.SetBorder(true).SetTitle("Mode")
	status.SetText("[blue]Normal")

	vd := &VimDecorator{
		Pages:     tview.NewPages(),
		mode:      NormalMode,
		statusBar: status,
		app:       app,
	}

	app.SetInputCapture(vd.inputHandler)
	return vd
}

func (vd *VimDecorator) Layout() tview.Primitive {
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(vd.Pages, 0, 20, true).
		AddItem(vd.statusBar, 0, 1, false)

	return flex
}

func (vd *VimDecorator) inputHandler(event *tcell.EventKey) *tcell.EventKey {
	if len(vd.focusStack) == 0 {
		focused := vd.app.GetFocus()
		vd.focusStack = append(vd.focusStack, focused)
	}

	curFocus := vd.peekFocus()
	vd.app.SetFocus(curFocus)
	nav, ok := curFocus.(Navigatable)

	switch vd.mode {
	case NormalMode:
		switch event.Rune() {
		case 'j':
			if ok {
				logger.Info("Navigating Down", "curFocus", fmt.Sprintf("%T", vd.peekFocus()))
				nav.NavigateDown()
				return nil
			}
		case 'k':
			if ok {
				logger.Info("Navigating Up", "curFocus", fmt.Sprintf("%T", vd.peekFocus()))
				nav.NavigateUp()
				return nil
			}
		case 'i', 'a':
			logger.Info("Switching to Insert Mode", "curFocus", fmt.Sprintf("%T", vd.peekFocus()))
			vd.statusBar.SetText("[green]Insert")
			vd.mode = InsertMode
			return nil
		}

		switch event.Key() {
		case tcell.KeyEscape:
			logger.Info("Returning to parent", "curFocus", fmt.Sprintf("%T", vd.peekFocus()), "nextFocus", fmt.Sprintf("%T", vd.peekParent()))
			vd.popFocus()
			return nil
		case tcell.KeyEnter:
			if ok {
				child := nav.GetCurrent()
				logger.Info("Entering Child", "curFocus", fmt.Sprintf("%T", vd.peekFocus()), "nextFocus", fmt.Sprintf("%T", child))

				if child == nil {
					return event
				}

				vd.pushFocus(nav.GetCurrent())
				return nil
			}
		}
		return nil
	case InsertMode:
		switch event.Key() {
		case tcell.KeyEnter, tcell.KeyEscape:
			logger.Info("Switching to Normal Mode", "curFocus", fmt.Sprintf("%T", vd.peekFocus()))
			vd.statusBar.SetText("[blue]Normal")
			vd.mode = NormalMode
			return nil
		default:
			return event
		}
	}
	return event
}

func (vd *VimDecorator) pushFocus(p tview.Primitive) {
	vd.focusStack = append(vd.focusStack, p)
}

func (vd *VimDecorator) popFocus() {
	if len(vd.focusStack) > 0 {
		last := vd.focusStack[len(vd.focusStack)-1]
		vd.focusStack = vd.focusStack[:len(vd.focusStack)-1]
		vd.app.SetFocus(last)
	}
}

func (vd *VimDecorator) peekFocus() tview.Primitive {
	if len(vd.focusStack) > 0 {
		return vd.focusStack[len(vd.focusStack)-1]
	}
	return nil
}

func (vd *VimDecorator) peekParent() tview.Primitive {
	if len(vd.focusStack) > 1 {
		return vd.focusStack[len(vd.focusStack)-2]
	}
	return nil
}
