package main

import (
	"Spacetraders/src/Agents"
	"Spacetraders/src/Fleet"
	"Spacetraders/src/General"
	"Spacetraders/src/Server"

	"github.com/rivo/tview"
)

type MenuItem struct {
	Name     string
	Action   func()
	FocusOut bool
}

func main() {
	// Must initialize the DB connection the first time here.
	if err := General.DB(); err != nil {
		General.LogErr(err.Error())
		panic(err)
	}

	app := tview.NewApplication()

	// MAIN MENU (Top-Left)
	mainMenu := tview.NewList()
	mainMenu.ShowSecondaryText(false).SetBorder(true).SetTitle(" Main Menu ")
	mainMenu.AddItem("Server", "", '1', nil)
	mainMenu.AddItem("Agents", "", '2', nil)
	mainMenu.AddItem("Ships", "", '3', nil)
	mainMenu.AddItem("Systems", "", '4', nil)
	mainMenu.AddItem("Contracts", "", '5', nil)
	mainMenu.AddItem("Quit", "", 'q', func() { app.Stop() })

	// SUB MENU (Bottom-Left)
	subMenu := tview.NewList()
	subMenu.ShowSecondaryText(false).SetBorder(true).SetTitle(" Submenu ")

	// OUTPUT BOX (Right Side)
	output := tview.NewFlex()
	output.SetBorder(false)

	// LEFT SIDE = main menu (top) + sub menu (bottom)
	leftSide := tview.NewFlex()
	leftSide.SetDirection(tview.FlexRow)
	leftSide.AddItem(mainMenu, 0, 1, true)
	leftSide.AddItem(subMenu, 0, 1, false)

	// FULL WINDOW = left column + output box
	window := tview.NewFlex()
	window.AddItem(leftSide, 30, 1, true)
	window.AddItem(output, 0, 2, false)

	// Function to load submenu items
	loadSubmenu := func(category string) {
		subMenu.Clear()

		var opts []MenuItem // ← FIXED HERE

		switch category {
		case "Server":
			opts = []MenuItem{
				{
					Name: "Get Server Status",
					Action: func() {
						output.Clear()
						output.AddItem(Server.DisplayGameServerState(), 0, 1, false)
					},
					FocusOut: false,
				},
			}

		case "Agents":
			opts = []MenuItem{
				{
					Name: "NULLSKY",
					Action: func() {
						output.Clear()
						output.AddItem(Agents.DisplayAgentState(), 0, 1, false)
					},
					FocusOut: false,
				},
			}

		case "Ships":
			opts = []MenuItem{
				{
					Name: "All",
					Action: func() {
						output.Clear()
						output.AddItem(Fleet.DisplayShipState(), 0, 1, true)
					},
					FocusOut: false,
				},
			}
		}

		// Add submenu choices
		for _, item := range opts {
			mi := item
			subMenu.AddItem(mi.Name, "", 0, func() {
				mi.Action()
				if mi.FocusOut {
					app.SetFocus(output)
				}
			})
		}

		subMenu.AddItem("Back", "", 'b', func() {
			subMenu.Clear()
			output.Clear()
			app.SetFocus(mainMenu)
		})

		app.SetFocus(subMenu)
	}

	// MAIN MENU HANDLER
	mainMenu.SetSelectedFunc(func(i int, name, second string, r rune) {
		if name == "Quit" {
			app.Stop()
			return
		}
		loadSubmenu(name)
	})

	// RUN APP
	if err := app.SetRoot(window, true).Run(); err != nil {
		General.LogErr(err.Error())
		panic(err)
	}
}
