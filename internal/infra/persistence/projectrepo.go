package persistence

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jgr142/zeno/internal/domain"
)

type ProjectRepo struct {
	projects []domain.Project
}

func New() *ProjectRepo {
	projects := searchProjects()
	return &ProjectRepo{projects: projects}
}

func searchProjects() []domain.Project {
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

	projects := make([]domain.Project, 0)
	for _, project := range projectDirs {
		lastSeparator := strings.LastIndexByte(project, os.PathSeparator)
		projName := project[lastSeparator+1:]
		projects = append(
			projects,
			domain.Project{
				Name: projName,
				Path: project,
			})
	}

	return projects
}

func (pr *ProjectRepo) Get() []domain.Project {
	return pr.projects
}
