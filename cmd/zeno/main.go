package main

import (
	"github.com/jgr142/zeno/internal/domain"
	"github.com/jgr142/zeno/internal/infra/persistence"
	"github.com/jgr142/zeno/internal/infra/ui"
)

func main() {
	projectRepo := persistence.New()
	projects := make([]domain.Project, 0)
	for key := range projectRepo.Get() {
		projects = append(projects, key)
	}
	ui.Init(projects)
}
