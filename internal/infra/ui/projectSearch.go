package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ProjectSearch struct {
	*tview.InputField
	onEscape func()
}

func NewProjectSearch(pj *projectList, onEscape func()) *ProjectSearch {
	searchInput := tview.NewInputField()
	projectSearch := &ProjectSearch{searchInput, onEscape}
	projectSearch.
		SetLabel("\\").
		SetPlaceholder("projname").
		SetChangedFunc(func(text string) {
			pj.Filter(text)
		}).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyEnter, tcell.KeyEsc:
				if projectSearch.onEscape != nil {
					projectSearch.onEscape()
				}
				return nil
			}
			return event
		})

	return projectSearch
}

func (ps *ProjectSearch) SetOnEscape(fn func()) {
	ps.onEscape = fn
}
