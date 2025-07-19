package ui

import (
	"log"

	"github.com/rivo/tview"
)

func Init(projectsInDir []string) {
	app := tview.NewApplication()
	projectList := NewProjectList(&projectsInDir)
	if err := app.SetRoot(projectList, true).Run(); err != nil {
		log.Fatal(err.Error())
	}
}
