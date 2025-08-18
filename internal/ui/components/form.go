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

func (f *Form) NavigateDown() {
	idx := -1
	item, button := f.form.GetFocusedItemIndex()
	if item != -1 {
		idx = item
	}

	if button != -1 {
		idx = item
	}

	if idx == -1 || idx == f.form.GetFormItemCount()-1 {
		return
	}

	f.form.SetFocus(idx + 1)
}

func (f *Form) NavigateUp() {
	idx := -1
	item, button := f.form.GetFocusedItemIndex()
	if item != -1 {
		idx = item
	}

	if button != -1 {
		idx = item
	}

	if idx == -1 || idx == 0 {
		return
	}

	f.form.SetFocus(idx - 1)
}
