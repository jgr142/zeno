package main

import (
	"github.com/jgr142/zeno/internal/infra/persistence"
	"github.com/jgr142/zeno/internal/infra/ui"
)

func main() {
	projRepo := persistence.New()
	ui.Init(projRepo.Get())
}
