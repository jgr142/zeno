package components

import (
	"log"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/jgr142/zeno/internal/project"
	"github.com/jgr142/zeno/internal/ui/theme"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/rivo/tview"
)

type ProjectList struct {
	*tview.List
	project  *project.ProjectRepo
	projects []project.Project
	onSearch func()
}

func NewProjectList(project *project.ProjectRepo, onSearch func()) *ProjectList {
	t := theme.New()
	projects := project.GetAll()
	projectChoices := tview.NewList()
	projectList := &ProjectList{projectChoices, project, projects, onSearch}

	projectList.
		ShowSecondaryText(false).
		SetSelectedFocusOnly(false).
		SetSelectedFunc(projectList.selectedFunc).
		SetInputCapture(projectList.vimMotions)

	for _, project := range projects {
		projectChoices.AddItem(project.Name, project.Path, 0, nil)
	}

	projectList.SetBackgroundColor(t.Background)

	return projectList
}

func (pj *ProjectList) Filter(filter string) {
	pj.Clear()

	for _, project := range pj.projects {
		if len(strings.Trim(filter, " ")) == 0 || fuzzy.Match(filter, project.Name) {
			pj.AddItem(project.Name, project.Path, 0, func() {})
		}
	}
}

func (pj *ProjectList) SetOnSearch(fn func()) {
	pj.onSearch = fn
}

func (pj *ProjectList) selectedFunc(idx int, mainText string, secondaryText string, shortcut rune) {
	// Open the project
	for _, proj := range pj.projects {
		if proj.Name == mainText {
			pj.project.Open(proj)
			return
		}
	}

	log.Fatalf("Project Not Found")
}

func (pj *ProjectList) vimMotions(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'j':
		current := pj.GetCurrentItem()
		if current < pj.GetItemCount()-1 {
			pj.SetCurrentItem(current + 1)
		}
		return nil
	case 'k':
		current := pj.GetCurrentItem()
		if current > 0 {
			pj.SetCurrentItem(current - 1)
		}
		return nil
	case 'i', 'a':
		if pj.onSearch != nil {
			pj.onSearch()
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
}
