package components

import (
	"log/slog"
	"os/exec"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GithubTools struct {
	*tview.TextView
	projectPath string
}

func NewGithubTools(projectPath string) *GithubTools {
	gitToolsText := `
	Git Tools:
	-----------
	
	(r) Rebase
	(p) Add + Commit + Push
	(c) Checkout Branch
	(b) Create Branch
	(s) Status
	(d) Diff
	(l) Log
	(u) Pull
	(x) Stash
	(z) Pop Stash
	(q) Quit
	`

	// TextView UI
	gt := &GithubTools{
		tview.NewTextView(),
		projectPath,
	}

	gt.SetText(gitToolsText).
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
		SetBorder(true).
		SetTitle(" GitHub Tools ").
		SetInputCapture(gt.actionListeners)

	return gt
}

func (gt *GithubTools) actionListeners(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	// TODO: Implement handlers
	case 'r':
		cmd := exec.Command(
			"git",
			"pull",
			"origin",
			"main",
			"&&",
			"git",
			"rebase",
			"main",
		)
		_, err := cmd.Output()
		if err != nil {
			slog.Info("Error Running Rebase", "Error", err.Error())
			return nil
		}
	case 'p':
	case 'c':
	case 'b':
	case 's':
	case 'd':
	case 'l':
	case 'u':
	case 'x':
	case 'z':
	case 'q':
	}
	return event
}
