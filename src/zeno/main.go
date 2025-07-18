package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

// type Project struct {
// 	name     string
// 	metaData MetaData
// }

// type MetaData struct {
// 	createdBy      string
// 	createdAt      time.Time
// 	lastModifiedBy string
// 	lastModifiedAt time.Time
// }

func main() {
	projectDirs := searchProjects()
	for i, proj := range projectDirs {
		fmt.Printf("%d: %s\n", i, proj)
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
