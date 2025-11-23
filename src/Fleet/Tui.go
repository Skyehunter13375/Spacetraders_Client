package Fleet

import "Spacetraders/src/General"
import "github.com/rivo/tview"

func DisplayShipState() tview.Primitive {
	window := tview.NewFlex()
	window.SetBorder(false)
	window.SetDirection(tview.FlexRow)

	row_1 := tview.NewFlex()
	row_1.SetBorder(false)

	var box_1 tview.Primitive = BuildShipForm("NULL-SKY-1")
	var box_2 tview.Primitive = BuildShipForm("NULL-SKY-2")

	row_1.AddItem(box_1, 0, 1, false)
	row_1.AddItem(box_2, 0, 1, false)

	window.AddItem(row_1, 0, 1, false)
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
