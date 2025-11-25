package Agents

import (
	"fmt"
	"Spacetraders/src/General"

	"github.com/rivo/tview"
)

func ShowAgentsMenu(app *General.App) {
	app.UIState.SubMenu.Clear()
	app.UIState.Output.Clear()
	app.UIState.Output.AddItem(DisplayAgentState(), 0, 1, false)
	app.UI.SetFocus(app.UIState.MainMenu)
}

func DisplayAgentState() tview.Primitive {
	agent := GetAgentState("NULL-SKY")

	window := tview.NewFlex()
	window.SetBorder(false)
	window.SetDirection(tview.FlexRow)

	stats_1 := tview.NewFlex()
	stats_1.SetBorder(false)
	box_1 := tview.NewTextView()
	box_1.SetBorder(true)
	fmt.Fprintf(box_1, "Account ID:   %s\n", agent.Data.AccountID)
	fmt.Fprintf(box_1, "Agent Symbol: %s\n", agent.Data.Symbol)
	fmt.Fprintf(box_1, "Faction:      %s\n", agent.Data.Faction)
	fmt.Fprintf(box_1, "Headquarters: %s\n", agent.Data.HQ)
	fmt.Fprintf(box_1, "Ship Count:   %d\n", agent.Data.Ships)
	fmt.Fprintf(box_1, "Credits:      %d\n", agent.Data.Credits)
	stats_1.AddItem(box_1, 0, 1, false)

	window.AddItem(stats_1, 0, 1, false)

	return window
}
