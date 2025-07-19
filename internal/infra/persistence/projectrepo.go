package persistence

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/jgr142/zeno/internal/domain"
)

type ProjectRepo struct {
	projects     map[domain.Project]struct{}
	projectPaths map[string]struct{}
}

func New() *ProjectRepo {
	projects, projectPaths := searchProjects()
	return &ProjectRepo{
		projects:     projects,
		projectPaths: projectPaths,
	}
}

func (pr *ProjectRepo) Get() map[domain.Project]struct{} {
	return pr.projects
}

func searchProjects() (map[domain.Project]struct{}, map[string]struct{}) {
	projectsRoot := "/Users/joshuagisiger/projects"
	projectPaths := make(map[string]struct{})
	err := filepath.WalkDir(
		projectsRoot,
		func(path string, d fs.DirEntry, err error) error {
			// If we find a .git dir add it to proj dirs
			if d.Name() == ".git" && d.IsDir() {
				projectPaths[filepath.Dir(path)] = struct{}{}
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

	projects := make(map[domain.Project]struct{})
	for projectPath, _ := range projectPaths {
		lastSeparator := strings.LastIndexByte(projectPath, os.PathSeparator)
		projectName := projectPath[lastSeparator+1:]
		projects[domain.Project{
			Name: projectName,
			Path: projectPath,
		}] = struct{}{}
	}

	return projects, projectPaths
}

func createWatchers(projectPaths map[string]struct{}) {
	projectsRoot := "/Users/joshuagisiger/projects"

	filepath.WalkDir(projectsRoot, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			return nil
		}

		if strings.HasPrefix(d.Name(), ".") && d.IsDir() {
			return filepath.SkipDir
		}

		if _, ok := projectPaths[path]; ok {
			return filepath.SkipDir
		}

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			return err
		}

		watcher.Add(path)

		return nil
	},
	)
}
