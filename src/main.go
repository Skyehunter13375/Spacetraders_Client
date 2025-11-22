package main

// import "fmt"
// import "os"
import "Spacetraders/src/Agents"
import "Spacetraders/src/Contracts"
import "Spacetraders/src/Fleet"
import "Spacetraders/src/General"
import "Spacetraders/src/Server"
// import "Spacetraders/src/Registration"
import "Spacetraders/src/Systems"
import "github.com/rivo/tview"

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
	mainMenu.AddItem("Server",    "", '1', nil)
	mainMenu.AddItem("Agents",    "", '2', nil)
	mainMenu.AddItem("Ships",     "", '3', nil)
	mainMenu.AddItem("Systems",   "", '4', nil)
	mainMenu.AddItem("Contracts", "", '5', nil)
	mainMenu.AddItem("Quit",      "", 'q', func() { app.Stop() })

	// SUB MENU (Bottom-Left)
	subMenu := tview.NewList()
	subMenu.ShowSecondaryText(false).SetBorder(true).SetTitle(" Submenu ")

	// OUTPUT BOX (Right Side)
	output := tview.NewFlex()
	output.SetBorder(false)

	// LEFT SIDE = main menu (top) + sub menu (bottom)
	leftSide := tview.NewFlex()
	leftSide.SetDirection(tview.FlexRow)
	leftSide.AddItem(mainMenu, 8, 1, true)
	leftSide.AddItem(subMenu,  0, 1, false)

	// FULL WINDOW = left column + output box
	window := tview.NewFlex()
	window.AddItem(leftSide, 30, 1, true)
	window.AddItem(output,    0, 2, false)

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
				// NewAgentToken, err := Registration.RegisterNewAgent("NULL_SKY", "VOID")
				// if err != nil { 
					// General.LogErr("Failed to register new agent: " + err.Error()) 
					// fmt.Println("Failed to register new agent: " + err.Error())
					// os.Exit(1)
				// } else {
					// fmt.Println("New agent registered: Replace your agent token in the config.yaml file: \n" + NewAgentToken)
					// os.Exit(0)
				// }
				// app.Stop()

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

			case "Systems":
				ids,_ := General.PG.Query(`SELECT symbol FROM systems`)
				var systems []string
				for ids.Next() {
					var s string
					ids.Scan(&s)
					systems = append(systems, s)
				}

				for i,v := range systems {
					opts = append(opts, MenuItem{
						Name:     systems[i],
						Action:   func() {
							output.Clear()
							output.AddItem(Waypoints.DisplaySystem(v), 0, 1, false)
						},
						FocusOut: false,
					})
				}

			case "Contracts":
				// Contracts.NegotiateNewContract("NULL-SKY-1")
				// Contracts.UpdateContracts()
				ids,_ := General.PG.Query(`SELECT id FROM contracts`)
				var contracts []string
				for ids.Next() {
					var s string
					ids.Scan(&s)
					contracts = append(contracts, s)
				}

				for i,v := range contracts {
					opts = append(opts, MenuItem{
						Name:     contracts[i],
						Action:   func() {
							output.Clear()
							output.AddItem(Contracts.DisplayContract(Contracts.GetContract(v)), 0, 1, false)
						},
						FocusOut: false,
					})
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
