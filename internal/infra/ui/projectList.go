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
}

func NewProjectList(projects []domain.Project) *projectList {
	projectChoices := tview.NewList().
		ShowSecondaryText(false).
		SetSelectedFocusOnly(false)

	ifSelected := func() {

	}

	for _, project := range projects {
		projectChoices.AddItem(project.Name, project.Path, 0, ifSelected)
	}

	projectChoices.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
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
		}
		return event
	})
	return &projectList{projectChoices, projects} // I changed from *projectChoices to projectChoices and that fixed my j and k logic, but idk why
}

func (pj *projectList) Filter(filter string) {
	pj.Clear()

	for _, project := range pj.projects {
		if len(strings.Trim(filter, " ")) == 0 || fuzzy.Match(filter, project.Name) {
			pj.AddItem(project.Name, project.Path, 0, func() {})
		}
	}
}
