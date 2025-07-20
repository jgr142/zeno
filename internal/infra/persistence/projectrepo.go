package persistence

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jgr142/zeno/internal/domain"
	"github.com/jgr142/zeno/pkg/set"
)

type ProjectRepo struct {
	projects set.Set[domain.Project]
}

func New() *ProjectRepo {
	projects := searchProjects()
	return &ProjectRepo{projects: projects}
}

func searchProjects() set.Set[domain.Project] {
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

	projects := set.NewSet[domain.Project]()
	for _, project := range projectDirs {
		lastSeparator := strings.LastIndexByte(project, os.PathSeparator)
		projName := project[lastSeparator+1:]
		projects.Add(
			domain.Project{
				Name: projName,
				Path: project,
			})
	}

	return *projects
}

func (pr *ProjectRepo) GetAll() set.Set[domain.Project] {
	return pr.projects
}
