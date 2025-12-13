package Tui 

import "fmt"
import "Spacetraders/src/Task"
import "Spacetraders/src/Model"
import "github.com/rivo/tview"

func ShowAgentsMenu(app *Model.App) {
	app.UIState.SubMenu.Clear()
	app.UIState.Output.Clear()
	app.UIState.Output.AddItem(DisplayAgentState(), 0, 1, false)
	app.UI.SetFocus(app.UIState.MainMenu)
}

func DisplayAgentState() tview.Primitive {
	// TASK: Get agent name from config file
	CFG, _ := Task.GetConfig()

	// TASK: Display the agent data
	agent  := Task.GetAgentState(CFG.API.AgentName)

	window := tview.NewFlex()
	window.SetBorder(false)
	window.SetDirection(tview.FlexRow)

	stats_1 := tview.NewFlex()
	stats_1.SetBorder(false)
	box_1 := tview.NewTextView()
	box_1.SetBorder(true)
	box_1.SetBorderColor(Model.Theme.BgBorder)
	fmt.Fprintf(box_1, "Account ID:   %s\n", agent.AccountID)
	fmt.Fprintf(box_1, "Agent Symbol: %s\n", agent.Symbol)
	fmt.Fprintf(box_1, "Faction:      %s\n", agent.Faction)
	fmt.Fprintf(box_1, "Headquarters: %s\n", agent.HQ)
	fmt.Fprintf(box_1, "Ship Count:   %d\n", agent.Ships)
	fmt.Fprintf(box_1, "Credits:      %d\n", agent.Credits)
	stats_1.AddItem(box_1, 0, 1, false)

	graph_1 := tview.NewTextArea()
	graph_1.SetBorder(true)
	graph_1.SetTitle("  Cash Flow History  ")
	graph_1.SetBorderColor(Model.Theme.BgBorder)
	graph_1.SetText("Coming Soon...", false)

	window.AddItem(stats_1, 0, 1, false)
	window.AddItem(graph_1, 0, 1, false)

	return window
}
