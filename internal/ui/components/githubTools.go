package components

import (
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
	// pages       *tview.Pages
	projectPath string
	notify      func(msg string, isErr bool)
}

func NewGithubTools(projectPath string) *GithubTools {
	// Set Up Flex Structure
	mainDisplay := tview.NewTextView()
	notif := NewNotification()
	skltn := tview.
		NewFlex().
		AddItem(notif, 1, 0, false).
		AddItem(mainDisplay, 0, 1, true)

		// Configure Pages that you can switch to
	gt := &GithubTools{
		skltn,
		// pages,
		projectPath,
		notif.DisplayNotification,
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
	case 'p':
		err := git.AddCommitPush("feat: new feature")
		if err != nil {
			gt.notify("ACP Process Failed", true)
		} else {
			gt.notify("ACP Success", false)
		}
		return nil
	case 'c':
		// Make a text input modal show up
	case 'b':
		branchForm := NewForm()
		form := branchForm.Form()

		form.
			AddDropDown("Branch Type", []string{"feat", "fix", "refactor"}, 0, nil).
			AddInputField("Branch Name", "", 20, nil, nil).
			AddButton("Submit", func() {
				_, branchType := form.GetFormItemByLabel("Branch Type").(*tview.DropDown).GetCurrentOption()
				branchName := form.GetFormItemByLabel("Branch Name").(*tview.InputField).GetText()

				if branchName == "" {
					branchForm.Notify("Please enter a branch name", true)
					return
				}

				err := git.CreateBranch("jgr142/" + branchType + "/" + branchName)
				if err != nil {
					branchForm.notify("Branch Creation Failed: "+err.Error(), true)
					return
				}
				// gt.pages.SwitchToPage("github tools")
				// gt.pages.RemovePage("new branch")
			})
		// gt.pages.AddAndSwitchToPage("new branch", branchForm, true)
	case 's':
		// status, err := git.Status()
		return nil
	case 'l':
		// log, err := git.Log()
		// if err != nil {
		// 	gt.notify(
		// 		"Error fetching logs",
		// 		true,
		// 	)
		// 	return nil
		// }
		//
		//
		// return nil
	case 'u':
		// assuming we set origin everywhere we go...
		err := git.Pull()
		if err != nil {
			gt.notify(
				"Could not pull from "+gt.projectPath,
				true,
			)
		} else {
			gt.notify(
				"Successfully pulled from "+gt.projectPath,
				false,
			)
		}
		return nil
	case 'x':
		err := git.Stash()
		if err != nil {
			gt.notify(
				"Could not stash changes",
				true,
			)
		} else {
			gt.notify(
				"Successfully stashed changes",
				false,
			)
		}
		return nil
	case 'z':
		err := git.Pop()
		if err != nil {
			gt.notify(
				"Could not pop the stashed changes",
				true,
			)
		} else {
			gt.notify(
				"Could not pop the stashed changes",
				false,
			)
		}
		return nil
	case 'q':
		// gt.pages.SwitchToPage("projects search")
	}
	return event
}
