package Fleet

import "Spacetraders/src/General"
import "github.com/rivo/tview"
import "fmt"

func DisplayFleetMenu(app *General.App) tview.Primitive {
	app.UIState.SubMenu.Clear()

	ShipList, err := General.PG.Query("SELECT symbol FROM ships")
	if err != nil { General.LogErr("DisplayFleetMenu: " + err.Error()) }

	window    := tview.NewFlex()
	ShipsMenu := tview.NewList()
	ShipsMenu.SetBorder(true)
	ShipsMenu.SetTitle("  Current Fleet  ")

	// Create a button for each ship in the fleet
	for ShipList.Next() {
		var symbol string
		ShipList.Scan(&symbol)
		sym  := symbol
		data := GetShipState(sym)
		label := fmt.Sprintf("%-10s | Type: %-10s | Status: %-10s", 
			data.Symbol,
			data.Frame.Name,
			data.Nav.Status,
		)

		ShipsMenu.AddItem(label, "", 0, func() {
			app.UIState.SubMenu.Clear()
			app.UIState.SubMenu.AddItem("Show Details", "", 0, func() {
				app.UIState.Output.Clear()
				app.UIState.Output.AddItem(BuildShipForm(symbol), 0, 1, false)
			})
			app.UIState.SubMenu.AddItem("Nav to waypoint", "", 0, nil)
			app.UIState.SubMenu.AddItem("Scan Waypoint",   "", 0, nil)
			app.UIState.SubMenu.AddItem("Repair ship",     "", 0, nil)
			app.UIState.SubMenu.AddItem("Unload cargo",    "", 0, nil)
			app.UIState.SubMenu.AddItem("Back",            "", 0, func() { app.UIState.Output.Clear(); DisplayFleetMenu(app) } )
			app.UI.SetFocus(app.UIState.SubMenu)
		})
	}
	
	// Make sure we always have a back button at the end that takes us to the main menu
	ShipsMenu.AddItem("Back", "", 0, func() { app.UIState.SubMenu.Clear(); app.UIState.Output.Clear(); app.UI.SetFocus(app.UIState.MainMenu) })

	// Add the menu to the window, add the window to the output and set focus to the output
	window.AddItem(ShipsMenu, 0, 1, true)
	app.UIState.Output.AddItem(window, 0, 1, true)
	app.UI.SetFocus(app.UIState.Output)

	return window
}

func BuildShipForm(symbol string) tview.Primitive {
	box := tview.NewForm()
	box.SetBorder(true)
	box.SetTitle("  " + symbol + "  ")

	ship := GetShipState(symbol)

	box.AddTextView("Role",     ship.Registration.Role,  0, 1, true, true)
	box.AddTextView("Status",   ship.Nav.Status,         0, 1, true, true)
	box.AddTextView("Frame",    ship.Frame.Name,         0, 1, true, true)
	box.AddTextView("Reactor",  ship.Reactor.Name,       0, 1, true, true)
	box.AddTextView("Engine",   ship.Engine.Name,        0, 1, true, true)
	box.AddTextView("Mode",     ship.Nav.FlightMode,     0, 1, true, true)
	box.AddTextView("Waypoint", ship.Nav.WaypointSymbol, 0, 1, true, true)
	box.AddTextView("Crew",     General.ProgressBar(ship.Crew.Current, ship.Crew.Required, ship.Crew.Capacity), 0, 1, true, true)
	box.AddTextView("Fuel",     General.ProgressBar(ship.Fuel.Current, 0,                  ship.Fuel.Capacity), 0, 1, true, true)
	box.AddTextView("Morale",   General.ProgressBar(ship.Crew.Morale,  0,                  100),                0, 1, true, true)

	return box
}
