package Waypoints

import (
	"Spacetraders/src/General"
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

func ShowSystemsMenu(app *General.App) {
	ui := app.UIState
	ui.SubMenu.Clear()
	ui.SubMenu.AddItem("Select System", "", 0, func() {
		ui.Output.Clear()
		ui.Output.AddItem(tview.NewTextView().
			SetText("Systems Placeholder"), 0, 1, false)
	})
	ui.SubMenu.AddItem("Back", "", 'b', func() { General.FocusMain(app) } )
}

func DisplaySystem(id string) tview.Primitive {
    System := GetSystem(id)

    window := tview.NewFlex().SetDirection(tview.FlexRow)

    // Title box
    title := tview.NewTextView()
    title.SetBorder(true)
    title.SetTitle(fmt.Sprintf("  %s  ", id))
    title.SetText(fmt.Sprintf(
        "Name: %s | Constellation: %s | Type: %s | %d:%d",
        System.Name, System.Constellation, System.Type,
        System.Xcoord, System.Ycoord,
    ))

    window.AddItem(title, 3, 0, false)

    // Query waypoint symbols
    ids, _ := General.PG.Query(`SELECT symbol FROM waypoints WHERE system = $1`, id)
    var symbols []string
    for ids.Next() {
        var s string
        ids.Scan(&s)
        symbols = append(symbols, s)
    }

    // --- Build grid layout (4 per row) ---
    row := tview.NewFlex().SetDirection(tview.FlexColumn)
    count := 0

    for i, sym := range symbols {
		if i > 11 { continue }

        wpBox := DisplayWaypoint(GetWaypoint(sym))

        row.AddItem(wpBox, 0, 1, false)
        count++

        // If row is full → append to window, start new row
        if count%4 == 0 {
            window.AddItem(row, 0, 1, false)
            row = tview.NewFlex().SetDirection(tview.FlexColumn)
        }
    }

    // Add final row if it has leftovers
    if count%4 != 0 {
        window.AddItem(row, 0, 1, false)
    }

    return window
}

func DisplayWaypoint(data Waypoint) tview.Primitive {

	box := tview.NewForm()
	box.SetBorder(true)
	box.SetTitle(fmt.Sprintf("  %s  ", data.Symbol))

	box.AddTextView("Type",          data.Type, 0, 1, true, true)
	box.AddTextView("Coords:",       fmt.Sprintf("%d:%d", data.X, data.Y), 0, 1, true, true)
	box.AddTextView("Orbits:",       data.Orbits, 0, 1, true, true)
	box.AddTextView("Construction:", strconv.FormatBool(data.Construction), 0, 1, true, true)
	box.AddTextView("Traits:",       "", 0, 1, true, true)

	return box
}
