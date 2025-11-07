package Contracts

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	menuItems     []string
	selectedIndex int
	content       string
	width, height int
	quitting      bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {

		case "up", "k":
			if m.selectedIndex > 0 {
				m.selectedIndex--
			}

		case "down", "j":
			if m.selectedIndex < len(m.menuItems)-1 {
				m.selectedIndex++
			}

		case "enter":
			// item := m.menuItems[m.selectedIndex]

		case "q", "esc":
			m.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return ""
	}

	var menuParts []string
	for i, item := range m.menuItems {
		style := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 2).
			BorderForeground(lipgloss.Color("240")).
			Foreground(lipgloss.Color("250"))
		if i == m.selectedIndex {
			style = style.
				Foreground(lipgloss.Color("46")).
				BorderForeground(lipgloss.Color("46"))
		}
		menuParts = append(menuParts, style.Render(item))
	}

	menuBar := lipgloss.JoinVertical(lipgloss.Left, menuParts...)

	body := lipgloss.JoinVertical(
		lipgloss.Left,
		menuBar,
		m.content,
	)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Left, lipgloss.Top,
		body,
	)
}

// =======================================================================================================
func ProgressBar(value, max int) string {
	bar := progress.New(
		progress.WithGradient("#045f04ff", "#28b328ff"),
	)
	bar.Width = 30

	var ratio float64
	if max > 0 {
		ratio = float64(value) / float64(max)
	} else {
		ratio = 1
	}

	// Clamp between 0 and 1 just in case
	if ratio < 0 {
		ratio = 0
	} else if ratio > 1 {
		ratio = 1
	}

	return bar.ViewAs(ratio)
}

func ContractView2(width int) string {
	CStruct := GetAllContracts()

	m := model{
		menuItems: []string{
			CStruct.Data[0].Faction + " | " +
				CStruct.Data[0].Type + " | " +
				CStruct.Data[0].DeadlineToAccept,
			"Example 2",
			"Example 3",
		},
		content: "Use ↑ ↓ to navigate, Enter to select.",
		width:   width,
		// height:  20, // optional default height
	}

	// Return static render
	return m.View()
}

func ContractView(width int) string {
	var sb strings.Builder
	CStruct := GetAllContracts()
	fmt.Fprintf(&sb, "ID:        %s\n", CStruct.Data[0].ID)
	fmt.Fprintf(&sb, "UpFront:   %d\n", CStruct.Data[0].Terms.Payment.OnAccepted)
	fmt.Fprintf(&sb, "OnCompl:   %d\n", CStruct.Data[0].Terms.Payment.OnFulfilled)
	fmt.Fprintf(&sb, "Accepted:  %t\n", CStruct.Data[0].Accepted)
	fmt.Fprintf(&sb, "Fulfilled: %t\n", CStruct.Data[0].Fulfilled)
	fmt.Fprintf(&sb, "Deadline:  %s\n", CStruct.Data[0].Terms.Deadline)
	fmt.Fprintf(&sb, "Expires:   %s\n", CStruct.Data[0].Expiration)
	fmt.Fprintf(&sb, "Material:  %d/%d (%s → %s)",
		CStruct.Data[0].Terms.Deliver[0].UnitsFulfilled,
		CStruct.Data[0].Terms.Deliver[0].UnitsRequired,
		CStruct.Data[0].Terms.Deliver[0].Material,
		CStruct.Data[0].Terms.Deliver[0].Destination,
	)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		Width(width / 2).
		BorderForeground(lipgloss.Color("240"))

	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("63")).
		Bold(true).
		Underline(true).
		Render(CStruct.Data[0].Faction + " | " + CStruct.Data[0].Type + " | " + CStruct.Data[0].DeadlineToAccept)

	return boxStyle.Render(title + "\n\n" + sb.String())
}
