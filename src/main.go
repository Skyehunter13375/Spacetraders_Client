package main

import "Spacetraders/src/Fleet"
import "Spacetraders/src/General"
import "Spacetraders/src/Server"
import "Spacetraders/src/Agents"
import "Spacetraders/src/Systems"
import "Spacetraders/src/Contracts"
import "Spacetraders/src/Settings"
import "github.com/rivo/tview"

// ┌──────────────────────────────────────────────────────────────────────────────┐
// │                              Main Layout Shell                               │
// └──────────────────────────────────────────────────────────────────────────────┘
func NewApp() *General.App {
	return &General.App{
		UI:      tview.NewApplication(),
		UIState: &General.UIState{},
		State:   &General.GlobalState{},
	}
}

// INFO: Register all of the standard menus with the app then we can just call and focus them as needed
func BuildLayoutShell(app *General.App) tview.Primitive {
	// MAIN MENU (top-left)
	mainMenu := tview.NewList()
	mainMenu.ShowSecondaryText(false)
	mainMenu.SetBorder(true)
	mainMenu.SetTitle("  Main Menu  ")
	mainMenu.SetBorderColor(General.Theme.BgBorder)
	mainMenu.AddItem("Server Status", "", 0, func() { Server.ShowServerMenu(app)         })
	mainMenu.AddItem("Agent Status",  "", 0, func() { Agents.ShowAgentsMenu(app)         })
	mainMenu.AddItem("Fleet Status",  "", 0, func() { Fleet.DisplayFleetMenu(app)        })
	mainMenu.AddItem("Systems",       "", 0, func() { Waypoints.DisplaySystemMenu(app)   })
	mainMenu.AddItem("Contracts",     "", 0, func() { Contracts.DisplayContractMenu(app) })
	mainMenu.AddItem("Settings",      "", 0, func() { Settings.ShowSettingsMenu(app)     })
	mainMenu.AddItem("Quit",          "", 0, func() { app.UI.Stop()                      })

	// SUBMENU (bottom-left, dynamic)
	subMenu := tview.NewList()
	subMenu.ShowSecondaryText(false)
	subMenu.SetBorder(true)
	subMenu.SetTitle("  Submenu  ")
	subMenu.SetBorderColor(General.Theme.BgBorder)

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
	if err := General.DB(); err != nil {
		General.LogErr(err.Error())
		panic(err)
	}

	app := NewApp()
	layout := BuildLayoutShell(app)

	if err := app.UI.SetRoot(layout, true).Run(); err != nil {
		General.LogErr(err.Error())
		panic(err)
	}
}

