package Server

import (
	"fmt"

	"github.com/rivo/tview"
)

func DisplayGameServerState() tview.Primitive {
	g := GetGameServerState()
	l := GetLeaderboards()

	window := tview.NewFlex()
	window.SetBorder(false)
	window.SetDirection(tview.FlexRow)

	stats_1 := tview.NewFlex()
	stats_1.SetBorder(false)

	stats_2 := tview.NewFlex()
	stats_2.SetBorder(false)

	StateBox := tview.NewFlex()
	StateBox.SetBorder(true)
	StateBox.SetTitle("Game Server Status")
	StateInfo := tview.NewTextView()
	fmt.Fprintf(StateInfo, "Status:    %s\n", g.Status)
	fmt.Fprintf(StateInfo, "Version:   %s\n", g.Version)
	fmt.Fprintf(StateInfo, "ResetFreq: %s\n", g.ServerResets.ResetFreq)
	fmt.Fprintf(StateInfo, "LastReset: %s\n", g.LastReset)
	fmt.Fprintf(StateInfo, "NextReset: %s\n", g.ServerResets.NextReset)
	fmt.Fprintf(StateInfo, "LastCheck: %s\n", g.LastCheckIn)
	StateBox.AddItem(StateInfo, 0, 1, false)
	stats_1.AddItem(StateBox, 0, 1, false)

	PlayerBox := tview.NewFlex()
	PlayerBox.SetBorder(true)
	PlayerBox.SetTitle("Registered Player Stats")
	PlayerInfo := tview.NewTextView()
	fmt.Fprintf(PlayerInfo, "Accounts:  %d\n", g.Stats.Accounts)
	fmt.Fprintf(PlayerInfo, "Agents:    %d\n", g.Stats.Agents)
	fmt.Fprintf(PlayerInfo, "Ships:     %d\n", g.Stats.Ships)
	fmt.Fprintf(PlayerInfo, "Systems:   %d\n", g.Stats.Systems)
	fmt.Fprintf(PlayerInfo, "Waypoints: %d\n", g.Stats.Waypoints)
	PlayerBox.AddItem(PlayerInfo, 0, 1, false)
	stats_1.AddItem(PlayerBox, 0, 1, false)

	// INFO: Leaderboard credits box
	LeaderBox1 := tview.NewFlex()
	LeaderBox1.SetBorder(true)
	LeaderBox1.SetTitle("  Leaderboard (Credits)  ")
	CredsGrid := tview.NewGrid()
	CredsGrid.SetRows(1,1,1,1,1,1,1,1,1,1)        // Position | Agent | Credits
	CredsGrid.SetColumns(5, 25, 25) // Position | Agent | Credits
	CredsGrid.SetBorders(true)

	for i, v := range l.MostCredits {
		row := i // row index (0-based)

		// Position (1-based ranking)
		pos := tview.NewTextView().
			SetText(fmt.Sprintf("%d.", i+1)).
			SetTextAlign(tview.AlignRight)

		// Agent name
		agent := tview.NewTextView().
			SetText(v.Agent)

		// Credit count
		creds := tview.NewTextView().
			SetText(fmt.Sprintf("%d", v.Creds)).
			SetTextAlign(tview.AlignLeft)

		// Add the 3 columns in this row
		CredsGrid.AddItem(pos,   row, 0, 1, 1, 0, 0, false)
		CredsGrid.AddItem(agent, row, 1, 1, 1, 0, 0, false)
		CredsGrid.AddItem(creds, row, 2, 1, 1, 0, 0, false)
	}

	LeaderBox1.AddItem(CredsGrid, 0, 1, false)
	stats_2.AddItem(LeaderBox1, 0, 1, false)

	// INFO: Leaderboard charts box
	LeaderBox2 := tview.NewFlex()
	LeaderBox2.SetBorder(true)
	LeaderBox2.SetTitle("  Leaderboard (Charts)  ")
	ChartsGrid := tview.NewGrid()
	ChartsGrid.SetRows(1,1,1,1,1,1,1,1,1,1)
	ChartsGrid.SetColumns(5, 25, 20) // position | agent | credits
	ChartsGrid.SetBorders(true)

	for i, v := range l.MostCharted {
		row := i

		// Position (1-based ranking)
		pos := tview.NewTextView().
			SetText(fmt.Sprintf("%d.", i+1)).
			SetTextAlign(tview.AlignLeft)

		// Agent name
		agent := tview.NewTextView().
			SetText(v.Agent)

		// Chart count
		charts := tview.NewTextView().
			SetText(fmt.Sprintf("%d", v.Charts)).
			SetTextAlign(tview.AlignRight)

		// Add the 3 columns in this row
		ChartsGrid.AddItem(pos,    row, 0, 1, 1, 0, 0, false)
		ChartsGrid.AddItem(agent,  row, 1, 1, 1, 0, 0, false)
		ChartsGrid.AddItem(charts, row, 2, 1, 1, 0, 0, false)
	}

	LeaderBox2.AddItem(ChartsGrid, 0, 1, false)
	stats_2.AddItem(LeaderBox2, 0, 1, false)



	window.AddItem(stats_1, 0, 1, false)
	window.AddItem(stats_2, 0, 1, false)
	return window
}
