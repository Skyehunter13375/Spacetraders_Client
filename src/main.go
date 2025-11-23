package main

import (
	"Spacetraders/src/Fleet"
	"Spacetraders/src/General"
	"Spacetraders/src/Server"
	"Spacetraders/src/Agents"
	"Spacetraders/src/Settings"
	"github.com/rivo/tview"
)

// ────────────────────────────────────────────────────────────────────────────────
// APPLICATION STRUCTURES
// ────────────────────────────────────────────────────────────────────────────────
// Holds dynamic UI component references that screens will update.
type UIState struct {
	MainMenu *tview.List
	SubMenu  *tview.List
	Output   *tview.Flex
}

// Global data shared across screens (expand over time).
type GlobalState struct {
	// Ships map[string]*Fleet.Ship
	// Systems map[string]*System
	// etc.
}

type App struct {
	UI      *tview.Application
	UIState *UIState
	State   *GlobalState
}

func NewApp() *App {
	return &App{
		UI:      tview.NewApplication(),
		UIState: &UIState{},
		State:   &GlobalState{},
	}
}

// ────────────────────────────────────────────────────────────────────────────────
// MAIN LAYOUT SHELL
// ────────────────────────────────────────────────────────────────────────────────
func BuildLayoutShell(app *App) tview.Primitive {
	// MAIN MENU (top-left)
	mainMenu := tview.NewList()
	mainMenu.ShowSecondaryText(false)
	mainMenu.SetBorder(true)
	mainMenu.SetTitle(" Main Menu ")
	mainMenu.AddItem("Server",    "", '1', func() { app.UI.SetFocus(app.UIState.SubMenu); ShowServerMenu(app) })
	mainMenu.AddItem("Agents",    "", '2', func() { app.UI.SetFocus(app.UIState.SubMenu); ShowAgentsMenu(app) })
	mainMenu.AddItem("Ships",     "", '3', func() { app.UI.SetFocus(app.UIState.SubMenu); ShowShipsMenu(app) })
	mainMenu.AddItem("Systems",   "", '4', func() { app.UI.SetFocus(app.UIState.SubMenu); ShowSystemsMenu(app) })
	mainMenu.AddItem("Contracts", "", '5', func() { app.UI.SetFocus(app.UIState.SubMenu); ShowContractsMenu(app) })
	mainMenu.AddItem("Settings",  "", '9', func() { app.UI.SetFocus(app.UIState.SubMenu); ShowSettingsMenu(app) })
	mainMenu.AddItem("Quit",      "", 'q', func() { app.UI.Stop() })

	// SUBMENU (bottom-left, dynamic)
	subMenu := tview.NewList()
	subMenu.ShowSecondaryText(false)
	subMenu.SetBorder(true)
	subMenu.SetTitle(" Submenu ")

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
// SCREEN HANDLERS (UPDATE SUBMENU + OUTPUT)
// ────────────────────────────────────────────────────────────────────────────────
func FocusMain(app *App) {
	ui := app.UIState
	ui.SubMenu.Clear()
	ui.Output.Clear()
	app.UI.SetFocus(app.UIState.MainMenu)
}

func FocusSub(app *App) {
	ui := app.UIState
	ui.Output.Clear()
	app.UI.SetFocus(app.UIState.MainMenu)
}

func ShowServerMenu(app *App) {
	ui := app.UIState
	ui.SubMenu.Clear()
	ui.SubMenu.AddItem("Get Server Status", "", 0, func() {
		ui.Output.Clear()
		ui.Output.AddItem(Server.DisplayGameServerState() , 0, 1, false)
	})
	ui.SubMenu.AddItem("Back", "", 'b', func() { FocusMain(app) } )
}

func ShowAgentsMenu(app *App) {
	ui := app.UIState
	ui.SubMenu.Clear()
	ui.SubMenu.AddItem("Show Agent Info", "", 0, func() {
		ui.Output.Clear()
		ui.Output.AddItem(Agents.DisplayAgentState(), 0, 1, false)
	})
	ui.SubMenu.AddItem("Back", "", 'b', func() { FocusMain(app) } )
}

func ShowShipsMenu(app *App) {
	ui := app.UIState
	ui.SubMenu.Clear()
	ui.SubMenu.AddItem("List Ships", "", 0, func() {
		ui.Output.Clear()
		ui.Output.AddItem(Fleet.DisplayShipState() , 0, 1, false)
	})
	ui.SubMenu.AddItem("Back", "", 'b', func() { FocusMain(app) } )
}

func ShowSystemsMenu(app *App) {
	ui := app.UIState
	ui.SubMenu.Clear()
	ui.SubMenu.AddItem("Select System", "", 0, func() {
		ui.Output.Clear()
		ui.Output.AddItem(tview.NewTextView().
			SetText("Systems Placeholder"), 0, 1, false)
	})
	ui.SubMenu.AddItem("Back", "", 'b', func() { FocusMain(app) } )
}

func ShowContractsMenu(app *App) {
	ui := app.UIState
	ui.SubMenu.Clear()
	ui.SubMenu.AddItem("List Contracts", "", 0, func() {
		ui.Output.Clear()
		ui.Output.AddItem(tview.NewTextView().
			SetText("Contracts Placeholder"), 0, 1, false)
	})
	ui.SubMenu.AddItem("Back", "", 'b', func() { FocusMain(app) } )
}

func ShowSettingsMenu(app *App) {
	ui := app.UIState
	ui.SubMenu.Clear()
	ui.SubMenu.AddItem("Toggle Something", "", 0, func() {
		ui.Output.Clear()
		ui.Output.AddItem(Settings.DisplaySettings(), 0, 1, false)
	})
	ui.SubMenu.AddItem("Back", "", 'b', func() { FocusMain(app) } )
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

