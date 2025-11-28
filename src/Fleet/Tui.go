package Fleet

import "Spacetraders/src/General"
import "github.com/rivo/tview"
import "github.com/gdamore/tcell/v2"
// import "fmt"

func BuildShipForm(symbol string) tview.Primitive {
	box  := tview.NewForm()
	ship := GetShipState(symbol)

	box.AddTextView("Name:",     ship.Symbol, 0, 1, true, true)
	box.AddTextView("Role:",     ship.Registration.Role + " (" + ship.Frame.Name + ")", 0, 1, true, true)
	box.AddTextView("Status:",   ship.Nav.Status + " (" + ship.Nav.FlightMode + ")",    0, 1, true, true)
	if ship.Nav.WaypointSymbol == ship.Nav.Route.Destination.Symbol {
		box.AddTextView("Waypoint:", ship.Nav.WaypointSymbol, 0, 1, true, true)
	} else {
		box.AddTextView("Waypoint:", ship.Nav.WaypointSymbol + " -> " + ship.Nav.Route.Destination.Symbol, 0, 1, true, true)
	}
	box.AddTextView("Crew:",     General.ProgressBar(ship.Crew.Current, ship.Crew.Required, ship.Crew.Capacity), 0, 1, true, true)
	box.AddTextView("Fuel:",     General.ProgressBar(ship.Fuel.Current, 0,                  ship.Fuel.Capacity), 0, 1, true, true)
	box.AddTextView("Morale:",   General.ProgressBar(ship.Crew.Morale,  0,                  100),                0, 1, true, true)

	return box
}

func DisplayFleetMenu(app *General.App) tview.Primitive {
	app.UIState.SubMenu.Clear()
	app.UIState.Output.Clear()
	window := tview.NewFlex().SetDirection(tview.FlexRow)

	// First we need a list of ships to build cards for
	ShipList, err := General.PG.Query("SELECT symbol FROM ships")
	if err != nil { General.LogErr("DisplayFleetMenu: " + err.Error()) }

	var symbols []string
	for ShipList.Next() {
		var sym string
		ShipList.Scan(&sym)
		symbols = append(symbols, sym)
	}

	// Defining some default parameters for how the cards will be displayed in the menu
	const cardsPerRow = 5
	const cardHeight  = 19
	const cardWidth   = 43

	// Define the submenu for each card when selected
	var cards []*General.CardButton
	for _, sym := range symbols {
		localSym := sym // capture loop variable
		card := General.NewCardButton(
			BuildShipForm(localSym),
			"",
			func() {
				app.UI.SetFocus(app.UIState.SubMenu)
				app.UIState.SubMenu.Clear()
				app.UIState.SubMenu.AddItem("Move to waypoint", "", 0, nil)
				app.UIState.SubMenu.AddItem("Scan waypoint",    "", 0, nil)
				app.UIState.SubMenu.AddItem("Unload cargo",     "", 0, nil)
				app.UIState.SubMenu.AddItem("Eject cargo",      "", 0, nil)
				app.UIState.SubMenu.AddItem("Repair ship",      "", 0, nil)
				app.UIState.SubMenu.AddItem("Sell ship",        "", 0, nil)
				app.UIState.SubMenu.AddItem("Back",             "", 0, func() { app.UIState.Output.Clear(); DisplayFleetMenu(app) })
			},
		)
		cards = append(cards, card)
	}

	// Define the grid for the cards to live in, this way they can be selected with arrow keys
	var grid [][]*General.CardButton
	for i := 0; i < len(cards); i += cardsPerRow {
		end := i + cardsPerRow
		end  = min(end, len(cards)) // if end > len(cards) { end = len(cards) }
		grid = append(grid, cards[i:end])
	}

	// Insert the built cards into their rows
	for _, rowCards := range grid {
		rowFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
		rowFlex.SetBorder(false)

		for _, card := range rowCards {
			rowFlex.AddItem(card, cardWidth, 0, false)
		}

		window.AddItem(rowFlex, cardHeight, 0, false)
	}

	// Create input capture rules because by default flex objects cannot be focused so we have to define that logic manually
	row, col := 0, 0
	window.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
		switch ev.Key() {
			case tcell.KeyRight:
				if col < len(grid[row])-1 {
					grid[row][col].Blur()
					col++
					app.UI.SetFocus(grid[row][col])
				}
				return nil

			case tcell.KeyLeft:
				if col > 0 {
					grid[row][col].Blur()
					col--
					app.UI.SetFocus(grid[row][col])
				}
				return nil

			case tcell.KeyDown:
				if row < len(grid)-1 {
					if col >= len(grid[row+1]) {
						col = len(grid[row+1]) - 1
					}
					grid[row][col].Blur()
					row++
					app.UI.SetFocus(grid[row][col])
				}
				return nil

			case tcell.KeyUp:
				if row > 0 {
					if col >= len(grid[row-1]) {
						col = len(grid[row-1]) - 1
					}
					grid[row][col].Blur()
					row--
					app.UI.SetFocus(grid[row][col])
				}

			case tcell.KeyF1:
				app.UIState.SubMenu.Clear()
				app.UIState.Output.Clear()
				app.UI.SetFocus(app.UIState.MainMenu)
				return nil
		}

		return ev
	})

	// Add the window to the output field, auto-select the first card so input capture works, and return control to the user
	app.UIState.Output.AddItem(window, 0, 1, true)

	first := grid[0][0]
	app.UI.SetFocus(first)

	return window
}
