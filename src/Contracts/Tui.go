package Contracts

import (
	"fmt"
	"strconv"
	"Spacetraders/src/General"

	"github.com/rivo/tview"
)

func ShowContractsMenu(app *General.App) {
	ui := app.UIState
	ui.SubMenu.Clear()
	ui.SubMenu.AddItem("List Contracts", "", 0, func() {
		ui.Output.Clear()
		ui.Output.AddItem(tview.NewTextView().
			SetText("Contracts Placeholder"), 0, 1, false)
	})
	ui.SubMenu.AddItem("Back", "", 'b', func() { General.FocusMain(app) } )
}

func DisplayContract(data Contract) tview.Primitive {
	box := tview.NewForm()
	box.SetBorder(true)
	box.SetTitle("  " + data.ID + "  ")

	box.AddTextView("Contract:",  data.ID,       0, 1, true, true)
	box.AddTextView("Faction:",   data.Faction,  0, 1, true, true)
	box.AddTextView("Type:",      data.Type,     0, 1, true, true)
	box.AddTextView("Payment:",   strconv.Itoa(data.Terms.Payment.OnAccepted + data.Terms.Payment.OnFulfilled),  0, 1, true, true)
	box.AddTextView("Accepted:",  strconv.FormatBool(data.Accepted),  0, 1, true, true)
	box.AddTextView("Fulfilled:", strconv.FormatBool(data.Fulfilled), 0, 1, true, true)
	box.AddTextView("Expires:",   data.Expiration, 0, 1, true, true)
	box.AddTextView("Deadline:",  data.DeadlineToAccept, 0, 1, true, true)

	for i,v := range data.Terms.Deliver {
		box.AddTextView("Material " + strconv.Itoa(i+1) + ":", fmt.Sprintf("%d/%d | %s | %s", v.UnitsFulfilled, v.UnitsRequired, v.Material, v.Destination), 0, 1, true, true)
	}

	return box
}
