package Server

import (
	"fmt"

	"github.com/rivo/tview"
)

func DisplayGameServerState() tview.Primitive {
	g := GetGameServerState()

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

	LeaderBox1 := tview.NewFlex()
	LeaderBox1.SetBorder(true)
	LeaderBox1.SetTitle("  Leaderboard (Credits)  ")
	stats_2.AddItem(LeaderBox1, 0, 1, false)

	window.AddItem(stats_1, 0, 1, false)
	window.AddItem(stats_2, 0, 1, false)
	return window
}
