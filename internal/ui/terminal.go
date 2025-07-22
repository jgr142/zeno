package ui

import (
	"log"

	"github.com/jgr142/zeno/internal/project"
	"github.com/jgr142/zeno/internal/ui/components"
	"github.com/rivo/tview"
)

func Init(project *project.ProjectRepo) {
	app := tview.NewApplication()
	projectsDisplay := components.NewProjectsDisplay(app, project)
	if err := app.SetRoot(projectsDisplay, true).Run(); err != nil {
		log.Fatal(err.Error())
	}
}
