package Ships

import (
	"Spacetraders/src/General"
	"fmt"

	"github.com/rivo/tview"
)

func DisplayShipState() tview.Primitive {
	window := tview.NewFlex()
	window.SetBorder(false)
	window.SetDirection(tview.FlexRow)

	row_1 := tview.NewFlex()
	row_1.SetBorder(false)
	row_2 := tview.NewFlex()
	row_2.SetBorder(false)

	var box_1 tview.Primitive = BuildShipForm("NULL_SKY-1")
	var box_2 tview.Primitive = BuildShipForm("NULL_SKY-2")

	box_3 := tview.NewTextView().SetBorder(true).SetTitle("  Box 3  ")
	box_4 := tview.NewTextView().SetBorder(true).SetTitle("  Box 4  ")

	row_1.AddItem(box_1, 0, 1, false)
	row_1.AddItem(box_2, 0, 1, false)

	row_2.AddItem(box_3, 0, 1, false)
	row_2.AddItem(box_4, 0, 1, false)

	window.AddItem(row_1, 0, 1, false)
	window.AddItem(row_2, 0, 1, false)
	return window
}

func BuildShipForm(symbol string) tview.Primitive {
	box := tview.NewForm()
	box.SetBorder(true)
	box.SetTitle("  " + symbol + "  ")

	ship := GetShipState(symbol)

	box.AddTextView("Role", ship.Registration.Role, 0, 1, true, true)
	box.AddTextView("Status", ship.Nav.Status, 0, 1, true, true)
	box.AddTextView("Frame", ship.Frame.Name, 0, 1, true, true)
	box.AddTextView("Reactor", ship.Reactor.Name, 0, 1, true, true)
	box.AddTextView("Engine", ship.Engine.Name, 0, 1, true, true)
	box.AddTextView("Mode", ship.Nav.FlightMode, 0, 1, true, true)
	box.AddTextView("Waypoint", ship.Nav.WaypointSymbol, 0, 1, true, true)
	box.AddTextView("Crew", fmt.Sprintf("%d (Min: %d | Max: %d)", ship.Crew.Current, ship.Crew.Required, ship.Crew.Capacity), 0, 1, true, true)
	box.AddTextView("Fuel", General.ProgressBar(ship.Fuel.Current, ship.Fuel.Capacity), 0, 1, true, true)
	box.AddTextView("Morale", General.ProgressBar(ship.Crew.Morale, 100), 0, 1, true, true)
	// box.AddTextView("Cargo", ProgressBar(ship.Cargo.), 0, 1, true, true)

	return box
}
