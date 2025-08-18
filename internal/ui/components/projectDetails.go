package components

import (
	"github.com/jgr142/zeno/internal/ui/theme"
	"github.com/rivo/tview"
)

type ProjectDetails struct {
	*tview.TextView
}

func NewProjectDetails() *ProjectDetails {
	projectDetails := &ProjectDetails{tview.NewTextView()}

	projectDetails.
		SetDynamicColors(true).
		SetWrap(true).
		SetTextAlign(tview.AlignLeft).
		SetBackgroundColor(theme.Background)

	projectDetails.SetBorder(true).
		SetTitle(" [::b]Details ")

	return projectDetails
}
