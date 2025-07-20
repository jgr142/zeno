package ui

import (
	"github.com/jgr142/zeno/internal/domain"
	"github.com/rivo/tview"
)

type ProjectsDisplay struct {
	*tview.Flex
}

func NewProjectsDisplay(projects []domain.Project) *ProjectsDisplay {
	// TODO: Add Frames
	projectList := NewProjectList(projects)
	projectSearch := NewProjectSearch(projectList)
	leftPane := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(projectSearch, 1, 0, true).
		AddItem(projectList, 0, 1, false)

	projectDetails := tview.NewTextView().
		SetText("Select a project to see details").
		SetDynamicColors(true)

	mainLayout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftPane, 0, 1, true).
		AddItem(projectDetails, 0, 2, false)

	return &ProjectsDisplay{mainLayout}
}
