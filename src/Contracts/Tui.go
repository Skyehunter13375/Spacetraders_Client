package Contracts

import (
	"fmt"
	"strconv"
	"Spacetraders/src/General"

	"github.com/rivo/tview"
)

func DisplayContractMenu(app *General.App) tview.Primitive {
	app.UIState.SubMenu.Clear()
	app.UIState.Output.Clear()

	CList, err := General.PG.Query("SELECT id FROM contracts")
	if err != nil { General.LogErr("DisplayFleetMenu: " + err.Error()) }

	window := tview.NewFlex()
	CMenu  := tview.NewList()
	CMenu.SetBorder(true)
	CMenu.SetTitle("  Current Contracts  ")

	// Create a button for each ship in the fleet
	for CList.Next() {
		var id string
		CList.Scan(&id)
		sym   := id
		data  := GetContract(sym)

		CStatus := "Unknown"
		if (data.Accepted == true && data.Fulfilled == false) {
			CStatus = "Accepted"
		} else if (data.Accepted == true && data.Fulfilled == true) {
			CStatus = "Fulfilled"
		} else {
			CStatus = "Available"
		}

		label := fmt.Sprintf("[ %s ] Faction: %-10s | Type: %-10s | Payment: %-10d", CStatus, data.Faction, data.Type, data.Terms.Payment.OnAccepted + data.Terms.Payment.OnFulfilled )

		CMenu.AddItem(label, "", 0, func() {
			app.UIState.SubMenu.Clear()
			app.UIState.SubMenu.AddItem("Show Details", "", 0, func() { app.UIState.Output.Clear() })
			app.UIState.SubMenu.AddItem("Back",         "", 0, func() { app.UIState.Output.Clear(); DisplayContractMenu(app) } )
			app.UI.SetFocus(app.UIState.SubMenu)
		})
	}
	
	// Make sure we always have a back button at the end that takes us to the main menu
	CMenu.AddItem("Back", "", 0, func() { app.UIState.SubMenu.Clear(); app.UIState.Output.Clear(); app.UI.SetFocus(app.UIState.MainMenu) })

	// Add the menu to the window, add the window to the output and set focus to the output
	window.AddItem(CMenu, 0, 1, true)
	app.UIState.Output.AddItem(window, 0, 1, true)
	app.UI.SetFocus(app.UIState.Output)

	return window
}

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
