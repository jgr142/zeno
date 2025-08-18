package components

import (
	"github.com/jgr142/zeno/internal/ui/theme"
	"github.com/rivo/tview"
)

type ProjectSearch struct {
	*tview.InputField
	onEscape func()
}

func NewProjectSearch(pj *ProjectList, onEscape func()) *ProjectSearch {
	searchInput := tview.NewInputField()
	projectSearch := &ProjectSearch{searchInput, onEscape}

	projectSearch.
		SetPlaceholder("Search projects...").
		// SetFieldWidth(0).
		SetChangedFunc(func(text string) {
			pj.Filter(text)
		}).
		SetLabel("/").
		SetFieldBackgroundColor(theme.Primary)

	return projectSearch
}
