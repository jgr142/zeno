package main

import (
	"github.com/jgr142/zeno/internal/infra/persistence"
	"github.com/jgr142/zeno/internal/infra/ui"
)

func main() {
	projectRepo := persistence.New()
	projects := projectRepo.Get()
	projectList := projects.Items()
	ui.Init(projectList)
}
