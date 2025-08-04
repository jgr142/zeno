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
	pages *tview.Pages,
	project *project.ProjectRepo,
) *ProjectsDisplay {
	t := theme.New()
	leftPane := defineLeftPane(pages, project)
	projectDetails := defineProjectDetails(t)

	mainLayout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftPane, 30, 1, true).
		AddItem(projectDetails, 0, 3, false)

	mainLayout.SetBackgroundColor(t.Background)

	pd := &ProjectsDisplay{mainLayout, app, leftPane, 0}
	motions := inputs.NewVim(pd)
	pd.SetInputCapture(motions.VimInputHandler)

	return pd
}

func defineLeftPane(pages *tview.Pages, project *project.ProjectRepo) *tview.Flex {
	projectList := NewProjectList(
		project,
		nil,
		func(idx int, projectName string, projectPath string, shortcut rune) {
			project.Open(projectPath)
			pages.AddAndSwitchToPage("github tools", NewGithubTools(pages, projectPath), true)
		},
	)
	projectSearch := NewProjectSearch(projectList, nil)

	left := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(projectSearch, 1, 0, true).
		AddItem(projectList, 0, 1, false)

	left.SetBorder(true).
		SetTitle(" Projects ")

	return left
}

func defineProjectDetails(t *theme.Theme) *tview.Box {
	return tview.NewTextView().
		SetText("[::b]Select a project to see details").
		SetDynamicColors(true).
		SetWrap(true).
		SetTextAlign(tview.AlignLeft).
		SetBackgroundColor(t.Background)

}

func (pd *ProjectsDisplay) NavigateDown() {
	if pd.left.GetItemCount()-1 > pd.focusedItem {
		pd.focusedItem++
		pd.app.SetFocus(pd.left.GetItem(pd.focusedItem))
	}
}

func (pd *ProjectsDisplay) NavigateUp() {
	if 0 < pd.focusedItem {
		pd.focusedItem--
		pd.app.SetFocus(pd.left.GetItem(pd.focusedItem))
	}
}
