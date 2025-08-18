package components

import (
	"strings"

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
	onSelect func(int, string, string, rune)
}

func NewProjectList(
	project *project.ProjectRepo,
	onSearch func(),
	onSelect func(int, string, string, rune),
) *ProjectList {
	t := theme.New()
	projects := project.GetAll()
	projectChoices := tview.NewList()

	projectList := &ProjectList{
		projectChoices,
		project,
		projects,
		onSearch,
		onSelect,
	}

	projectList.
		ShowSecondaryText(false).
		SetSelectedFocusOnly(false).
		SetSelectedFunc(onSelect).
		SetMainTextColor(t.PrimaryText).
		SetSelectedTextColor(t.Accent).
		SetBorder(false).
		SetBackgroundColor(t.Background)

	for _, p := range projects {
		projectChoices.AddItem(p.Name, p.Path, 0, nil)
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

func (pj *ProjectList) NavigateUp() {
	if 0 < pj.GetCurrentItem() {
		pj.SetCurrentItem(pj.GetCurrentItem() - 1)
	}
}

func (pj *ProjectList) NavigateDown() {
	if pj.GetItemCount()-1 > pj.GetCurrentItem() {
		pj.SetCurrentItem(pj.GetCurrentItem() + 1)
	}
}

func (pj *ProjectList) GetCurrent() tview.Primitive {
	return nil
}
