package components

import "github.com/rivo/tview"

// Box is a basic component that wraps tview.Box.
type Box struct {
	*tview.Box
}

// NewBox creates a new Box component.
func NewBox() *Box {
	return &Box{
		Box: tview.NewBox(),
	}
}