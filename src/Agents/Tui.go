package Agents

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func BuildAgentBox(title, content string, width int) string {
	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("63")).
		Bold(true).
		Underline(true)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		Width(width).
		BorderForeground(lipgloss.Color("240"))

	return boxStyle.Render(titleStyle.Render(title) + "\n\n" + content)
}

func AgentView(width int) string {
	agent := GetAgentState("NULL-SKY")

	// Box 1 – Agent Info
	var box1 strings.Builder
	fmt.Fprintf(&box1, "Account ID:   %s\n", agent.Data.AccountID)
	fmt.Fprintf(&box1, "Agent Symbol: %s\n", agent.Data.Symbol)
	fmt.Fprintf(&box1, "Faction:      %s\n", agent.Data.Faction)
	fmt.Fprintf(&box1, "Headquarters: %s\n", agent.Data.HQ)
	fmt.Fprintf(&box1, "Ship Count:   %d\n", agent.Data.Ships)
	fmt.Fprintf(&box1, "Credits:      %d", agent.Data.Credits)

	// Create up to 9 boxes and arrange them horizontally/vertically
	row1 := lipgloss.JoinHorizontal(
		lipgloss.Left,
		BuildAgentBox("Agent Details", box1.String(), width/2))

	return lipgloss.JoinVertical(lipgloss.Left, row1)
}
