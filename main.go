// Demo code for the Form primitive.
package main

import (
	"github.com/rivo/tview"
)

func main() {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	main := newPrimitive("Main content")

	grid := tview.NewGrid().
		SetRows(2, 0, 1).
		SetColumns(55, 0).
		SetBorders(true).
		AddItem(newPrimitive("Header"), 0, 0, 1, 2, 0, 0, false).
		AddItem(tview.NewTextView().SetTextAlign(tview.AlignLeft).SetText("F1 ===> HELP   F3 ===> BACK   F12 ===> START"), 2, 0, 1, 2, 0, 0, false)

	app := tview.NewApplication()
	form := tview.NewForm().SetItemPadding(0).
		AddInputField("Starting System", "", 30, nil, nil).
		AddInputField("Funds Available", "", 20, nil, nil).
		AddInputField("Time Available", "", 10, nil, nil).
		AddCheckbox("Citadel Allowed", true, nil).
		AddCheckbox("Lowsec Allowed", false, nil).
		AddCheckbox("0.0 Allowed", false, nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)

	grid.AddItem(form, 1, 0, 1, 1, 0, 100, true).
		AddItem(main, 1, 1, 1, 1, 0, 100, false)
	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
