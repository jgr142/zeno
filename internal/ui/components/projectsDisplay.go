package components

import (
	"github.com/jgr142/zeno/internal/project"
	"github.com/jgr142/zeno/internal/ui/inputs"
	"github.com/jgr142/zeno/internal/ui/theme"
	"github.com/rivo/tview"
)

type ProjectsDisplay struct {
	*tview.Flex
	app         *tview.Application
	left        *tview.Flex
	focusedItem int
}

func NewProjectsDisplay(
	app *tview.Application,
	pages *inputs.VimDecorator,
	project *project.ProjectRepo,
) *ProjectsDisplay {
	leftPane := defineLeftPane(pages, project)
	projectDetails := NewProjectDetails()

	mainLayout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftPane, 0, 1, false).
		AddItem(projectDetails, 0, 2, false)

	mainLayout.SetBackgroundColor(theme.Background)

	pd := &ProjectsDisplay{mainLayout, app, leftPane, 0}

	return pd
}

func defineLeftPane(pages *inputs.VimDecorator, project *project.ProjectRepo) *tview.Flex {
	projectList := NewProjectList(
		project,
		nil,
		func(idx int, projectName string, projectPath string, shortcut rune) {
			project.Open(projectPath)
			pages.AddAndSwitchToPage("github tools", NewGithubTools(projectPath), true)
		},
	)
	projectSearch := NewProjectSearch(projectList, nil)

	left := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(projectSearch, 1, 0, true).
		AddItem(projectList, 0, 1, false)

	left.SetBorder(true).
		SetTitle(" [::b]Projects ")

	return left
}

func (pd *ProjectsDisplay) NavigateDown() {
	if pd.focusedItem < pd.left.GetItemCount()-1 {
		pd.focusedItem++
	}
}

func (pd *ProjectsDisplay) NavigateUp() {
	if 0 < pd.focusedItem {
		pd.focusedItem--
	}
}

func (pd *ProjectsDisplay) GetCurrent() tview.Primitive {
	return pd.left.GetItem(pd.focusedItem)
}
