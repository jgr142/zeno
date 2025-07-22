package components

import (
	"github.com/jgr142/zeno/internal/project"
	"github.com/jgr142/zeno/internal/ui/theme"
	"github.com/rivo/tview"
)

type ProjectsDisplay struct {
	*tview.Flex
}

func NewProjectsDisplay(app *tview.Application, project *project.ProjectRepo) *ProjectsDisplay {
	t := theme.New()
	// TODO: Add Frames
	leftPane := defineLeftPane(app, project)
	projectDetails := defineProjectDetails()

	mainLayout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftPane, 0, 1, true).
		AddItem(projectDetails, 0, 2, false)

	mainLayout.SetBackgroundColor(t.Primary)

	return &ProjectsDisplay{mainLayout}
}

func defineLeftPane(app *tview.Application, project *project.ProjectRepo) *tview.Flex {
	projectList := NewProjectList(project, nil)
	projectSearch := NewProjectSearch(projectList, nil)

	projectList.SetOnSearch(func() {
		app.SetFocus(projectSearch)
	})

	projectSearch.SetOnEscape(func() {
		app.SetFocus(projectList)
	})

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(projectSearch, 1, 0, true).
		AddItem(projectList, 0, 1, false)

}

func defineProjectDetails() *tview.TextView {
	return tview.NewTextView().
		SetText("Select a project to see details").
		SetDynamicColors(true)
}
