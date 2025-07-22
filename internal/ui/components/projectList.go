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
		SetMainTextColor(t.PrimaryText).
		SetSelectedTextColor(t.Accent).
		SetBorder(false).
		SetBackgroundColor(t.Background).
		SetInputCapture(projectList.vimMotions)

	for _, p := range projects {
		projectChoices.AddItem(p.Name, "", 0, nil)
	}

	return projectList
}

func (pj *ProjectList) Filter(filter string) {
	pj.Clear()

	for _, p := range pj.projects {
		if len(strings.TrimSpace(filter)) == 0 || fuzzy.Match(filter, p.Name) {
			pj.AddItem(p.Name, "", 0, func() {})
		}
	}
}

func (pj *ProjectList) SetOnSearch(fn func()) {
	pj.onSearch = fn
}

func (pj *ProjectList) selectedFunc(idx int, mainText string, secondaryText string, shortcut rune) {
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
		cur := pj.GetCurrentItem()
		if cur < pj.GetItemCount()-1 {
			pj.SetCurrentItem(cur + 1)
		}
		return nil
	case 'k':
		cur := pj.GetCurrentItem()
		if cur > 0 {
			pj.SetCurrentItem(cur - 1)
		}
		return nil
	case 'i', 'a':
		if pj.onSearch != nil {
			pj.onSearch()
			return nil
		}
	}
	return event
}
