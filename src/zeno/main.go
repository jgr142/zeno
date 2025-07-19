package main

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	projectsInDir := searchProjects()
	projectChoices := tview.NewList()
	ifSelected := func() {

	}
	for _, project := range projectsInDir {
		projectChoices.AddItem(project, "", 0, ifSelected)
	}

	projectChoices.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'j':
			current := projectChoices.GetCurrentItem()
			if current < projectChoices.GetItemCount()-1 {
				projectChoices.SetCurrentItem(current + 1)
			}
			return nil
		case 'k':
			current := projectChoices.GetCurrentItem()
			if current > 0 {
				projectChoices.SetCurrentItem(current - 1)
			}
			return nil
		}
		return event
	})

	if err := app.SetRoot(projectChoices, true).Run(); err != nil {
		log.Fatal(err.Error())
	}

}

func searchProjects() []string {
	projectsRoot := "/Users/joshuagisiger/projects"
	projectDirs := make([]string, 0)
	err := filepath.WalkDir(
		projectsRoot,
		func(path string, d fs.DirEntry, err error) error {
			// If we find a .git dir add it to proj dirs
			if d.Name() == ".git" && d.IsDir() {
				projectDirs = append(projectDirs, filepath.Dir(path))
				return filepath.SkipDir
			}

			// Prune Hidden dirs
			if strings.HasPrefix(d.Name(), ".") && d.IsDir() {
				return filepath.SkipDir
			}

			return nil
		},
	)

	if err != nil {
		log.Fatal("Project Search Failed")
	}

	return projectDirs
}
