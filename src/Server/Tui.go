package Server

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type infoBox struct {
	Title string
	Body  string
}

func renderBox(b infoBox, width int) string {
	style := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		Width(width / 3).
		BorderForeground(lipgloss.Color("240"))

	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("63")).
		Bold(true).
		Underline(true).
		Render(b.Title)

	return style.Render(title + "\n\n" + b.Body)
}

func ServerView(width int) string {
	g := GetGameServerState()

	// Build box content with fmt.Sprintf for clarity
	state := fmt.Sprintf(
		"Status:    %s\nVersion:   %s\nResetFreq: %s\nLastReset: %s\nNextReset: %s\nLastCheck: %s",
		g.Status, g.Version, g.ServerResets.ResetFreq, g.LastReset, g.ServerResets.NextReset, g.LastCheckIn,
	)
	players := fmt.Sprintf(
		"Accounts:  %d\nAgents:    %d\nShips:     %d\nSystems:   %d\nWaypoints: %d\n",
		g.Stats.Accounts, g.Stats.Agents, g.Stats.Ships, g.Stats.Systems, g.Stats.Waypoints,
	)

	boxes := []infoBox{
		{"Game Server Status", state},
		{"Registered Agents", players},
		{"Leaderboard Credits", "Coming Soon..."},
		{"Leaderboard Charts", "Coming Soon..."},
	}

	// Render all boxes
	var rendered []string
	for _, b := range boxes {
		rendered = append(rendered, renderBox(b, width))
	}

	// Group boxes into rows of 3
	const cols = 2
	var rows []string
	for i := 0; i < len(rendered); i += cols {
		end := i + cols
		if end > len(rendered) {
			end = len(rendered)
		}
		row := lipgloss.JoinHorizontal(lipgloss.Top, rendered[i:end]...)
		rows = append(rows, row)
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
