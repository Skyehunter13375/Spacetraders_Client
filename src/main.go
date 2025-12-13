package main

import "Spacetraders/src/Tui"
import "Spacetraders/src/Model"
import "Spacetraders/src/Task"
import "github.com/rivo/tview"

// ┌──────────────────────────────────────────────────────────────────────────────┐
// │                              Main Layout Shell                               │
// └──────────────────────────────────────────────────────────────────────────────┘
// FEAT: Create TUI app to track state
func NewApp() *Model.App {
	return &Model.App{
		UI:      tview.NewApplication(),
		UIState: &Model.UIState{},
		State:   &Model.GlobalState{},
	}
}

// FEAT: Register all of the standard menus with the app then we can just call and focus them as needed
func BuildLayoutShell(app *Model.App) tview.Primitive {
	// MAIN MENU (top-left)
	mainMenu := tview.NewList()
	mainMenu.ShowSecondaryText(false)
	mainMenu.SetBorder(true)
	mainMenu.SetTitle("  Main Menu  ")
	mainMenu.SetBorderColor(Model.Theme.BgBorder)
	mainMenu.AddItem("Server Status", "", 0, func() { Tui.ShowServerMenu(app)      })
	mainMenu.AddItem("Agent Status",  "", 0, func() { Tui.ShowAgentsMenu(app)      })
	mainMenu.AddItem("Fleet Status",  "", 0, func() { Tui.DisplayFleetMenu(app)    })
	mainMenu.AddItem("Systems",       "", 0, func() { Tui.DisplaySystemMenu(app)   })
	mainMenu.AddItem("Contracts",     "", 0, func() { Tui.DisplayContractMenu(app) })
	mainMenu.AddItem("Settings",      "", 0, func() { Tui.ShowSettingsMenu(app)    })
	mainMenu.AddItem("Quit",          "", 0, func() { app.UI.Stop()                })

	// SUBMENU (bottom-left, dynamic)
	subMenu := tview.NewList()
	subMenu.ShowSecondaryText(false)
	subMenu.SetBorder(true)
	subMenu.SetTitle("  Submenu  ")
	subMenu.SetBorderColor(Model.Theme.BgBorder)

	// OUTPUT PANEL (right)
	output := tview.NewFlex().SetDirection(tview.FlexRow)
	output.SetBorder(false)

	// LEFT COLUMN (main menu + submenu)
	left := tview.NewFlex()
	left.SetDirection(tview.FlexRow)
	left.AddItem(mainMenu, 0, 2, true)
	left.AddItem(subMenu,  0, 3, false)

	// FULL WINDOW LAYOUT
	window := tview.NewFlex()
	window.AddItem(left,   30, 1, true)
	window.AddItem(output, 0,  3, false)

	// Store UI references in App.State so other screens can update them
	app.UIState.MainMenu = mainMenu
	app.UIState.SubMenu  = subMenu
	app.UIState.Output   = output

	return window
}

// ┌──────────────────────────────────────────────────────────────────────────────┐
// │                                     Main                                     │
// └──────────────────────────────────────────────────────────────────────────────┘
func main() {
	if err := Task.CheckDB(); err != nil {
		Task.LogErr(err.Error())
		panic(err)
	}

	if err := Task.DbLite(); err != nil {
		Task.LogErr(err.Error())
		panic(err)
	}

	app := NewApp()
	layout := BuildLayoutShell(app)

	if err := app.UI.SetRoot(layout, true).Run(); err != nil {
		Task.LogErr(err.Error())
		panic(err)
	}
}

