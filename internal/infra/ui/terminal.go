package ui

import (
	"log"

	"github.com/jgr142/zeno/internal/domain"
	"github.com/rivo/tview"
)

func Init(projectsInDir []domain.Project) {
	app := tview.NewApplication()
	projectsDisplay := NewProjectsDisplay(projectsInDir)
	if err := app.SetRoot(projectsDisplay, true).Run(); err != nil {
		log.Fatal(err.Error())
	}
}
