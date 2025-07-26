package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Notification struct {
	*tview.TextView
}

func NewNotification() *Notification {
	return &Notification{tview.NewTextView()}
}

func (n *Notification) DisplaySuccessorErr(
	successMsg string,
	errMsg string,
	isErr bool,
) {
	if isErr {
		n.setErrorUI()
		n.SetText(errMsg)
	} else {
		n.setSuccessUI()
		n.SetText(successMsg)
	}
}

func (n *Notification) setErrorUI() {
	n.
		SetTextColor(tcell.ColorWhite).
		SetBackgroundColor(tcell.ColorDarkRed)
}

func (n *Notification) setSuccessUI() {
	n.
		SetTextColor(tcell.ColorWhite).
		SetBackgroundColor(tcell.ColorDarkGreen)

}
