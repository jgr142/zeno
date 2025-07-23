package project

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ProjectRepo struct{}

func New() *ProjectRepo {
	return &ProjectRepo{}
}

func (pr *ProjectRepo) GetAll() []Project {
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

	projects := make([]Project, 0)
	for _, projectPath := range projectDirs {
		lastSeparator := strings.LastIndexByte(projectPath, os.PathSeparator)
		projectName := projectPath[lastSeparator+1:]
		projects = append(projects, Project{projectName, projectPath})
	}

	return projects
}

func (pr *ProjectRepo) Open(projectPath string) {
	// Open the project
	// TODO: should add check that path is real
	cmd := exec.Command("code", projectPath)
	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}
