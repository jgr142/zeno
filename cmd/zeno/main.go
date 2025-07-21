package main

import (
	"github.com/jgr142/zeno/internal/project"
	"github.com/jgr142/zeno/internal/ui"
)

func main() {
	projectRepo := project.New()
	ui.Init(projectRepo)
}
