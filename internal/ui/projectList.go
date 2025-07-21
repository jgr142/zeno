package ui

import (
	"log"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/jgr142/zeno/internal/project"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/rivo/tview"
)

type projectList struct {
	*tview.List
	project  ProjectRepo
	projects []project.Project
	onSearch func()
}

func NewProjectList(project ProjectRepo, onSearch func()) *projectList {
	projects := project.GetAll()
	projectChoices := tview.NewList()
	projectList := &projectList{projectChoices, project, projects, onSearch} // I changed from *projectChoices to projectChoices and that fixed my j and k logic, but idk why

	projectList.
		ShowSecondaryText(false).
		SetSelectedFocusOnly(false).
		SetSelectedFunc(projectList.selectedFunc).
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
				// TODO: add support for vim like commands
				// case ' ', 'l':
				// 	project.Open(projectList.GetCurrentItem())
			}

			// switch event.Key() {
			// case tcell.KeyEnter, tcell.KeyRight:
			// }

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

func (pj *projectList) selectedFunc(idx int, mainText string, secondaryText string, shortcut rune) {
	for _, proj := range pj.projects {
		if proj.Name == mainText {
			pj.project.Open(proj)
			return
		}
	}
	log.Fatalf("Project Not Found")
}
