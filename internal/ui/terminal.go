package ui

import (
	"log"

	"github.com/jgr142/zeno/internal/project"
	"github.com/rivo/tview"
)

type ProjectRepo interface {
	GetAll() []project.Project
	Open(project project.Project)
}

func Init(project ProjectRepo) {
	app := tview.NewApplication()
	projectsDisplay := NewProjectsDisplay(app, project)
	if err := app.SetRoot(projectsDisplay, true).Run(); err != nil {
		log.Fatal(err.Error())
	}
}
