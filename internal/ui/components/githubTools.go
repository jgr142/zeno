package components

import (
	"log/slog"
	"os"
	"os/exec"

	"github.com/gdamore/tcell/v2"
	"github.com/jgr142/zeno/internal/git"
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
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
		slog.SetDefault(logger)

		cmd := exec.Command(
			"git",
			"fetch",
			"origin",
			"&&",
			"git",
			"rebase",
			"origin/main",
		)
		out, err := cmd.Output()
		if err != nil {
			slog.Error("Error Running Rebase", "Error", err.Error())
			gt.SetBackgroundColor(tcell.ColorRed)
			return nil
		}
		gt.SetBackgroundColor(tcell.ColorGreen)
		slog.Info("Out Data", "StdOut", out)
	case 'p':
		git.AddCommitPush("feat: new feature")
	case 'c':
		slog.Info("poop")
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
