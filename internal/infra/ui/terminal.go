package ui

import (
	"log"

	"github.com/jgr142/zeno/internal/domain"
	"github.com/rivo/tview"
)

func Init(projectsInDir []domain.Project) {
	app := tview.NewApplication()
	projectList := NewProjectList(projectsInDir)
	if err := app.SetRoot(projectList, true).Run(); err != nil {
		log.Fatal(err.Error())
	}
}
