package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jgr142/zeno/internal/ui/theme"
	"github.com/rivo/tview"
)

type ProjectSearch struct {
	*tview.InputField
	onEscape func()
}

func NewProjectSearch(pj *ProjectList, onEscape func()) *ProjectSearch {
	t := theme.New()
	searchInput := tview.NewInputField()
	projectSearch := &ProjectSearch{searchInput, onEscape}

	projectSearch.
		SetPlaceholder("Search projects...").
		// SetFieldWidth(0).
		SetChangedFunc(func(text string) {
			pj.Filter(text)
		}).
		SetLabel("/").
		SetFieldBackgroundColor(t.Primary).
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
	// projectSearch.SetBorder(true)

	return projectSearch
}

func (ps *ProjectSearch) SetOnEscape(fn func()) {
	ps.onEscape = fn
}
