package components

import "github.com/rivo/tview"

type Form struct {
	*tview.Flex
	form   *tview.Form
	notify func(msg string, isError bool)
}

func NewForm() *Form {
	notif := NewNotification()
	form := tview.NewForm()

	skltn := tview.NewFlex().
		AddItem(notif, 3, 0, false).
		AddItem(form, 0, 1, true)

	wrapper := &Form{skltn, form, notif.DisplayNotification}
	return wrapper
}

func (f *Form) Form() *tview.Form {
	return f.form
}

func (f *Form) Notify(msg string, isError bool) {
	f.notify(msg, isError)
}
