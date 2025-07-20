package project

import (
	"log"
	"os/exec"

	"github.com/jgr142/zeno/internal/domain"
)

func Open(project domain.Project) {
	// Open the project
	cmd := exec.Command("code", project.Path)
	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}
