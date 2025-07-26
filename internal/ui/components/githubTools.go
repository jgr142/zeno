package components

import (
	"log/slog"
	"os"
	"os/exec"

	"github.com/gdamore/tcell/v2"
	"github.com/jgr142/zeno/internal/git"
	"github.com/rivo/tview"
)

var (
	gitToolsText = `
	Git Tools:
	-----------
	
	(r) Rebase
	(p) Add + Commit + Push
	(c) Checkout Branch
	(b) Create Branch
	(s) Status
	(l) Log
	(u) Pull
	(x) Stash
	(z) Pop Stash
	(q) Quit
	`
)

type GithubTools struct {
	*tview.Flex
	pages       *tview.Pages
	projectPath string
	notify      func(successMsg, errMsg string, isErr bool)
}

func NewGithubTools(pages *tview.Pages, projectPath string) *GithubTools {
	// Set Up Flex Structure
	mainDisplay := tview.NewTextView()
	notif := NewNotification()
	skltn := tview.NewFlex().
		AddItem(notif, 1, 0, false).
		AddItem(mainDisplay, 0, 1, true)

	gt := &GithubTools{
		skltn,
		pages,
		projectPath,
		notif.DisplaySuccessorErr,
	}

	// Configure mainDisplay
	mainDisplay.SetText(gitToolsText).
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
		// Make a text input modal show up
	case 'b':
		// Again make a text intput modal show up
	case 's':
		// status, err := git.Status()
		return nil
	case 'l':
		// log, err := git.Log()
		// if err != nil {
		// 	// show err message
		// 	gt.pages.AddAndSwitchToPage("error modal", NewErrorModal(err.Error()), true)
		// } else {
		// 	//show git status page
		// 	gt.pages.AddAndSwitchToPage("non-err page", NewThemedTextModal(string(status)), true)
		// }
		return nil
		// Show logs
	case 'u':
		// assuming we set origin everywhere we go...
		err := git.Pull()
		gt.notify(
			"Successfully pulled from "+gt.projectPath,
			"Could not pull from "+gt.projectPath,
			err != nil,
		)
	case 'x':
		err := git.Stash()
		gt.notify(
			"Successfully stashed changes",
			"Could not stash changes",
			err != nil,
		)
	case 'z':
		err := git.Pop()
		gt.notify(
			"Successfully popped stashed changes",
			"Could not pop the stashed changes",
			err != nil,
		)
	case 'q':
		gt.pages.SwitchToPage("project search")
	}
	return event
}
