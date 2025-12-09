package Waypoints

// import "fmt"
import "strconv"
import "Spacetraders/src/General"
import "github.com/rivo/tview"
import "github.com/gdamore/tcell/v2"

func DisplaySystem(id string) tview.Primitive {
	System := GetSystem(id)
	var Waypoints int64
	var Markets   int64
	var Shipyards int64
	General.PG.QueryRow("SELECT COUNT(*) FROM waypoints WHERE system = $1", id).Scan(&Waypoints)
	General.PG.QueryRow("SELECT COUNT(*) FROM waypoints WHERE system = $1 AND 'MARKETPLACE' = ANY(traits)", id).Scan(&Markets)
	General.PG.QueryRow("SELECT COUNT(*) FROM waypoints WHERE system = $1 AND 'SHIPYARD'    = ANY(traits)", id).Scan(&Shipyards)
	form := tview.NewForm()
	form.SetBorder(false)
	form.SetBackgroundColor(General.Theme.BgBase)
	form.AddTextView("Symbol:",    System.Symbol, 0, 1, true, true)
	form.AddTextView("Type:",      System.Type,   0, 1, true, true)
	form.AddTextView("Waypoints:", strconv.FormatInt(Waypoints, 10), 0, 1, true, true)
	form.AddTextView("Markets:",   strconv.FormatInt(Markets,   10), 0, 1, true, true)
	form.AddTextView("Shipyards:", strconv.FormatInt(Shipyards, 10), 0, 1, true, true)
	return form
}

func DisplayWaypoint(id string) tview.Primitive {
	Waypoint := GetWaypoint(id)
	form := tview.NewForm()
	form.SetBorder(false)
	form.SetBackgroundColor(General.Theme.BgBase)
	form.AddTextView("Symbol",        Waypoint.Symbol,                              0, 1, true, true)
	form.AddTextView("Type",          Waypoint.Type,                                0, 1, true, true)
	// form.AddTextView("Coords:",       fmt.Sprintf("%d:%d", Waypoint.X, Waypoint.Y), 0, 1, true, true)
	// form.AddTextView("Orbits:",       Waypoint.Orbits,                              0, 1, true, true)
	// form.AddTextView("Construction:", strconv.FormatBool(Waypoint.Construction),    0, 1, true, true)
	// form.AddTextView("Traits:",       "",                                           0, 1, true, true)
	return form 
}

func DisplaySystemMenu(app *General.App) tview.Primitive {
	app.UIState.SubMenu.Clear()
	app.UIState.Output.Clear()
	window := tview.NewFlex().SetDirection(tview.FlexRow)

	// TASK: First check if there is any data to show
	var count int
	err := General.PG.QueryRow("SELECT COUNT(*) FROM systems").Scan(&count)
	if err != nil { General.LogErr("DisplaySystemMenu: " + err.Error()) }
	// DEBG: Need a function here to capture starting system
	if count == 0 { panic("Need a function to capture starting system here") }

	// TASK: Get a list of ships to build cards for
	SysList, err := General.PG.Query("SELECT symbol FROM systems")
	if err != nil { General.LogErr("DisplaySystemMenu: " + err.Error()) }

	var symbols []string
	for SysList.Next() {
		var sym string
		SysList.Scan(&sym)
		symbols = append(symbols, sym)
	}

	// Defining some default parameters for how the cards will be displayed in the menu
	const cardsPerRow = 5
	const cardHeight  = 13
	const cardWidth   = 43

	// Define the submenu for each card when selected
	var cards []*General.CardButton
	for _, sym := range symbols {
		localSym := sym // capture loop variable
		card := General.NewCardButton(
			DisplaySystem(localSym),
			"",
			func() { 
				DisplayWaypointMenu(app, localSym)
				// app.UI.SetFocus(app.UIState.SubMenu)
				// app.UIState.SubMenu.Clear()
				// app.UIState.SubMenu.AddItem("Back", "", 0, func() { app.UIState.Output.Clear(); DisplaySystemMenu(app) })
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

func DisplayWaypointMenu(app *General.App, System string) tview.Primitive {
	app.UIState.SubMenu.Clear()
	app.UIState.Output.Clear()
	window := tview.NewFlex().SetDirection(tview.FlexRow)

	// FEATURE: Should split data into segments by type
	// Planet  := tview.NewFlex().SetDirection(tview.FlexRow)
	// Fuel    := tview.NewFlex().SetDirection(tview.FlexRow)
	// Jump    := tview.NewFlex().SetDirection(tview.FlexRow)
	// Moon    := tview.NewFlex().SetDirection(tview.FlexRow)
	// Station := tview.NewFlex().SetDirection(tview.FlexRow)
	// pcontent := tview.NewTextView()
	// pcontent.SetText("Planets")
	// Planet.AddItem(pcontent, 0, 1, false)
	// window.AddItem(pcontent, 1, 1, false)

	// First we need a list of ships to build cards for
	WayList, err := General.PG.Query("SELECT symbol FROM waypoints WHERE system = $1", System)
	if err != nil { General.LogErr("DisplaySystemMenu: " + err.Error()) }

	var symbols []string
	for WayList.Next() {
		var sym string
		WayList.Scan(&sym)
		symbols = append(symbols, sym)
	}

	// Defining some default parameters for how the cards will be displayed in the menu
	const cardsPerRow = 6 
	const cardHeight  = 7 
	const cardWidth   = 32

	// Define the submenu for each card when selected
	var cards []*General.CardButton
	for _, sym := range symbols {
		localSym := sym // capture loop variable
		card := General.NewCardButton(
			DisplayWaypoint(localSym),
			"",
			func() {
				app.UI.SetFocus(app.UIState.SubMenu)
				app.UIState.SubMenu.Clear()
				app.UIState.SubMenu.AddItem("Back", "", 0, func() { app.UIState.Output.Clear(); DisplaySystemMenu(app) })
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
