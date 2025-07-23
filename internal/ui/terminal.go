package ui

import (
	"log"

	"github.com/jgr142/zeno/internal/project"
	"github.com/jgr142/zeno/internal/ui/components"
	"github.com/rivo/tview"
)

func Init(project *project.ProjectRepo) {
	app := tview.NewApplication()
	pages := tview.NewPages()

	projectsDisplay := components.NewProjectsDisplay(app, pages, project)

	pages.AddPage("projects search", projectsDisplay, true, true)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		log.Fatal(err.Error())
	}
}
