package ui

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/jgr142/zeno/internal/project"
	"github.com/jgr142/zeno/internal/ui/components"
	"github.com/jgr142/zeno/internal/ui/inputs"
	"github.com/rivo/tview"
)

type InputReceiver interface {
	inputs.Navigatable
	SetInputCapture(func(*tcell.EventKey) *tcell.EventKey) *tview.Box
}

func Init(project *project.ProjectRepo) {
	app := tview.NewApplication()
	dec := inputs.NewVimDecorator(app)

	projectsDisplay := components.NewProjectsDisplay(app, dec, project)

	dec.AddPage("projects search", projectsDisplay, true, true)

	if err := app.SetRoot(dec.Layout(), true).Run(); err != nil {
		log.Fatal(err.Error())
	}
}
