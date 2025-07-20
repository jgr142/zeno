package ui

import "github.com/rivo/tview"

type ProjectSearch struct {
	*tview.InputField
}

func NewProjectSearch(pj *projectList) *ProjectSearch {
	searchInput := tview.NewInputField().
		SetLabel("\\").
		SetPlaceholder("projname").
		SetChangedFunc(func(text string) {
			pj.Filter(text)
		})

	return &ProjectSearch{searchInput}
}
