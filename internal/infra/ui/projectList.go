package ui

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/jgr142/zeno/internal/domain"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/rivo/tview"
)

type projectList struct {
	*tview.List
	projects []domain.Project
	onSearch func()
}

func NewProjectList(projects []domain.Project, onSearch func()) *projectList {
	projectChoices := tview.NewList()
	projectList := &projectList{projectChoices, projects, onSearch} // I changed from *projectChoices to projectChoices and that fixed my j and k logic, but idk why

	projectList.
		ShowSecondaryText(false).
		SetSelectedFocusOnly(false).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Rune() {
			case 'j':
				current := projectChoices.GetCurrentItem()
				if current < projectChoices.GetItemCount()-1 {
					projectChoices.SetCurrentItem(current + 1)
				}
				return nil
			case 'k':
				current := projectChoices.GetCurrentItem()
				if current > 0 {
					projectChoices.SetCurrentItem(current - 1)
				}
				return nil
			case 'i', 'a':
				if projectList.onSearch != nil {
					projectList.onSearch()
					return nil
				}
			}

			return event
		})

	ifSelected := func() {

	}

	for _, project := range projects {
		projectChoices.AddItem(project.Name, project.Path, 0, ifSelected)
	}

	return projectList
}

func (pj *projectList) Filter(filter string) {
	pj.Clear()

	for _, project := range pj.projects {
		if len(strings.Trim(filter, " ")) == 0 || fuzzy.Match(filter, project.Name) {
			pj.AddItem(project.Name, project.Path, 0, func() {})
		}
	}
}

func (pj *projectList) SetOnSearch(fn func()) {
	pj.onSearch = fn
}
