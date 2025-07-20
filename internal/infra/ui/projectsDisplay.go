package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jgr142/zeno/internal/domain"
	"github.com/rivo/tview"
)

type ProjectsDisplay struct {
	*tview.Flex
}

func NewProjectsDisplay(app *tview.Application, projects []domain.Project) *ProjectsDisplay {
	// TODO: Add Frames
	projectList := NewProjectList(projects)
	projectSearch := NewProjectSearch(projectList)
	leftPane := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(projectSearch, 1, 0, true).
		AddItem(projectList, 0, 1, false)

	leftPane.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter, tcell.KeyEsc:
			app.SetFocus(projectList)
			return nil
		}

		switch event.Rune() {
		case 'a', 'i':
			app.SetFocus(projectSearch)
			return nil
		}
		return event
	})
	projectDetails := tview.NewTextView().
		SetText("Select a project to see details").
		SetDynamicColors(true)

	mainLayout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftPane, 0, 1, true).
		AddItem(projectDetails, 0, 2, false)

	return &ProjectsDisplay{mainLayout}
}
