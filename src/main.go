package main

import (
	"Spacetraders/src/Fleet"
	"Spacetraders/src/General"
	"Spacetraders/src/Server"
	"Spacetraders/src/Agents"
	"Spacetraders/src/Systems"
	"Spacetraders/src/Contracts"
	"Spacetraders/src/Settings"

	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

// ────────────────────────────────────────────────────────────────────────────────
// MAIN LAYOUT SHELL
// ────────────────────────────────────────────────────────────────────────────────
func NewApp() *General.App {
	return &General.App{
		UI:      tview.NewApplication(),
		UIState: &General.UIState{},
		State:   &General.GlobalState{
			Theme_BgBase:        tcell.GetColor("#22272E"),
			Theme_BgPanel:       tcell.GetColor("#444C56"),
			Theme_BgAlt:         tcell.GetColor("#2D333B"),
			Theme_FgBase:        tcell.GetColor("#C9D1D9"),
			Theme_FgMuted:       tcell.GetColor("#636E7B"),
			Theme_AccentBlue:    tcell.GetColor("#58A6FF"),
			Theme_AccentGreen:   tcell.GetColor("#56D364"),
			Theme_AccentYellow:  tcell.GetColor("#D29922"),
			Theme_AccentRed:     tcell.GetColor("#F85149"),
			Theme_AccentMagenta: tcell.GetColor("#BF4D80"),
		},
	}
}

// INFO: Register all of the standard menus with the app then we can just call and focus them as needed
func BuildLayoutShell(app *General.App) tview.Primitive {
	// MAIN MENU (top-left)
	mainMenu := tview.NewList()
	mainMenu.ShowSecondaryText(false)
	mainMenu.SetBorder(true)
	mainMenu.SetTitle("  Main Menu  ")
	mainMenu.AddItem("Server Status", "", 0, func() { Server.ShowServerMenu(app)         })
	mainMenu.AddItem("Agent Status",  "", 0, func() { Agents.ShowAgentsMenu(app)         })
	mainMenu.AddItem("Fleet Status",  "", 0, func() { Fleet.DisplayFleetMenu(app)        })
	mainMenu.AddItem("Systems",       "", 0, func() { Waypoints.ShowSystemsMenu(app)     })
	mainMenu.AddItem("Contracts",     "", 0, func() { Contracts.DisplayContractMenu(app) })
	mainMenu.AddItem("Settings",      "", 0, func() { Settings.ShowSettingsMenu(app)     })
	mainMenu.AddItem("Quit",          "", 0, func() { app.UI.Stop()                      })

	// SUBMENU (bottom-left, dynamic)
	subMenu := tview.NewList()
	subMenu.ShowSecondaryText(false)
	subMenu.SetBorder(true)
	subMenu.SetTitle("  Submenu  ")

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


// ────────────────────────────────────────────────────────────────────────────────
// MAIN
// ────────────────────────────────────────────────────────────────────────────────
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

