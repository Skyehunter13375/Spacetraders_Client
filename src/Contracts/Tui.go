package Contracts

import "Spacetraders/src/General"
import "github.com/rivo/tview"
import "github.com/gdamore/tcell/v2"
import "time"

func DisplayContractMenu(app *General.App) tview.Primitive {
	CFG, _ := General.GetConfig()
	app.UIState.SubMenu.Clear()
	app.UIState.Output.Clear()
	window := tview.NewFlex().SetDirection(tview.FlexRow)

	// TASK: First check if there is any data to show
	var count int
	err := General.PG.QueryRow("SELECT COUNT(*) FROM contracts").Scan(&count)
	if err != nil { General.LogErr("DisplayContractMenu: " + err.Error()) }
	if count == 0 { NegotiateNewContract(CFG.API.AgentName + "-1") }

	// TASK: Get a list of ships to build cards for
	ContList, err := General.PG.Query("SELECT id FROM contracts")
	if err != nil { General.LogErr("DisplayContractMenu: " + err.Error()) }

	var symbols []string
	for ContList.Next() {
		var sym string
		ContList.Scan(&sym)
		symbols = append(symbols, sym)
	}

	// Defining some default parameters for how the cards will be displayed in the menu
	const cardsPerRow = 5
	const cardHeight  = 16 
	const cardWidth   = 43 

	// Define the submenu for each card when selected
	var cards []*General.CardButton
	for _, sym := range symbols {
		localSym := sym // capture loop variable
		card := General.NewCardButton(
			BuildContractForm(localSym),
			"",
			func() {
				app.UI.SetFocus(app.UIState.SubMenu)
				app.UIState.SubMenu.Clear()
				app.UIState.SubMenu.AddItem("Update Contract", "", 0, nil)
				app.UIState.SubMenu.AddItem("Accept Contract", "", 0, nil)
				app.UIState.SubMenu.AddItem("Drop Contract",   "", 0, nil)
				app.UIState.SubMenu.AddItem("Delete Contract", "", 0, nil)
				app.UIState.SubMenu.AddItem("Back",            "", 0, func() { app.UIState.Output.Clear(); DisplayContractMenu(app) })
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

func BuildContractForm(id string) tview.Primitive {
	box  := tview.NewForm()
	box.SetBackgroundColor(tcell.GetColor("#2A2E2A"))
	data := GetContract(id)

	status := "Available"
	if data.Accepted == true && data.Fulfilled == true {
		status = "Fulfilled"
	} else if data.Accepted == true && data.Fulfilled == false {
		status = "Accepted"
	}

	accepttimestamp, err := time.Parse(time.RFC3339Nano, data.DeadlineToAccept)
	if err != nil { General.LogErr("BuildContractForm: accepttimestampconvert: " + err.Error()) }

	deadlinetimestamp, err := time.Parse(time.RFC3339Nano, data.Terms.Deadline)
	if err != nil { General.LogErr("BuildContractForm: deadlinetimestampconvert: " + err.Error()) }


	box.AddTextView("ID:",           data.ID,               0, 1, true, true)
	box.AddTextView("Faction:",      data.Faction,          0, 1, true, true)
	box.AddTextView("Type:",         data.Type,             0, 1, true, true)
	box.AddTextView("Status:",       status,                0, 1, true, true)
	box.AddTextView("MustAcceptBy:", accepttimestamp.Format("01/02/2006 @ 15:04:05"), 0, 1, true, true)
	box.AddTextView("Deadline:",     deadlinetimestamp.Format("01/02/2006 @ 15:04:05"),   0, 1, true, true)

	// box.AddTextView("Fuel:",     General.ProgressBar(ship.Fuel.Current, 0,                  ship.Fuel.Capacity), 0, 1, true, true)

	return box
}
